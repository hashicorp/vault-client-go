// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// LeasesRenewLeaseRequest struct for LeasesRenewLeaseRequest
type LeasesRenewLeaseRequest struct {
	// The desired increment in seconds to the lease
	Increment string `json:"increment,omitempty"`

	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id,omitempty"`
}
