// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiWriteRolesRoleAcmeRevokeCertRequest struct for PkiWriteRolesRoleAcmeRevokeCertRequest
type PkiWriteRolesRoleAcmeRevokeCertRequest struct {
	// ACME request 'payload' value
	Payload string `json:"payload,omitempty"`

	// ACME request 'protected' value
	Protected string `json:"protected,omitempty"`

	// ACME request 'signature' value
	Signature string `json:"signature,omitempty"`
}

// NewPkiWriteRolesRoleAcmeRevokeCertRequestWithDefaults instantiates a new PkiWriteRolesRoleAcmeRevokeCertRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiWriteRolesRoleAcmeRevokeCertRequestWithDefaults() *PkiWriteRolesRoleAcmeRevokeCertRequest {
	var this PkiWriteRolesRoleAcmeRevokeCertRequest

	return &this
}
