// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// OidcWriteProviderRequest struct for OidcWriteProviderRequest
type OidcWriteProviderRequest struct {
	// The client IDs that are permitted to use the provider
	AllowedClientIds []string `json:"allowed_client_ids,omitempty"`

	// Specifies what will be used for the iss claim of ID tokens.
	Issuer string `json:"issuer,omitempty"`

	// The scopes supported for requesting on the provider
	ScopesSupported []string `json:"scopes_supported,omitempty"`
}
