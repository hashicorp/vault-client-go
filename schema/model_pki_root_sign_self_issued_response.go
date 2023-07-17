// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiRootSignSelfIssuedResponse struct for PkiRootSignSelfIssuedResponse
type PkiRootSignSelfIssuedResponse struct {
	// Certificate
	Certificate string `json:"certificate,omitempty"`

	// Issuing CA
	IssuingCa string `json:"issuing_ca,omitempty"`
}
