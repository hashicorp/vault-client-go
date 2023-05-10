// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// QueryTokenSelfCapabilitiesRequest struct for QueryTokenSelfCapabilitiesRequest
type QueryTokenSelfCapabilitiesRequest struct {
	// Use 'paths' instead.
	// Deprecated
	Path []string `json:"path,omitempty"`

	// Paths on which capabilities are being queried.
	Paths []string `json:"paths,omitempty"`

	// Token for which capabilities are being queried.
	Token string `json:"token,omitempty"`
}

// NewQueryTokenSelfCapabilitiesRequestWithDefaults instantiates a new QueryTokenSelfCapabilitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryTokenSelfCapabilitiesRequestWithDefaults() *QueryTokenSelfCapabilitiesRequest {
	var this QueryTokenSelfCapabilitiesRequest

	return &this
}
