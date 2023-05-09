// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PoliciesListPasswordPoliciesResponse struct for PoliciesListPasswordPoliciesResponse
type PoliciesListPasswordPoliciesResponse struct {
	Keys []string `json:"keys,omitempty"`
}

// NewPoliciesListPasswordPoliciesResponseWithDefaults instantiates a new PoliciesListPasswordPoliciesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesListPasswordPoliciesResponseWithDefaults() *PoliciesListPasswordPoliciesResponse {
	var this PoliciesListPasswordPoliciesResponse

	return &this
}
