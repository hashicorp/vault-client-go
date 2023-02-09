// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RevokePrefixRequest struct for RevokePrefixRequest
type RevokePrefixRequest struct {
	// Whether or not to perform the revocation synchronously
	Sync bool `json:"sync"`
}

// NewRevokePrefixRequestWithDefaults instantiates a new RevokePrefixRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokePrefixRequestWithDefaults() *RevokePrefixRequest {
	var this RevokePrefixRequest

	this.Sync = true

	return &this
}

func (o RevokePrefixRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
