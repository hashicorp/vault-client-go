// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RootTokenGenerationUpdateRequest struct for RootTokenGenerationUpdateRequest
type RootTokenGenerationUpdateRequest struct {
	// Specifies a single unseal key share.
	Key string `json:"key"`

	// Specifies the nonce of the attempt.
	Nonce string `json:"nonce"`
}

// NewRootTokenGenerationUpdateRequestWithDefaults instantiates a new RootTokenGenerationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootTokenGenerationUpdateRequestWithDefaults() *RootTokenGenerationUpdateRequest {
	var this RootTokenGenerationUpdateRequest

	return &this
}

func (o RootTokenGenerationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["key"] = o.Key
	toSerialize["nonce"] = o.Nonce

	return json.Marshal(toSerialize)
}
