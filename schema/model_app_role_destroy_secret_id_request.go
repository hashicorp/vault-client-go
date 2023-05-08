// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleDestroySecretIdRequest struct for AppRoleDestroySecretIdRequest
type AppRoleDestroySecretIdRequest struct {
	// SecretID attached to the role.
	SecretId string `json:"secret_id,omitempty"`
}

// NewAppRoleDestroySecretIdRequestWithDefaults instantiates a new AppRoleDestroySecretIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleDestroySecretIdRequestWithDefaults() *AppRoleDestroySecretIdRequest {
	var this AppRoleDestroySecretIdRequest

	return &this
}
