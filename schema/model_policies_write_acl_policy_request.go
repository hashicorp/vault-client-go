// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PoliciesWriteAclPolicyRequest struct for PoliciesWriteAclPolicyRequest
type PoliciesWriteAclPolicyRequest struct {
	// The rules of the policy.
	Policy string `json:"policy,omitempty"`
}

// NewPoliciesWriteAclPolicyRequestWithDefaults instantiates a new PoliciesWriteAclPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesWriteAclPolicyRequestWithDefaults() *PoliciesWriteAclPolicyRequest {
	var this PoliciesWriteAclPolicyRequest

	return &this
}
