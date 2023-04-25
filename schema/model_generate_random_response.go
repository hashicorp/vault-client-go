// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GenerateRandomResponse struct for GenerateRandomResponse
type GenerateRandomResponse struct {
	RandomBytes string `json:"random_bytes"`
}

// NewGenerateRandomResponseWithDefaults instantiates a new GenerateRandomResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateRandomResponseWithDefaults() *GenerateRandomResponse {
	var this GenerateRandomResponse

	return &this
}

func (o GenerateRandomResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["random_bytes"] = o.RandomBytes

	return json.Marshal(toSerialize)
}
