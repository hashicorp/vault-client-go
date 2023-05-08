// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AwsTidyIdentityWhitelistRequest struct for AwsTidyIdentityWhitelistRequest
type AwsTidyIdentityWhitelistRequest struct {
	// The amount of extra time that must have passed beyond the identity's expiration, before it is removed from the backend storage.
	SafetyBuffer int32 `json:"safety_buffer,omitempty"`
}

// NewAwsTidyIdentityWhitelistRequestWithDefaults instantiates a new AwsTidyIdentityWhitelistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsTidyIdentityWhitelistRequestWithDefaults() *AwsTidyIdentityWhitelistRequest {
	var this AwsTidyIdentityWhitelistRequest

	this.SafetyBuffer = 259200

	return &this
}
