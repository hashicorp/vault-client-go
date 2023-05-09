// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// OciLoginRequest struct for OciLoginRequest
type OciLoginRequest struct {
	// The signed headers of the client
	RequestHeaders string `json:"request_headers,omitempty"`
}

// NewOciLoginRequestWithDefaults instantiates a new OciLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOciLoginRequestWithDefaults() *OciLoginRequest {
	var this OciLoginRequest

	return &this
}
