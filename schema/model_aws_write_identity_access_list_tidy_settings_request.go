/*

HashiCorp Vault API



HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.



API version: 1.13.0


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AWSWriteIdentityAccessListTidySettingsRequest struct for AWSWriteIdentityAccessListTidySettingsRequest
type AWSWriteIdentityAccessListTidySettingsRequest struct {
	// The amount of extra time that must have passed beyond the identity's expiration, before it is removed from the backend storage.

	SafetyBuffer int32 `json:"safety_buffer"`
}

// NewAWSWriteIdentityAccessListTidySettingsRequestWithDefaults instantiates a new AWSWriteIdentityAccessListTidySettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSWriteIdentityAccessListTidySettingsRequestWithDefaults() *AWSWriteIdentityAccessListTidySettingsRequest {
	var this AWSWriteIdentityAccessListTidySettingsRequest

	this.SafetyBuffer = 259200

	return &this
}

func (o AWSWriteIdentityAccessListTidySettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["safety_buffer"] = o.SafetyBuffer

	return json.Marshal(toSerialize)
}
