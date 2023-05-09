// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RateLimitQuotasReadResponse struct for RateLimitQuotasReadResponse
type RateLimitQuotasReadResponse struct {
	BlockInterval int32 `json:"block_interval,omitempty"`

	Interval int32 `json:"interval,omitempty"`

	Name string `json:"name,omitempty"`

	Path string `json:"path,omitempty"`

	Rate float32 `json:"rate,omitempty"`

	Role string `json:"role,omitempty"`

	Type string `json:"type,omitempty"`
}

// NewRateLimitQuotasReadResponseWithDefaults instantiates a new RateLimitQuotasReadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitQuotasReadResponseWithDefaults() *RateLimitQuotasReadResponse {
	var this RateLimitQuotasReadResponse

	return &this
}
