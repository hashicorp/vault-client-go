// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// ActiveDirectoryCheckOutLibraryRequest struct for ActiveDirectoryCheckOutLibraryRequest
type ActiveDirectoryCheckOutLibraryRequest struct {
	// The length of time before the check-out will expire, in seconds.
	Ttl int32 `json:"ttl"`
}

// NewActiveDirectoryCheckOutLibraryRequestWithDefaults instantiates a new ActiveDirectoryCheckOutLibraryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveDirectoryCheckOutLibraryRequestWithDefaults() *ActiveDirectoryCheckOutLibraryRequest {
	var this ActiveDirectoryCheckOutLibraryRequest

	return &this
}

func (o ActiveDirectoryCheckOutLibraryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
