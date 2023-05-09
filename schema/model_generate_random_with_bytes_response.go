// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GenerateRandomWithBytesResponse struct for GenerateRandomWithBytesResponse
type GenerateRandomWithBytesResponse struct {
	RandomBytes string `json:"random_bytes,omitempty"`
}

// NewGenerateRandomWithBytesResponseWithDefaults instantiates a new GenerateRandomWithBytesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateRandomWithBytesResponseWithDefaults() *GenerateRandomWithBytesResponse {
	var this GenerateRandomWithBytesResponse

	return &this
}
