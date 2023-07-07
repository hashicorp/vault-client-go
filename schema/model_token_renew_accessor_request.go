// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TokenRenewAccessorRequest struct for TokenRenewAccessorRequest
type TokenRenewAccessorRequest struct {
	// Accessor of the token to renew (request body)
	Accessor string `json:"accessor,omitempty"`

	// The desired increment in seconds to the token expiration
	Increment string `json:"increment,omitempty"`
}

// NewTokenRenewAccessorRequestWithDefaults instantiates a new TokenRenewAccessorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRenewAccessorRequestWithDefaults() *TokenRenewAccessorRequest {
	var this TokenRenewAccessorRequest

	this.Increment = "0"

	return &this
}
