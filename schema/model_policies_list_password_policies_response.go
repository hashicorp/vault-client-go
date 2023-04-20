// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PoliciesListPasswordPoliciesResponse struct for PoliciesListPasswordPoliciesResponse
type PoliciesListPasswordPoliciesResponse struct {
	Keys []string `json:"keys"`
}

// NewPoliciesListPasswordPoliciesResponseWithDefaults instantiates a new PoliciesListPasswordPoliciesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesListPasswordPoliciesResponseWithDefaults() *PoliciesListPasswordPoliciesResponse {
	var this PoliciesListPasswordPoliciesResponse

	return &this
}

func (o PoliciesListPasswordPoliciesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["keys"] = o.Keys

	return json.Marshal(toSerialize)
}
