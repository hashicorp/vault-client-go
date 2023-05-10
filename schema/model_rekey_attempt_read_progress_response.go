// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RekeyAttemptReadProgressResponse struct for RekeyAttemptReadProgressResponse
type RekeyAttemptReadProgressResponse struct {
	Backup bool `json:"backup"`

	N int32 `json:"n"`

	Nounce string `json:"nounce"`

	PgpFingerprints []string `json:"pgp_fingerprints"`

	Progress int32 `json:"progress"`

	Required int32 `json:"required"`

	Started string `json:"started"`

	T int32 `json:"t"`

	VerificationNonce string `json:"verification_nonce"`

	VerificationRequired bool `json:"verification_required"`
}

// NewRekeyAttemptReadProgressResponseWithDefaults instantiates a new RekeyAttemptReadProgressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRekeyAttemptReadProgressResponseWithDefaults() *RekeyAttemptReadProgressResponse {
	var this RekeyAttemptReadProgressResponse

	return &this
}

func (o RekeyAttemptReadProgressResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["backup"] = o.Backup
	toSerialize["n"] = o.N
	toSerialize["nounce"] = o.Nounce
	toSerialize["pgp_fingerprints"] = o.PgpFingerprints
	toSerialize["progress"] = o.Progress
	toSerialize["required"] = o.Required
	toSerialize["started"] = o.Started
	toSerialize["t"] = o.T
	toSerialize["verification_nonce"] = o.VerificationNonce
	toSerialize["verification_required"] = o.VerificationRequired

	return json.Marshal(toSerialize)
}
