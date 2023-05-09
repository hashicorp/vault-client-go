// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleWriteSecretIdBoundCidrsRequest struct for AppRoleWriteSecretIdBoundCidrsRequest
type AppRoleWriteSecretIdBoundCidrsRequest struct {
	// Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can perform the login operation.
	SecretIdBoundCidrs []string `json:"secret_id_bound_cidrs,omitempty"`
}

// NewAppRoleWriteSecretIdBoundCidrsRequestWithDefaults instantiates a new AppRoleWriteSecretIdBoundCidrsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteSecretIdBoundCidrsRequestWithDefaults() *AppRoleWriteSecretIdBoundCidrsRequest {
	var this AppRoleWriteSecretIdBoundCidrsRequest

	return &this
}
