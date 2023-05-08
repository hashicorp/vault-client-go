// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// LeasesLookUpWithPrefixResponse struct for LeasesLookUpWithPrefixResponse
type LeasesLookUpWithPrefixResponse struct {
	// A list of lease ids
	Keys []string `json:"keys,omitempty"`
}

// NewLeasesLookUpWithPrefixResponseWithDefaults instantiates a new LeasesLookUpWithPrefixResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeasesLookUpWithPrefixResponseWithDefaults() *LeasesLookUpWithPrefixResponse {
	var this LeasesLookUpWithPrefixResponse

	return &this
}
