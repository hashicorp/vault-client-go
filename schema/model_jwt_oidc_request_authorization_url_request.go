// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// JwtOidcRequestAuthorizationUrlRequest struct for JwtOidcRequestAuthorizationUrlRequest
type JwtOidcRequestAuthorizationUrlRequest struct {
	// Optional client-provided nonce that must match during callback, if present.
	ClientNonce string `json:"client_nonce,omitempty"`

	// The OAuth redirect_uri to use in the authorization URL.
	RedirectUri string `json:"redirect_uri,omitempty"`

	// The role to issue an OIDC authorization URL against.
	Role string `json:"role,omitempty"`
}
