// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RawWriteRequest struct for RawWriteRequest
type RawWriteRequest struct {
	Compressed bool `json:"compressed,omitempty"`

	CompressionType string `json:"compression_type,omitempty"`

	Encoding string `json:"encoding,omitempty"`

	Path string `json:"path,omitempty"`

	Value string `json:"value,omitempty"`
}

// NewRawWriteRequestWithDefaults instantiates a new RawWriteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRawWriteRequestWithDefaults() *RawWriteRequest {
	var this RawWriteRequest

	return &this
}
