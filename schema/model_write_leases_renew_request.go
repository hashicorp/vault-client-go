// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteLeasesRenewRequest struct for WriteLeasesRenewRequest
type WriteLeasesRenewRequest struct {
	// The desired increment in seconds to the lease
	Increment int32 `json:"increment"`

	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id"`

	// The lease identifier to renew. This is included with a lease.
	UrlLeaseId string `json:"url_lease_id"`
}

// NewWriteLeasesRenewRequestWithDefaults instantiates a new WriteLeasesRenewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteLeasesRenewRequestWithDefaults() *WriteLeasesRenewRequest {
	var this WriteLeasesRenewRequest

	return &this
}

func (o WriteLeasesRenewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
