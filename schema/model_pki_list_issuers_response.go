// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiListIssuersResponse struct for PkiListIssuersResponse
type PkiListIssuersResponse struct {
	// Key info with issuer name
	KeyInfo map[string]interface{} `json:"key_info,omitempty"`

	// A list of keys
	Keys []string `json:"keys,omitempty"`
}

// NewPkiListIssuersResponseWithDefaults instantiates a new PkiListIssuersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiListIssuersResponseWithDefaults() *PkiListIssuersResponse {
	var this PkiListIssuersResponse

	return &this
}
