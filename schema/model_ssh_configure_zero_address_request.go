// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// SshConfigureZeroAddressRequest struct for SshConfigureZeroAddressRequest
type SshConfigureZeroAddressRequest struct {
	// [Required] Comma separated list of role names which allows credentials to be requested for any IP address. CIDR blocks previously registered under these roles will be ignored.
	Roles []string `json:"roles,omitempty"`
}
