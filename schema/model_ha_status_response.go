// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// HaStatusResponse struct for HaStatusResponse
type HaStatusResponse struct {
	Nodes []map[string]interface{} `json:"nodes"`
}

// NewHaStatusResponseWithDefaults instantiates a new HaStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHaStatusResponseWithDefaults() *HaStatusResponse {
	var this HaStatusResponse

	return &this
}

func (o HaStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["nodes"] = o.Nodes

	return json.Marshal(toSerialize)
}
