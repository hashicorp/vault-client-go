// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiRotateDeltaCrlResponse struct for PkiRotateDeltaCrlResponse
type PkiRotateDeltaCrlResponse struct {
	// Whether rotation was successful
	Success bool `json:"success"`
}

// NewPkiRotateDeltaCrlResponseWithDefaults instantiates a new PkiRotateDeltaCrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiRotateDeltaCrlResponseWithDefaults() *PkiRotateDeltaCrlResponse {
	var this PkiRotateDeltaCrlResponse

	return &this
}

func (o PkiRotateDeltaCrlResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["success"] = o.Success

	return json.Marshal(toSerialize)
}
