// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiImportKeyRequest struct for PkiImportKeyRequest
type PkiImportKeyRequest struct {
	// Optional name to be used for this key
	KeyName string `json:"key_name"`

	// PEM-format, unencrypted secret key
	PemBundle string `json:"pem_bundle"`
}

// NewPkiImportKeyRequestWithDefaults instantiates a new PkiImportKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiImportKeyRequestWithDefaults() *PkiImportKeyRequest {
	var this PkiImportKeyRequest

	return &this
}

func (o PkiImportKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["key_name"] = o.KeyName
	toSerialize["pem_bundle"] = o.PemBundle

	return json.Marshal(toSerialize)
}