// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LdapLoginRequest struct for LdapLoginRequest
type LdapLoginRequest struct {
	// Password for this user.
	Password string `json:"password"`
}

// NewLdapLoginRequestWithDefaults instantiates a new LdapLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapLoginRequestWithDefaults() *LdapLoginRequest {
	var this LdapLoginRequest

	return &this
}

func (o LdapLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["password"] = o.Password

	return json.Marshal(toSerialize)
}
