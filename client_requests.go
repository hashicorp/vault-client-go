package vault

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"

	"golang.org/x/exp/slices"
)

// sendStructuredRequestParseResponse is a helper function to construct a
// json.Marshaler encoded request, send it to Vault and parse the response
func sendStructuredRequestParseResponse[ResponseT any](ctx context.Context, client *Client, method, path string, body json.Marshaler, parameters url.Values) (*Response[ResponseT], error) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, fmt.Errorf("could not encode request body: %w", err)
	}

	return sendRequestParseResponse[ResponseT](ctx, client, method, path, &buf, parameters)
}

// sendRequestParseResponse is a helper function to construct a request, send
// it to Vault and parse the response
func sendRequestParseResponse[ResponseT any](ctx context.Context, client *Client, method, path string, body io.Reader, parameters url.Values) (*Response[ResponseT], error) {
	req, err := client.newRequest(ctx, method, path, body, parameters)
	if err != nil {
		return nil, err
	}

	resp, err := client.do(ctx, req, true)
	if err != nil || resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	return parseResponse[ResponseT](resp.Body)
}

// newRequest returns a new request with vault-specific headers
func (c *Client) newRequest(ctx context.Context, method, path string, body io.Reader, parameters url.Values) (*http.Request, error) {
	// concatenate the base address with the given path
	url, err := c.parsedBaseAddress.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("could not join %q with the base address: %w", path, err)
	}

	// If SRV records exist, lookup the SRV record and take the highest match.
	// This is not designed for high availability, just discovery. Internet
	// draft (https://datatracker.ietf.org/doc/html/draft-andrews-http-srv-02)
	// specifies that the SRV record is ignored if a port is given.
	if c.configuration.EnableSRVLookup && url.Port() == "" {
		_, addrs, err := net.LookupSRV("http", "tcp", url.Hostname())
		// don't return the error to the user, address might not have an SRV record
		if err == nil && len(addrs) > 0 {
			url.Host = fmt.Sprintf("%s:%d", addrs[0].Target, addrs[0].Port)
		}
	}

	// add query parameters (if any)
	if len(parameters) != 0 {
		url.RawQuery = parameters.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, url.String(), body)
	if err != nil {
		return nil, fmt.Errorf("could not create '%s %s' request: %w", method, url.String(), err)
	}

	m := c.cloneRequestModifiers()

	if m.validationError != nil {
		return nil, m.validationError
	}

	if m.headers.token != "" {
		req.Header.Set("X-Vault-Token", m.headers.token)
	}

	if m.headers.namespace != "" {
		req.Header.Set("X-Vault-Namespace", m.headers.namespace)
	}

	for _, credentials := range m.headers.mfaCredentials {
		req.Header.Add("X-Vault-MFA", credentials)
	}

	switch m.headers.replicationForwardingMode {
	// unconditionally forwarding (see 'ReplicationForwardAlways' docs)
	case ReplicationForwardAlways:
		req.Header.Set("X-Vault-Forward", "active-node")

	// conditional formwarding (see 'ReplicationForwardInconsistent' docs)
	case ReplicationForwardInconsistent:
		req.Header.Set("X-Vault-Inconsistent", "forward-active-node")
	}

	return req, nil
}

// do sends the given request to Vault, handling retries, redirects, and rate limiting
func (c *Client) do(ctx context.Context, req *http.Request, retry bool) (*http.Response, error) {
	// block on the rate limiter, if set
	if c.configuration.RateLimiter != nil {
		c.configuration.RateLimiter.Wait(ctx)
	}

	if c.configuration.EnforceReadYourWritesConsistency {
		c.replicationStates.requireReplicationStates(req)
	}

	// clone request modifiers behind a lock
	m := c.cloneRequestModifiers()

	// invoke request callbacks
	for _, callback := range m.requestCallbacks {
		callback(req)
	}

	var (
		resp *http.Response
		err  error
	)

	// allow at most one redirect
	redirectCount := 0

	for {
		resp, err = c.doWithRetries(req, retry)
		if err != nil {
			return resp, err
		}

		redirect, err := handleRedirect(req, resp, &redirectCount)
		if err != nil {
			return resp, fmt.Errorf("redirect error: %w", err)
		}
		if !redirect {
			break
		}
	}

	// invoke response callbacks
	for _, callback := range m.responseCallbacks {
		callback(req, resp)
	}

	if c.configuration.EnforceReadYourWritesConsistency && resp != nil {
		c.replicationStates.recordReplicationState(resp)
	}

	return resp, nil
}

// doWithRetries is a helper function that wraps http.client.Do / retryablehttp.client.Do
func (c *Client) doWithRetries(req *http.Request, retry bool) (*http.Response, error) {
	if !retry {
		return c.client.Do(req)
	}

	retryableReq, err := retryablehttp.FromRequest(req)
	if err != nil {
		return nil, fmt.Errorf("could not form a retryable request: %w", err)
	}

	return c.clientWithRetries.Do(retryableReq)
}

// handleRedirect checks the given response for a redirect status
//  returns:
//    true & modifies the request accordingly if the redirect is needed
//    false otherwise
func handleRedirect(req *http.Request, resp *http.Response, redirectCount *int) (bool, error) {
	// allow at most one redirect
	if *redirectCount != 0 {
		return false, nil
	}

	*redirectCount++

	redirectStatuses := [...]int{
		http.StatusMovedPermanently,  // 301
		http.StatusFound,             // 302
		http.StatusTemporaryRedirect, // 307
	}

	if !slices.Contains(redirectStatuses[:], resp.StatusCode) {
		return false, nil
	}

	redirectTo, err := resp.Location()
	if err != nil {
		return false, fmt.Errorf("could not read the redirect location: %w", err)
	}

	if req.URL.Scheme == "https" && redirectTo.Scheme != "https" {
		return false, fmt.Errorf("redirect would cause a protocol downgrade")
	}

	req.URL = redirectTo

	// restore the original request body (if any) since it had been consumed by client.Do
	if req.GetBody != nil {
		b, err := req.GetBody()
		if err != nil {
			return false, fmt.Errorf("could not restore request body: %w", err)
		}
		req.Body = b
	}

	return true, nil
}
