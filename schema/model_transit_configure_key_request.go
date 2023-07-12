// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitConfigureKeyRequest struct for TransitConfigureKeyRequest
type TransitConfigureKeyRequest struct {
	// Enables taking a backup of the named key in plaintext format. Once set, this cannot be disabled.
	AllowPlaintextBackup bool `json:"allow_plaintext_backup,omitempty"`

	// Amount of time the key should live before being automatically rotated. A value of 0 disables automatic rotation for the key.
	AutoRotatePeriod string `json:"auto_rotate_period,omitempty"`

	// Whether to allow deletion of the key
	DeletionAllowed bool `json:"deletion_allowed,omitempty"`

	// Enables export of the key. Once set, this cannot be disabled.
	Exportable bool `json:"exportable,omitempty"`

	// If set, the minimum version of the key allowed to be decrypted. For signing keys, the minimum version allowed to be used for verification.
	MinDecryptionVersion int32 `json:"min_decryption_version,omitempty"`

	// If set, the minimum version of the key allowed to be used for encryption; or for signing keys, to be used for signing. If set to zero, only the latest version of the key is allowed.
	MinEncryptionVersion int32 `json:"min_encryption_version,omitempty"`
}
