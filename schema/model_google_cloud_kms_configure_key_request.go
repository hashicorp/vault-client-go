// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudKmsConfigureKeyRequest struct for GoogleCloudKmsConfigureKeyRequest
type GoogleCloudKmsConfigureKeyRequest struct {
	// Maximum allowed crypto key version. If set to a positive value, key versions greater than the given value are not permitted to be used. If set to 0 or a negative value, there is no maximum key version.
	MaxVersion int32 `json:"max_version,omitempty"`

	// Minimum allowed crypto key version. If set to a positive value, key versions less than the given value are not permitted to be used. If set to 0 or a negative value, there is no minimum key version. This value only affects encryption/re-encryption, not decryption. To restrict old values from being decrypted, increase this value and then perform a trim operation.
	MinVersion int32 `json:"min_version,omitempty"`
}
