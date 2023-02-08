// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// UserpassLoginRequest struct for UserpassLoginRequest
type UserpassLoginRequest struct {
	// Password for this user.
	Password string `json:"password"`
}

// NewUserpassLoginRequestWithDefaults instantiates a new UserpassLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserpassLoginRequestWithDefaults() *UserpassLoginRequest {
	var this UserpassLoginRequest

	return &this
}

func (o UserpassLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
