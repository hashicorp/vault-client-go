// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RateLimitQuotasReadResponse struct for RateLimitQuotasReadResponse
type RateLimitQuotasReadResponse struct {
	BlockInterval int32 `json:"block_interval"`

	Interval int32 `json:"interval"`

	Name string `json:"name"`

	Path string `json:"path"`

	Rate float32 `json:"rate"`

	Role string `json:"role"`

	Type string `json:"type"`
}

// NewRateLimitQuotasReadResponseWithDefaults instantiates a new RateLimitQuotasReadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitQuotasReadResponseWithDefaults() *RateLimitQuotasReadResponse {
	var this RateLimitQuotasReadResponse

	return &this
}

func (o RateLimitQuotasReadResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["block_interval"] = o.BlockInterval
	toSerialize["interval"] = o.Interval
	toSerialize["name"] = o.Name
	toSerialize["path"] = o.Path
	toSerialize["rate"] = o.Rate
	toSerialize["role"] = o.Role
	toSerialize["type"] = o.Type

	return json.Marshal(toSerialize)
}
