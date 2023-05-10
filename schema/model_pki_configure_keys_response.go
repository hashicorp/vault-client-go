// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiConfigureKeysResponse struct for PkiConfigureKeysResponse
type PkiConfigureKeysResponse struct {
	// Reference (name or identifier) to the default issuer.
	Default string `json:"default"`
}

// NewPkiConfigureKeysResponseWithDefaults instantiates a new PkiConfigureKeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiConfigureKeysResponseWithDefaults() *PkiConfigureKeysResponse {
	var this PkiConfigureKeysResponse

	return &this
}

func (o PkiConfigureKeysResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["default"] = o.Default

	return json.Marshal(toSerialize)
}
