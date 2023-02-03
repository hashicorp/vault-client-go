// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

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
