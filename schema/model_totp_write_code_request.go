// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TOTPWriteCodeRequest struct for TOTPWriteCodeRequest
type TOTPWriteCodeRequest struct {
	// TOTP code to be validated.
	Code string `json:"code"`
}

// NewTOTPWriteCodeRequestWithDefaults instantiates a new TOTPWriteCodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTOTPWriteCodeRequestWithDefaults() *TOTPWriteCodeRequest {
	var this TOTPWriteCodeRequest

	return &this
}
