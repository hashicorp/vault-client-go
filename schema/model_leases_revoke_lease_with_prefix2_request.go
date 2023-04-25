// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LeasesRevokeLeaseWithPrefix2Request struct for LeasesRevokeLeaseWithPrefix2Request
type LeasesRevokeLeaseWithPrefix2Request struct {
	// Whether or not to perform the revocation synchronously
	Sync bool `json:"sync"`
}

// NewLeasesRevokeLeaseWithPrefix2RequestWithDefaults instantiates a new LeasesRevokeLeaseWithPrefix2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeasesRevokeLeaseWithPrefix2RequestWithDefaults() *LeasesRevokeLeaseWithPrefix2Request {
	var this LeasesRevokeLeaseWithPrefix2Request

	this.Sync = true

	return &this
}

func (o LeasesRevokeLeaseWithPrefix2Request) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["sync"] = o.Sync

	return json.Marshal(toSerialize)
}
