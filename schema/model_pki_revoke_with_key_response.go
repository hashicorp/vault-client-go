// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
	"time"
)

// PkiRevokeWithKeyResponse struct for PkiRevokeWithKeyResponse
type PkiRevokeWithKeyResponse struct {
	// Revocation Time
	RevocationTime int32 `json:"revocation_time"`

	// Revocation Time
	RevocationTimeRfc3339 time.Time `json:"revocation_time_rfc3339"`

	// Revocation State
	State string `json:"state"`
}

// NewPkiRevokeWithKeyResponseWithDefaults instantiates a new PkiRevokeWithKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiRevokeWithKeyResponseWithDefaults() *PkiRevokeWithKeyResponse {
	var this PkiRevokeWithKeyResponse

	return &this
}

func (o PkiRevokeWithKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["revocation_time"] = o.RevocationTime
	toSerialize["revocation_time_rfc3339"] = o.RevocationTimeRfc3339
	toSerialize["state"] = o.State

	return json.Marshal(toSerialize)
}
