// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// OidcProviderAuthorizeWithParametersRequest struct for OidcProviderAuthorizeWithParametersRequest
type OidcProviderAuthorizeWithParametersRequest struct {
	// The ID of the requesting client.
	ClientId string `json:"client_id"`

	// The code challenge derived from the code verifier.
	CodeChallenge string `json:"code_challenge,omitempty"`

	// The method that was used to derive the code challenge. The following methods are supported: 'S256', 'plain'. Defaults to 'plain'.
	CodeChallengeMethod string `json:"code_challenge_method,omitempty"`

	// The allowable elapsed time in seconds since the last time the end-user was actively authenticated.
	MaxAge int32 `json:"max_age,omitempty"`

	// The value that will be returned in the ID token nonce claim after a token exchange.
	Nonce string `json:"nonce,omitempty"`

	// The redirection URI to which the response will be sent.
	RedirectUri string `json:"redirect_uri"`

	// The OIDC authentication flow to be used. The following response types are supported: 'code'
	ResponseType string `json:"response_type"`

	// A space-delimited, case-sensitive list of scopes to be requested. The 'openid' scope is required.
	Scope string `json:"scope"`

	// The value used to maintain state between the authentication request and client.
	State string `json:"state,omitempty"`
}
