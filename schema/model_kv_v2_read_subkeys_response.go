// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KvV2ReadSubkeysResponse struct for KvV2ReadSubkeysResponse
type KvV2ReadSubkeysResponse struct {
	Metadata map[string]interface{} `json:"metadata"`

	Subkeys map[string]interface{} `json:"subkeys"`
}

// NewKvV2ReadSubkeysResponseWithDefaults instantiates a new KvV2ReadSubkeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvV2ReadSubkeysResponseWithDefaults() *KvV2ReadSubkeysResponse {
	var this KvV2ReadSubkeysResponse

	return &this
}

func (o KvV2ReadSubkeysResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["metadata"] = o.Metadata
	toSerialize["subkeys"] = o.Subkeys

	return json.Marshal(toSerialize)
}
