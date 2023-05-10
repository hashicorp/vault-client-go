// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiIssuerReadCrlDeltaDerResponse struct for PkiIssuerReadCrlDeltaDerResponse
type PkiIssuerReadCrlDeltaDerResponse struct {
	Crl string `json:"crl,omitempty"`
}

// NewPkiIssuerReadCrlDeltaDerResponseWithDefaults instantiates a new PkiIssuerReadCrlDeltaDerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerReadCrlDeltaDerResponseWithDefaults() *PkiIssuerReadCrlDeltaDerResponse {
	var this PkiIssuerReadCrlDeltaDerResponse

	return &this
}
