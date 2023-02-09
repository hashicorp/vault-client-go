// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadRoleIDResponse struct for AppRoleReadRoleIDResponse
type AppRoleReadRoleIDResponse struct {
	// Identifier of the role. Defaults to a UUID.
	RoleId string `json:"role_id"`
}

// NewAppRoleReadRoleIDResponseWithDefaults instantiates a new AppRoleReadRoleIDResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadRoleIDResponseWithDefaults() *AppRoleReadRoleIDResponse {
	var this AppRoleReadRoleIDResponse

	return &this
}

func (o AppRoleReadRoleIDResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
