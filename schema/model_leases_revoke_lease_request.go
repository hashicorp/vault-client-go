// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LeasesRevokeLeaseRequest struct for LeasesRevokeLeaseRequest
type LeasesRevokeLeaseRequest struct {
	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id"`

	// Whether or not to perform the revocation synchronously
	Sync bool `json:"sync"`

	// The lease identifier to renew. This is included with a lease.
	UrlLeaseId string `json:"url_lease_id"`
}

// NewLeasesRevokeLeaseRequestWithDefaults instantiates a new LeasesRevokeLeaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeasesRevokeLeaseRequestWithDefaults() *LeasesRevokeLeaseRequest {
	var this LeasesRevokeLeaseRequest

	this.Sync = true

	return &this
}

func (o LeasesRevokeLeaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["lease_id"] = o.LeaseId
	toSerialize["sync"] = o.Sync
	toSerialize["url_lease_id"] = o.UrlLeaseId

	return json.Marshal(toSerialize)
}
