// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RekeyVerificationUpdateResponse struct for RekeyVerificationUpdateResponse
type RekeyVerificationUpdateResponse struct {
	Complete bool `json:"complete"`

	Nounce string `json:"nounce"`
}

// NewRekeyVerificationUpdateResponseWithDefaults instantiates a new RekeyVerificationUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRekeyVerificationUpdateResponseWithDefaults() *RekeyVerificationUpdateResponse {
	var this RekeyVerificationUpdateResponse

	return &this
}

func (o RekeyVerificationUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["complete"] = o.Complete
	toSerialize["nounce"] = o.Nounce

	return json.Marshal(toSerialize)
}
