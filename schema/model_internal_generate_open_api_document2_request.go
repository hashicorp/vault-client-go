// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// InternalGenerateOpenApiDocument2Request struct for InternalGenerateOpenApiDocument2Request
type InternalGenerateOpenApiDocument2Request struct {
	// Context string appended to every operationId
	Context string `json:"context"`
}

// NewInternalGenerateOpenApiDocument2RequestWithDefaults instantiates a new InternalGenerateOpenApiDocument2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalGenerateOpenApiDocument2RequestWithDefaults() *InternalGenerateOpenApiDocument2Request {
	var this InternalGenerateOpenApiDocument2Request

	return &this
}

func (o InternalGenerateOpenApiDocument2Request) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["context"] = o.Context

	return json.Marshal(toSerialize)
}
