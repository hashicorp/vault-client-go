// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiListRolesResponse struct for PkiListRolesResponse
type PkiListRolesResponse struct {
	// List of keys
	Keys map[string]interface{} `json:"keys"`
}

// NewPkiListRolesResponseWithDefaults instantiates a new PkiListRolesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiListRolesResponseWithDefaults() *PkiListRolesResponse {
	var this PkiListRolesResponse

	return &this
}

func (o PkiListRolesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["keys"] = o.Keys

	return json.Marshal(toSerialize)
}
