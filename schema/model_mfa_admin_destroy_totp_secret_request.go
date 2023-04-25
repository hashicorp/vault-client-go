// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MfaAdminDestroyTotpSecretRequest struct for MfaAdminDestroyTotpSecretRequest
type MfaAdminDestroyTotpSecretRequest struct {
	// Identifier of the entity from which the MFA method secret needs to be removed.
	EntityId string `json:"entity_id"`

	// The unique identifier for this MFA method.
	MethodId string `json:"method_id"`
}

// NewMfaAdminDestroyTotpSecretRequestWithDefaults instantiates a new MfaAdminDestroyTotpSecretRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaAdminDestroyTotpSecretRequestWithDefaults() *MfaAdminDestroyTotpSecretRequest {
	var this MfaAdminDestroyTotpSecretRequest

	return &this
}

func (o MfaAdminDestroyTotpSecretRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["entity_id"] = o.EntityId
	toSerialize["method_id"] = o.MethodId

	return json.Marshal(toSerialize)
}
