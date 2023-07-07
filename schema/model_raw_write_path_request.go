// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RawWritePathRequest struct for RawWritePathRequest
type RawWritePathRequest struct {
	Compressed bool `json:"compressed,omitempty"`

	CompressionType string `json:"compression_type,omitempty"`

	Encoding string `json:"encoding,omitempty"`

	Value string `json:"value,omitempty"`
}

// NewRawWritePathRequestWithDefaults instantiates a new RawWritePathRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRawWritePathRequestWithDefaults() *RawWritePathRequest {
	var this RawWritePathRequest

	return &this
}
