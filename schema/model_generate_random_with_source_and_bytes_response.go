// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GenerateRandomWithSourceAndBytesResponse struct for GenerateRandomWithSourceAndBytesResponse
type GenerateRandomWithSourceAndBytesResponse struct {
	RandomBytes string `json:"random_bytes,omitempty"`
}

// NewGenerateRandomWithSourceAndBytesResponseWithDefaults instantiates a new GenerateRandomWithSourceAndBytesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateRandomWithSourceAndBytesResponseWithDefaults() *GenerateRandomWithSourceAndBytesResponse {
	var this GenerateRandomWithSourceAndBytesResponse

	return &this
}
