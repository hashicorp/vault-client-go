// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GenerateHashResponse struct for GenerateHashResponse
type GenerateHashResponse struct {
	Sum string `json:"sum"`
}

// NewGenerateHashResponseWithDefaults instantiates a new GenerateHashResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateHashResponseWithDefaults() *GenerateHashResponse {
	var this GenerateHashResponse

	return &this
}

func (o GenerateHashResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["sum"] = o.Sum

	return json.Marshal(toSerialize)
}
