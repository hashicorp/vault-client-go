// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiListRevokedCertsResponse struct for PkiListRevokedCertsResponse
type PkiListRevokedCertsResponse struct {
	// List of Keys
	Keys []string `json:"keys"`
}

// NewPkiListRevokedCertsResponseWithDefaults instantiates a new PkiListRevokedCertsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiListRevokedCertsResponseWithDefaults() *PkiListRevokedCertsResponse {
	var this PkiListRevokedCertsResponse

	return &this
}

func (o PkiListRevokedCertsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["keys"] = o.Keys

	return json.Marshal(toSerialize)
}
