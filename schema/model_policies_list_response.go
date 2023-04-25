// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PoliciesListResponse struct for PoliciesListResponse
type PoliciesListResponse struct {
	Keys []string `json:"keys"`

	Policies []string `json:"policies"`
}

// NewPoliciesListResponseWithDefaults instantiates a new PoliciesListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesListResponseWithDefaults() *PoliciesListResponse {
	var this PoliciesListResponse

	return &this
}

func (o PoliciesListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["keys"] = o.Keys
	toSerialize["policies"] = o.Policies

	return json.Marshal(toSerialize)
}
