// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiReplaceRootRequest struct for PkiReplaceRootRequest
type PkiReplaceRootRequest struct {
	// Reference (name or identifier) to the default issuer.
	Default string `json:"default,omitempty"`
}

// NewPkiReplaceRootRequestWithDefaults instantiates a new PkiReplaceRootRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiReplaceRootRequestWithDefaults() *PkiReplaceRootRequest {
	var this PkiReplaceRootRequest

	this.Default = "next"

	return &this
}
