// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PoliciesListAclPoliciesResponse struct for PoliciesListAclPoliciesResponse
type PoliciesListAclPoliciesResponse struct {
	Keys []string `json:"keys,omitempty"`

	Policies []string `json:"policies,omitempty"`
}

// NewPoliciesListAclPoliciesResponseWithDefaults instantiates a new PoliciesListAclPoliciesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesListAclPoliciesResponseWithDefaults() *PoliciesListAclPoliciesResponse {
	var this PoliciesListAclPoliciesResponse

	return &this
}
