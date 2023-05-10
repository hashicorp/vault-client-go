// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OidcProviderTokenRequest struct for OidcProviderTokenRequest
type OidcProviderTokenRequest struct {
	// The ID of the requesting client.
	ClientId string `json:"client_id"`

	// The secret of the requesting client.
	ClientSecret string `json:"client_secret"`

	// The authorization code received from the provider's authorization endpoint.
	Code string `json:"code"`

	// The code verifier associated with the authorization code.
	CodeVerifier string `json:"code_verifier"`

	// The authorization grant type. The following grant types are supported: 'authorization_code'.
	GrantType string `json:"grant_type"`

	// The callback location where the authentication response was sent.
	RedirectUri string `json:"redirect_uri"`
}

// NewOidcProviderTokenRequestWithDefaults instantiates a new OidcProviderTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcProviderTokenRequestWithDefaults() *OidcProviderTokenRequest {
	var this OidcProviderTokenRequest

	return &this
}

func (o OidcProviderTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["client_id"] = o.ClientId
	toSerialize["client_secret"] = o.ClientSecret
	toSerialize["code"] = o.Code
	toSerialize["code_verifier"] = o.CodeVerifier
	toSerialize["grant_type"] = o.GrantType
	toSerialize["redirect_uri"] = o.RedirectUri

	return json.Marshal(toSerialize)
}
