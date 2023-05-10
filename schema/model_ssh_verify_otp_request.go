// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// SshVerifyOtpRequest struct for SshVerifyOtpRequest
type SshVerifyOtpRequest struct {
	// [Required] One-Time-Key that needs to be validated
	Otp string `json:"otp"`
}

// NewSshVerifyOtpRequestWithDefaults instantiates a new SshVerifyOtpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshVerifyOtpRequestWithDefaults() *SshVerifyOtpRequest {
	var this SshVerifyOtpRequest

	return &this
}

func (o SshVerifyOtpRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["otp"] = o.Otp

	return json.Marshal(toSerialize)
}
