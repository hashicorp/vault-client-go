// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleReadBindSecretIdResponse struct for AppRoleReadBindSecretIdResponse
type AppRoleReadBindSecretIdResponse struct {
	// Impose secret_id to be presented when logging in using this role. Defaults to 'true'.
	BindSecretId bool `json:"bind_secret_id,omitempty"`
}
