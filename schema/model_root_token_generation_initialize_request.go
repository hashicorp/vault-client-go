// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RootTokenGenerationInitializeRequest struct for RootTokenGenerationInitializeRequest
type RootTokenGenerationInitializeRequest struct {
	// Specifies a base64-encoded PGP public key.
	PgpKey string `json:"pgp_key,omitempty"`
}

// NewRootTokenGenerationInitializeRequestWithDefaults instantiates a new RootTokenGenerationInitializeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootTokenGenerationInitializeRequestWithDefaults() *RootTokenGenerationInitializeRequest {
	var this RootTokenGenerationInitializeRequest

	return &this
}
