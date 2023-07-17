// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudKmsSignRequest struct for GoogleCloudKmsSignRequest
type GoogleCloudKmsSignRequest struct {
	// Digest to sign. This digest must use the same SHA algorithm as the underlying Cloud KMS key. The digest must be the base64-encoded binary value. This field is required.
	Digest string `json:"digest,omitempty"`

	// Integer version of the crypto key version to use for signing. This field is required.
	KeyVersion int32 `json:"key_version,omitempty"`
}
