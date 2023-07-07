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

// NewLdapLibraryCheckOutRequestWithDefaults instantiates a new LdapLibraryCheckOutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapLibraryCheckOutRequestWithDefaults() *LdapLibraryCheckOutRequest {
	var this LdapLibraryCheckOutRequest

	return &this
}
