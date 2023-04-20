// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LeasesRevokeLeaseWithIdRequest struct for LeasesRevokeLeaseWithIdRequest
type LeasesRevokeLeaseWithIdRequest struct {
	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id"`

	// Whether or not to perform the revocation synchronously
	Sync bool `json:"sync"`
}

// NewLeasesRevokeLeaseWithIdRequestWithDefaults instantiates a new LeasesRevokeLeaseWithIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeasesRevokeLeaseWithIdRequestWithDefaults() *LeasesRevokeLeaseWithIdRequest {
	var this LeasesRevokeLeaseWithIdRequest

	this.Sync = true

	return &this
}

func (o LeasesRevokeLeaseWithIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["lease_id"] = o.LeaseId
	toSerialize["sync"] = o.Sync

	return json.Marshal(toSerialize)
}
