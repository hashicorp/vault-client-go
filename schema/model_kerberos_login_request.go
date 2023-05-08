// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// KerberosLoginRequest struct for KerberosLoginRequest
type KerberosLoginRequest struct {
	// SPNEGO Authorization header. Required.
	Authorization string `json:"authorization,omitempty"`
}

// NewKerberosLoginRequestWithDefaults instantiates a new KerberosLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosLoginRequestWithDefaults() *KerberosLoginRequest {
	var this KerberosLoginRequest

	return &this
}
