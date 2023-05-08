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

// NewRadiusLoginWithUsernameRequestWithDefaults instantiates a new RadiusLoginWithUsernameRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadiusLoginWithUsernameRequestWithDefaults() *RadiusLoginWithUsernameRequest {
	var this RadiusLoginWithUsernameRequest

	return &this
}
