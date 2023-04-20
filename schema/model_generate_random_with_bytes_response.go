// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GenerateRandomWithBytesResponse struct for GenerateRandomWithBytesResponse
type GenerateRandomWithBytesResponse struct {
	RandomBytes string `json:"random_bytes"`
}

// NewGenerateRandomWithBytesResponseWithDefaults instantiates a new GenerateRandomWithBytesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateRandomWithBytesResponseWithDefaults() *GenerateRandomWithBytesResponse {
	var this GenerateRandomWithBytesResponse

	return &this
}

func (o GenerateRandomWithBytesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["random_bytes"] = o.RandomBytes

	return json.Marshal(toSerialize)
}
