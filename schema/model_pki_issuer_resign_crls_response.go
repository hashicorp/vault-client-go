// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiIssuerResignCrlsResponse struct for PkiIssuerResignCrlsResponse
type PkiIssuerResignCrlsResponse struct {
	// CRL
	Crl string `json:"crl,omitempty"`
}

// NewPkiIssuerResignCrlsResponseWithDefaults instantiates a new PkiIssuerResignCrlsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerResignCrlsResponseWithDefaults() *PkiIssuerResignCrlsResponse {
	var this PkiIssuerResignCrlsResponse

	return &this
}
