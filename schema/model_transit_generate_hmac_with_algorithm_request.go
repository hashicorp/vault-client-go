// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitGenerateHMACWithAlgorithmRequest struct for TransitGenerateHMACWithAlgorithmRequest
type TransitGenerateHMACWithAlgorithmRequest struct {
	// Algorithm to use (POST body parameter). Valid values are: * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 Defaults to \"sha2-256\".
	Algorithm string `json:"algorithm"`

	// The base64-encoded input data
	Input string `json:"input"`

	// The version of the key to use for generating the HMAC. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key.
	KeyVersion int32 `json:"key_version"`
}

// NewTransitGenerateHMACWithAlgorithmRequestWithDefaults instantiates a new TransitGenerateHMACWithAlgorithmRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitGenerateHMACWithAlgorithmRequestWithDefaults() *TransitGenerateHMACWithAlgorithmRequest {
	var this TransitGenerateHMACWithAlgorithmRequest

	this.Algorithm = "sha2-256"

	return &this
}

func (o TransitGenerateHMACWithAlgorithmRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
