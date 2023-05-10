// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PoliciesReadPasswordPolicyResponse struct for PoliciesReadPasswordPolicyResponse
type PoliciesReadPasswordPolicyResponse struct {
	Policy string `json:"policy,omitempty"`
}

// NewPoliciesReadPasswordPolicyResponseWithDefaults instantiates a new PoliciesReadPasswordPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesReadPasswordPolicyResponseWithDefaults() *PoliciesReadPasswordPolicyResponse {
	var this PoliciesReadPasswordPolicyResponse

	return &this
}
