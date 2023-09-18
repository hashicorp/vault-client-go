// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// SshVerifyOtpRequest struct for SshVerifyOtpRequest
type SshVerifyOtpRequest struct {
	// [Required] One-Time-Key that needs to be validated
	Otp string `json:"otp,omitempty"`
}
