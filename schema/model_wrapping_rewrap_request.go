// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WrappingRewrapRequest struct for WrappingRewrapRequest
type WrappingRewrapRequest struct {
	Token string `json:"token"`
}

// NewWrappingRewrapRequestWithDefaults instantiates a new WrappingRewrapRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWrappingRewrapRequestWithDefaults() *WrappingRewrapRequest {
	var this WrappingRewrapRequest

	return &this
}

func (o WrappingRewrapRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
