// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// OidcRotateKeyRequest struct for OidcRotateKeyRequest
type OidcRotateKeyRequest struct {
	// Controls how long the public portion of a key will be available for verification after being rotated. Setting verification_ttl here will override the verification_ttl set on the key.
	VerificationTtl string `json:"verification_ttl,omitempty"`
}

// NewOidcRotateKeyRequestWithDefaults instantiates a new OidcRotateKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcRotateKeyRequestWithDefaults() *OidcRotateKeyRequest {
	var this OidcRotateKeyRequest

	return &this
}
