// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// LeasesListResponse struct for LeasesListResponse
type LeasesListResponse struct {
	// Number of matching leases per mount
	Counts int32 `json:"counts,omitempty"`

	// Number of matching leases
	LeaseCount int32 `json:"lease_count,omitempty"`
}
