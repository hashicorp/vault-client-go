// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

// RequestOption is a functional parameter used to modify a request
type RequestOption func(*requestModifiers) error

// WithToken sets the token for the next request; it takes precedence over the
// client-level token.
//
// See https://developer.hashicorp.com/vault/docs/concepts/tokens for more info
// on tokens.
func WithToken(token string) RequestOption {
	return func(m *requestModifiers) error {
		if err := validateToken(token); err != nil {
			return err
		}
		m.headers.token = token
		return nil
	}
}

// WithNamespace sets the namespace for the next request; it takes precedence
// over the client-level namespace. Use an empty string to clear the namespace
// from the next request.
//
// See https://developer.hashicorp.com/vault/docs/enterprise/namespaces for
// more info on namespaces.
func WithNamespace(namespace string) RequestOption {
	return func(m *requestModifiers) error {
		if err := validateNamespace(namespace); err != nil {
			return err
		}
		m.headers.namespace = namespace
		return nil
	}
}

// WithMFACredentials sets the multi-factor authentication credentials for the
// next request, it takes precedence over the client-level MFA credentials.
//
// See https://learn.hashicorp.com/tutorials/vault/multi-factor-authentication
// for more information on multi-factor authentication.
func WithMFACredentials(credentials ...string) RequestOption {
	return func(m *requestModifiers) error {
		m.headers.mfaCredentials = credentials
		return nil
	}
}

// WithResponseWrapping sets the response-wrapping TTL to the given duration
// for the next request; it takes precedence over the client-level
// response-wrapping TTL. A non-zero duration will tell Vault to wrap the
// response and return a response-wrapping token instead. Set `ttl` to zero
// to clear the response-wrapping header from the next request.
//
// See https://developer.hashicorp.com/vault/docs/concepts/response-wrapping
// for more information on response wrapping.
func WithResponseWrapping(ttl time.Duration) RequestOption {
	return func(m *requestModifiers) error {
		m.headers.responseWrappingTTL = ttl
		return nil
	}
}

// WithCustomHeaders sets custom headers for the next request; these headers
// will be appended to any custom headers set at the client level. The internal
// prefix 'X-Vault-' is not permitted for the header keys.
func WithCustomHeaders(headers http.Header) RequestOption {
	return func(m *requestModifiers) error {
		if err := validateCustomHeaders(headers); err != nil {
			return err
		}
		m.headers.customHeaders = headers
		return nil
	}
}

// WithRequestCallbacks sets callbacks which will be invoked before the next
// request; these callbacks will be appended to the list of the callbacks set
// at the client-level.
func WithRequestCallbacks(callbacks ...RequestCallback) RequestOption {
	return func(m *requestModifiers) error {
		m.requestCallbacks = callbacks
		return nil
	}
}

// WithResponseCallbacks sets callbacks which will be invoked after a
// successful response within the next request; these callbacks will be
// appended to the list of the callbacks set at the client-level.
func WithResponseCallbacks(callbacks ...ResponseCallback) RequestOption {
	return func(m *requestModifiers) error {
		m.responseCallbacks = callbacks
		return nil
	}
}

// WithMountPath overwrites the default mount path in client.Auth and
// client.Secrets requests with the given mount path string.
func WithMountPath(path string) RequestOption {
	return func(m *requestModifiers) error {
		// the mount path is expected to have no leading/trailing "/" characters
		m.mountPath = strings.Trim(path, "/")
		return nil
	}
}

// WithQueryParameters appends the given list of query parameters to the request.
// This is primarily intended to be used with client.Read, client.ReadRaw,
// client.List, and client.Delete methods.
func WithQueryParameters(parameters url.Values) RequestOption {
	return func(m *requestModifiers) error {
		m.additionalQueryParameters = parameters
		return nil
	}
}

// WithReplicationForwardingMode sets a replication forwarding header to the
// given value for the next request; it takes precedence over the client-level
// replication forwarding header.
//
//	ReplicationForwardNone         - no forwarding headers
//	ReplicationForwardAlways       - 'X-Vault-Forward'
//	ReplicationForwardInconsistent - 'X-Vault-Inconsistent'
//
// See https://developer.hashicorp.com/vault/docs/enterprise/consistency#vault-1-7-mitigations
func WithReplicationForwardingMode(mode ReplicationForwardingMode) RequestOption {
	return func(m *requestModifiers) error {
		m.headers.replicationForwardingMode = mode
		return nil
	}
}

// requestOptionsToRequestModifiers constructs `requestModifiers` propagating the errors, if any
func requestOptionsToRequestModifiers(options []RequestOption) (requestModifiers, error) {
	var modifiers requestModifiers

	for _, option := range options {
		if err := option(&modifiers); err != nil {
			return requestModifiers{}, err
		}
	}

	return modifiers, nil
}
