// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LDAPLoginRequest struct for LDAPLoginRequest
type LDAPLoginRequest struct {
	// Password for this user.
	Password string `json:"password"`
}

// NewLDAPLoginRequestWithDefaults instantiates a new LDAPLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPLoginRequestWithDefaults() *LDAPLoginRequest {
	var this LDAPLoginRequest

	return &this
}

func (o LDAPLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
