// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadRoleIdResponse struct for AppRoleReadRoleIdResponse
type AppRoleReadRoleIdResponse struct {
	// Identifier of the role. Defaults to a UUID.
	RoleId string `json:"role_id"`
}

// NewAppRoleReadRoleIdResponseWithDefaults instantiates a new AppRoleReadRoleIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadRoleIdResponseWithDefaults() *AppRoleReadRoleIdResponse {
	var this AppRoleReadRoleIdResponse

	return &this
}

func (o AppRoleReadRoleIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["role_id"] = o.RoleId

	return json.Marshal(toSerialize)
}
