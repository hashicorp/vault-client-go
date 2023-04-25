// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PoliciesGeneratePasswordFromPasswordPolicyResponse struct for PoliciesGeneratePasswordFromPasswordPolicyResponse
type PoliciesGeneratePasswordFromPasswordPolicyResponse struct {
	Password string `json:"password"`
}

// NewPoliciesGeneratePasswordFromPasswordPolicyResponseWithDefaults instantiates a new PoliciesGeneratePasswordFromPasswordPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesGeneratePasswordFromPasswordPolicyResponseWithDefaults() *PoliciesGeneratePasswordFromPasswordPolicyResponse {
	var this PoliciesGeneratePasswordFromPasswordPolicyResponse

	return &this
}

func (o PoliciesGeneratePasswordFromPasswordPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["password"] = o.Password

	return json.Marshal(toSerialize)
}
