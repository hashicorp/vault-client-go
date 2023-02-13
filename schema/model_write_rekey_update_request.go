// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteRekeyUpdateRequest struct for WriteRekeyUpdateRequest
type WriteRekeyUpdateRequest struct {
	// Specifies a single unseal key share.
	Key string `json:"key"`

	// Specifies the nonce of the rekey attempt.
	Nonce string `json:"nonce"`
}

// NewWriteRekeyUpdateRequestWithDefaults instantiates a new WriteRekeyUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteRekeyUpdateRequestWithDefaults() *WriteRekeyUpdateRequest {
	var this WriteRekeyUpdateRequest

	return &this
}

func (o WriteRekeyUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["key"] = o.Key
	toSerialize["nonce"] = o.Nonce

	return json.Marshal(toSerialize)
}
