// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiImportKeyResponse struct for PkiImportKeyResponse
type PkiImportKeyResponse struct {
	// ID assigned to this key.
	KeyId string `json:"key_id"`

	// Name assigned to this key.
	KeyName string `json:"key_name"`

	// The type of key to use; defaults to RSA. \"rsa\" \"ec\" and \"ed25519\" are the only valid values.
	KeyType string `json:"key_type"`
}

// NewPkiImportKeyResponseWithDefaults instantiates a new PkiImportKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiImportKeyResponseWithDefaults() *PkiImportKeyResponse {
	var this PkiImportKeyResponse

	return &this
}

func (o PkiImportKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["key_id"] = o.KeyId
	toSerialize["key_name"] = o.KeyName
	toSerialize["key_type"] = o.KeyType

	return json.Marshal(toSerialize)
}
