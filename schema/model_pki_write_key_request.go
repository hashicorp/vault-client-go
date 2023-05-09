// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiWriteKeyRequest struct for PkiWriteKeyRequest
type PkiWriteKeyRequest struct {
	// Human-readable name for this key.
	KeyName string `json:"key_name,omitempty"`
}

// NewPkiWriteKeyRequestWithDefaults instantiates a new PkiWriteKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiWriteKeyRequestWithDefaults() *PkiWriteKeyRequest {
	var this PkiWriteKeyRequest

	return &this
}
