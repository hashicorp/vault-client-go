// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// QueryTokenAccessorCapabilitiesRequest struct for QueryTokenAccessorCapabilitiesRequest
type QueryTokenAccessorCapabilitiesRequest struct {
	// Accessor of the token for which capabilities are being queried.
	Accessor string `json:"accessor,omitempty"`

	// Use 'paths' instead.
	// Deprecated
	Path []string `json:"path,omitempty"`

	// Paths on which capabilities are being queried.
	Paths []string `json:"paths,omitempty"`
}

// NewQueryTokenAccessorCapabilitiesRequestWithDefaults instantiates a new QueryTokenAccessorCapabilitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryTokenAccessorCapabilitiesRequestWithDefaults() *QueryTokenAccessorCapabilitiesRequest {
	var this QueryTokenAccessorCapabilitiesRequest

	return &this
}
