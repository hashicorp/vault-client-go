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
