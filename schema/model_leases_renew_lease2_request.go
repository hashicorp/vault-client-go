// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LeasesRenewLease2Request struct for LeasesRenewLease2Request
type LeasesRenewLease2Request struct {
	// The desired increment in seconds to the lease
	Increment int32 `json:"increment"`

	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id"`

	// The lease identifier to renew. This is included with a lease.
	UrlLeaseId string `json:"url_lease_id"`
}

// NewLeasesRenewLease2RequestWithDefaults instantiates a new LeasesRenewLease2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeasesRenewLease2RequestWithDefaults() *LeasesRenewLease2Request {
	var this LeasesRenewLease2Request

	return &this
}

func (o LeasesRenewLease2Request) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["increment"] = o.Increment
	toSerialize["lease_id"] = o.LeaseId
	toSerialize["url_lease_id"] = o.UrlLeaseId

	return json.Marshal(toSerialize)
}