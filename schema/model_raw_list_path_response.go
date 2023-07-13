// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RawListPathResponse struct for RawListPathResponse
type RawListPathResponse struct {
	Keys []string `json:"keys,omitempty"`
}

// NewRawListPathResponseWithDefaults instantiates a new RawListPathResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRawListPathResponseWithDefaults() *RawListPathResponse {
	var this RawListPathResponse

	return &this
}
