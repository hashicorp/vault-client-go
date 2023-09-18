// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RekeyAttemptReadProgressResponse struct for RekeyAttemptReadProgressResponse
type RekeyAttemptReadProgressResponse struct {
	Backup bool `json:"backup,omitempty"`

	N int32 `json:"n,omitempty"`

	Nounce string `json:"nounce,omitempty"`

	PgpFingerprints []string `json:"pgp_fingerprints,omitempty"`

	Progress int32 `json:"progress,omitempty"`

	Required int32 `json:"required,omitempty"`

	Started string `json:"started,omitempty"`

	T int32 `json:"t,omitempty"`

	VerificationNonce string `json:"verification_nonce,omitempty"`

	VerificationRequired bool `json:"verification_required,omitempty"`
}
