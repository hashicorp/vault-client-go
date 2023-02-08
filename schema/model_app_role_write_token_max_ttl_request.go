// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteTokenMaxTTLRequest struct for AppRoleWriteTokenMaxTTLRequest
type AppRoleWriteTokenMaxTTLRequest struct {
	// The maximum lifetime of the generated token
	TokenMaxTtl int32 `json:"token_max_ttl"`
}

// NewAppRoleWriteTokenMaxTTLRequestWithDefaults instantiates a new AppRoleWriteTokenMaxTTLRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteTokenMaxTTLRequestWithDefaults() *AppRoleWriteTokenMaxTTLRequest {
	var this AppRoleWriteTokenMaxTTLRequest

	return &this
}

func (o AppRoleWriteTokenMaxTTLRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
