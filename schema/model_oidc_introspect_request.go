// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// OidcIntrospectRequest struct for OidcIntrospectRequest
type OidcIntrospectRequest struct {
	// Optional client_id to verify
	ClientId string `json:"client_id,omitempty"`

	// Token to verify
	Token string `json:"token,omitempty"`
}
