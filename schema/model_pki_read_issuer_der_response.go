// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiReadIssuerDerResponse struct for PkiReadIssuerDerResponse
type PkiReadIssuerDerResponse struct {
	// CA Chain
	CaChain []string `json:"ca_chain,omitempty"`

	// Certificate
	Certificate string `json:"certificate,omitempty"`

	// Issuer Id
	IssuerId string `json:"issuer_id,omitempty"`

	// Issuer Name
	IssuerName string `json:"issuer_name,omitempty"`
}
