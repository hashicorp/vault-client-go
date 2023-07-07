// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RawReadResponse struct for RawReadResponse
type RawReadResponse struct {
	Value string `json:"value,omitempty"`
}

// NewRawReadResponseWithDefaults instantiates a new RawReadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRawReadResponseWithDefaults() *RawReadResponse {
	var this RawReadResponse

	return &this
}
