// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteTokenBoundCIDRsRequest struct for AppRoleWriteTokenBoundCIDRsRequest
type AppRoleWriteTokenBoundCIDRsRequest struct {
	// Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token.
	TokenBoundCidrs []string `json:"token_bound_cidrs"`
}

// NewAppRoleWriteTokenBoundCIDRsRequestWithDefaults instantiates a new AppRoleWriteTokenBoundCIDRsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteTokenBoundCIDRsRequestWithDefaults() *AppRoleWriteTokenBoundCIDRsRequest {
	var this AppRoleWriteTokenBoundCIDRsRequest

	return &this
}

func (o AppRoleWriteTokenBoundCIDRsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token_bound_cidrs"] = o.TokenBoundCidrs

	return json.Marshal(toSerialize)
}
