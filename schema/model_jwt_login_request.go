// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// JwtLoginRequest struct for JwtLoginRequest
type JwtLoginRequest struct {
	// The signed JWT to validate.
	Jwt string `json:"jwt,omitempty"`

	// The role to log in against.
	Role string `json:"role,omitempty"`
}
