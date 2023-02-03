// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteGenerateRootUpdateRequest struct for WriteGenerateRootUpdateRequest
type WriteGenerateRootUpdateRequest struct {
	// Specifies a single unseal key share.

	Key string `json:"key"`

	// Specifies the nonce of the attempt.

	Nonce string `json:"nonce"`
}

// NewWriteGenerateRootUpdateRequestWithDefaults instantiates a new WriteGenerateRootUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteGenerateRootUpdateRequestWithDefaults() *WriteGenerateRootUpdateRequest {
	var this WriteGenerateRootUpdateRequest

	return &this
}

func (o WriteGenerateRootUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["key"] = o.Key

	toSerialize["nonce"] = o.Nonce

	return json.Marshal(toSerialize)
}
