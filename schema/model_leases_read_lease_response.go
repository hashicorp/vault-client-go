// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import "time"

// LeasesReadLeaseResponse struct for LeasesReadLeaseResponse
type LeasesReadLeaseResponse struct {
	// Optional lease expiry time
	ExpireTime time.Time `json:"expire_time,omitempty"`

	// Lease id
	Id string `json:"id,omitempty"`

	// Timestamp for the lease's issue time
	IssueTime time.Time `json:"issue_time,omitempty"`

	// Optional Timestamp of the last time the lease was renewed
	LastRenewal time.Time `json:"last_renewal,omitempty"`

	// True if the lease is able to be renewed
	Renewable bool `json:"renewable,omitempty"`

	// Time to Live set for the lease, returns 0 if unset
	Ttl int32 `json:"ttl,omitempty"`
}
