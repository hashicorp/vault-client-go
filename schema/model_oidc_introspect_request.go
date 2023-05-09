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

// NewOidcIntrospectRequestWithDefaults instantiates a new OidcIntrospectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcIntrospectRequestWithDefaults() *OidcIntrospectRequest {
	var this OidcIntrospectRequest

	return &this
}
