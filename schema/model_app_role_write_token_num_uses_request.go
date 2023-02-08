// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteTokenNumUsesRequest struct for AppRoleWriteTokenNumUsesRequest
type AppRoleWriteTokenNumUsesRequest struct {
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses int32 `json:"token_num_uses"`
}

// NewAppRoleWriteTokenNumUsesRequestWithDefaults instantiates a new AppRoleWriteTokenNumUsesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteTokenNumUsesRequestWithDefaults() *AppRoleWriteTokenNumUsesRequest {
	var this AppRoleWriteTokenNumUsesRequest

	return &this
}

func (o AppRoleWriteTokenNumUsesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
