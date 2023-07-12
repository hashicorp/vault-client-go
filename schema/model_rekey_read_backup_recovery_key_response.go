// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RekeyReadBackupRecoveryKeyResponse struct for RekeyReadBackupRecoveryKeyResponse
type RekeyReadBackupRecoveryKeyResponse struct {
	Keys map[string]interface{} `json:"keys,omitempty"`

	KeysBase64 map[string]interface{} `json:"keys_base64,omitempty"`

	Nonce string `json:"nonce,omitempty"`
}
