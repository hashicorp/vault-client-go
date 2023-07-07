// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiWriteIssuerIssuerRefRolesRoleAcmeNewAccountRequest struct for PkiWriteIssuerIssuerRefRolesRoleAcmeNewAccountRequest
type PkiWriteIssuerIssuerRefRolesRoleAcmeNewAccountRequest struct {
	// ACME request 'payload' value
	Payload string `json:"payload,omitempty"`

	// ACME request 'protected' value
	Protected string `json:"protected,omitempty"`

	// ACME request 'signature' value
	Signature string `json:"signature,omitempty"`
}

// NewPkiWriteIssuerIssuerRefRolesRoleAcmeNewAccountRequestWithDefaults instantiates a new PkiWriteIssuerIssuerRefRolesRoleAcmeNewAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiWriteIssuerIssuerRefRolesRoleAcmeNewAccountRequestWithDefaults() *PkiWriteIssuerIssuerRefRolesRoleAcmeNewAccountRequest {
	var this PkiWriteIssuerIssuerRefRolesRoleAcmeNewAccountRequest

	return &this
}
