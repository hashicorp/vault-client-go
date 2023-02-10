// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AWSConfigWriteRoleTagDenyListRequest struct for AWSConfigWriteRoleTagDenyListRequest
type AWSConfigWriteRoleTagDenyListRequest struct {
	// If set to 'true', disables the periodic tidying of deny listed entries.
	DisablePeriodicTidy bool `json:"disable_periodic_tidy"`

	// The amount of extra time that must have passed beyond the roletag expiration, before it is removed from the backend storage. Defaults to 4320h (180 days).
	SafetyBuffer int32 `json:"safety_buffer"`
}

// NewAWSConfigWriteRoleTagDenyListRequestWithDefaults instantiates a new AWSConfigWriteRoleTagDenyListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSConfigWriteRoleTagDenyListRequestWithDefaults() *AWSConfigWriteRoleTagDenyListRequest {
	var this AWSConfigWriteRoleTagDenyListRequest

	this.DisablePeriodicTidy = false
	this.SafetyBuffer = 15552000

	return &this
}

func (o AWSConfigWriteRoleTagDenyListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["disable_periodic_tidy"] = o.DisablePeriodicTidy
	toSerialize["safety_buffer"] = o.SafetyBuffer

	return json.Marshal(toSerialize)
}
