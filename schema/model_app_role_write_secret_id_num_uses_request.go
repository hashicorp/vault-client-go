// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleWriteSecretIdNumUsesRequest struct for AppRoleWriteSecretIdNumUsesRequest
type AppRoleWriteSecretIdNumUsesRequest struct {
	// Number of times a SecretID can access the role, after which the SecretID will expire.
	SecretIdNumUses int32 `json:"secret_id_num_uses,omitempty"`
}
