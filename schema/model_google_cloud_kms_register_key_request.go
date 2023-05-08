// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudKmsRegisterKeyRequest struct for GoogleCloudKmsRegisterKeyRequest
type GoogleCloudKmsRegisterKeyRequest struct {
	// Full resource ID of the crypto key including the project, location, key ring, and crypto key like \"projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s\". This crypto key must already exist in Google Cloud KMS unless verify is set to \"false\".
	CryptoKey string `json:"crypto_key,omitempty"`

	// Verify that the given Google Cloud KMS crypto key exists and is accessible before creating the storage entry in Vault. Set this to \"false\" if the key will not exist at creation time.
	Verify bool `json:"verify,omitempty"`
}

// NewGoogleCloudKmsRegisterKeyRequestWithDefaults instantiates a new GoogleCloudKmsRegisterKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudKmsRegisterKeyRequestWithDefaults() *GoogleCloudKmsRegisterKeyRequest {
	var this GoogleCloudKmsRegisterKeyRequest

	this.Verify = true

	return &this
}
