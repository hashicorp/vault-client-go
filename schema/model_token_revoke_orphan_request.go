// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TokenRevokeOrphanRequest struct for TokenRevokeOrphanRequest
type TokenRevokeOrphanRequest struct {
	// Token to revoke (request body)
	Token string `json:"token"`
}

// NewTokenRevokeOrphanRequestWithDefaults instantiates a new TokenRevokeOrphanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRevokeOrphanRequestWithDefaults() *TokenRevokeOrphanRequest {
	var this TokenRevokeOrphanRequest

	return &this
}

func (o TokenRevokeOrphanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
