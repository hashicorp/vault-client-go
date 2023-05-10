// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RekeyVerificationCancelResponse struct for RekeyVerificationCancelResponse
type RekeyVerificationCancelResponse struct {
	N int32 `json:"n"`

	Nounce string `json:"nounce"`

	Progress int32 `json:"progress"`

	Started string `json:"started"`

	T int32 `json:"t"`
}

// NewRekeyVerificationCancelResponseWithDefaults instantiates a new RekeyVerificationCancelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRekeyVerificationCancelResponseWithDefaults() *RekeyVerificationCancelResponse {
	var this RekeyVerificationCancelResponse

	return &this
}

func (o RekeyVerificationCancelResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["n"] = o.N
	toSerialize["nounce"] = o.Nounce
	toSerialize["progress"] = o.Progress
	toSerialize["started"] = o.Started
	toSerialize["t"] = o.T

	return json.Marshal(toSerialize)
}
