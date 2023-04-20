// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RootTokenGenerationInitialize2Request struct for RootTokenGenerationInitialize2Request
type RootTokenGenerationInitialize2Request struct {
	// Specifies a base64-encoded PGP public key.
	PgpKey string `json:"pgp_key"`
}

// NewRootTokenGenerationInitialize2RequestWithDefaults instantiates a new RootTokenGenerationInitialize2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootTokenGenerationInitialize2RequestWithDefaults() *RootTokenGenerationInitialize2Request {
	var this RootTokenGenerationInitialize2Request

	return &this
}

func (o RootTokenGenerationInitialize2Request) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["pgp_key"] = o.PgpKey

	return json.Marshal(toSerialize)
}
