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

// NewOidcWriteProviderRequestWithDefaults instantiates a new OidcWriteProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcWriteProviderRequestWithDefaults() *OidcWriteProviderRequest {
	var this OidcWriteProviderRequest

	return &this
}
