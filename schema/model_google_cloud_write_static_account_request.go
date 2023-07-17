// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudWriteStaticAccountRequest struct for GoogleCloudWriteStaticAccountRequest
type GoogleCloudWriteStaticAccountRequest struct {
	// Bindings configuration string.
	Bindings string `json:"bindings,omitempty"`

	// Type of secret generated for this account. Cannot be updated. Defaults to \"access_token\"
	SecretType string `json:"secret_type,omitempty"`

	// Required. Email of the GCP service account to manage. Cannot be updated.
	ServiceAccountEmail string `json:"service_account_email,omitempty"`

	// List of OAuth scopes to assign to access tokens generated under this account. Ignored if \"secret_type\" is not \"\"access_token\"\"
	TokenScopes []string `json:"token_scopes,omitempty"`
}
