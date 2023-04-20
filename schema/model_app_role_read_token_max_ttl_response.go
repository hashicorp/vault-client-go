// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadTokenMaxTtlResponse struct for AppRoleReadTokenMaxTtlResponse
type AppRoleReadTokenMaxTtlResponse struct {
	// The maximum lifetime of the generated token
	TokenMaxTtl int32 `json:"token_max_ttl"`
}

// NewAppRoleReadTokenMaxTtlResponseWithDefaults instantiates a new AppRoleReadTokenMaxTtlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadTokenMaxTtlResponseWithDefaults() *AppRoleReadTokenMaxTtlResponse {
	var this AppRoleReadTokenMaxTtlResponse

	return &this
}

func (o AppRoleReadTokenMaxTtlResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token_max_ttl"] = o.TokenMaxTtl

	return json.Marshal(toSerialize)
}
