// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiReadCertCaChainResponse struct for PkiReadCertCaChainResponse
type PkiReadCertCaChainResponse struct {
	// Issuing CA Chain
	CaChain string `json:"ca_chain,omitempty"`

	// Certificate
	Certificate string `json:"certificate,omitempty"`

	// ID of the issuer
	IssuerId string `json:"issuer_id,omitempty"`

	// Revocation time
	RevocationTime int64 `json:"revocation_time,omitempty"`

	// Revocation time RFC 3339 formatted
	RevocationTimeRfc3339 string `json:"revocation_time_rfc3339,omitempty"`
}

// NewPkiReadCertCaChainResponseWithDefaults instantiates a new PkiReadCertCaChainResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiReadCertCaChainResponseWithDefaults() *PkiReadCertCaChainResponse {
	var this PkiReadCertCaChainResponse

	return &this
}
