// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// LeasesRenewLeaseWithIdRequest struct for LeasesRenewLeaseWithIdRequest
type LeasesRenewLeaseWithIdRequest struct {
	// The desired increment in seconds to the lease
	Increment int32 `json:"increment,omitempty"`

	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id,omitempty"`
}

// NewLeasesRenewLeaseWithIdRequestWithDefaults instantiates a new LeasesRenewLeaseWithIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeasesRenewLeaseWithIdRequestWithDefaults() *LeasesRenewLeaseWithIdRequest {
	var this LeasesRenewLeaseWithIdRequest

	return &this
}
