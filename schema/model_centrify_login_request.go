// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// CentrifyLoginRequest struct for CentrifyLoginRequest
type CentrifyLoginRequest struct {
	// Auth mode ('ro' for resource owner, 'cc' for credential client).
	Mode string `json:"mode,omitempty"`

	// Password for this user.
	Password string `json:"password,omitempty"`

	// Username of the user.
	Username string `json:"username,omitempty"`
}
