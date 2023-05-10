// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AwsConfigureIdentityAccessListTidyOperationRequest struct for AwsConfigureIdentityAccessListTidyOperationRequest
type AwsConfigureIdentityAccessListTidyOperationRequest struct {
	// If set to 'true', disables the periodic tidying of the 'identity-accesslist/<instance_id>' entries.
	DisablePeriodicTidy bool `json:"disable_periodic_tidy"`

	// The amount of extra time that must have passed beyond the identity's expiration, before it is removed from the backend storage.
	SafetyBuffer int32 `json:"safety_buffer"`
}

// NewAwsConfigureIdentityAccessListTidyOperationRequestWithDefaults instantiates a new AwsConfigureIdentityAccessListTidyOperationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsConfigureIdentityAccessListTidyOperationRequestWithDefaults() *AwsConfigureIdentityAccessListTidyOperationRequest {
	var this AwsConfigureIdentityAccessListTidyOperationRequest

	this.DisablePeriodicTidy = false
	this.SafetyBuffer = 259200

	return &this
}

func (o AwsConfigureIdentityAccessListTidyOperationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["disable_periodic_tidy"] = o.DisablePeriodicTidy
	toSerialize["safety_buffer"] = o.SafetyBuffer

	return json.Marshal(toSerialize)
}
