// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteRawRequest struct for WriteRawRequest
type WriteRawRequest struct {
	Compressed bool `json:"compressed"`

	CompressionType string `json:"compression_type"`

	Encoding string `json:"encoding"`

	Path string `json:"path"`

	Value string `json:"value"`
}

// NewWriteRawRequestWithDefaults instantiates a new WriteRawRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteRawRequestWithDefaults() *WriteRawRequest {
	var this WriteRawRequest

	return &this
}

func (o WriteRawRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
