// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitGenerateHmacWithAlgorithmRequest struct for TransitGenerateHmacWithAlgorithmRequest
type TransitGenerateHmacWithAlgorithmRequest struct {
	// Algorithm to use (POST body parameter). Valid values are: * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 Defaults to \"sha2-256\".
	Algorithm string `json:"algorithm"`

	// Specifies a list of items to be processed in a single batch. When this parameter is set, if the parameter 'input' is also set, it will be ignored. Any batch output will preserve the order of the batch input.
	BatchInput []map[string]interface{} `json:"batch_input"`

	// The base64-encoded input data
	Input string `json:"input"`

	// The version of the key to use for generating the HMAC. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key.
	KeyVersion int32 `json:"key_version"`
}

// NewTransitGenerateHmacWithAlgorithmRequestWithDefaults instantiates a new TransitGenerateHmacWithAlgorithmRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitGenerateHmacWithAlgorithmRequestWithDefaults() *TransitGenerateHmacWithAlgorithmRequest {
	var this TransitGenerateHmacWithAlgorithmRequest

	this.Algorithm = "sha2-256"

	return &this
}

func (o TransitGenerateHmacWithAlgorithmRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["algorithm"] = o.Algorithm
	toSerialize["batch_input"] = o.BatchInput
	toSerialize["input"] = o.Input
	toSerialize["key_version"] = o.KeyVersion

	return json.Marshal(toSerialize)
}
