// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadSecretIDTTLResponse struct for AppRoleReadSecretIDTTLResponse
type AppRoleReadSecretIDTTLResponse struct {
	// Duration in seconds after which the issued secret ID should expire. Defaults to 0, meaning no expiration.
	SecretIdTtl int32 `json:"secret_id_ttl"`
}

// NewAppRoleReadSecretIDTTLResponseWithDefaults instantiates a new AppRoleReadSecretIDTTLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadSecretIDTTLResponseWithDefaults() *AppRoleReadSecretIDTTLResponse {
	var this AppRoleReadSecretIDTTLResponse

	return &this
}

func (o AppRoleReadSecretIDTTLResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
