// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MfaValidateRequest struct for MfaValidateRequest
type MfaValidateRequest struct {
	// A map from MFA method ID to a slice of passcodes or an empty slice if the method does not use passcodes
	MfaPayload map[string]interface{} `json:"mfa_payload"`

	// ID for this MFA request
	MfaRequestId string `json:"mfa_request_id"`
}

// NewMfaValidateRequestWithDefaults instantiates a new MfaValidateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaValidateRequestWithDefaults() *MfaValidateRequest {
	var this MfaValidateRequest

	return &this
}

func (o MfaValidateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["mfa_payload"] = o.MfaPayload
	toSerialize["mfa_request_id"] = o.MfaRequestId

	return json.Marshal(toSerialize)
}
