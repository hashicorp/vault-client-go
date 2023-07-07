// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import "time"

// PkiGenerateEabKeyForRoleResponse struct for PkiGenerateEabKeyForRoleResponse
type PkiGenerateEabKeyForRoleResponse struct {
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

// NewPkiGenerateEabKeyForRoleResponseWithDefaults instantiates a new PkiGenerateEabKeyForRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiGenerateEabKeyForRoleResponseWithDefaults() *PkiGenerateEabKeyForRoleResponse {
	var this PkiGenerateEabKeyForRoleResponse

	return &this
}
