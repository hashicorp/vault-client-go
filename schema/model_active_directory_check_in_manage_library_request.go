// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// ActiveDirectoryCheckInManageLibraryRequest struct for ActiveDirectoryCheckInManageLibraryRequest
type ActiveDirectoryCheckInManageLibraryRequest struct {
	// The username/logon name for the service accounts to check in.
	ServiceAccountNames []string `json:"service_account_names"`
}

// NewActiveDirectoryCheckInManageLibraryRequestWithDefaults instantiates a new ActiveDirectoryCheckInManageLibraryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveDirectoryCheckInManageLibraryRequestWithDefaults() *ActiveDirectoryCheckInManageLibraryRequest {
	var this ActiveDirectoryCheckInManageLibraryRequest

	return &this
}

func (o ActiveDirectoryCheckInManageLibraryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
