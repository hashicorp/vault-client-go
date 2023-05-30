// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitHashWithAlgorithmRequest struct for TransitHashWithAlgorithmRequest
type TransitHashWithAlgorithmRequest struct {
	// Algorithm to use (POST body parameter). Valid values are: * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 Defaults to \"sha2-256\".
	Algorithm string `json:"algorithm,omitempty"`

	// Encoding format to use. Can be \"hex\" or \"base64\". Defaults to \"hex\".
	Format string `json:"format,omitempty"`

	// The base64-encoded input data
	Input string `json:"input,omitempty"`
}

// NewTransitHashWithAlgorithmRequestWithDefaults instantiates a new TransitHashWithAlgorithmRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitHashWithAlgorithmRequestWithDefaults() *TransitHashWithAlgorithmRequest {
	var this TransitHashWithAlgorithmRequest

	this.Algorithm = "sha2-256"
	this.Format = "hex"

	return &this
}