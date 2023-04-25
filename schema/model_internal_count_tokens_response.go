// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// InternalCountTokensResponse struct for InternalCountTokensResponse
type InternalCountTokensResponse struct {
	Counters map[string]interface{} `json:"counters"`
}

// NewInternalCountTokensResponseWithDefaults instantiates a new InternalCountTokensResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalCountTokensResponseWithDefaults() *InternalCountTokensResponse {
	var this InternalCountTokensResponse

	return &this
}

func (o InternalCountTokensResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["counters"] = o.Counters

	return json.Marshal(toSerialize)
}
