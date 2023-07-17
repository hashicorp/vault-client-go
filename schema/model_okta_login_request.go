// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// OktaLoginRequest struct for OktaLoginRequest
type OktaLoginRequest struct {
	// Nonce provided if performing login that requires number verification challenge. Logins through the vault login CLI command will automatically generate a nonce.
	Nonce string `json:"nonce,omitempty"`

	// Password for this user.
	Password string `json:"password,omitempty"`

	// Preferred factor provider.
	Provider string `json:"provider,omitempty"`

	// TOTP passcode.
	Totp string `json:"totp,omitempty"`
}
