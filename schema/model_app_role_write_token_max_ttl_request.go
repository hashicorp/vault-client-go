// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleWriteTokenMaxTtlRequest struct for AppRoleWriteTokenMaxTtlRequest
type AppRoleWriteTokenMaxTtlRequest struct {
	// The maximum lifetime of the generated token
	TokenMaxTtl string `json:"token_max_ttl,omitempty"`
}

// NewAppRoleWriteTokenMaxTtlRequestWithDefaults instantiates a new AppRoleWriteTokenMaxTtlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteTokenMaxTtlRequestWithDefaults() *AppRoleWriteTokenMaxTtlRequest {
	var this AppRoleWriteTokenMaxTtlRequest

	return &this
}
