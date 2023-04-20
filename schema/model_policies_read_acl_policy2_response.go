// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PoliciesReadAclPolicy2Response struct for PoliciesReadAclPolicy2Response
type PoliciesReadAclPolicy2Response struct {
	Name string `json:"name"`

	Policy string `json:"policy"`

	Rules string `json:"rules"`
}

// NewPoliciesReadAclPolicy2ResponseWithDefaults instantiates a new PoliciesReadAclPolicy2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesReadAclPolicy2ResponseWithDefaults() *PoliciesReadAclPolicy2Response {
	var this PoliciesReadAclPolicy2Response

	return &this
}

func (o PoliciesReadAclPolicy2Response) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["name"] = o.Name
	toSerialize["policy"] = o.Policy
	toSerialize["rules"] = o.Rules

	return json.Marshal(toSerialize)
}
