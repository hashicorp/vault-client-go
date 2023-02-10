// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteRawRequest struct for WriteRawRequest
type WriteRawRequest struct {
	Compressed bool `json:"compressed"`

	CompressionType string `json:"compression_type"`

	Encoding string `json:"encoding"`

	Path string `json:"path"`

	Value string `json:"value"`
}

// NewWriteRawRequestWithDefaults instantiates a new WriteRawRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteRawRequestWithDefaults() *WriteRawRequest {
	var this WriteRawRequest

	return &this
}

func (o WriteRawRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["compressed"] = o.Compressed
	toSerialize["compression_type"] = o.CompressionType
	toSerialize["encoding"] = o.Encoding
	toSerialize["path"] = o.Path
	toSerialize["value"] = o.Value

	return json.Marshal(toSerialize)
}
