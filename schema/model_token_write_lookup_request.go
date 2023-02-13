// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TokenWriteLookupRequest struct for TokenWriteLookupRequest
type TokenWriteLookupRequest struct {
	// Token to lookup (POST request body)
	Token string `json:"token"`
}

// NewTokenWriteLookupRequestWithDefaults instantiates a new TokenWriteLookupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWriteLookupRequestWithDefaults() *TokenWriteLookupRequest {
	var this TokenWriteLookupRequest

	return &this
}

func (o TokenWriteLookupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token"] = o.Token

	return json.Marshal(toSerialize)
}
