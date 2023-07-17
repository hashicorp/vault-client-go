// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiRootSignIntermediateResponse struct for PkiRootSignIntermediateResponse
type PkiRootSignIntermediateResponse struct {
	// CA Chain
	CaChain []string `json:"ca_chain,omitempty"`

	// Certificate
	Certificate string `json:"certificate,omitempty"`

	// Expiration Time
	Expiration int64 `json:"expiration,omitempty"`

	// Issuing CA
	IssuingCa string `json:"issuing_ca,omitempty"`

	// Serial Number
	SerialNumber string `json:"serial_number,omitempty"`
}
