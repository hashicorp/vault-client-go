// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TokenRevokeAccessorRequest struct for TokenRevokeAccessorRequest
type TokenRevokeAccessorRequest struct {
	// Accessor of the token (request body)
	Accessor string `json:"accessor"`
}

// NewTokenRevokeAccessorRequestWithDefaults instantiates a new TokenRevokeAccessorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRevokeAccessorRequestWithDefaults() *TokenRevokeAccessorRequest {
	var this TokenRevokeAccessorRequest

	return &this
}

func (o TokenRevokeAccessorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
