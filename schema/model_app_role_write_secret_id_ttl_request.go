// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleWriteSecretIdTtlRequest struct for AppRoleWriteSecretIdTtlRequest
type AppRoleWriteSecretIdTtlRequest struct {
	// Duration in seconds after which the issued SecretID should expire. Defaults to 0, meaning no expiration.
	SecretIdTtl string `json:"secret_id_ttl,omitempty"`
}
