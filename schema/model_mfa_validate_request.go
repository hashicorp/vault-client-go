// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MFAValidateRequest struct for MFAValidateRequest
type MFAValidateRequest struct {
	// A map from MFA method ID to a slice of passcodes or an empty slice if the method does not use passcodes
	MfaPayload map[string]interface{} `json:"mfa_payload"`

	// ID for this MFA request
	MfaRequestId string `json:"mfa_request_id"`
}

// NewMFAValidateRequestWithDefaults instantiates a new MFAValidateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAValidateRequestWithDefaults() *MFAValidateRequest {
	var this MFAValidateRequest

	return &this
}

func (o MFAValidateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["mfa_payload"] = o.MfaPayload
	toSerialize["mfa_request_id"] = o.MfaRequestId

	return json.Marshal(toSerialize)
}
