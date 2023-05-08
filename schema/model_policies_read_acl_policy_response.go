// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PoliciesReadAclPolicyResponse struct for PoliciesReadAclPolicyResponse
type PoliciesReadAclPolicyResponse struct {
	Name string `json:"name,omitempty"`

	Policy string `json:"policy,omitempty"`

	Rules string `json:"rules,omitempty"`
}

// NewPoliciesReadAclPolicyResponseWithDefaults instantiates a new PoliciesReadAclPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesReadAclPolicyResponseWithDefaults() *PoliciesReadAclPolicyResponse {
	var this PoliciesReadAclPolicyResponse

	return &this
}
