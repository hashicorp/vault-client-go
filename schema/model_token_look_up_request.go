// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TokenLookUpRequest struct for TokenLookUpRequest
type TokenLookUpRequest struct {
	// Token to lookup (POST request body)
	Token string `json:"token"`
}

// NewTokenLookUpRequestWithDefaults instantiates a new TokenLookUpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenLookUpRequestWithDefaults() *TokenLookUpRequest {
	var this TokenLookUpRequest

	return &this
}

func (o TokenLookUpRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token"] = o.Token

	return json.Marshal(toSerialize)
}
