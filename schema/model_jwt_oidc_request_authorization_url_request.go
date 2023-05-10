// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// JwtOidcRequestAuthorizationUrlRequest struct for JwtOidcRequestAuthorizationUrlRequest
type JwtOidcRequestAuthorizationUrlRequest struct {
	// Optional client-provided nonce that must match during callback, if present.
	ClientNonce string `json:"client_nonce"`

	// The OAuth redirect_uri to use in the authorization URL.
	RedirectUri string `json:"redirect_uri"`

	// The role to issue an OIDC authorization URL against.
	Role string `json:"role"`
}

// NewJwtOidcRequestAuthorizationUrlRequestWithDefaults instantiates a new JwtOidcRequestAuthorizationUrlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJwtOidcRequestAuthorizationUrlRequestWithDefaults() *JwtOidcRequestAuthorizationUrlRequest {
	var this JwtOidcRequestAuthorizationUrlRequest

	return &this
}

func (o JwtOidcRequestAuthorizationUrlRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["client_nonce"] = o.ClientNonce
	toSerialize["redirect_uri"] = o.RedirectUri
	toSerialize["role"] = o.Role

	return json.Marshal(toSerialize)
}
