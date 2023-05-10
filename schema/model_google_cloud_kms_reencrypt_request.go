// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GoogleCloudKmsReencryptRequest struct for GoogleCloudKmsReencryptRequest
type GoogleCloudKmsReencryptRequest struct {
	// Optional data that, if specified, must also be provided during decryption.
	AdditionalAuthenticatedData string `json:"additional_authenticated_data"`

	// Ciphertext to be re-encrypted to the latest key version. This must be ciphertext that Vault previously generated for this named key.
	Ciphertext string `json:"ciphertext"`

	// Integer version of the crypto key version to use for the new encryption. If unspecified, this defaults to the latest active crypto key version.
	KeyVersion int32 `json:"key_version"`
}

// NewGoogleCloudKmsReencryptRequestWithDefaults instantiates a new GoogleCloudKmsReencryptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudKmsReencryptRequestWithDefaults() *GoogleCloudKmsReencryptRequest {
	var this GoogleCloudKmsReencryptRequest

	return &this
}

func (o GoogleCloudKmsReencryptRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["additional_authenticated_data"] = o.AdditionalAuthenticatedData
	toSerialize["ciphertext"] = o.Ciphertext
	toSerialize["key_version"] = o.KeyVersion

	return json.Marshal(toSerialize)
}
