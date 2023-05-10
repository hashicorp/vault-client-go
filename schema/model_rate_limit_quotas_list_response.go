// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RateLimitQuotasListResponse struct for RateLimitQuotasListResponse
type RateLimitQuotasListResponse struct {
	Keys []string `json:"keys"`
}

// NewRateLimitQuotasListResponseWithDefaults instantiates a new RateLimitQuotasListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitQuotasListResponseWithDefaults() *RateLimitQuotasListResponse {
	var this RateLimitQuotasListResponse

	return &this
}

func (o RateLimitQuotasListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["keys"] = o.Keys

	return json.Marshal(toSerialize)
}
