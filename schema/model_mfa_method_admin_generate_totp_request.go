// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MFAMethodAdminGenerateTOTPRequest struct for MFAMethodAdminGenerateTOTPRequest
type MFAMethodAdminGenerateTOTPRequest struct {
	// Entity ID on which the generated secret needs to get stored.

	EntityId string `json:"entity_id"`

	// The unique identifier for this MFA method.

	MethodId string `json:"method_id"`
}

// NewMFAMethodAdminGenerateTOTPRequestWithDefaults instantiates a new MFAMethodAdminGenerateTOTPRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAMethodAdminGenerateTOTPRequestWithDefaults() *MFAMethodAdminGenerateTOTPRequest {
	var this MFAMethodAdminGenerateTOTPRequest

	return &this
}

func (o MFAMethodAdminGenerateTOTPRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["entity_id"] = o.EntityId

	toSerialize["method_id"] = o.MethodId

	return json.Marshal(toSerialize)
}
