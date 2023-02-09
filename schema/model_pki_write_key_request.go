// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIWriteKeyRequest struct for PKIWriteKeyRequest
type PKIWriteKeyRequest struct {
	// Human-readable name for this key.
	KeyName string `json:"key_name"`
}

// NewPKIWriteKeyRequestWithDefaults instantiates a new PKIWriteKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIWriteKeyRequestWithDefaults() *PKIWriteKeyRequest {
	var this PKIWriteKeyRequest

	return &this
}

func (o PKIWriteKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
