// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// MfaConfigureTotpMethodRequest struct for MfaConfigureTotpMethodRequest
type MfaConfigureTotpMethodRequest struct {
	// The hashing algorithm used to generate the TOTP token. Options include SHA1, SHA256 and SHA512.
	Algorithm string `json:"algorithm,omitempty"`

	// The number of digits in the generated TOTP token. This value can either be 6 or 8.
	Digits int32 `json:"digits,omitempty"`

	// The name of the key's issuing organization.
	Issuer string `json:"issuer,omitempty"`

	// Determines the size in bytes of the generated key.
	KeySize int32 `json:"key_size,omitempty"`

	// Max number of allowed validation attempts.
	MaxValidationAttempts int32 `json:"max_validation_attempts,omitempty"`

	// The unique name identifier for this MFA method.
	MethodName string `json:"method_name,omitempty"`

	// The length of time used to generate a counter for the TOTP token calculation.
	Period int32 `json:"period,omitempty"`

	// The pixel size of the generated square QR code.
	QrSize int32 `json:"qr_size,omitempty"`

	// The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
	Skew int32 `json:"skew,omitempty"`
}

// NewMfaConfigureTotpMethodRequestWithDefaults instantiates a new MfaConfigureTotpMethodRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaConfigureTotpMethodRequestWithDefaults() *MfaConfigureTotpMethodRequest {
	var this MfaConfigureTotpMethodRequest

	this.Algorithm = "SHA1"
	this.Digits = 6
	this.KeySize = 20
	this.Period = 30
	this.QrSize = 200
	this.Skew = 1

	return &this
}
