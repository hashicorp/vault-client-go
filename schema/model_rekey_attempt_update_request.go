// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RekeyAttemptUpdateRequest struct for RekeyAttemptUpdateRequest
type RekeyAttemptUpdateRequest struct {
	// Specifies a single unseal key share.
	Key string `json:"key"`

	// Specifies the nonce of the rekey attempt.
	Nonce string `json:"nonce"`
}

// NewRekeyAttemptUpdateRequestWithDefaults instantiates a new RekeyAttemptUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRekeyAttemptUpdateRequestWithDefaults() *RekeyAttemptUpdateRequest {
	var this RekeyAttemptUpdateRequest

	return &this
}

func (o RekeyAttemptUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["key"] = o.Key
	toSerialize["nonce"] = o.Nonce

	return json.Marshal(toSerialize)
}
