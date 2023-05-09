// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleWriteRoleIdRequest struct for AppRoleWriteRoleIdRequest
type AppRoleWriteRoleIdRequest struct {
	// Identifier of the role. Defaults to a UUID.
	RoleId string `json:"role_id,omitempty"`
}

// NewAppRoleWriteRoleIdRequestWithDefaults instantiates a new AppRoleWriteRoleIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteRoleIdRequestWithDefaults() *AppRoleWriteRoleIdRequest {
	var this AppRoleWriteRoleIdRequest

	return &this
}
