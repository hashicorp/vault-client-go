// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleWriteBindSecretIdRequest struct for AppRoleWriteBindSecretIdRequest
type AppRoleWriteBindSecretIdRequest struct {
	// Impose secret_id to be presented when logging in using this role.
	BindSecretId bool `json:"bind_secret_id,omitempty"`
}
