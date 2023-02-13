// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WritePoliciesACLRequest struct for WritePoliciesACLRequest
type WritePoliciesACLRequest struct {
	// The rules of the policy.
	Policy string `json:"policy"`
}

// NewWritePoliciesACLRequestWithDefaults instantiates a new WritePoliciesACLRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritePoliciesACLRequestWithDefaults() *WritePoliciesACLRequest {
	var this WritePoliciesACLRequest

	return &this
}

func (o WritePoliciesACLRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["policy"] = o.Policy

	return json.Marshal(toSerialize)
}
