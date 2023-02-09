// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GoogleCloudWriteRolesetKeyRequest struct for GoogleCloudWriteRolesetKeyRequest
type GoogleCloudWriteRolesetKeyRequest struct {
	// Private key algorithm for service account key - defaults to KEY_ALG_RSA_2048\"
	KeyAlgorithm string `json:"key_algorithm"`

	// Private key type for service account key - defaults to TYPE_GOOGLE_CREDENTIALS_FILE\"
	KeyType string `json:"key_type"`

	// Lifetime of the service account key
	Ttl int32 `json:"ttl"`
}

// NewGoogleCloudWriteRolesetKeyRequestWithDefaults instantiates a new GoogleCloudWriteRolesetKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudWriteRolesetKeyRequestWithDefaults() *GoogleCloudWriteRolesetKeyRequest {
	var this GoogleCloudWriteRolesetKeyRequest

	this.KeyAlgorithm = "KEY_ALG_RSA_2048"
	this.KeyType = "TYPE_GOOGLE_CREDENTIALS_FILE"

	return &this
}

func (o GoogleCloudWriteRolesetKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
