// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteTokenTTLRequest struct for AppRoleWriteTokenTTLRequest
type AppRoleWriteTokenTTLRequest struct {
	// The initial ttl of the token to generate
	TokenTtl int32 `json:"token_ttl"`
}

// NewAppRoleWriteTokenTTLRequestWithDefaults instantiates a new AppRoleWriteTokenTTLRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteTokenTTLRequestWithDefaults() *AppRoleWriteTokenTTLRequest {
	var this AppRoleWriteTokenTTLRequest

	return &this
}

func (o AppRoleWriteTokenTTLRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
