// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// LeasesReadLeaseRequest struct for LeasesReadLeaseRequest
type LeasesReadLeaseRequest struct {
	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id,omitempty"`
}

// NewLeasesReadLeaseRequestWithDefaults instantiates a new LeasesReadLeaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeasesReadLeaseRequestWithDefaults() *LeasesReadLeaseRequest {
	var this LeasesReadLeaseRequest

	return &this
}
