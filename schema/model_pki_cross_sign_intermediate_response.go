// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiCrossSignIntermediateResponse struct for PkiCrossSignIntermediateResponse
type PkiCrossSignIntermediateResponse struct {
	// Certificate signing request.
	Csr string `json:"csr,omitempty"`

	// Id of the key.
	KeyId string `json:"key_id,omitempty"`

	// Generated private key.
	PrivateKey string `json:"private_key,omitempty"`

	// Specifies the format used for marshaling the private key.
	PrivateKeyType string `json:"private_key_type,omitempty"`
}

// NewPkiCrossSignIntermediateResponseWithDefaults instantiates a new PkiCrossSignIntermediateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiCrossSignIntermediateResponseWithDefaults() *PkiCrossSignIntermediateResponse {
	var this PkiCrossSignIntermediateResponse

	return &this
}
