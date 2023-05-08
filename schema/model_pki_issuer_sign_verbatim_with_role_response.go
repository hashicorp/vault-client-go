// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiIssuerSignVerbatimWithRoleResponse struct for PkiIssuerSignVerbatimWithRoleResponse
type PkiIssuerSignVerbatimWithRoleResponse struct {
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

// NewPkiIssuerSignVerbatimWithRoleResponseWithDefaults instantiates a new PkiIssuerSignVerbatimWithRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiIssuerSignVerbatimWithRoleResponseWithDefaults() *PkiIssuerSignVerbatimWithRoleResponse {
	var this PkiIssuerSignVerbatimWithRoleResponse

	return &this
}
