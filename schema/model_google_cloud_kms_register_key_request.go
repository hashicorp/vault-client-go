// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GoogleCloudKMSRegisterKeyRequest struct for GoogleCloudKMSRegisterKeyRequest
type GoogleCloudKMSRegisterKeyRequest struct {
	// Full resource ID of the crypto key including the project, location, key ring, and crypto key like \"projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s\". This crypto key must already exist in Google Cloud KMS unless verify is set to \"false\".
	CryptoKey string `json:"crypto_key"`

	// Verify that the given Google Cloud KMS crypto key exists and is accessible before creating the storage entry in Vault. Set this to \"false\" if the key will not exist at creation time.
	Verify bool `json:"verify"`
}

// NewGoogleCloudKMSRegisterKeyRequestWithDefaults instantiates a new GoogleCloudKMSRegisterKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudKMSRegisterKeyRequestWithDefaults() *GoogleCloudKMSRegisterKeyRequest {
	var this GoogleCloudKMSRegisterKeyRequest

	this.Verify = true

	return &this
}

func (o GoogleCloudKMSRegisterKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["crypto_key"] = o.CryptoKey
	toSerialize["verify"] = o.Verify

	return json.Marshal(toSerialize)
}
