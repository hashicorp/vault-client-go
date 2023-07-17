// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudWriteRolesetRequest struct for GoogleCloudWriteRolesetRequest
type GoogleCloudWriteRolesetRequest struct {
	// Bindings configuration string.
	Bindings string `json:"bindings,omitempty"`

	// Name of the GCP project that this roleset's service account will belong to.
	Project string `json:"project,omitempty"`

	// Type of secret generated for this role set. Defaults to 'access_token'
	SecretType string `json:"secret_type,omitempty"`

	// List of OAuth scopes to assign to credentials generated under this role set
	TokenScopes []string `json:"token_scopes,omitempty"`
}
