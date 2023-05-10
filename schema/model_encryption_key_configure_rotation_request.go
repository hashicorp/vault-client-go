// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// EncryptionKeyConfigureRotationRequest struct for EncryptionKeyConfigureRotationRequest
type EncryptionKeyConfigureRotationRequest struct {
	// Whether automatic rotation is enabled.
	Enabled bool `json:"enabled"`

	// How long after installation of an active key term that the key will be automatically rotated.
	Interval int32 `json:"interval"`

	// The number of encryption operations performed before the barrier key is automatically rotated.
	MaxOperations int64 `json:"max_operations"`
}

// NewEncryptionKeyConfigureRotationRequestWithDefaults instantiates a new EncryptionKeyConfigureRotationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionKeyConfigureRotationRequestWithDefaults() *EncryptionKeyConfigureRotationRequest {
	var this EncryptionKeyConfigureRotationRequest

	return &this
}

func (o EncryptionKeyConfigureRotationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["enabled"] = o.Enabled
	toSerialize["interval"] = o.Interval
	toSerialize["max_operations"] = o.MaxOperations

	return json.Marshal(toSerialize)
}
