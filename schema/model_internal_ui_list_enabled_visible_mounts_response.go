// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// InternalUiListEnabledVisibleMountsResponse struct for InternalUiListEnabledVisibleMountsResponse
type InternalUiListEnabledVisibleMountsResponse struct {
	// auth mounts
	Auth map[string]interface{} `json:"auth"`

	// secret mounts
	Secret map[string]interface{} `json:"secret"`
}

// NewInternalUiListEnabledVisibleMountsResponseWithDefaults instantiates a new InternalUiListEnabledVisibleMountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalUiListEnabledVisibleMountsResponseWithDefaults() *InternalUiListEnabledVisibleMountsResponse {
	var this InternalUiListEnabledVisibleMountsResponse

	return &this
}

func (o InternalUiListEnabledVisibleMountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["auth"] = o.Auth
	toSerialize["secret"] = o.Secret

	return json.Marshal(toSerialize)
}
