// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GenerateHashWithAlgorithmRequest struct for GenerateHashWithAlgorithmRequest
type GenerateHashWithAlgorithmRequest struct {
	// Algorithm to use (POST body parameter). Valid values are: * sha2-224 * sha2-256 * sha2-384 * sha2-512 Defaults to \"sha2-256\".
	Algorithm string `json:"algorithm"`

	// Encoding format to use. Can be \"hex\" or \"base64\". Defaults to \"hex\".
	Format string `json:"format"`

	// The base64-encoded input data
	Input string `json:"input"`
}

// NewGenerateHashWithAlgorithmRequestWithDefaults instantiates a new GenerateHashWithAlgorithmRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateHashWithAlgorithmRequestWithDefaults() *GenerateHashWithAlgorithmRequest {
	var this GenerateHashWithAlgorithmRequest

	this.Algorithm = "sha2-256"
	this.Format = "hex"

	return &this
}

func (o GenerateHashWithAlgorithmRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["algorithm"] = o.Algorithm
	toSerialize["format"] = o.Format
	toSerialize["input"] = o.Input

	return json.Marshal(toSerialize)
}
