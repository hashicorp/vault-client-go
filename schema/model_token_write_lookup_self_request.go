// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TokenWriteLookupSelfRequest struct for TokenWriteLookupSelfRequest
type TokenWriteLookupSelfRequest struct {
	// Token to look up (unused, does not need to be set)
	Token string `json:"token"`
}

// NewTokenWriteLookupSelfRequestWithDefaults instantiates a new TokenWriteLookupSelfRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWriteLookupSelfRequestWithDefaults() *TokenWriteLookupSelfRequest {
	var this TokenWriteLookupSelfRequest

	return &this
}

func (o TokenWriteLookupSelfRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
