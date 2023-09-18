// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitRewrapRequest struct for TransitRewrapRequest
type TransitRewrapRequest struct {
	// Specifies a list of items to be re-encrypted in a single batch. When this parameter is set, if the parameters 'ciphertext', 'context' and 'nonce' are also set, they will be ignored. Any batch output will preserve the order of the batch input.
	BatchInput []map[string]interface{} `json:"batch_input,omitempty"`

	// Ciphertext value to rewrap
	Ciphertext string `json:"ciphertext,omitempty"`

	// Base64 encoded context for key derivation. Required for derived keys.
	Context string `json:"context,omitempty"`

	// The version of the key to use for encryption. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key.
	KeyVersion int32 `json:"key_version,omitempty"`

	// Nonce for when convergent encryption is used
	Nonce string `json:"nonce,omitempty"`
}
