// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import "time"

// PkiRevokeResponse struct for PkiRevokeResponse
type PkiRevokeResponse struct {
	// Revocation Time
	RevocationTime int32 `json:"revocation_time,omitempty"`

	// Revocation Time
	RevocationTimeRfc3339 time.Time `json:"revocation_time_rfc3339,omitempty"`

	// Revocation State
	State string `json:"state,omitempty"`
}

// NewPkiRevokeResponseWithDefaults instantiates a new PkiRevokeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiRevokeResponseWithDefaults() *PkiRevokeResponse {
	var this PkiRevokeResponse

	return &this
}
