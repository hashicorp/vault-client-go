// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// UiHeadersConfigureRequest struct for UiHeadersConfigureRequest
type UiHeadersConfigureRequest struct {
	// Returns multiple values if true
	Multivalue bool `json:"multivalue,omitempty"`

	// The values to set the header.
	Values []string `json:"values,omitempty"`
}
