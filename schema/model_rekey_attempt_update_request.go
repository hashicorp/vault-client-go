// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RekeyAttemptUpdateRequest struct for RekeyAttemptUpdateRequest
type RekeyAttemptUpdateRequest struct {
	// Specifies a single unseal key share.
	Key string `json:"key,omitempty"`

	// Specifies the nonce of the rekey attempt.
	Nonce string `json:"nonce,omitempty"`
}
