// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// InternalGenerateOpenApiDocumentWithParametersRequest struct for InternalGenerateOpenApiDocumentWithParametersRequest
type InternalGenerateOpenApiDocumentWithParametersRequest struct {
	// Context string appended to every operationId
	Context string `json:"context,omitempty"`

	// Use generic mount paths
	GenericMountPaths bool `json:"generic_mount_paths,omitempty"`
}
