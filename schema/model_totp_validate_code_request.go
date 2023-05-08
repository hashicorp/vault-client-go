// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TotpValidateCodeRequest struct for TotpValidateCodeRequest
type TotpValidateCodeRequest struct {
	// TOTP code to be validated.
	Code string `json:"code,omitempty"`
}

// NewTotpValidateCodeRequestWithDefaults instantiates a new TotpValidateCodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTotpValidateCodeRequestWithDefaults() *TotpValidateCodeRequest {
	var this TotpValidateCodeRequest

	return &this
}
