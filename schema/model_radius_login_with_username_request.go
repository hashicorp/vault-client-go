// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RadiusLoginWithUsernameRequest struct for RadiusLoginWithUsernameRequest
type RadiusLoginWithUsernameRequest struct {
	// Password for this user.
	Password string `json:"password,omitempty"`

	// Username to be used for login. (POST request body)
	Username string `json:"username,omitempty"`
}
