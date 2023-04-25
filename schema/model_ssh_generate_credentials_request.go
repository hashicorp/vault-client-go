// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// SshGenerateCredentialsRequest struct for SshGenerateCredentialsRequest
type SshGenerateCredentialsRequest struct {
	// [Required] IP of the remote host
	Ip string `json:"ip"`

	// [Optional] Username in remote host
	Username string `json:"username"`
}

// NewSshGenerateCredentialsRequestWithDefaults instantiates a new SshGenerateCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshGenerateCredentialsRequestWithDefaults() *SshGenerateCredentialsRequest {
	var this SshGenerateCredentialsRequest

	return &this
}

func (o SshGenerateCredentialsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["ip"] = o.Ip
	toSerialize["username"] = o.Username

	return json.Marshal(toSerialize)
}
