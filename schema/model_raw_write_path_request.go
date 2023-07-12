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
