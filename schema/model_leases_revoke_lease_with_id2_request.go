// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LeasesRevokeLeaseWithId2Request struct for LeasesRevokeLeaseWithId2Request
type LeasesRevokeLeaseWithId2Request struct {
	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id"`

	// Whether or not to perform the revocation synchronously
	Sync bool `json:"sync"`
}

// NewLeasesRevokeLeaseWithId2RequestWithDefaults instantiates a new LeasesRevokeLeaseWithId2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeasesRevokeLeaseWithId2RequestWithDefaults() *LeasesRevokeLeaseWithId2Request {
	var this LeasesRevokeLeaseWithId2Request

	this.Sync = true

	return &this
}

func (o LeasesRevokeLeaseWithId2Request) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["lease_id"] = o.LeaseId
	toSerialize["sync"] = o.Sync

	return json.Marshal(toSerialize)
}