// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// MfaWriteLoginEnforcementRequest struct for MfaWriteLoginEnforcementRequest
type MfaWriteLoginEnforcementRequest struct {
	// Array of auth mount accessor IDs
	AuthMethodAccessors []string `json:"auth_method_accessors,omitempty"`

	// Array of auth mount types
	AuthMethodTypes []string `json:"auth_method_types,omitempty"`

	// Array of identity entity IDs
	IdentityEntityIds []string `json:"identity_entity_ids,omitempty"`

	// Array of identity group IDs
	IdentityGroupIds []string `json:"identity_group_ids,omitempty"`

	// Array of Method IDs that determine what methods will be enforced
	MfaMethodIds []string `json:"mfa_method_ids"`
}
