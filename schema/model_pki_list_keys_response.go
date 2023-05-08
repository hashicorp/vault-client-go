// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiListKeysResponse struct for PkiListKeysResponse
type PkiListKeysResponse struct {
	// Key info with issuer name
	KeyInfo map[string]interface{} `json:"key_info,omitempty"`

	// A list of keys
	Keys []string `json:"keys,omitempty"`
}

// NewPkiListKeysResponseWithDefaults instantiates a new PkiListKeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiListKeysResponseWithDefaults() *PkiListKeysResponse {
	var this PkiListKeysResponse

	return &this
}
