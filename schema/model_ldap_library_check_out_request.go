// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// LdapLibraryCheckOutRequest struct for LdapLibraryCheckOutRequest
type LdapLibraryCheckOutRequest struct {
	// The length of time before the check-out will expire, in seconds.
	Ttl string `json:"ttl,omitempty"`
}
