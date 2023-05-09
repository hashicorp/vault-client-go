// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// UserpassUpdatePoliciesRequest struct for UserpassUpdatePoliciesRequest
type UserpassUpdatePoliciesRequest struct {
	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies,omitempty"`

	// Comma-separated list of policies
	TokenPolicies []string `json:"token_policies,omitempty"`
}

// NewUserpassUpdatePoliciesRequestWithDefaults instantiates a new UserpassUpdatePoliciesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserpassUpdatePoliciesRequestWithDefaults() *UserpassUpdatePoliciesRequest {
	var this UserpassUpdatePoliciesRequest

	return &this
}
