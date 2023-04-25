// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PluginsReloadBackendsResponse struct for PluginsReloadBackendsResponse
type PluginsReloadBackendsResponse struct {
	ReloadId string `json:"reload_id"`
}

// NewPluginsReloadBackendsResponseWithDefaults instantiates a new PluginsReloadBackendsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginsReloadBackendsResponseWithDefaults() *PluginsReloadBackendsResponse {
	var this PluginsReloadBackendsResponse

	return &this
}

func (o PluginsReloadBackendsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["reload_id"] = o.ReloadId

	return json.Marshal(toSerialize)
}
