// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GoogleCloudWriteStaticAccountRequest struct for GoogleCloudWriteStaticAccountRequest
type GoogleCloudWriteStaticAccountRequest struct {
	// Bindings configuration string.
	Bindings string `json:"bindings"`

	// Type of secret generated for this account. Cannot be updated. Defaults to \"access_token\"
	SecretType string `json:"secret_type"`

	// Required. Email of the GCP service account to manage. Cannot be updated.
	ServiceAccountEmail string `json:"service_account_email"`

	// List of OAuth scopes to assign to access tokens generated under this account. Ignored if \"secret_type\" is not \"\"access_token\"\"
	TokenScopes []string `json:"token_scopes"`
}

// NewGoogleCloudWriteStaticAccountRequestWithDefaults instantiates a new GoogleCloudWriteStaticAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudWriteStaticAccountRequestWithDefaults() *GoogleCloudWriteStaticAccountRequest {
	var this GoogleCloudWriteStaticAccountRequest

	this.SecretType = "access_token"

	return &this
}

func (o GoogleCloudWriteStaticAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
