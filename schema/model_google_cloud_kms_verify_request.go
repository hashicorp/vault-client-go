// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GoogleCloudKmsVerifyRequest struct for GoogleCloudKmsVerifyRequest
type GoogleCloudKmsVerifyRequest struct {
	// Digest to verify. This digest must use the same SHA algorithm as the underlying Cloud KMS key. The digest must be the base64-encoded binary value. This field is required.
	Digest string `json:"digest"`

	// Integer version of the crypto key version to use for verification. This field is required.
	KeyVersion int32 `json:"key_version"`

	// Base64-encoded signature to use for verification. This field is required.
	Signature string `json:"signature"`
}

// NewGoogleCloudKmsVerifyRequestWithDefaults instantiates a new GoogleCloudKmsVerifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudKmsVerifyRequestWithDefaults() *GoogleCloudKmsVerifyRequest {
	var this GoogleCloudKmsVerifyRequest

	return &this
}

func (o GoogleCloudKmsVerifyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["digest"] = o.Digest
	toSerialize["key_version"] = o.KeyVersion
	toSerialize["signature"] = o.Signature

	return json.Marshal(toSerialize)
}
