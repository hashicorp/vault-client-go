// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiRotateCrlResponse struct for PkiRotateCrlResponse
type PkiRotateCrlResponse struct {
	// Whether rotation was successful
	Success bool `json:"success"`
}

// NewPkiRotateCrlResponseWithDefaults instantiates a new PkiRotateCrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiRotateCrlResponseWithDefaults() *PkiRotateCrlResponse {
	var this PkiRotateCrlResponse

	return &this
}

func (o PkiRotateCrlResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["success"] = o.Success

	return json.Marshal(toSerialize)
}
