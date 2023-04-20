// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiIssuersGenerateIntermediateResponse struct for PkiIssuersGenerateIntermediateResponse
type PkiIssuersGenerateIntermediateResponse struct {
	// Certificate signing request.
	Csr string `json:"csr"`

	// Id of the key.
	KeyId string `json:"key_id"`

	// Generated private key.
	PrivateKey string `json:"private_key"`

	// Specifies the format used for marshaling the private key.
	PrivateKeyType string `json:"private_key_type"`
}

// NewPkiIssuersGenerateIntermediateResponseWithDefaults instantiates a new PkiIssuersGenerateIntermediateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuersGenerateIntermediateResponseWithDefaults() *PkiIssuersGenerateIntermediateResponse {
	var this PkiIssuersGenerateIntermediateResponse

	return &this
}

func (o PkiIssuersGenerateIntermediateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["csr"] = o.Csr
	toSerialize["key_id"] = o.KeyId
	toSerialize["private_key"] = o.PrivateKey
	toSerialize["private_key_type"] = o.PrivateKeyType

	return json.Marshal(toSerialize)
}
