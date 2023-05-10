// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiIssuerReadCrlDerResponse struct for PkiIssuerReadCrlDerResponse
type PkiIssuerReadCrlDerResponse struct {
	Crl string `json:"crl,omitempty"`
}

// NewPkiIssuerReadCrlDerResponseWithDefaults instantiates a new PkiIssuerReadCrlDerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerReadCrlDerResponseWithDefaults() *PkiIssuerReadCrlDerResponse {
	var this PkiIssuerReadCrlDerResponse

	return &this
}
