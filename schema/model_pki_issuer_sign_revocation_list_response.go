// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiIssuerSignRevocationListResponse struct for PkiIssuerSignRevocationListResponse
type PkiIssuerSignRevocationListResponse struct {
	// CRL
	Crl string `json:"crl"`
}

// NewPkiIssuerSignRevocationListResponseWithDefaults instantiates a new PkiIssuerSignRevocationListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerSignRevocationListResponseWithDefaults() *PkiIssuerSignRevocationListResponse {
	var this PkiIssuerSignRevocationListResponse

	return &this
}

func (o PkiIssuerSignRevocationListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["crl"] = o.Crl

	return json.Marshal(toSerialize)
}
