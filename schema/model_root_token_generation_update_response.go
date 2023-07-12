// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RootTokenGenerationUpdateResponse struct for RootTokenGenerationUpdateResponse
type RootTokenGenerationUpdateResponse struct {
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
