// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitGenerateDataKeyRequest struct for TransitGenerateDataKeyRequest
type TransitGenerateDataKeyRequest struct {
	// Number of bits for the key; currently 128, 256, and 512 bits are supported. Defaults to 256.
	Bits int32 `json:"bits,omitempty"`

	// Context for key derivation. Required for derived keys.
	Context string `json:"context,omitempty"`

	// The version of the Vault key to use for encryption of the data key. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key.
	KeyVersion int32 `json:"key_version,omitempty"`

	// Nonce for when convergent encryption v1 is used (only in Vault 0.6.1)
	Nonce string `json:"nonce,omitempty"`
}
