// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// SSHWriteCredentialsRequest struct for SSHWriteCredentialsRequest
type SSHWriteCredentialsRequest struct {
	// [Required] IP of the remote host
	Ip string `json:"ip"`

	// [Optional] Username in remote host
	Username string `json:"username"`
}

// NewSSHWriteCredentialsRequestWithDefaults instantiates a new SSHWriteCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHWriteCredentialsRequestWithDefaults() *SSHWriteCredentialsRequest {
	var this SSHWriteCredentialsRequest

	return &this
}

func (o SSHWriteCredentialsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
