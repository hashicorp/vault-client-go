// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// JWTWriteOIDCCallbackRequest struct for JWTWriteOIDCCallbackRequest
type JWTWriteOIDCCallbackRequest struct {
	ClientNonce string `json:"client_nonce"`

	Code string `json:"code"`

	IdToken string `json:"id_token"`

	State string `json:"state"`
}

// NewJWTWriteOIDCCallbackRequestWithDefaults instantiates a new JWTWriteOIDCCallbackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJWTWriteOIDCCallbackRequestWithDefaults() *JWTWriteOIDCCallbackRequest {
	var this JWTWriteOIDCCallbackRequest

	return &this
}

func (o JWTWriteOIDCCallbackRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["client_nonce"] = o.ClientNonce

	toSerialize["code"] = o.Code

	toSerialize["id_token"] = o.IdToken

	toSerialize["state"] = o.State

	return json.Marshal(toSerialize)
}
