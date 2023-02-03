// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AWSWriteRoleTagBlackListTidySettingsRequest struct for AWSWriteRoleTagBlackListTidySettingsRequest
type AWSWriteRoleTagBlackListTidySettingsRequest struct {
	// The amount of extra time that must have passed beyond the roletag expiration, before it is removed from the backend storage.
	SafetyBuffer int32 `json:"safety_buffer"`
}

// NewAWSWriteRoleTagBlackListTidySettingsRequestWithDefaults instantiates a new AWSWriteRoleTagBlackListTidySettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSWriteRoleTagBlackListTidySettingsRequestWithDefaults() *AWSWriteRoleTagBlackListTidySettingsRequest {
	var this AWSWriteRoleTagBlackListTidySettingsRequest

	this.SafetyBuffer = 259200

	return &this
}
