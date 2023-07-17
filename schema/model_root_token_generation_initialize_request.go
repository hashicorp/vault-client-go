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
