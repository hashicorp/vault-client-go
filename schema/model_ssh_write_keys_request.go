// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// SSHWriteKeysRequest struct for SSHWriteKeysRequest
type SSHWriteKeysRequest struct {
	// [Required] SSH private key with super user privileges in host
	Key string `json:"key"`
}

// NewSSHWriteKeysRequestWithDefaults instantiates a new SSHWriteKeysRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHWriteKeysRequestWithDefaults() *SSHWriteKeysRequest {
	var this SSHWriteKeysRequest

	return &this
}

func (o SSHWriteKeysRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["key"] = o.Key

	return json.Marshal(toSerialize)
}
