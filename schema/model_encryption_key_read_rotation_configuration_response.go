// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// EncryptionKeyReadRotationConfigurationResponse struct for EncryptionKeyReadRotationConfigurationResponse
type EncryptionKeyReadRotationConfigurationResponse struct {
	Enabled bool `json:"enabled"`

	Interval int32 `json:"interval"`

	MaxOperations int64 `json:"max_operations"`
}

// NewEncryptionKeyReadRotationConfigurationResponseWithDefaults instantiates a new EncryptionKeyReadRotationConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionKeyReadRotationConfigurationResponseWithDefaults() *EncryptionKeyReadRotationConfigurationResponse {
	var this EncryptionKeyReadRotationConfigurationResponse

	return &this
}

func (o EncryptionKeyReadRotationConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["enabled"] = o.Enabled
	toSerialize["interval"] = o.Interval
	toSerialize["max_operations"] = o.MaxOperations

	return json.Marshal(toSerialize)
}
