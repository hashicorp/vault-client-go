// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OIDCWriteClientRequest struct for OIDCWriteClientRequest
type OIDCWriteClientRequest struct {
	// The time-to-live for access tokens obtained by the client.
	AccessTokenTtl int32 `json:"access_token_ttl"`

	// Comma separated string or array of assignment resources.
	Assignments []string `json:"assignments"`

	// The client type based on its ability to maintain confidentiality of credentials. The following client types are supported: 'confidential', 'public'. Defaults to 'confidential'.
	ClientType string `json:"client_type"`

	// The time-to-live for ID tokens obtained by the client.
	IdTokenTtl int32 `json:"id_token_ttl"`

	// A reference to a named key resource. Cannot be modified after creation. Defaults to the 'default' key.
	Key string `json:"key"`

	// Comma separated string or array of redirect URIs used by the client. One of these values must exactly match the redirect_uri parameter value used in each authentication request.
	RedirectUris []string `json:"redirect_uris"`
}

// NewOIDCWriteClientRequestWithDefaults instantiates a new OIDCWriteClientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWriteClientRequestWithDefaults() *OIDCWriteClientRequest {
	var this OIDCWriteClientRequest

	this.ClientType = "confidential"
	this.Key = "default"

	return &this
}

func (o OIDCWriteClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
