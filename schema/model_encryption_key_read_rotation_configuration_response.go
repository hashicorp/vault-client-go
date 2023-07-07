// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// EncryptionKeyReadRotationConfigurationResponse struct for EncryptionKeyReadRotationConfigurationResponse
type EncryptionKeyReadRotationConfigurationResponse struct {
	Enabled bool `json:"enabled,omitempty"`

	Interval string `json:"interval,omitempty"`

	MaxOperations int64 `json:"max_operations,omitempty"`
}

// NewEncryptionKeyReadRotationConfigurationResponseWithDefaults instantiates a new EncryptionKeyReadRotationConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionKeyReadRotationConfigurationResponseWithDefaults() *EncryptionKeyReadRotationConfigurationResponse {
	var this EncryptionKeyReadRotationConfigurationResponse

	return &this
}
