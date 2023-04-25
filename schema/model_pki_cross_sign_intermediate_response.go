// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiCrossSignIntermediateResponse struct for PkiCrossSignIntermediateResponse
type PkiCrossSignIntermediateResponse struct {
	// Certificate signing request.
	Csr string `json:"csr"`

	// Id of the key.
	KeyId string `json:"key_id"`

	// Generated private key.
	PrivateKey string `json:"private_key"`

	// Specifies the format used for marshaling the private key.
	PrivateKeyType string `json:"private_key_type"`
}

// NewPkiCrossSignIntermediateResponseWithDefaults instantiates a new PkiCrossSignIntermediateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiCrossSignIntermediateResponseWithDefaults() *PkiCrossSignIntermediateResponse {
	var this PkiCrossSignIntermediateResponse

	return &this
}

func (o PkiCrossSignIntermediateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["csr"] = o.Csr
	toSerialize["key_id"] = o.KeyId
	toSerialize["private_key"] = o.PrivateKey
	toSerialize["private_key_type"] = o.PrivateKeyType

	return json.Marshal(toSerialize)
}
