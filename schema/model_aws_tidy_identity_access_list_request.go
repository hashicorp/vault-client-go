// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AwsTidyIdentityAccessListRequest struct for AwsTidyIdentityAccessListRequest
type AwsTidyIdentityAccessListRequest struct {
	// The amount of extra time that must have passed beyond the identity's expiration, before it is removed from the backend storage.
	SafetyBuffer string `json:"safety_buffer,omitempty"`
}
