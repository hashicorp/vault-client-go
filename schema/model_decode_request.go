// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// DecodeRequest struct for DecodeRequest
type DecodeRequest struct {
	// Specifies the encoded token (result from generate-root).
	EncodedToken string `json:"encoded_token,omitempty"`

	// Specifies the otp code for decode.
	Otp string `json:"otp,omitempty"`
}

// NewDecodeRequestWithDefaults instantiates a new DecodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodeRequestWithDefaults() *DecodeRequest {
	var this DecodeRequest

	return &this
}
