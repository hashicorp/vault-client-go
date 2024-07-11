// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleWriteCustomSecretIdResponse struct for AppRoleWriteCustomSecretIdResponse
type AppRoleWriteCustomSecretIdResponse struct {
	// Secret ID attached to the role.
	SecretId string `json:"secret_id,omitempty"`

	// Accessor of the secret ID
	SecretIdAccessor string `json:"secret_id_accessor,omitempty"`

	// Number of times a secret ID can access the role, after which the secret ID will expire.
	SecretIdNumUses int32 `json:"secret_id_num_uses,omitempty"`

	// Duration in seconds after which the issued secret ID expires.
	SecretIdTtl int64 `json:"secret_id_ttl,omitempty"`
}
