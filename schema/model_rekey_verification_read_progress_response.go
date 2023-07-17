// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RekeyVerificationReadProgressResponse struct for RekeyVerificationReadProgressResponse
type RekeyVerificationReadProgressResponse struct {
	N int32 `json:"n,omitempty"`

	Nounce string `json:"nounce,omitempty"`

	Progress int32 `json:"progress,omitempty"`

	Started string `json:"started,omitempty"`

	T int32 `json:"t,omitempty"`
}
