// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OidcProviderAuthorize2Request struct for OidcProviderAuthorize2Request
type OidcProviderAuthorize2Request struct {
	// The ID of the requesting client.
	ClientId string `json:"client_id"`

	// The code challenge derived from the code verifier.
	CodeChallenge string `json:"code_challenge"`

	// The method that was used to derive the code challenge. The following methods are supported: 'S256', 'plain'. Defaults to 'plain'.
	CodeChallengeMethod string `json:"code_challenge_method"`

	// The allowable elapsed time in seconds since the last time the end-user was actively authenticated.
	MaxAge int32 `json:"max_age"`

	// The value that will be returned in the ID token nonce claim after a token exchange.
	Nonce string `json:"nonce"`

	// The redirection URI to which the response will be sent.
	RedirectUri string `json:"redirect_uri"`

	// The OIDC authentication flow to be used. The following response types are supported: 'code'
	ResponseType string `json:"response_type"`

	// A space-delimited, case-sensitive list of scopes to be requested. The 'openid' scope is required.
	Scope string `json:"scope"`

	// The value used to maintain state between the authentication request and client.
	State string `json:"state"`
}

// NewOidcProviderAuthorize2RequestWithDefaults instantiates a new OidcProviderAuthorize2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcProviderAuthorize2RequestWithDefaults() *OidcProviderAuthorize2Request {
	var this OidcProviderAuthorize2Request

	this.CodeChallengeMethod = "plain"

	return &this
}

func (o OidcProviderAuthorize2Request) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["client_id"] = o.ClientId
	toSerialize["code_challenge"] = o.CodeChallenge
	toSerialize["code_challenge_method"] = o.CodeChallengeMethod
	toSerialize["max_age"] = o.MaxAge
	toSerialize["nonce"] = o.Nonce
	toSerialize["redirect_uri"] = o.RedirectUri
	toSerialize["response_type"] = o.ResponseType
	toSerialize["scope"] = o.Scope
	toSerialize["state"] = o.State

	return json.Marshal(toSerialize)
}