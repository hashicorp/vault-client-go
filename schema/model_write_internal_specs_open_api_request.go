// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteInternalSpecsOpenAPIRequest struct for WriteInternalSpecsOpenAPIRequest
type WriteInternalSpecsOpenAPIRequest struct {
	// Context string appended to every operationId
	Context string `json:"context"`
}

// NewWriteInternalSpecsOpenAPIRequestWithDefaults instantiates a new WriteInternalSpecsOpenAPIRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteInternalSpecsOpenAPIRequestWithDefaults() *WriteInternalSpecsOpenAPIRequest {
	var this WriteInternalSpecsOpenAPIRequest

	return &this
}

func (o WriteInternalSpecsOpenAPIRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
