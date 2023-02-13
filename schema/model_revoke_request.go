// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RevokeRequest struct for RevokeRequest
type RevokeRequest struct {
	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id"`

	// Whether or not to perform the revocation synchronously
	Sync bool `json:"sync"`

	// The lease identifier to renew. This is included with a lease.
	UrlLeaseId string `json:"url_lease_id"`
}

// NewRevokeRequestWithDefaults instantiates a new RevokeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokeRequestWithDefaults() *RevokeRequest {
	var this RevokeRequest

	this.Sync = true

	return &this
}

func (o RevokeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["lease_id"] = o.LeaseId
	toSerialize["sync"] = o.Sync
	toSerialize["url_lease_id"] = o.UrlLeaseId

	return json.Marshal(toSerialize)
}
