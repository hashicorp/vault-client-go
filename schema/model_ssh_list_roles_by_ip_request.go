// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// SshListRolesByIpRequest struct for SshListRolesByIpRequest
type SshListRolesByIpRequest struct {
	// [Required] IP address of remote host
	Ip string `json:"ip,omitempty"`
}

// NewSshListRolesByIpRequestWithDefaults instantiates a new SshListRolesByIpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshListRolesByIpRequestWithDefaults() *SshListRolesByIpRequest {
	var this SshListRolesByIpRequest

	return &this
}
