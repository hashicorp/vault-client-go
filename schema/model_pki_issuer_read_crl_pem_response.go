// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiIssuerReadCrlPemResponse struct for PkiIssuerReadCrlPemResponse
type PkiIssuerReadCrlPemResponse struct {
	Crl string `json:"crl,omitempty"`
}

// NewPkiIssuerReadCrlPemResponseWithDefaults instantiates a new PkiIssuerReadCrlPemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerReadCrlPemResponseWithDefaults() *PkiIssuerReadCrlPemResponse {
	var this PkiIssuerReadCrlPemResponse

	return &this
}
