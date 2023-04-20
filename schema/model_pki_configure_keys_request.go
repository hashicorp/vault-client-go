// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiConfigureKeysRequest struct for PkiConfigureKeysRequest
type PkiConfigureKeysRequest struct {
	// Reference (name or identifier) of the default key.
	Default string `json:"default"`
}

// NewPkiConfigureKeysRequestWithDefaults instantiates a new PkiConfigureKeysRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiConfigureKeysRequestWithDefaults() *PkiConfigureKeysRequest {
	var this PkiConfigureKeysRequest

	return &this
}

func (o PkiConfigureKeysRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["default"] = o.Default

	return json.Marshal(toSerialize)
}
