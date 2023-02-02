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

// WrappingWriteLookupRequest struct for WrappingWriteLookupRequest
type WrappingWriteLookupRequest struct {
	Token string `json:"token"`
}

// NewWrappingWriteLookupRequestWithDefaults instantiates a new WrappingWriteLookupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWrappingWriteLookupRequestWithDefaults() *WrappingWriteLookupRequest {
	var this WrappingWriteLookupRequest

	return &this
}

func (o WrappingWriteLookupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token"] = o.Token

	return json.Marshal(toSerialize)
}
