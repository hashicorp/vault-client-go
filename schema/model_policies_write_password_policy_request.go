// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PoliciesWritePasswordPolicyRequest struct for PoliciesWritePasswordPolicyRequest
type PoliciesWritePasswordPolicyRequest struct {
	// The password policy
	Policy string `json:"policy,omitempty"`
}

// NewPoliciesWritePasswordPolicyRequestWithDefaults instantiates a new PoliciesWritePasswordPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesWritePasswordPolicyRequestWithDefaults() *PoliciesWritePasswordPolicyRequest {
	var this PoliciesWritePasswordPolicyRequest

	return &this
}
