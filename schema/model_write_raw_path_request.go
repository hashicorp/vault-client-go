// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteRawPathRequest struct for WriteRawPathRequest
type WriteRawPathRequest struct {
	Compressed bool `json:"compressed"`

	CompressionType string `json:"compression_type"`

	Encoding string `json:"encoding"`

	Value string `json:"value"`
}

// NewWriteRawPathRequestWithDefaults instantiates a new WriteRawPathRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteRawPathRequestWithDefaults() *WriteRawPathRequest {
	var this WriteRawPathRequest

	return &this
}

func (o WriteRawPathRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
