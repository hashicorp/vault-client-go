package vault

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/hashicorp/go-multierror"
)

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

	// mountPath, if specified, will overwrite the default mount path used in
	// client.Auth & client.Secrets requests
	mountPath string
}

// requestHeaders contains headers that will be added to requests
type requestHeaders struct {
	userAgent                 string                    // 'User-Agent'
	token                     string                    // 'X-Vault-Token'
	namespace                 string                    // 'X-Vault-Namespace'
	mfaCredentials            []string                  // 'X-Vault-MFA'
	responseWrappingTTL       time.Duration             // 'X-Vault-Wrap-TTL'
	replicationForwardingMode ReplicationForwardingMode // 'X-Vault-Forward' or 'X-Vault-Inconsistent'
	customHeaders             http.Header
}

// SetToken sets the token to be used with all subsequent requests.
//
// See https://developer.hashicorp.com/vault/docs/concepts/tokens for more info
// on tokens.
func (c *Client) SetToken(token string) error {
	if err := validateToken(token); err != nil {
		return err
	}

	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.token = token
	c.clientRequestModifiersLock.Unlock()

	return nil
}

// ClearToken clears the token for all subsequent requests.
//
// See https://developer.hashicorp.com/vault/docs/concepts/tokens for more info
// on tokens.
func (c *Client) ClearToken() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.token = ""
	c.clientRequestModifiersLock.Unlock()
}

// SetNamespace sets the namespace to be used with all subsequent requests.
// Use an empty string to clear the namespace.
//
// See https://developer.hashicorp.com/vault/docs/enterprise/namespaces for
// more info on namespaces.
func (c *Client) SetNamespace(namespace string) error {
	if err := validateNamespace(namespace); err != nil {
		return err
	}

	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.namespace = namespace
	c.clientRequestModifiersLock.Unlock()

	return nil
}

// ClearNamespace clears the namespace from all subsequent requests.
//
// See https://developer.hashicorp.com/vault/docs/enterprise/namespaces for
// more info on namespaces.
func (c *Client) ClearNamespace() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.namespace = ""
	c.clientRequestModifiersLock.Unlock()
}


// SetMFACredentials sets multi-factor authentication credentials to be used
// with all subsequent requests.
//
// See https://learn.hashicorp.com/tutorials/vault/multi-factor-authentication
// for more information on multi-factor authentication.
func (c *Client) SetMFACredentials(credentials ...string) error {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.mfaCredentials = credentials
	c.clientRequestModifiersLock.Unlock()

	return nil
}

// ClearMFACredentials clears multi-factor authentication credentials from all
// subsequent requests.
//
// See https://learn.hashicorp.com/tutorials/vault/multi-factor-authentication
// for more information on multi-factor authentication.
func (c *Client) ClearMFACredentials() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.mfaCredentials = nil
	c.clientRequestModifiersLock.Unlock()
}

// SetResponseWrapping sets the response-wrapping TTL to the given duration for
// all subsequent requests, telling Vault to wrap responses and return
// response-wrapping tokens instead.
//
// See https://developer.hashicorp.com/vault/docs/concepts/response-wrapping
// for more information on response wrapping.
func (c *Client) SetResponseWrapping(ttl time.Duration) error {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.responseWrappingTTL = ttl
	c.clientRequestModifiersLock.Unlock()

	return nil
}

// ClearResponseWrapping clears the response-wrapping header from all
// subsequent requests.
//
// See https://developer.hashicorp.com/vault/docs/concepts/response-wrapping
// for more information on response wrapping.
func (c *Client) ClearResponseWrapping() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.responseWrappingTTL = 0
	c.clientRequestModifiersLock.Unlock()
}

// SetCustomHeaders sets custom headers to be used in all subsequent requests.
// The internal prefix 'X-Vault-' is not permitted for the header keys.
func (c *Client) SetCustomHeaders(headers http.Header) error {
	if err := validateCustomHeaders(headers); err != nil {
		return err
	}

	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.customHeaders = headers.Clone()
	c.clientRequestModifiersLock.Unlock()

	return nil
}

// ClearsCustomHeaders clears all custom headers from the subsequent requests.
func (c *Client) ClearCustomHeaders() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.customHeaders = nil
	c.clientRequestModifiersLock.Unlock()
}


// SetRequestCallbacks sets callbacks which will be invoked before each request.
func (c *Client) SetRequestCallbacks(callbacks ...RequestCallback) error {
	c.clientRequestModifiersLock.Lock()
	copy(c.clientRequestModifiers.requestCallbacks, callbacks)
	c.clientRequestModifiersLock.Unlock()

	return nil
}

// ClearRequestCallbacks clears all request callbacks.
func (c *Client) ClearRequestCallbacks() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.requestCallbacks = nil
	c.clientRequestModifiersLock.Unlock()
}

// SetResponseCallbacks sets callbacks which will be invoked after each
// successful response.
func (c *Client) SetResponseCallbacks(callbacks ...ResponseCallback) error {
	c.clientRequestModifiersLock.Lock()
	copy(c.clientRequestModifiers.responseCallbacks, callbacks)
	c.clientRequestModifiersLock.Unlock()

	return nil
}

// ClearResponseCallbacks clears all response callbacks.
func (c *Client) ClearResponseCallbacks() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.responseCallbacks = nil
	c.clientRequestModifiersLock.Unlock()
}

func (m *requestModifiers) mountPathOr(defaultMountPath string) string {
	if m.mountPath == "" {
		return defaultMountPath
	}
	return m.mountPath
}

// mergeRequestModifiers merges the two objects, preferring the per-request modifiers
func mergeRequestModifiers(perClient, perRequest requestModifiers) requestModifiers {
	merged := perClient

	if perRequest.headers.userAgent != "" {
		merged.headers.userAgent = perRequest.headers.userAgent
	}

	if perRequest.headers.token != "" {
		merged.headers.token = perRequest.headers.token
	}

	if perRequest.headers.namespace != "" {
		merged.headers.namespace = perRequest.headers.namespace
	}

	if len(perRequest.headers.mfaCredentials) != 0 {
		merged.headers.mfaCredentials = perRequest.headers.mfaCredentials
	}

	if perRequest.headers.responseWrappingTTL != 0 {
		merged.headers.responseWrappingTTL = perRequest.headers.responseWrappingTTL
	}

	if perRequest.headers.replicationForwardingMode != ReplicationForwardNone {
		merged.headers.replicationForwardingMode = perRequest.headers.replicationForwardingMode
	}

	if len(perRequest.headers.customHeaders) != 0 {
		merged.headers.customHeaders = perRequest.headers.customHeaders
	}

	if len(perRequest.requestCallbacks) != 0 {
		merged.requestCallbacks = perRequest.requestCallbacks
	}

	if len(perRequest.responseCallbacks) != 0 {
		merged.responseCallbacks = perRequest.responseCallbacks
	}

	return merged
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

