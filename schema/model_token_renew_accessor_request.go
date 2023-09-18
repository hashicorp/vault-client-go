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
