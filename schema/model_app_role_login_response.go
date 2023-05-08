// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleLoginResponse struct for AppRoleLoginResponse
type AppRoleLoginResponse struct {
	Role string `json:"role,omitempty"`
}

// NewAppRoleLoginResponseWithDefaults instantiates a new AppRoleLoginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleLoginResponseWithDefaults() *AppRoleLoginResponse {
	var this AppRoleLoginResponse

	return &this
}
