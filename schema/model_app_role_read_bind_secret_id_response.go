// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadBindSecretIdResponse struct for AppRoleReadBindSecretIdResponse
type AppRoleReadBindSecretIdResponse struct {
	// Impose secret_id to be presented when logging in using this role. Defaults to 'true'.
	BindSecretId bool `json:"bind_secret_id"`
}

// NewAppRoleReadBindSecretIdResponseWithDefaults instantiates a new AppRoleReadBindSecretIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadBindSecretIdResponseWithDefaults() *AppRoleReadBindSecretIdResponse {
	var this AppRoleReadBindSecretIdResponse

	return &this
}

func (o AppRoleReadBindSecretIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["bind_secret_id"] = o.BindSecretId

	return json.Marshal(toSerialize)
}
