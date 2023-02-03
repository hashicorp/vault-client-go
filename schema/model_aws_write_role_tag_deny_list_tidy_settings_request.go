// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AWSWriteRoleTagDenyListTidySettingsRequest struct for AWSWriteRoleTagDenyListTidySettingsRequest
type AWSWriteRoleTagDenyListTidySettingsRequest struct {
	// The amount of extra time that must have passed beyond the roletag expiration, before it is removed from the backend storage.
	SafetyBuffer int32 `json:"safety_buffer"`
}

// NewAWSWriteRoleTagDenyListTidySettingsRequestWithDefaults instantiates a new AWSWriteRoleTagDenyListTidySettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSWriteRoleTagDenyListTidySettingsRequestWithDefaults() *AWSWriteRoleTagDenyListTidySettingsRequest {
	var this AWSWriteRoleTagDenyListTidySettingsRequest

	this.SafetyBuffer = 259200

	return &this
}

func (o AWSWriteRoleTagDenyListTidySettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
