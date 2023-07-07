// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import "time"

// PkiGenerateEabKeyForIssuerAndRoleResponse struct for PkiGenerateEabKeyForIssuerAndRoleResponse
type PkiGenerateEabKeyForIssuerAndRoleResponse struct {
	// The ACME directory to which the key belongs
	AcmeDirectory string `json:"acme_directory,omitempty"`

	// An RFC3339 formatted date time when the EAB token was created
	CreatedOn time.Time `json:"created_on,omitempty"`

	// The EAB key identifier
	Id string `json:"id,omitempty"`

	// The EAB hmac key
	Key string `json:"key,omitempty"`

	// The EAB key type
	KeyType string `json:"key_type,omitempty"`
}

// NewPkiGenerateEabKeyForIssuerAndRoleResponseWithDefaults instantiates a new PkiGenerateEabKeyForIssuerAndRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiGenerateEabKeyForIssuerAndRoleResponseWithDefaults() *PkiGenerateEabKeyForIssuerAndRoleResponse {
	var this PkiGenerateEabKeyForIssuerAndRoleResponse

	return &this
}
