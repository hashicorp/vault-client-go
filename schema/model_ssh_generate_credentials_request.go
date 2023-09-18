// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// SshGenerateCredentialsRequest struct for SshGenerateCredentialsRequest
type SshGenerateCredentialsRequest struct {
	// [Required] IP of the remote host
	Ip string `json:"ip,omitempty"`

	// [Optional] Username in remote host
	Username string `json:"username,omitempty"`
}
