// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MFAMethodGenerateTOTPRequest struct for MFAMethodGenerateTOTPRequest
type MFAMethodGenerateTOTPRequest struct {
	// The unique identifier for this MFA method.

	MethodId string `json:"method_id"`
}

// NewMFAMethodGenerateTOTPRequestWithDefaults instantiates a new MFAMethodGenerateTOTPRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAMethodGenerateTOTPRequestWithDefaults() *MFAMethodGenerateTOTPRequest {
	var this MFAMethodGenerateTOTPRequest

	return &this
}

func (o MFAMethodGenerateTOTPRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["method_id"] = o.MethodId

	return json.Marshal(toSerialize)
}
