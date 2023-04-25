// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LeasesRenewLeaseWithId2Request struct for LeasesRenewLeaseWithId2Request
type LeasesRenewLeaseWithId2Request struct {
	// The desired increment in seconds to the lease
	Increment int32 `json:"increment"`

	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id"`
}

// NewLeasesRenewLeaseWithId2RequestWithDefaults instantiates a new LeasesRenewLeaseWithId2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeasesRenewLeaseWithId2RequestWithDefaults() *LeasesRenewLeaseWithId2Request {
	var this LeasesRenewLeaseWithId2Request

	return &this
}

func (o LeasesRenewLeaseWithId2Request) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["increment"] = o.Increment
	toSerialize["lease_id"] = o.LeaseId

	return json.Marshal(toSerialize)
}
