// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudWriteImpersonatedAccountRequest struct for GoogleCloudWriteImpersonatedAccountRequest
type GoogleCloudWriteImpersonatedAccountRequest struct {
	// Required. Email of the GCP service account to manage. Cannot be updated.
	ServiceAccountEmail string `json:"service_account_email,omitempty"`

	// List of OAuth scopes to assign to access tokens generated under this account.
	TokenScopes []string `json:"token_scopes,omitempty"`

	// Lifetime of the token for the impersonated account.
	Ttl string `json:"ttl,omitempty"`
}
