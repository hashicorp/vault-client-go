// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WritePoliciesPasswordRequest struct for WritePoliciesPasswordRequest
type WritePoliciesPasswordRequest struct {
	// The password policy
	Policy string `json:"policy"`
}

// NewWritePoliciesPasswordRequestWithDefaults instantiates a new WritePoliciesPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritePoliciesPasswordRequestWithDefaults() *WritePoliciesPasswordRequest {
	var this WritePoliciesPasswordRequest

	return &this
}
