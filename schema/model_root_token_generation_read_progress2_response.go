// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RootTokenGenerationReadProgress2Response struct for RootTokenGenerationReadProgress2Response
type RootTokenGenerationReadProgress2Response struct {
	Complete bool `json:"complete"`

	EncodedRootToken string `json:"encoded_root_token"`

	EncodedToken string `json:"encoded_token"`

	Nonce string `json:"nonce"`

	Otp string `json:"otp"`

	OtpLength int32 `json:"otp_length"`

	PgpFingerprint string `json:"pgp_fingerprint"`

	Progress int32 `json:"progress"`

	Required int32 `json:"required"`

	Started bool `json:"started"`
}

// NewRootTokenGenerationReadProgress2ResponseWithDefaults instantiates a new RootTokenGenerationReadProgress2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootTokenGenerationReadProgress2ResponseWithDefaults() *RootTokenGenerationReadProgress2Response {
	var this RootTokenGenerationReadProgress2Response

	return &this
}

func (o RootTokenGenerationReadProgress2Response) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["complete"] = o.Complete
	toSerialize["encoded_root_token"] = o.EncodedRootToken
	toSerialize["encoded_token"] = o.EncodedToken
	toSerialize["nonce"] = o.Nonce
	toSerialize["otp"] = o.Otp
	toSerialize["otp_length"] = o.OtpLength
	toSerialize["pgp_fingerprint"] = o.PgpFingerprint
	toSerialize["progress"] = o.Progress
	toSerialize["required"] = o.Required
	toSerialize["started"] = o.Started

	return json.Marshal(toSerialize)
}
