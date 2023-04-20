// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MfaGenerateTotpSecretRequest struct for MfaGenerateTotpSecretRequest
type MfaGenerateTotpSecretRequest struct {
	// The unique identifier for this MFA method.
	MethodId string `json:"method_id"`
}

// NewMfaGenerateTotpSecretRequestWithDefaults instantiates a new MfaGenerateTotpSecretRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaGenerateTotpSecretRequestWithDefaults() *MfaGenerateTotpSecretRequest {
	var this MfaGenerateTotpSecretRequest

	return &this
}

func (o MfaGenerateTotpSecretRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["method_id"] = o.MethodId

	return json.Marshal(toSerialize)
}
