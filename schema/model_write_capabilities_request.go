// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteCapabilitiesRequest struct for WriteCapabilitiesRequest
type WriteCapabilitiesRequest struct {
	// Use 'paths' instead.
	// Deprecated
	Path []string `json:"path"`

	// Paths on which capabilities are being queried.
	Paths []string `json:"paths"`

	// Token for which capabilities are being queried.
	Token string `json:"token"`
}

// NewWriteCapabilitiesRequestWithDefaults instantiates a new WriteCapabilitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteCapabilitiesRequestWithDefaults() *WriteCapabilitiesRequest {
	var this WriteCapabilitiesRequest

	return &this
}

func (o WriteCapabilitiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
