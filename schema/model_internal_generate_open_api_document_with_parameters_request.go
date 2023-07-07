// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// InternalGenerateOpenApiDocumentWithParametersRequest struct for InternalGenerateOpenApiDocumentWithParametersRequest
type InternalGenerateOpenApiDocumentWithParametersRequest struct {
	// Context string appended to every operationId
	Context string `json:"context,omitempty"`
}

// NewInternalGenerateOpenApiDocumentWithParametersRequestWithDefaults instantiates a new InternalGenerateOpenApiDocumentWithParametersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalGenerateOpenApiDocumentWithParametersRequestWithDefaults() *InternalGenerateOpenApiDocumentWithParametersRequest {
	var this InternalGenerateOpenApiDocumentWithParametersRequest

	return &this
}
