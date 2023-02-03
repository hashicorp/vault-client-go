// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MFAMethodAdminDestroyTOTPRequest struct for MFAMethodAdminDestroyTOTPRequest
type MFAMethodAdminDestroyTOTPRequest struct {
	// Identifier of the entity from which the MFA method secret needs to be removed.
	EntityId string `json:"entity_id"`

	// The unique identifier for this MFA method.
	MethodId string `json:"method_id"`
}

// NewMFAMethodAdminDestroyTOTPRequestWithDefaults instantiates a new MFAMethodAdminDestroyTOTPRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAMethodAdminDestroyTOTPRequestWithDefaults() *MFAMethodAdminDestroyTOTPRequest {
	var this MFAMethodAdminDestroyTOTPRequest

	return &this
}

func (o MFAMethodAdminDestroyTOTPRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
