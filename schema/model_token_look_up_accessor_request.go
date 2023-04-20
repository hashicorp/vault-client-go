// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TokenLookUpAccessorRequest struct for TokenLookUpAccessorRequest
type TokenLookUpAccessorRequest struct {
	// Accessor of the token to look up (request body)
	Accessor string `json:"accessor"`
}

// NewTokenLookUpAccessorRequestWithDefaults instantiates a new TokenLookUpAccessorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenLookUpAccessorRequestWithDefaults() *TokenLookUpAccessorRequest {
	var this TokenLookUpAccessorRequest

	return &this
}

func (o TokenLookUpAccessorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["accessor"] = o.Accessor

	return json.Marshal(toSerialize)
}
