// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIWriteKeysConfigRequest struct for PKIWriteKeysConfigRequest
type PKIWriteKeysConfigRequest struct {
	// Reference (name or identifier) of the default key.
	Default string `json:"default"`
}

// NewPKIWriteKeysConfigRequestWithDefaults instantiates a new PKIWriteKeysConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIWriteKeysConfigRequestWithDefaults() *PKIWriteKeysConfigRequest {
	var this PKIWriteKeysConfigRequest

	return &this
}

func (o PKIWriteKeysConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
