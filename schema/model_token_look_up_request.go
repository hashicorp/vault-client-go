// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TokenLookUpRequest struct for TokenLookUpRequest
type TokenLookUpRequest struct {
	// Token to lookup (POST request body)
	Token string `json:"token,omitempty"`
}

// NewTokenLookUpRequestWithDefaults instantiates a new TokenLookUpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenLookUpRequestWithDefaults() *TokenLookUpRequest {
	var this TokenLookUpRequest

	return &this
}
