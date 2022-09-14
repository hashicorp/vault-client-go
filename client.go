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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"unicode"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-retryablehttp"

	"golang.org/x/exp/slices"
)

// Client manages communication with the HashiCorp Vault API v1.12.0
// In most cases there should be only one, shared, Client.
type Client struct {
	configuration Configuration

	parsedBaseAddress url.URL

	client            *http.Client
	clientWithRetries *retryablehttp.Client

	// headers, callbacks, etc. that will be added to each request
	requestModifiers     requestModifiers
	requestModifiersLock sync.RWMutex

	// API wrappers
	Auth     Auth
	Identity Identity
	Secrets  Secrets
	System   System
}

type (
	RequestCallback  func(*http.Request)
	ResponseCallback func(*http.Request, *http.Response)
)

// requestModifiers contains headers, callbacks, etc. that will be added to
// each request
type requestModifiers struct {
	headers requestHeaders

	requestCallbacks  []RequestCallback
	responseCallbacks []ResponseCallback

	// This error is set in client.WithX methods and checked in client.newRequest.
	// Since client.WithX methods are used for method chaining, they cannot
	// return errors.
	validationError error
}

// requestHeaders contains headers that will be added to each request
type requestHeaders struct {
	token         string // request header 'X-Vault-Token'
	namespace     string // request header 'X-Vault-Namespace'
	customHeaders http.Header
}

// NewClient returns a new Vault client with a copy of the given configuration
func NewClient(configuration Configuration) (*Client, error) {
	// Ensure that the configuration fields are initialized
	configuration.SetDefaultsForUninitialized()

	c := Client{
		configuration: configuration,

		// configured or default HTTP client
		client: configuration.BaseClient,

		// retryablehttp wrapper around the HTTP client
		clientWithRetries: &retryablehttp.Client{
			HTTPClient:   configuration.BaseClient,
			Logger:       configuration.Retry.Logger,
			RetryWaitMin: configuration.Retry.RetryWaitMin,
			RetryWaitMax: configuration.Retry.RetryWaitMax,
			RetryMax:     configuration.Retry.RetryMax,
			CheckRetry:   configuration.Retry.CheckRetry,
			Backoff:      configuration.Retry.Backoff,
			ErrorHandler: configuration.Retry.ErrorHandler,
		},

		requestHeaders: requestHeaders{
			token:         configuration.InitialToken,
			namespace:     configuration.InitialNamespace,
			customHeaders: nil,
		},
		requestHeadersLock: sync.RWMutex{},
	}

	a, err := url.Parse(configuration.BaseAddress)
	if err != nil {
		return nil, err
	}
	c.parsedBaseAddress = *a

	transport, ok := c.client.Transport.(*http.Transport)
	if !ok {
		return nil, fmt.Errorf("the configured base client's transport (%T) is not of type *http.Transport", c.client.Transport)
	}

	if err := configuration.TLS.applyTo(transport.TLSClientConfig); err != nil {
		return nil, err
	}

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

// Clone creates a new Vault client with the same configuration as the original
// client. Note that the cloned Vault client will point to the same base
// http.Client and retryablehttp.Client objects.
func (c *Client) Clone() *Client {
	clone := Client{
		configuration:     c.configuration,
		parsedBaseAddress: c.parsedBaseAddress,
		client:            c.client,
		clientWithRetries: c.clientWithRetries,
	}

	clone.requestModifiers = c.cloneRequestModifiers()

	clone.Auth = Auth{
		client: &clone,
	}
	clone.Identity = Identity{
		client: &clone,
	}
	clone.Secrets = Secrets{
		client: &clone,
	}
	clone.System = System{
		client: &clone,
	}

	return &clone
}

// cloneRequestModifiers returns a copy of the request modifiers behind a mutex
func (c *Client) cloneRequestModifiers() requestModifiers {
	/* */ c.requestModifiersLock.RLock()
	defer c.requestModifiersLock.RUnlock()

	var clone requestModifiers

	copy(clone.requestCallbacks, c.requestModifiers.requestCallbacks)
	copy(clone.responseCallbacks, c.requestModifiers.responseCallbacks)

	clone.headers = requestHeaders{
		token:         c.requestModifiers.headers.token,
		namespace:     c.requestModifiers.headers.namespace,
		customHeaders: c.requestModifiers.headers.customHeaders.Clone(),
	}

	return clone
}

// Configuration returns a copy of the configuration object used to initialize
// this client
func (c *Client) Configuration() Configuration {
	return c.configuration
}

// SetToken sets the token to be used with all subsequent requests.
// See https://www.vaultproject.io/docs/concepts/tokens for more info on
// tokens.
func (c *Client) SetToken(token string) error {
	if err := validateToken(token); err != nil {
		return err
	}

	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.token = token
	c.requestModifiersLock.Unlock()

	return nil
}

// ClearToken clears the token for all subsequent requests.
// See https://www.vaultproject.io/docs/concepts/tokens for more info on
// tokens.
func (c *Client) ClearToken() {
	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.token = ""
	c.requestModifiersLock.Unlock()
}

// WithToken returns a shallow copy of the client with the token set to the
// given value.
// See https://www.vaultproject.io/docs/concepts/tokens for more info on
// tokens.
func (c *Client) WithToken(token string) *Client {
	clone := c.Clone()

	if err := validateToken(token); err != nil {
		clone.requestModifiers.validationError = multierror.Append(clone.requestModifiers.validationError, err)
	} else {
		clone.requestModifiers.headers.token = token
	}

	return clone
}

// SetNamespace sets the namespace to be used with all subsequent requests,
// set to "" to clear the namespace.
// See https://www.vaultproject.io/docs/enterprise/namespaces for more info on
// namespaces.
func (c *Client) SetNamespace(namespace string) error {
	if err := validateNamespace(namespace); err != nil {
		return err
	}

	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.namespace = namespace
	c.requestModifiersLock.Unlock()

	return nil
}

// ClearNamespace clears the namespace from all subsequent requests.
// See https://www.vaultproject.io/docs/enterprise/namespaces for more info on
// namespaces.
func (c *Client) ClearNamespace() {
	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.namespace = ""
	c.requestModifiersLock.Unlock()
}

// WithNamespace returns a shallow copy of the client with the namespace set to
// the given value, use "" to clear the namespace.
// See https://www.vaultproject.io/docs/enterprise/namespaces for more info on
// namespaces.
func (c *Client) WithNamespace(namespace string) *Client {
	clone := c.Clone()

	if err := validateNamespace(namespace); err != nil {
		clone.requestModifiers.validationError = multierror.Append(clone.requestModifiers.validationError, err)
	} else {
		clone.requestModifiers.headers.namespace = namespace
	}

	return clone
}

// SetCustomHeaders sets custom headers to be used in all subsequent requests.
// The internal prefix 'X-Vault-' is not permitted for the header keys.
func (c *Client) SetCustomHeaders(headers http.Header) error {
	if err := validateCustomHeaders(headers); err != nil {
		return err
	}

	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.customHeaders = headers.Clone()
	c.requestModifiersLock.Unlock()

	return nil
}

// ClearsCustomHeaders clears all custom headers from the subsequent requests.
func (c *Client) ClearCustomHeaders() {
	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.customHeaders = nil
	c.requestModifiersLock.Unlock()
}

// WithCustomHeaders returns a shallow copy of the client with custom headers
// set to the given value, use nil to clear out the headers.
func (c *Client) WithCustomHeaders(headers http.Header) *Client {
	clone := c.Clone()

	if err := validateCustomHeaders(headers); err != nil {
		clone.requestModifiers.validationError = multierror.Append(clone.requestModifiers.validationError, err)
	} else {
		clone.requestModifiers.headers.customHeaders = headers.Clone()
	}

	return clone
}

// SetRequestCallbacks sets callbacks which will be invoked before each request.
func (c *Client) SetRequestCallbacks(callbacks ...RequestCallback) error {
	c.requestModifiersLock.Lock()
	copy(c.requestModifiers.requestCallbacks, callbacks)
	c.requestModifiersLock.Unlock()

	return nil
}

// ClearRequestCallbacks clears all request callbacks.
func (c *Client) ClearRequestCallbacks() {
	c.requestModifiersLock.Lock()
	c.requestModifiers.requestCallbacks = nil
	c.requestModifiersLock.Unlock()
}

// WithRequestCallbacks returns a shallow copy of the client with request
// callbacks set to the given value, use nil to clear out the callbacks.
func (c *Client) WithRequestCallbacks(callbacks ...RequestCallback) *Client {
	clone := c.Clone()
	copy(clone.requestModifiers.requestCallbacks, callbacks)

	return clone
}

// SetResponseCallbacks sets callbacks which will be invoked after each
// successful response.
func (c *Client) SetResponseCallbacks(callbacks ...ResponseCallback) error {
	c.requestModifiersLock.Lock()
	copy(c.requestModifiers.responseCallbacks, callbacks)
	c.requestModifiersLock.Unlock()

	return nil
}

// ClearResponseCallbacks clears all response callbacks.
func (c *Client) ClearResponseCallbacks() {
	c.requestModifiersLock.Lock()
	c.requestModifiers.responseCallbacks = nil
	c.requestModifiersLock.Unlock()
}

// WithResponseCallbacks returns a shallow copy of the client with response
// callbacks set to the given value, use nil to clear out the callbacks.
func (c *Client) WithResponseCallbacks(callbacks ...ResponseCallback) *Client {
	clone := c.Clone()
	copy(clone.requestModifiers.responseCallbacks, callbacks)

	return clone
}

// newStructuredRequest expects json.Marshaler encoded request body and returns a new request with vault-specific headers
func (c *Client) newStructuredRequest(ctx context.Context, method, path string, body json.Marshaler) (*http.Request, error) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, fmt.Errorf("could not encode request body: %w", err)
	}

	return c.newRequest(ctx, method, path, &buf)
}

// newRequest returns a new request with vault-specific headers
func (c *Client) newRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	// concatenate the base address with the given path
	url, err := c.parsedBaseAddress.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("could not join %q with the base address: %w", path, err)
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

	return req, nil
}

// do sends the given request to Vault, handling retries, redirects, and rate limiting
func (c *Client) do(ctx context.Context, req *http.Request, retry bool) (*http.Response, error) {
	// block on the rate limiter, if set
	if c.configuration.RateLimiter != nil {
		c.configuration.RateLimiter.Wait(ctx)
	}

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

func validateToken(token string) error {
	if !printable(token) {
		return fmt.Errorf("vault token contains non-printable characters")
	}

	return nil
}

func validateNamespace(namespace string) error {
	if !printable(namespace) {
		return fmt.Errorf("vault namespace %q contains non-printable characters", namespace)
	}

	return nil
}

func validateCustomHeaders(headers http.Header) (errs error) {
	for key := range headers {
		if strings.HasPrefix(strings.ToLower(key), "x-vault-") {
			errs = multierror.Append(
				errs,
				fmt.Errorf("custom header key %q is not allowed: 'X-Vault-' prefix is for internal use only", key),
			)
		}
	}

	return errs
}

// printable returns true if the given string has no non-printable characters
func printable(str string) bool {
	idx := strings.IndexFunc(str, func(c rune) bool {
		return !unicode.IsPrint(c)
	})

	return idx == -1
}
