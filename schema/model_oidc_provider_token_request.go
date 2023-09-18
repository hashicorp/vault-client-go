// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// OidcProviderTokenRequest struct for OidcProviderTokenRequest
type OidcProviderTokenRequest struct {
	// The ID of the requesting client.
	ClientId string `json:"client_id,omitempty"`

	// The secret of the requesting client.
	ClientSecret string `json:"client_secret,omitempty"`

	// The authorization code received from the provider's authorization endpoint.
	Code string `json:"code"`

	// The code verifier associated with the authorization code.
	CodeVerifier string `json:"code_verifier,omitempty"`

	// The authorization grant type. The following grant types are supported: 'authorization_code'.
	GrantType string `json:"grant_type"`

	// The callback location where the authentication response was sent.
	RedirectUri string `json:"redirect_uri"`
}
