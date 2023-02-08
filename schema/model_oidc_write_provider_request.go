// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OIDCWriteProviderRequest struct for OIDCWriteProviderRequest
type OIDCWriteProviderRequest struct {
	// The client IDs that are permitted to use the provider
	AllowedClientIds []string `json:"allowed_client_ids"`

	// Specifies what will be used for the iss claim of ID tokens.
	Issuer string `json:"issuer"`

	// The scopes supported for requesting on the provider
	ScopesSupported []string `json:"scopes_supported"`
}

// NewOIDCWriteProviderRequestWithDefaults instantiates a new OIDCWriteProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWriteProviderRequestWithDefaults() *OIDCWriteProviderRequest {
	var this OIDCWriteProviderRequest

	return &this
}

func (o OIDCWriteProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
