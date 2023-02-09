// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitGenerateRandomSourceBytesRequest struct for TransitGenerateRandomSourceBytesRequest
type TransitGenerateRandomSourceBytesRequest struct {
	// The number of bytes to generate (POST body parameter). Defaults to 32 (256 bits).
	Bytes int32 `json:"bytes"`

	// Encoding format to use. Can be \"hex\" or \"base64\". Defaults to \"base64\".
	Format string `json:"format"`
}

// NewTransitGenerateRandomSourceBytesRequestWithDefaults instantiates a new TransitGenerateRandomSourceBytesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitGenerateRandomSourceBytesRequestWithDefaults() *TransitGenerateRandomSourceBytesRequest {
	var this TransitGenerateRandomSourceBytesRequest

	this.Bytes = 32
	this.Format = "base64"

	return &this
}

func (o TransitGenerateRandomSourceBytesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
