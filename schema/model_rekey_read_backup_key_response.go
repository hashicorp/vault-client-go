// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RekeyReadBackupKeyResponse struct for RekeyReadBackupKeyResponse
type RekeyReadBackupKeyResponse struct {
	Keys map[string]interface{} `json:"keys"`

	KeysBase64 map[string]interface{} `json:"keys_base64"`

	Nonce string `json:"nonce"`
}

// NewRekeyReadBackupKeyResponseWithDefaults instantiates a new RekeyReadBackupKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRekeyReadBackupKeyResponseWithDefaults() *RekeyReadBackupKeyResponse {
	var this RekeyReadBackupKeyResponse

	return &this
}

func (o RekeyReadBackupKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["keys"] = o.Keys
	toSerialize["keys_base64"] = o.KeysBase64
	toSerialize["nonce"] = o.Nonce

	return json.Marshal(toSerialize)
}
