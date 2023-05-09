// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RootTokenGenerationReadProgressResponse struct for RootTokenGenerationReadProgressResponse
type RootTokenGenerationReadProgressResponse struct {
	Complete bool `json:"complete,omitempty"`

	EncodedRootToken string `json:"encoded_root_token,omitempty"`

	EncodedToken string `json:"encoded_token,omitempty"`

	Nonce string `json:"nonce,omitempty"`

	Otp string `json:"otp,omitempty"`

	OtpLength int32 `json:"otp_length,omitempty"`

	PgpFingerprint string `json:"pgp_fingerprint,omitempty"`

	Progress int32 `json:"progress,omitempty"`

	Required int32 `json:"required,omitempty"`

	Started bool `json:"started,omitempty"`
}

// NewRootTokenGenerationReadProgressResponseWithDefaults instantiates a new RootTokenGenerationReadProgressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootTokenGenerationReadProgressResponseWithDefaults() *RootTokenGenerationReadProgressResponse {
	var this RootTokenGenerationReadProgressResponse

	return &this
}
