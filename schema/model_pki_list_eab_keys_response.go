// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiListEabKeysResponse struct for PkiListEabKeysResponse
type PkiListEabKeysResponse struct {
	// EAB details keyed by the eab key id
	KeyInfo map[string]interface{} `json:"key_info,omitempty"`

	// A list of unused eab keys
	Keys []string `json:"keys,omitempty"`
}

// NewPkiListEabKeysResponseWithDefaults instantiates a new PkiListEabKeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiListEabKeysResponseWithDefaults() *PkiListEabKeysResponse {
	var this PkiListEabKeysResponse

	return &this
}
