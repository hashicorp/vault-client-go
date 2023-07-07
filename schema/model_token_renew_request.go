// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TokenRenewRequest struct for TokenRenewRequest
type TokenRenewRequest struct {
	// The desired increment in seconds to the token expiration
	Increment string `json:"increment,omitempty"`

	// Token to renew (request body)
	Token string `json:"token,omitempty"`
}

// NewTokenRenewRequestWithDefaults instantiates a new TokenRenewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRenewRequestWithDefaults() *TokenRenewRequest {
	var this TokenRenewRequest

	this.Increment = "0"

	return &this
}
