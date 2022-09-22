package vault

import (
	"fmt"
	"net/http"
	"strings"
	"unicode"

	"github.com/hashicorp/go-multierror"
)

// SetToken sets the token to be used with all subsequent requests.
//
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
//
// See https://www.vaultproject.io/docs/concepts/tokens for more info on
// tokens.
func (c *Client) ClearToken() {
	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.token = ""
	c.requestModifiersLock.Unlock()
}

// WithToken returns a shallow copy of the client with the token set to the
// given value.
//
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
//
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
//
// See https://www.vaultproject.io/docs/enterprise/namespaces for more info on
// namespaces.
func (c *Client) ClearNamespace() {
	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.namespace = ""
	c.requestModifiersLock.Unlock()
}

// WithNamespace returns a shallow copy of the client with the namespace set to
// the given value, use "" to clear the namespace.
//
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

// SetMFACredentials sets multi-factor authentication credentials to be used
// with all subsequent requests.
//
// See https://learn.hashicorp.com/tutorials/vault/multi-factor-authentication
// for more information on multi-factor authentication.
func (c *Client) SetMFACredentials(credentials ...string) error {
	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.mfaCredentials = credentials
	c.requestModifiersLock.Unlock()

	return nil
}

// ClearMFACredentials clears multi-factor authentication credentials from all
// subsequent requests.
//
// See https://learn.hashicorp.com/tutorials/vault/multi-factor-authentication
// for more information on multi-factor authentication.
func (c *Client) ClearMFACredentials() {
	c.requestModifiersLock.Lock()
	c.requestModifiers.headers.mfaCredentials = nil
	c.requestModifiersLock.Unlock()
}

// WithMFACredentials returns a shallow copy of the client with multi-factor
// authentication credentials set to the given value(s); passing no parameters
// is equivalent to clearing the list.
//
// See https://learn.hashicorp.com/tutorials/vault/multi-factor-authentication
// for more information on multi-factor authentication.
func (c *Client) WithMFACredentials(credentials ...string) *Client {
	clone := c.Clone()
	clone.requestModifiers.headers.mfaCredentials = credentials

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
