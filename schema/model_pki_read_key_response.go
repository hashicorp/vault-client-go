// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiReadKeyResponse struct for PkiReadKeyResponse
type PkiReadKeyResponse struct {
	// Key Id
	KeyId string `json:"key_id"`

	// Key Name
	KeyName string `json:"key_name"`

	// Key Type
	KeyType string `json:"key_type"`

	// Managed Key Id
	ManagedKeyId string `json:"managed_key_id"`

	// Managed Key Name
	ManagedKeyName string `json:"managed_key_name"`
}

// NewPkiReadKeyResponseWithDefaults instantiates a new PkiReadKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiReadKeyResponseWithDefaults() *PkiReadKeyResponse {
	var this PkiReadKeyResponse

	return &this
}

func (o PkiReadKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["key_id"] = o.KeyId
	toSerialize["key_name"] = o.KeyName
	toSerialize["key_type"] = o.KeyType
	toSerialize["managed_key_id"] = o.ManagedKeyId
	toSerialize["managed_key_name"] = o.ManagedKeyName

	return json.Marshal(toSerialize)
}
