// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GenerateRandomWithSourceRequest struct for GenerateRandomWithSourceRequest
type GenerateRandomWithSourceRequest struct {
	// The number of bytes to generate (POST body parameter). Defaults to 32 (256 bits).
	Bytes int32 `json:"bytes,omitempty"`

	// Encoding format to use. Can be \"hex\" or \"base64\". Defaults to \"base64\".
	Format string `json:"format,omitempty"`

	// The number of bytes to generate (POST URL parameter)
	Urlbytes string `json:"urlbytes,omitempty"`
}

// NewGenerateRandomWithSourceRequestWithDefaults instantiates a new GenerateRandomWithSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateRandomWithSourceRequestWithDefaults() *GenerateRandomWithSourceRequest {
	var this GenerateRandomWithSourceRequest

	this.Bytes = 32
	this.Format = "base64"

	return &this
}
