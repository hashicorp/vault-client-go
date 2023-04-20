// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GenerateHashWithAlgorithmResponse struct for GenerateHashWithAlgorithmResponse
type GenerateHashWithAlgorithmResponse struct {
	Sum string `json:"sum"`
}

// NewGenerateHashWithAlgorithmResponseWithDefaults instantiates a new GenerateHashWithAlgorithmResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateHashWithAlgorithmResponseWithDefaults() *GenerateHashWithAlgorithmResponse {
	var this GenerateHashWithAlgorithmResponse

	return &this
}

func (o GenerateHashWithAlgorithmResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["sum"] = o.Sum

	return json.Marshal(toSerialize)
}
