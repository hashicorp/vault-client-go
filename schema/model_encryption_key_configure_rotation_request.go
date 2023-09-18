// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// EncryptionKeyConfigureRotationRequest struct for EncryptionKeyConfigureRotationRequest
type EncryptionKeyConfigureRotationRequest struct {
	// Whether automatic rotation is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// How long after installation of an active key term that the key will be automatically rotated.
	Interval string `json:"interval,omitempty"`

	// The number of encryption operations performed before the barrier key is automatically rotated.
	MaxOperations int64 `json:"max_operations,omitempty"`
}
