// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GoogleCloudWriteRolesetRequest struct for GoogleCloudWriteRolesetRequest
type GoogleCloudWriteRolesetRequest struct {
	// Bindings configuration string.
	Bindings string `json:"bindings"`

	// Name of the GCP project that this roleset's service account will belong to.
	Project string `json:"project"`

	// Type of secret generated for this role set. Defaults to 'access_token'
	SecretType string `json:"secret_type"`

	// List of OAuth scopes to assign to credentials generated under this role set
	TokenScopes []string `json:"token_scopes"`
}

// NewGoogleCloudWriteRolesetRequestWithDefaults instantiates a new GoogleCloudWriteRolesetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudWriteRolesetRequestWithDefaults() *GoogleCloudWriteRolesetRequest {
	var this GoogleCloudWriteRolesetRequest

	this.SecretType = "access_token"

	return &this
}

func (o GoogleCloudWriteRolesetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
