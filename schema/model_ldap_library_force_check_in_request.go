// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LdapLibraryForceCheckInRequest struct for LdapLibraryForceCheckInRequest
type LdapLibraryForceCheckInRequest struct {
	// The username/logon name for the service accounts to check in.
	ServiceAccountNames []string `json:"service_account_names"`
}

// NewLdapLibraryForceCheckInRequestWithDefaults instantiates a new LdapLibraryForceCheckInRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapLibraryForceCheckInRequestWithDefaults() *LdapLibraryForceCheckInRequest {
	var this LdapLibraryForceCheckInRequest

	return &this
}

func (o LdapLibraryForceCheckInRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["service_account_names"] = o.ServiceAccountNames

	return json.Marshal(toSerialize)
}
