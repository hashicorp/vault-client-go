// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// LeasesRevokeLeaseWithIdRequest struct for LeasesRevokeLeaseWithIdRequest
type LeasesRevokeLeaseWithIdRequest struct {
	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id,omitempty"`

	// Whether or not to perform the revocation synchronously
	Sync bool `json:"sync,omitempty"`
}
