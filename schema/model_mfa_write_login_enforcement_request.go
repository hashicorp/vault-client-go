// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MfaWriteLoginEnforcementRequest struct for MfaWriteLoginEnforcementRequest
type MfaWriteLoginEnforcementRequest struct {
	// Array of auth mount accessor IDs
	AuthMethodAccessors []string `json:"auth_method_accessors"`

	// Array of auth mount types
	AuthMethodTypes []string `json:"auth_method_types"`

	// Array of identity entity IDs
	IdentityEntityIds []string `json:"identity_entity_ids"`

	// Array of identity group IDs
	IdentityGroupIds []string `json:"identity_group_ids"`

	// Array of Method IDs that determine what methods will be enforced
	MfaMethodIds []string `json:"mfa_method_ids"`
}

// NewMfaWriteLoginEnforcementRequestWithDefaults instantiates a new MfaWriteLoginEnforcementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaWriteLoginEnforcementRequestWithDefaults() *MfaWriteLoginEnforcementRequest {
	var this MfaWriteLoginEnforcementRequest

	return &this
}

func (o MfaWriteLoginEnforcementRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["auth_method_accessors"] = o.AuthMethodAccessors
	toSerialize["auth_method_types"] = o.AuthMethodTypes
	toSerialize["identity_entity_ids"] = o.IdentityEntityIds
	toSerialize["identity_group_ids"] = o.IdentityGroupIds
	toSerialize["mfa_method_ids"] = o.MfaMethodIds

	return json.Marshal(toSerialize)
}
