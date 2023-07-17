// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TokenRenewSelfRequest struct for TokenRenewSelfRequest
type TokenRenewSelfRequest struct {
	// The desired increment in seconds to the token expiration
	Increment string `json:"increment,omitempty"`

	// Token to renew (unused, does not need to be set)
	Token string `json:"token,omitempty"`
}
