// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TokenRenewSelfRequest struct for TokenRenewSelfRequest
type TokenRenewSelfRequest struct {
	// The desired increment in seconds to the token expiration
	Increment int32 `json:"increment"`

	// Token to renew (unused, does not need to be set)
	Token string `json:"token"`
}

// NewTokenRenewSelfRequestWithDefaults instantiates a new TokenRenewSelfRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRenewSelfRequestWithDefaults() *TokenRenewSelfRequest {
	var this TokenRenewSelfRequest

	this.Increment = 0

	return &this
}

func (o TokenRenewSelfRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["increment"] = o.Increment
	toSerialize["token"] = o.Token

	return json.Marshal(toSerialize)
}
