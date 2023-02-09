// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteLeasesRenew2Request struct for WriteLeasesRenew2Request
type WriteLeasesRenew2Request struct {
	// The desired increment in seconds to the lease
	Increment int32 `json:"increment"`

	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id"`
}

// NewWriteLeasesRenew2RequestWithDefaults instantiates a new WriteLeasesRenew2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteLeasesRenew2RequestWithDefaults() *WriteLeasesRenew2Request {
	var this WriteLeasesRenew2Request

	return &this
}

func (o WriteLeasesRenew2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
