// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleLoginRequest struct for AppRoleLoginRequest
type AppRoleLoginRequest struct {
	// Unique identifier of the Role. Required to be supplied when the 'bind_secret_id' constraint is set.
	RoleId string `json:"role_id"`

	// SecretID belong to the App role
	SecretId string `json:"secret_id"`
}

// NewAppRoleLoginRequestWithDefaults instantiates a new AppRoleLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleLoginRequestWithDefaults() *AppRoleLoginRequest {
	var this AppRoleLoginRequest

	this.SecretId = ""

	return &this
}

func (o AppRoleLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
