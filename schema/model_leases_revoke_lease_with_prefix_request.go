// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// LeasesRevokeLeaseWithPrefixRequest struct for LeasesRevokeLeaseWithPrefixRequest
type LeasesRevokeLeaseWithPrefixRequest struct {
	// Whether or not to perform the revocation synchronously
	Sync bool `json:"sync,omitempty"`
}

// NewLeasesRevokeLeaseWithPrefixRequestWithDefaults instantiates a new LeasesRevokeLeaseWithPrefixRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeasesRevokeLeaseWithPrefixRequestWithDefaults() *LeasesRevokeLeaseWithPrefixRequest {
	var this LeasesRevokeLeaseWithPrefixRequest

	this.Sync = true

	return &this
}
