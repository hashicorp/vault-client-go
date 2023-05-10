// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PluginsReloadBackendsResponse struct for PluginsReloadBackendsResponse
type PluginsReloadBackendsResponse struct {
	ReloadId string `json:"reload_id,omitempty"`
}

// NewPluginsReloadBackendsResponseWithDefaults instantiates a new PluginsReloadBackendsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginsReloadBackendsResponseWithDefaults() *PluginsReloadBackendsResponse {
	var this PluginsReloadBackendsResponse

	return &this
}
