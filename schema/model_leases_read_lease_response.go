// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
	"time"
)

// LeasesReadLeaseResponse struct for LeasesReadLeaseResponse
type LeasesReadLeaseResponse struct {
	// Optional lease expiry time
	ExpireTime time.Time `json:"expire_time"`

	// Lease id
	Id string `json:"id"`

	// Timestamp for the lease's issue time
	IssueTime time.Time `json:"issue_time"`

	// Optional Timestamp of the last time the lease was renewed
	LastRenewal time.Time `json:"last_renewal"`

	// True if the lease is able to be renewed
	Renewable bool `json:"renewable"`

	// Time to Live set for the lease, returns 0 if unset
	Ttl int32 `json:"ttl"`
}

// NewLeasesReadLeaseResponseWithDefaults instantiates a new LeasesReadLeaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeasesReadLeaseResponseWithDefaults() *LeasesReadLeaseResponse {
	var this LeasesReadLeaseResponse

	return &this
}

func (o LeasesReadLeaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["expire_time"] = o.ExpireTime
	toSerialize["id"] = o.Id
	toSerialize["issue_time"] = o.IssueTime
	toSerialize["last_renewal"] = o.LastRenewal
	toSerialize["renewable"] = o.Renewable
	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
