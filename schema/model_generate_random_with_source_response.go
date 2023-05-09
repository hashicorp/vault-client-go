// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GenerateRandomWithSourceResponse struct for GenerateRandomWithSourceResponse
type GenerateRandomWithSourceResponse struct {
	RandomBytes string `json:"random_bytes,omitempty"`
}

// NewGenerateRandomWithSourceResponseWithDefaults instantiates a new GenerateRandomWithSourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateRandomWithSourceResponseWithDefaults() *GenerateRandomWithSourceResponse {
	var this GenerateRandomWithSourceResponse

	return &this
}
