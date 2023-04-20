// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GoogleCloudWriteImpersonatedAccountRequest struct for GoogleCloudWriteImpersonatedAccountRequest
type GoogleCloudWriteImpersonatedAccountRequest struct {
	// Required. Email of the GCP service account to manage. Cannot be updated.
	ServiceAccountEmail string `json:"service_account_email"`

	// List of OAuth scopes to assign to access tokens generated under this account.
	TokenScopes []string `json:"token_scopes"`

	// Lifetime of the token for the impersonated account.
	Ttl int32 `json:"ttl"`
}

// NewGoogleCloudWriteImpersonatedAccountRequestWithDefaults instantiates a new GoogleCloudWriteImpersonatedAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudWriteImpersonatedAccountRequestWithDefaults() *GoogleCloudWriteImpersonatedAccountRequest {
	var this GoogleCloudWriteImpersonatedAccountRequest

	return &this
}

func (o GoogleCloudWriteImpersonatedAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["service_account_email"] = o.ServiceAccountEmail
	toSerialize["token_scopes"] = o.TokenScopes
	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
