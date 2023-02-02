/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OIDCWriteAuthURLRequest struct for OIDCWriteAuthURLRequest
type OIDCWriteAuthURLRequest struct {
	// Optional client-provided nonce that must match during callback, if present.
	ClientNonce string `json:"client_nonce"`
	// The OAuth redirect_uri to use in the authorization URL.
	RedirectUri string `json:"redirect_uri"`
	// The role to issue an OIDC authorization URL against.
	Role string `json:"role"`
}

// NewOIDCWriteAuthURLRequestWithDefaults instantiates a new OIDCWriteAuthURLRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWriteAuthURLRequestWithDefaults() *OIDCWriteAuthURLRequest {
	var this OIDCWriteAuthURLRequest

	return &this
}

func (o OIDCWriteAuthURLRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["client_nonce"] = o.ClientNonce
	toSerialize["redirect_uri"] = o.RedirectUri
	toSerialize["role"] = o.Role

	return json.Marshal(toSerialize)
}