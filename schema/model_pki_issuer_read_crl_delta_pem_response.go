// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiIssuerReadCrlDeltaPemResponse struct for PkiIssuerReadCrlDeltaPemResponse
type PkiIssuerReadCrlDeltaPemResponse struct {
	Crl string `json:"crl"`
}

// NewPkiIssuerReadCrlDeltaPemResponseWithDefaults instantiates a new PkiIssuerReadCrlDeltaPemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerReadCrlDeltaPemResponseWithDefaults() *PkiIssuerReadCrlDeltaPemResponse {
	var this PkiIssuerReadCrlDeltaPemResponse

	return &this
}

func (o PkiIssuerReadCrlDeltaPemResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["crl"] = o.Crl

	return json.Marshal(toSerialize)
}
