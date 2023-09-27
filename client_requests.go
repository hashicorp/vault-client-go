// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

// Read attempts to read the value stored at the given Vault path.
func (c *Client) Read(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		c,
		http.MethodGet,
		v1Path(path),
		nil, // request body
		requestModifiers.additionalQueryParameters,
		requestModifiers,
	)
}

// ReadRaw attempts to read the value stored at the given Vault path and returns
// a raw *http.Response. Compared with `Read`, this function:
//   - does not parse the response
//   - does not check the response for errors
//   - does not apply the client-level request timeout
func (c *Client) ReadRaw(ctx context.Context, path string, options ...RequestOption) (*http.Response, error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	return sendRequestReturnRawResponse(
		ctx,
		c,
		http.MethodGet,
		v1Path(path),
		nil, // request body
		requestModifiers.additionalQueryParameters,
		requestModifiers,
	)
}

// Write attempts to write the given map to the given Vault path.
func (c *Client) Write(ctx context.Context, path string, body map[string]interface{}, options ...RequestOption) (*Response[map[string]interface{}], error) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, fmt.Errorf("could not encode request body: %w", err)
	}

	return c.WriteFromReader(ctx, path, &buf, options...)
}

// WriteFromBytes attempts to write the given byte slice to the given Vault path.
func (c *Client) WriteFromBytes(ctx context.Context, path string, body []byte, options ...RequestOption) (*Response[map[string]interface{}], error) {
	return c.WriteFromReader(ctx, path, bytes.NewReader(body), options...)
}

// WriteFromReader attempts to write the given io.Reader data to the given Vault path.
func (c *Client) WriteFromReader(ctx context.Context, path string, body io.Reader, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		c,
		http.MethodPost,
		v1Path(path),
		body,
		requestModifiers.additionalQueryParameters,
		requestModifiers,
	)
}

// List attempts to list the keys stored at the given Vault path.
func (c *Client) List(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	requestQueryParameters := requestModifiers.additionalQueryParametersOrDefault()
	requestQueryParameters.Add("list", "true")

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		c,
		http.MethodGet,
		v1Path(path),
		nil, // request body
		requestQueryParameters,
		requestModifiers,
	)
}

// Delete attempts to delete the value stored at the given Vault path.
func (c *Client) Delete(ctx context.Context, path string, options ...RequestOption) (*Response[map[string]interface{}], error) {
	requestModifiers, err := requestOptionsToRequestModifiers(options)
	if err != nil {
		return nil, err
	}

	return sendRequestParseResponse[map[string]interface{}](
		ctx,
		c,
		http.MethodDelete,
		v1Path(path),
		nil, // request body
		requestModifiers.additionalQueryParameters,
		requestModifiers,
	)
}

// sendStructuredRequestParseResponse constructs a structured request, sends it, and parses the response
func sendStructuredRequestParseResponse[ResponseT any](
	ctx context.Context,
	client *Client,
	method string,
	path string,
	body any,
	parameters url.Values,
	requestModifiers requestModifiers,
) (*Response[ResponseT], error) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, fmt.Errorf("could not encode request body: %w", err)
	}

	return sendRequestParseResponse[ResponseT](
		ctx,
		client,
		method,
		path,
		&buf,
		parameters,
		requestModifiers,
	)
}

// sendRequestParseResponse constructs a request, sends it, and parses the response
func sendRequestParseResponse[ResponseT any](
	ctx context.Context,
	client *Client,
	method string,
	path string,
	body io.Reader,
	parameters url.Values,
	requestModifiers requestModifiers,
) (*Response[ResponseT], error) {
	// apply the client-level request timeout, if set
	if client.configuration.RequestTimeout > 0 {
		var cancelContextFunc context.CancelFunc
		ctx, cancelContextFunc = context.WithTimeout(ctx, client.configuration.RequestTimeout)
		defer cancelContextFunc()
	}

	// clone the client-level request modifiers to prevent race conditions
	modifiers := client.cloneClientRequestModifiers()

	// merge in the request-level request modifiers
	mergeRequestModifiers(&modifiers, &requestModifiers)

	req, err := client.newRequest(ctx, method, path, body, parameters, modifiers.headers)
	if err != nil {
		return nil, err
	}

	resp, err := client.send(ctx, req, true, modifiers.requestCallbacks, modifiers.responseCallbacks)
	if err != nil || resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := isResponseError(req, resp); err != nil {
		return nil, err
	}

	return parseResponse[ResponseT](resp.Body)
}

// sendRequestReturnRawResponse constructs a request, sends it and returns a raw
// *http.Response. Compared with sendRequestParseResponse, this function:
//   - does not parse the response
//   - does not check the response for errors
//   - does not apply the client-level request timeout
func sendRequestReturnRawResponse(
	ctx context.Context,
	client *Client,
	method string,
	path string,
	body io.Reader,
	parameters url.Values,
	requestModifiers requestModifiers,
) (*http.Response, error) {
	// clone the client-level request modifiers to prevent race conditions
	modifiers := client.cloneClientRequestModifiers()

	// merge in the request-level request modifiers
	mergeRequestModifiers(&modifiers, &requestModifiers)

	req, err := client.newRequest(ctx, method, path, body, parameters, modifiers.headers)
	if err != nil {
		return nil, err
	}

	return client.send(ctx, req, true, modifiers.requestCallbacks, modifiers.responseCallbacks)
}

// newRequest constructs a new request with Vault-specific headers
func (c *Client) newRequest(
	ctx context.Context,
	method string,
	path string,
	body io.Reader,
	parameters url.Values,
	headers requestHeaders,
) (*http.Request, error) {
	// concatenate the base address with the given path
	url, err := c.parsedBaseAddress.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("could not join %q with the base address: %w", path, err)
	}

	// add query parameters (if any)
	if len(parameters) != 0 {
		url.RawQuery = parameters.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, url.String(), body)
	if err != nil {
		return nil, fmt.Errorf("could not create '%s %s' request: %w", method, url.String(), err)
	}

	// populate request headers
	if headers.customHeaders != nil {
		req.Header = headers.customHeaders
	}

	if headers.userAgent != "" {
		req.Header.Add("User-Agent", headers.userAgent)
	}

	if headers.token != "" {
		req.Header.Set("X-Vault-Token", headers.token)
	}

	if headers.namespace != "" {
		req.Header.Set("X-Vault-Namespace", headers.namespace)
	}

	for _, credentials := range headers.mfaCredentials {
		req.Header.Add("X-Vault-MFA", credentials)
	}

	if headers.responseWrappingTTL != 0 {
		req.Header.Set("X-Vault-Wrap-TTL", headers.responseWrappingTTL.String())
	}

	switch headers.replicationForwardingMode {
	// unconditional forwarding (see 'ReplicationForwardAlways' docs)
	case ReplicationForwardAlways:
		req.Header.Set("X-Vault-Forward", "active-node")

	// conditional formwarding (see 'ReplicationForwardInconsistent' docs)
	case ReplicationForwardInconsistent:
		req.Header.Set("X-Vault-Inconsistent", "forward-active-node")
	}

	return req, nil
}

// send sends the given request to Vault, handling retries, redirects, and rate limiting
func (c *Client) send(
	ctx context.Context,
	req *http.Request,
	retry bool,
	requestCallbacks []RequestCallback,
	responseCallbacks []ResponseCallback,
) (*http.Response, error) {
	// block on the rate limiter, if set
	if c.configuration.RateLimiter != nil {
		c.configuration.RateLimiter.Wait(ctx)
	}

	if c.configuration.EnforceReadYourWritesConsistency {
		c.replicationStates.requireReplicationStates(req)
	}

	// invoke request callbacks
	for _, callback := range requestCallbacks {
		callback(req)
	}

	var (
		resp *http.Response
		err  error
	)

	// allow at most one redirect to prevent redirect loops
	redirectCount := 1

	for {
		resp, err = c.do(req, retry)
		if err != nil {
			return resp, err
		}

		redirect, err := c.handleRedirect(req, resp, &redirectCount)
		if err != nil {
			return nil, err
		}
		if !redirect {
			break
		}
	}

	// invoke response callbacks
	for _, callback := range responseCallbacks {
		callback(req, resp)
	}

	if c.configuration.EnforceReadYourWritesConsistency && resp != nil {
		c.replicationStates.recordReplicationState(resp)
	}

	return resp, nil
}

// do is a helper function that wraps http.client.Do / retryablehttp.client.Do
func (c *Client) do(req *http.Request, retry bool) (*http.Response, error) {
	// In the vast majority of cases, the retryablehttp client will be used.
	// However, the retryablehttp client reads the entire request's body into
	// an internal byte array, which could cause particularly large requests
	// (e.g. raft snapshot requests) to run out of memory. For such requests,
	// a regular http client is used instead.
	if !retry {
		return c.client.Do(req)
	}

	retryableReq, err := retryablehttp.FromRequest(req)
	if err != nil {
		return nil, fmt.Errorf("could not form a retryable request: %w", err)
	}

	return c.clientWithRetries.Do(retryableReq)
}

// handleRedirect checks the given response for a redirect status & modifies
// the request accordingly if the redirect is needed
func (c *Client) handleRedirect(req *http.Request, resp *http.Response, redirectCount *int) (bool, *RedirectError) {
	switch resp.StatusCode {
	case
		http.StatusMovedPermanently,  // 301
		http.StatusFound,             // 302
		http.StatusTemporaryRedirect: // 307
	default:
		return false, nil
	}

	// a helper function to form a redirect error
	redirectError := func(redirectTo *url.URL, message string, args ...any) *RedirectError {
		var redirectLocation string
		if redirectTo != nil {
			redirectLocation = redirectTo.String()
		}
		return &RedirectError{
			StatusCode:       resp.StatusCode,
			Message:          fmt.Sprintf(message, args...),
			RedirectLocation: redirectLocation,
			OriginalRequest:  req,
		}
	}

	redirectTo, err := resp.Location()
	if err != nil {
		return false, redirectError(nil, "could not read the redirect location: %s", err.Error())
	}

	if c.configuration.DisableRedirects {
		return false, redirectError(redirectTo, "the redirects are disabled")
	}

	if *redirectCount <= 0 {
		return false, redirectError(redirectTo, "at most one redirect is allowed")
	}

	*redirectCount--

	if req.URL.Scheme == "https" && redirectTo.Scheme != "https" {
		return false, redirectError(redirectTo, "the redirect would cause a protocol downgrade")
	}

	// restore the original request body (if any) since it had been consumed by client.Do
	if req.GetBody != nil {
		b, err := req.GetBody()
		if err != nil {
			return false, redirectError(redirectTo, "could not restore the request body: %s", err.Error())
		}
		req.Body = b
	}

	req.URL = redirectTo

	return true, nil
}

// v1Path returns the path string prefixed with /v1/ if needed
func v1Path(path string) string {
	switch {
	case strings.HasPrefix(path, "/v1/"):
		return path
	case strings.HasPrefix(path, "v1/"):
		return "/" + path
	case strings.HasPrefix(path, "/"):
		return "/v1" + path
	default:
		return "/v1/" + path
	}
}
