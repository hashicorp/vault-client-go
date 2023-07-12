// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudKmsDecryptRequest struct for GoogleCloudKmsDecryptRequest
type GoogleCloudKmsDecryptRequest struct {
	// Optional data that was specified during encryption of this payload.
	AdditionalAuthenticatedData string `json:"additional_authenticated_data,omitempty"`

	// Ciphertext to decrypt as previously returned from an encrypt operation. This must be base64-encoded ciphertext as previously returned from an encrypt operation.
	Ciphertext string `json:"ciphertext,omitempty"`

	// Integer version of the crypto key version to use for decryption. This is required for asymmetric keys. For symmetric keys, Cloud KMS will choose the correct version automatically.
	KeyVersion int32 `json:"key_version,omitempty"`
}
