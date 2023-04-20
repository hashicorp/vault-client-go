// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PoliciesWriteAclPolicy2Request struct for PoliciesWriteAclPolicy2Request
type PoliciesWriteAclPolicy2Request struct {
	// The rules of the policy.
	Policy string `json:"policy"`

	// The rules of the policy.
	// Deprecated
	Rules string `json:"rules"`
}

// NewPoliciesWriteAclPolicy2RequestWithDefaults instantiates a new PoliciesWriteAclPolicy2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesWriteAclPolicy2RequestWithDefaults() *PoliciesWriteAclPolicy2Request {
	var this PoliciesWriteAclPolicy2Request

	return &this
}

func (o PoliciesWriteAclPolicy2Request) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["policy"] = o.Policy
	toSerialize["rules"] = o.Rules

	return json.Marshal(toSerialize)
}
