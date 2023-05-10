// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiReadKeyResponse struct for PkiReadKeyResponse
type PkiReadKeyResponse struct {
	// Key Id
	KeyId string `json:"key_id,omitempty"`

	// Key Name
	KeyName string `json:"key_name,omitempty"`

	// Key Type
	KeyType string `json:"key_type,omitempty"`

	// Managed Key Id
	ManagedKeyId string `json:"managed_key_id,omitempty"`

	// Managed Key Name
	ManagedKeyName string `json:"managed_key_name,omitempty"`
}

// NewPkiReadKeyResponseWithDefaults instantiates a new PkiReadKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiReadKeyResponseWithDefaults() *PkiReadKeyResponse {
	var this PkiReadKeyResponse

	return &this
}
