// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteTokenMaxTtlRequest struct for AppRoleWriteTokenMaxTtlRequest
type AppRoleWriteTokenMaxTtlRequest struct {
	// The maximum lifetime of the generated token
	TokenMaxTtl int32 `json:"token_max_ttl"`
}

// NewAppRoleWriteTokenMaxTtlRequestWithDefaults instantiates a new AppRoleWriteTokenMaxTtlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteTokenMaxTtlRequestWithDefaults() *AppRoleWriteTokenMaxTtlRequest {
	var this AppRoleWriteTokenMaxTtlRequest

	return &this
}

func (o AppRoleWriteTokenMaxTtlRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token_max_ttl"] = o.TokenMaxTtl

	return json.Marshal(toSerialize)
}
