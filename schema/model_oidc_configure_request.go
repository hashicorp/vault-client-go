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
