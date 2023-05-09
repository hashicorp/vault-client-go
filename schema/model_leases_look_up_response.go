// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// LeasesLookUpResponse struct for LeasesLookUpResponse
type LeasesLookUpResponse struct {
	// A list of lease ids
	Keys []string `json:"keys,omitempty"`
}

// NewLeasesLookUpResponseWithDefaults instantiates a new LeasesLookUpResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeasesLookUpResponseWithDefaults() *LeasesLookUpResponse {
	var this LeasesLookUpResponse

	return &this
}
