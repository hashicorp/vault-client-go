// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteConfigUIHeaderRequest struct for WriteConfigUIHeaderRequest
type WriteConfigUIHeaderRequest struct {
	// Returns multiple values if true
	Multivalue bool `json:"multivalue"`

	// The values to set the header.
	Values []string `json:"values"`
}

// NewWriteConfigUIHeaderRequestWithDefaults instantiates a new WriteConfigUIHeaderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteConfigUIHeaderRequestWithDefaults() *WriteConfigUIHeaderRequest {
	var this WriteConfigUIHeaderRequest

	return &this
}

func (o WriteConfigUIHeaderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["multivalue"] = o.Multivalue
	toSerialize["values"] = o.Values

	return json.Marshal(toSerialize)
}
