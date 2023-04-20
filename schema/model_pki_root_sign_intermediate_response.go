// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiRootSignIntermediateResponse struct for PkiRootSignIntermediateResponse
type PkiRootSignIntermediateResponse struct {
	// CA Chain
	CaChain []string `json:"ca_chain"`

	// Certificate
	Certificate string `json:"certificate"`

	// Expiration Time
	Expiration int64 `json:"expiration"`

	// Issuing CA
	IssuingCa string `json:"issuing_ca"`

	// Serial Number
	SerialNumber string `json:"serial_number"`
}

// NewPkiRootSignIntermediateResponseWithDefaults instantiates a new PkiRootSignIntermediateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiRootSignIntermediateResponseWithDefaults() *PkiRootSignIntermediateResponse {
	var this PkiRootSignIntermediateResponse

	return &this
}

func (o PkiRootSignIntermediateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["ca_chain"] = o.CaChain
	toSerialize["certificate"] = o.Certificate
	toSerialize["expiration"] = o.Expiration
	toSerialize["issuing_ca"] = o.IssuingCa
	toSerialize["serial_number"] = o.SerialNumber

	return json.Marshal(toSerialize)
}
