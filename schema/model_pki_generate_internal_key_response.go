// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiGenerateInternalKeyResponse struct for PkiGenerateInternalKeyResponse
type PkiGenerateInternalKeyResponse struct {
	// ID assigned to this key.
	KeyId string `json:"key_id"`

	// Name assigned to this key.
	KeyName string `json:"key_name"`

	// The type of key to use; defaults to RSA. \"rsa\" \"ec\" and \"ed25519\" are the only valid values.
	KeyType string `json:"key_type"`

	// The private key string
	PrivateKey string `json:"private_key"`
}

// NewPkiGenerateInternalKeyResponseWithDefaults instantiates a new PkiGenerateInternalKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiGenerateInternalKeyResponseWithDefaults() *PkiGenerateInternalKeyResponse {
	var this PkiGenerateInternalKeyResponse

	return &this
}

func (o PkiGenerateInternalKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["key_id"] = o.KeyId
	toSerialize["key_name"] = o.KeyName
	toSerialize["key_type"] = o.KeyType
	toSerialize["private_key"] = o.PrivateKey

	return json.Marshal(toSerialize)
}
