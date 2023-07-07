// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AwsTidyRoleTagDenyListRequest struct for AwsTidyRoleTagDenyListRequest
type AwsTidyRoleTagDenyListRequest struct {
	// The amount of extra time that must have passed beyond the roletag expiration, before it is removed from the backend storage.
	SafetyBuffer string `json:"safety_buffer,omitempty"`
}

// NewAwsTidyRoleTagDenyListRequestWithDefaults instantiates a new AwsTidyRoleTagDenyListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsTidyRoleTagDenyListRequestWithDefaults() *AwsTidyRoleTagDenyListRequest {
	var this AwsTidyRoleTagDenyListRequest

	this.SafetyBuffer = "259200"

	return &this
}
