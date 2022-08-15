/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vault

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/go-rootcerts"

	"golang.org/x/exp/slices"
)

// Client manages communication with the HashiCorp Vault API v1.12.0
// In most cases there should be only one, shared, Client.
type Client struct {
	configuration Configuration

	parsedBaseAddress url.URL

	client            *http.Client
	clientWithRetries *retryablehttp.Client

	// headers that will be added to each request
	requestHeaders     requestHeaders
	requestHeadersLock sync.RWMutex

	// API wrappers
	Auth     Auth
	Identity Identity
	Secrets  Secrets
	System   System
}

// requestHeaders contains headers that will be added to each request
type requestHeaders struct {
	token string
}

// NewClient returns a new Vault client with a copy of the given configuration
func NewClient(configuration Configuration) (*Client, error) {
	c := Client{
		configuration: configuration,

		// configured or default HTTP client
		client: configuration.HTTPClient,

		// retryablehttp wrapper around the HTTP client
		clientWithRetries: &retryablehttp.Client{
			HTTPClient:   configuration.HTTPClient,
			Logger:       configuration.RetryOptions.Logger,
			RetryWaitMin: configuration.RetryOptions.RetryWaitMin,
			RetryWaitMax: configuration.RetryOptions.RetryWaitMax,
			RetryMax:     configuration.RetryOptions.RetryMax,
			CheckRetry:   configuration.RetryOptions.CheckRetry,
			Backoff:      configuration.RetryOptions.Backoff,
			ErrorHandler: configuration.RetryOptions.ErrorHandler,
		},
	}

	a, err := url.Parse(configuration.BaseAddress)
	if err != nil {
		return nil, err
	}

	c.parsedBaseAddress = *a

	c.Auth = Auth{
		client: &c,
	}
	c.Identity = Identity{
		client: &c,
	}
	c.Secrets = Secrets{
		client: &c,
	}
	c.System = System{
		client: &c,
	}

	return &c, nil
}

// SetToken sets a token to be used with all subsequent requests
func (c *Client) SetToken(token string) {
	/* */ c.requestHeadersLock.Lock()
	defer c.requestHeadersLock.Unlock()

	c.requestHeaders.token = token
}

// WithToken returns a shallow copy of the client with the token set to the given value
func (c *Client) WithToken(token string) *Client {
	copy := c.shallowCopy()
	copy.requestHeaders.token = token

	return copy
}

// NewStructuredRequest expects json.Marshaler encoded request body and returns a new request with vault-specific headers
func (c *Client) NewStructuredRequest(method, path string, body json.Marshaler) (*http.Request, error) {
	if body == nil {
		return c.NewRequest(method, path, nil)
	}

	buf := bytes.Buffer{}

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, fmt.Errorf("could not encode request body: %w", err)
	}

	return c.NewRequest(method, path, &buf)
}

// NewRequest returns a new request with vault-specific headers
func (c *Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	// concatenate the base address with the given path
	url, err := c.parsedBaseAddress.Parse(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	headers := c.copyRequestHeaders()

	if headers.token != "" {
		req.Header.Set("X-Vault-Token", headers.token)
	}

	return req, nil
}

// Do sends the given request to Vault, handling retries, redirects, and rate limiting
func (c *Client) Do(ctx context.Context, req *http.Request, retry bool) (*http.Response, error) {
	// block on the rate limiter, if set
	if c.configuration.RateLimiter != nil {
		c.configuration.RateLimiter.Wait(ctx)
	}

	var (
		resp *http.Response
		err  error
	)

	req = req.WithContext(ctx)

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

// shallowCopy returns a shallow copy of the client
func (c *Client) shallowCopy() *Client {
	copy := Client{
		configuration:     c.configuration,
		parsedBaseAddress: c.parsedBaseAddress,
		client:            c.client,
		clientWithRetries: c.clientWithRetries,
	}

	copy.requestHeaders = c.copyRequestHeaders()
	copy.requestHeadersLock = sync.RWMutex{}

	copy.Auth = Auth{
		client: &copy,
	}
	copy.Identity = Identity{
		client: &copy,
	}
	copy.Secrets = Secrets{
		client: &copy,
	}
	copy.System = System{
		client: &copy,
	}

	return &copy
}

// copyRequestHeaders returns a copy of the request headers behind a mutex
func (c *Client) copyRequestHeaders() requestHeaders {
	/* */ c.requestHeadersLock.RLock()
	defer c.requestHeadersLock.RUnlock()

	return c.requestHeaders
}

// applyTLSOptions applies TLSOptions to the given tls.Config object
func applyTLSOptions(out *tls.Config, options TLSOptions) error {
	switch {
	case len(options.ClientCertFile) != 0 && len(options.ClientCertKey) != 0:
		clientCert, err := tls.LoadX509KeyPair(
			options.ClientCertFile,
			options.ClientCertKey,
		)
		if err != nil {
			return err
		}

		// We use this function to ignore the server's preferential list of
		// CAs, otherwise any CA used for the cert auth backend must be in the
		// server's CA pool
		out.GetClientCertificate = func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
			return &clientCert, nil
		}
	case len(options.ClientCertFile) != 0 || len(options.ClientCertKey) != 0:
		return fmt.Errorf("both client cert and client key must be provided")
	}

	if len(options.CAFile) != 0 || len(options.CACertificate) != 0 || len(options.CAPath) != 0 {
		rootCertConfig := &rootcerts.Config{
			CAFile:        options.CAFile,
			CACertificate: options.CACertificate,
			CAPath:        options.CAPath,
		}
		if err := rootcerts.ConfigureTLS(out, rootCertConfig); err != nil {
			return err
		}
	}

	if options.InsecureSkipVerify {
		out.InsecureSkipVerify = options.InsecureSkipVerify
	}

	if options.TLSServerName != "" {
		out.ServerName = options.TLSServerName
	}

	return nil
}
