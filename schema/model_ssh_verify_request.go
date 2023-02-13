// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// SSHVerifyRequest struct for SSHVerifyRequest
type SSHVerifyRequest struct {
	// [Required] One-Time-Key that needs to be validated
	Otp string `json:"otp"`
}

// NewSSHVerifyRequestWithDefaults instantiates a new SSHVerifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHVerifyRequestWithDefaults() *SSHVerifyRequest {
	var this SSHVerifyRequest

	return &this
}

func (o SSHVerifyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["otp"] = o.Otp

	return json.Marshal(toSerialize)
}
