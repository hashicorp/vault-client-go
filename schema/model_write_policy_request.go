// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WritePolicyRequest struct for WritePolicyRequest
type WritePolicyRequest struct {
	// The rules of the policy.

	Policy string `json:"policy"`

	// The rules of the policy.

	// Deprecated

	Rules string `json:"rules"`
}

// NewWritePolicyRequestWithDefaults instantiates a new WritePolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritePolicyRequestWithDefaults() *WritePolicyRequest {
	var this WritePolicyRequest

	return &this
}

func (o WritePolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["policy"] = o.Policy

	toSerialize["rules"] = o.Rules

	return json.Marshal(toSerialize)
}
