// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// NomadWriteRoleRequest struct for NomadWriteRoleRequest
type NomadWriteRoleRequest struct {
	// Boolean value describing if the token should be global or not. Defaults to false.
	Global bool `json:"global,omitempty"`

	// Comma-separated string or list of policies as previously created in Nomad. Required for 'client' token.
	Policies []string `json:"policies,omitempty"`

	// Which type of token to create: 'client' or 'management'. If a 'management' token, the \"policies\" parameter is not required. Defaults to 'client'.
	Type string `json:"type,omitempty"`
}
