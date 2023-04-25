// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LdapLibraryCheckOutRequest struct for LdapLibraryCheckOutRequest
type LdapLibraryCheckOutRequest struct {
	// The length of time before the check-out will expire, in seconds.
	Ttl int32 `json:"ttl"`
}

// NewLdapLibraryCheckOutRequestWithDefaults instantiates a new LdapLibraryCheckOutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapLibraryCheckOutRequestWithDefaults() *LdapLibraryCheckOutRequest {
	var this LdapLibraryCheckOutRequest

	return &this
}

func (o LdapLibraryCheckOutRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
