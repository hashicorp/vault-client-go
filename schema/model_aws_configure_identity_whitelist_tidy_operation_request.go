// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AwsConfigureIdentityWhitelistTidyOperationRequest struct for AwsConfigureIdentityWhitelistTidyOperationRequest
type AwsConfigureIdentityWhitelistTidyOperationRequest struct {
	// If set to 'true', disables the periodic tidying of the 'identity-accesslist/<instance_id>' entries.
	DisablePeriodicTidy bool `json:"disable_periodic_tidy,omitempty"`

	// The amount of extra time that must have passed beyond the identity's expiration, before it is removed from the backend storage.
	SafetyBuffer string `json:"safety_buffer,omitempty"`
}
