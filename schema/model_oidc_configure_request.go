// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// OidcConfigureRequest struct for OidcConfigureRequest
type OidcConfigureRequest struct {
	// Issuer URL to be used in the iss claim of the token. If not set, Vault's app_addr will be used.
	Issuer string `json:"issuer,omitempty"`
}

// NewOidcConfigureRequestWithDefaults instantiates a new OidcConfigureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcConfigureRequestWithDefaults() *OidcConfigureRequest {
	var this OidcConfigureRequest

	return &this
}
