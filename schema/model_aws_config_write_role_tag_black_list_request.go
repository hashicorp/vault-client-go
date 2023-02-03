// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AWSConfigWriteRoleTagBlackListRequest struct for AWSConfigWriteRoleTagBlackListRequest
type AWSConfigWriteRoleTagBlackListRequest struct {
	// If set to 'true', disables the periodic tidying of deny listed entries.
	DisablePeriodicTidy bool `json:"disable_periodic_tidy"`

	// The amount of extra time that must have passed beyond the roletag expiration, before it is removed from the backend storage. Defaults to 4320h (180 days).
	SafetyBuffer int32 `json:"safety_buffer"`
}

// NewAWSConfigWriteRoleTagBlackListRequestWithDefaults instantiates a new AWSConfigWriteRoleTagBlackListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSConfigWriteRoleTagBlackListRequestWithDefaults() *AWSConfigWriteRoleTagBlackListRequest {
	var this AWSConfigWriteRoleTagBlackListRequest

	this.DisablePeriodicTidy = false
	this.SafetyBuffer = 15552000

	return &this
}
