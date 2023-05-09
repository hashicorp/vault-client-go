// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// LeasesCountResponse struct for LeasesCountResponse
type LeasesCountResponse struct {
	// Number of matching leases per mount
	Counts int32 `json:"counts,omitempty"`

	// Number of matching leases
	LeaseCount int32 `json:"lease_count,omitempty"`
}

// NewLeasesCountResponseWithDefaults instantiates a new LeasesCountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeasesCountResponseWithDefaults() *LeasesCountResponse {
	var this LeasesCountResponse

	return &this
}
