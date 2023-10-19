// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"fmt"
	"maps"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"time"
	"unicode"
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
	// client.Auth & client.Secrets requests.
	mountPath string

	// additionalQueryParameters, if specified will be appended to the list of
	// query parameters included with the request.
	additionalQueryParameters url.Values
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
	c.clientRequestModifiers.requestCallbacks = slices.Clone(callbacks)
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
	c.clientRequestModifiers.responseCallbacks = slices.Clone(callbacks)
	c.clientRequestModifiersLock.Unlock()

	return nil
}

// ClearResponseCallbacks clears all response callbacks.
func (c *Client) ClearResponseCallbacks() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.responseCallbacks = nil
	c.clientRequestModifiersLock.Unlock()
}

// SetReplicationForwardingMode sets a replication forwarding header for all
// subsequent requests:
//
//	ReplicationForwardNone         - no forwarding header
//	ReplicationForwardAlways       - 'X-Vault-Forward'
//	ReplicationForwardInconsistent - 'X-Vault-Inconsistent'
//
// Note: this feature must be enabled in Vault's configuration.
//
// See https://developer.hashicorp.com/vault/docs/enterprise/consistency#vault-1-7-mitigations
func (c *Client) SetReplicationForwardingMode(mode ReplicationForwardingMode) {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.replicationForwardingMode = mode
	c.clientRequestModifiersLock.Unlock()
}

// ReplicationForwardingMode clears the X-Vault-Forward / X-Vault-Inconsistent
// headers from all subsequent requests.
//
// See https://developer.hashicorp.com/vault/docs/enterprise/consistency#vault-1-7-mitigations
func (c *Client) ClearReplicationForwardingMode() {
	c.clientRequestModifiersLock.Lock()
	c.clientRequestModifiers.headers.replicationForwardingMode = ReplicationForwardNone
	c.clientRequestModifiersLock.Unlock()
}

// mountPathOr returns object's mount path or the given default value.
func (m *requestModifiers) mountPathOr(defaultMountPath string) string {
	if m.mountPath == "" {
		return defaultMountPath
	}
	return m.mountPath
}

// additionalQueryParametersOrDefault returns object's query parameters or an empty map.
func (m *requestModifiers) additionalQueryParametersOrDefault() url.Values {
	if m.additionalQueryParameters == nil {
		return make(url.Values)
	}
	return m.additionalQueryParameters
}

// mergeRequestModifiers merges the values in *rhs into *lhs. The merging is
// done according the following rules:
//
//   - for scalars : the rhs values, if present, will overwrite the lhs values
//   - for slices  : the rhs values will be appended to the lhs values
//   - for maps
//     -- new keys      : the rhs values will be appended to the lhs values
//     -- existing keys : the rhs values will overwrite the corresponding lhs values
func mergeRequestModifiers(lhs, rhs *requestModifiers) {
	if rhs.headers.userAgent != "" {
		lhs.headers.userAgent = rhs.headers.userAgent
	}

	if rhs.headers.token != "" {
		lhs.headers.token = rhs.headers.token
	}

	if rhs.headers.namespace != "" {
		lhs.headers.namespace = rhs.headers.namespace
	}

	lhs.headers.mfaCredentials = append(
		lhs.headers.mfaCredentials,
		rhs.headers.mfaCredentials...,
	)

	if rhs.headers.responseWrappingTTL != 0 {
		lhs.headers.responseWrappingTTL = rhs.headers.responseWrappingTTL
	}

	if rhs.headers.replicationForwardingMode != ReplicationForwardNone {
		lhs.headers.replicationForwardingMode = rhs.headers.replicationForwardingMode
	}

	// in case of key collisions, the rhs keys will take precedence
	if lhs.headers.customHeaders != nil {
		maps.Copy(lhs.headers.customHeaders, rhs.headers.customHeaders)
	} else {
		lhs.headers.customHeaders = rhs.headers.customHeaders
	}

	lhs.requestCallbacks = append(
		lhs.requestCallbacks,
		rhs.requestCallbacks...,
	)

	lhs.responseCallbacks = append(
		lhs.responseCallbacks,
		rhs.responseCallbacks...,
	)
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

func validateCustomHeaders(headers http.Header) error {
	for key := range headers {
		if strings.HasPrefix(strings.ToLower(key), "x-vault-") {
			return fmt.Errorf("custom header key %q is not allowed: 'X-Vault-' prefix is for internal use only", key)
		}
	}

	return nil
}

// printable returns true if the given string has no non-printable characters
func printable(str string) bool {
	idx := strings.IndexFunc(str, func(c rune) bool {
		return !unicode.IsPrint(c)
	})

	return idx == -1
}
