// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LeasesLookUpResponse struct for LeasesLookUpResponse
type LeasesLookUpResponse struct {
	// A list of lease ids
	Keys []string `json:"keys"`
}

// NewLeasesLookUpResponseWithDefaults instantiates a new LeasesLookUpResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeasesLookUpResponseWithDefaults() *LeasesLookUpResponse {
	var this LeasesLookUpResponse

	return &this
}

func (o LeasesLookUpResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["keys"] = o.Keys

	return json.Marshal(toSerialize)
}
