// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleWriteTokenTtlRequest struct for AppRoleWriteTokenTtlRequest
type AppRoleWriteTokenTtlRequest struct {
	// The initial ttl of the token to generate
	TokenTtl int32 `json:"token_ttl,omitempty"`
}

// NewAppRoleWriteTokenTtlRequestWithDefaults instantiates a new AppRoleWriteTokenTtlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteTokenTtlRequestWithDefaults() *AppRoleWriteTokenTtlRequest {
	var this AppRoleWriteTokenTtlRequest

	return &this
}
