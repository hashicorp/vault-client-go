// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleLoginRequest struct for AppRoleLoginRequest
type AppRoleLoginRequest struct {
	// Unique identifier of the Role. Required to be supplied when the 'bind_secret_id' constraint is set.
	RoleId string `json:"role_id,omitempty"`

	// SecretID belong to the App role
	SecretId string `json:"secret_id,omitempty"`
}
