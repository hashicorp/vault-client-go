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

// NewRekeyReadBackupRecoveryKeyResponseWithDefaults instantiates a new RekeyReadBackupRecoveryKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRekeyReadBackupRecoveryKeyResponseWithDefaults() *RekeyReadBackupRecoveryKeyResponse {
	var this RekeyReadBackupRecoveryKeyResponse

	return &this
}
