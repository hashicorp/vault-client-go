// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RekeyVerificationReadProgressResponse struct for RekeyVerificationReadProgressResponse
type RekeyVerificationReadProgressResponse struct {
	N int32 `json:"n"`

	Nounce string `json:"nounce"`

	Progress int32 `json:"progress"`

	Started string `json:"started"`

	T int32 `json:"t"`
}

// NewRekeyVerificationReadProgressResponseWithDefaults instantiates a new RekeyVerificationReadProgressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRekeyVerificationReadProgressResponseWithDefaults() *RekeyVerificationReadProgressResponse {
	var this RekeyVerificationReadProgressResponse

	return &this
}

func (o RekeyVerificationReadProgressResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["n"] = o.N
	toSerialize["nounce"] = o.Nounce
	toSerialize["progress"] = o.Progress
	toSerialize["started"] = o.Started
	toSerialize["t"] = o.T

	return json.Marshal(toSerialize)
}
