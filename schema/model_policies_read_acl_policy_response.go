// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PoliciesReadAclPolicyResponse struct for PoliciesReadAclPolicyResponse
type PoliciesReadAclPolicyResponse struct {
	Name string `json:"name"`

	Policy string `json:"policy"`

	Rules string `json:"rules"`
}

// NewPoliciesReadAclPolicyResponseWithDefaults instantiates a new PoliciesReadAclPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesReadAclPolicyResponseWithDefaults() *PoliciesReadAclPolicyResponse {
	var this PoliciesReadAclPolicyResponse

	return &this
}

func (o PoliciesReadAclPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["name"] = o.Name
	toSerialize["policy"] = o.Policy
	toSerialize["rules"] = o.Rules

	return json.Marshal(toSerialize)
}
