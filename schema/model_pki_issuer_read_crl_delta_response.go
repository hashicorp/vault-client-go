// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiIssuerReadCrlDeltaResponse struct for PkiIssuerReadCrlDeltaResponse
type PkiIssuerReadCrlDeltaResponse struct {
	Crl string `json:"crl"`
}

// NewPkiIssuerReadCrlDeltaResponseWithDefaults instantiates a new PkiIssuerReadCrlDeltaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerReadCrlDeltaResponseWithDefaults() *PkiIssuerReadCrlDeltaResponse {
	var this PkiIssuerReadCrlDeltaResponse

	return &this
}

func (o PkiIssuerReadCrlDeltaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["crl"] = o.Crl

	return json.Marshal(toSerialize)
}
