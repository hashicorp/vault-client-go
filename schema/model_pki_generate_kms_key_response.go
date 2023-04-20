// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiGenerateKmsKeyResponse struct for PkiGenerateKmsKeyResponse
type PkiGenerateKmsKeyResponse struct {
	// ID assigned to this key.
	KeyId string `json:"key_id"`

	// Name assigned to this key.
	KeyName string `json:"key_name"`

	// The type of key to use; defaults to RSA. \"rsa\" \"ec\" and \"ed25519\" are the only valid values.
	KeyType string `json:"key_type"`

	// The private key string
	PrivateKey string `json:"private_key"`
}

// NewPkiGenerateKmsKeyResponseWithDefaults instantiates a new PkiGenerateKmsKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiGenerateKmsKeyResponseWithDefaults() *PkiGenerateKmsKeyResponse {
	var this PkiGenerateKmsKeyResponse

	return &this
}

func (o PkiGenerateKmsKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["key_id"] = o.KeyId
	toSerialize["key_name"] = o.KeyName
	toSerialize["key_type"] = o.KeyType
	toSerialize["private_key"] = o.PrivateKey

	return json.Marshal(toSerialize)
}
