// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TokenLookUpSelf2Request struct for TokenLookUpSelf2Request
type TokenLookUpSelf2Request struct {
	// Token to look up (unused, does not need to be set)
	Token string `json:"token"`
}

// NewTokenLookUpSelf2RequestWithDefaults instantiates a new TokenLookUpSelf2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenLookUpSelf2RequestWithDefaults() *TokenLookUpSelf2Request {
	var this TokenLookUpSelf2Request

	return &this
}

func (o TokenLookUpSelf2Request) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token"] = o.Token

	return json.Marshal(toSerialize)
}
