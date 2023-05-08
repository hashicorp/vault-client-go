// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiListRevokedCertsResponse struct for PkiListRevokedCertsResponse
type PkiListRevokedCertsResponse struct {
	// List of Keys
	Keys []string `json:"keys,omitempty"`
}

// NewPkiListRevokedCertsResponseWithDefaults instantiates a new PkiListRevokedCertsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiListRevokedCertsResponseWithDefaults() *PkiListRevokedCertsResponse {
	var this PkiListRevokedCertsResponse

	return &this
}
