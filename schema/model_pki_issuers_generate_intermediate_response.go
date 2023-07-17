// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiIssuersGenerateIntermediateResponse struct for PkiIssuersGenerateIntermediateResponse
type PkiIssuersGenerateIntermediateResponse struct {
	// Certificate signing request.
	Csr string `json:"csr,omitempty"`

	// Id of the key.
	KeyId string `json:"key_id,omitempty"`

	// Generated private key.
	PrivateKey string `json:"private_key,omitempty"`

	// Specifies the format used for marshaling the private key.
	PrivateKeyType string `json:"private_key_type,omitempty"`
}
