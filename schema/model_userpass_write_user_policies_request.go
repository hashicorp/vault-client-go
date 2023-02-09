// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// UserpassWriteUserPoliciesRequest struct for UserpassWriteUserPoliciesRequest
type UserpassWriteUserPoliciesRequest struct {
	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies"`

	// Comma-separated list of policies
	TokenPolicies []string `json:"token_policies"`
}

// NewUserpassWriteUserPoliciesRequestWithDefaults instantiates a new UserpassWriteUserPoliciesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserpassWriteUserPoliciesRequestWithDefaults() *UserpassWriteUserPoliciesRequest {
	var this UserpassWriteUserPoliciesRequest

	return &this
}

func (o UserpassWriteUserPoliciesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
