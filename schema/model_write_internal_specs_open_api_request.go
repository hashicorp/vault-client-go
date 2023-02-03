// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteInternalSpecsOpenAPIRequest struct for WriteInternalSpecsOpenAPIRequest
type WriteInternalSpecsOpenAPIRequest struct {
	// Context string appended to every operationId

	Context string `json:"context"`
}

// NewWriteInternalSpecsOpenAPIRequestWithDefaults instantiates a new WriteInternalSpecsOpenAPIRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteInternalSpecsOpenAPIRequestWithDefaults() *WriteInternalSpecsOpenAPIRequest {
	var this WriteInternalSpecsOpenAPIRequest

	return &this
}

func (o WriteInternalSpecsOpenAPIRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["context"] = o.Context

	return json.Marshal(toSerialize)
}
