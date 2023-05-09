// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiSignVerbatimResponse struct for PkiSignVerbatimResponse
type PkiSignVerbatimResponse struct {
	// Certificate Chain
	CaChain []string `json:"ca_chain,omitempty"`

	// Certificate
	Certificate string `json:"certificate,omitempty"`

	// Time of expiration
	Expiration string `json:"expiration,omitempty"`

	// Issuing Certificate Authority
	IssuingCa string `json:"issuing_ca,omitempty"`

	// Private key
	PrivateKey string `json:"private_key,omitempty"`

	// Private key type
	PrivateKeyType string `json:"private_key_type,omitempty"`

	// Serial Number
	SerialNumber string `json:"serial_number,omitempty"`
}

// NewPkiSignVerbatimResponseWithDefaults instantiates a new PkiSignVerbatimResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiSignVerbatimResponseWithDefaults() *PkiSignVerbatimResponse {
	var this PkiSignVerbatimResponse

	return &this
}
