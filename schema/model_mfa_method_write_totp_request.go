// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MFAMethodWriteTOTPRequest struct for MFAMethodWriteTOTPRequest
type MFAMethodWriteTOTPRequest struct {
	// The hashing algorithm used to generate the TOTP token. Options include SHA1, SHA256 and SHA512.
	Algorithm string `json:"algorithm"`

	// The number of digits in the generated TOTP token. This value can either be 6 or 8.
	Digits int32 `json:"digits"`

	// The name of the key's issuing organization.
	Issuer string `json:"issuer"`

	// Determines the size in bytes of the generated key.
	KeySize int32 `json:"key_size"`

	// Max number of allowed validation attempts.
	MaxValidationAttempts int32 `json:"max_validation_attempts"`

	// The unique identifier for this MFA method.
	MethodId string `json:"method_id"`

	// The length of time used to generate a counter for the TOTP token calculation.
	Period int32 `json:"period"`

	// The pixel size of the generated square QR code.
	QrSize int32 `json:"qr_size"`

	// The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
	Skew int32 `json:"skew"`
}

// NewMFAMethodWriteTOTPRequestWithDefaults instantiates a new MFAMethodWriteTOTPRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAMethodWriteTOTPRequestWithDefaults() *MFAMethodWriteTOTPRequest {
	var this MFAMethodWriteTOTPRequest

	this.Algorithm = "SHA1"
	this.Digits = 6
	this.KeySize = 20
	this.Period = 30
	this.QrSize = 200
	this.Skew = 1

	return &this
}

func (o MFAMethodWriteTOTPRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["algorithm"] = o.Algorithm
	toSerialize["digits"] = o.Digits
	toSerialize["issuer"] = o.Issuer
	toSerialize["key_size"] = o.KeySize
	toSerialize["max_validation_attempts"] = o.MaxValidationAttempts
	toSerialize["method_id"] = o.MethodId
	toSerialize["period"] = o.Period
	toSerialize["qr_size"] = o.QrSize
	toSerialize["skew"] = o.Skew

	return json.Marshal(toSerialize)
}
