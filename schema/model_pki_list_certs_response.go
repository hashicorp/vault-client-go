// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiListCertsResponse struct for PkiListCertsResponse
type PkiListCertsResponse struct {
	// A list of keys
	Keys []string `json:"keys,omitempty"`
}

// NewPkiListCertsResponseWithDefaults instantiates a new PkiListCertsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiListCertsResponseWithDefaults() *PkiListCertsResponse {
	var this PkiListCertsResponse

	return &this
}
