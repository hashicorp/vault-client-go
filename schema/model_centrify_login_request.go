// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// CentrifyLoginRequest struct for CentrifyLoginRequest
type CentrifyLoginRequest struct {
	// Auth mode ('ro' for resource owner, 'cc' for credential client).
	Mode string `json:"mode"`

	// Password for this user.
	Password string `json:"password"`

	// Username of the user.
	Username string `json:"username"`
}

// NewCentrifyLoginRequestWithDefaults instantiates a new CentrifyLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentrifyLoginRequestWithDefaults() *CentrifyLoginRequest {
	var this CentrifyLoginRequest

	this.Mode = "ro"

	return &this
}

func (o CentrifyLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
