// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RekeyVerificationUpdateRequest struct for RekeyVerificationUpdateRequest
type RekeyVerificationUpdateRequest struct {
	// Specifies a single unseal share key from the new set of shares.
	Key string `json:"key"`

	// Specifies the nonce of the rekey verification operation.
	Nonce string `json:"nonce"`
}

// NewRekeyVerificationUpdateRequestWithDefaults instantiates a new RekeyVerificationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRekeyVerificationUpdateRequestWithDefaults() *RekeyVerificationUpdateRequest {
	var this RekeyVerificationUpdateRequest

	return &this
}

func (o RekeyVerificationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["key"] = o.Key
	toSerialize["nonce"] = o.Nonce

	return json.Marshal(toSerialize)
}
