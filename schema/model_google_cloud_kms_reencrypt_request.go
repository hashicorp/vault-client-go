// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudKmsReencryptRequest struct for GoogleCloudKmsReencryptRequest
type GoogleCloudKmsReencryptRequest struct {
	// Optional data that, if specified, must also be provided during decryption.
	AdditionalAuthenticatedData string `json:"additional_authenticated_data,omitempty"`

	// Ciphertext to be re-encrypted to the latest key version. This must be ciphertext that Vault previously generated for this named key.
	Ciphertext string `json:"ciphertext,omitempty"`

	// Integer version of the crypto key version to use for the new encryption. If unspecified, this defaults to the latest active crypto key version.
	KeyVersion int32 `json:"key_version,omitempty"`
}
