// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiSignWithRoleResponse struct for PkiSignWithRoleResponse
type PkiSignWithRoleResponse struct {
	// Certificate Chain
	CaChain []string `json:"ca_chain,omitempty"`

	// Certificate
	Certificate string `json:"certificate,omitempty"`

	// Time of expiration
	Expiration int64 `json:"expiration,omitempty"`

	// Issuing Certificate Authority
	IssuingCa string `json:"issuing_ca,omitempty"`

	// Serial Number
	SerialNumber string `json:"serial_number,omitempty"`
}
