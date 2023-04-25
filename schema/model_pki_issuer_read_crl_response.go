// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiIssuerReadCrlResponse struct for PkiIssuerReadCrlResponse
type PkiIssuerReadCrlResponse struct {
	Crl string `json:"crl"`
}

// NewPkiIssuerReadCrlResponseWithDefaults instantiates a new PkiIssuerReadCrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerReadCrlResponseWithDefaults() *PkiIssuerReadCrlResponse {
	var this PkiIssuerReadCrlResponse

	return &this
}

func (o PkiIssuerReadCrlResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["crl"] = o.Crl

	return json.Marshal(toSerialize)
}
