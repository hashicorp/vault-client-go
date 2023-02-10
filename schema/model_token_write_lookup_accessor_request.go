// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TokenWriteLookupAccessorRequest struct for TokenWriteLookupAccessorRequest
type TokenWriteLookupAccessorRequest struct {
	// Accessor of the token to look up (request body)
	Accessor string `json:"accessor"`
}

// NewTokenWriteLookupAccessorRequestWithDefaults instantiates a new TokenWriteLookupAccessorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWriteLookupAccessorRequestWithDefaults() *TokenWriteLookupAccessorRequest {
	var this TokenWriteLookupAccessorRequest

	return &this
}

func (o TokenWriteLookupAccessorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["accessor"] = o.Accessor

	return json.Marshal(toSerialize)
}
