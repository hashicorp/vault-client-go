// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteLeasesLookupRequest struct for WriteLeasesLookupRequest
type WriteLeasesLookupRequest struct {
	// The lease identifier to renew. This is included with a lease.
	LeaseId string `json:"lease_id"`
}

// NewWriteLeasesLookupRequestWithDefaults instantiates a new WriteLeasesLookupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteLeasesLookupRequestWithDefaults() *WriteLeasesLookupRequest {
	var this WriteLeasesLookupRequest

	return &this
}

func (o WriteLeasesLookupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
