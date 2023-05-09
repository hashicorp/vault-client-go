// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiConfigureKeysResponse struct for PkiConfigureKeysResponse
type PkiConfigureKeysResponse struct {
	// Reference (name or identifier) to the default issuer.
	Default string `json:"default,omitempty"`
}

// NewPkiConfigureKeysResponseWithDefaults instantiates a new PkiConfigureKeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiConfigureKeysResponseWithDefaults() *PkiConfigureKeysResponse {
	var this PkiConfigureKeysResponse

	return &this
}
