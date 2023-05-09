// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitGenerateRandomWithSourceRequest struct for TransitGenerateRandomWithSourceRequest
type TransitGenerateRandomWithSourceRequest struct {
	// The number of bytes to generate (POST body parameter). Defaults to 32 (256 bits).
	Bytes int32 `json:"bytes,omitempty"`

	// Encoding format to use. Can be \"hex\" or \"base64\". Defaults to \"base64\".
	Format string `json:"format,omitempty"`

	// The number of bytes to generate (POST URL parameter)
	Urlbytes string `json:"urlbytes,omitempty"`
}

// NewTransitGenerateRandomWithSourceRequestWithDefaults instantiates a new TransitGenerateRandomWithSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitGenerateRandomWithSourceRequestWithDefaults() *TransitGenerateRandomWithSourceRequest {
	var this TransitGenerateRandomWithSourceRequest

	this.Bytes = 32
	this.Format = "base64"

	return &this
}
