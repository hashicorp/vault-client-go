// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PluginsReloadBackendsRequest struct for PluginsReloadBackendsRequest
type PluginsReloadBackendsRequest struct {
	// The mount paths of the plugin backends to reload.
	Mounts []string `json:"mounts,omitempty"`

	// The name of the plugin to reload, as registered in the plugin catalog.
	Plugin string `json:"plugin,omitempty"`

	Scope string `json:"scope,omitempty"`
}

// NewPluginsReloadBackendsRequestWithDefaults instantiates a new PluginsReloadBackendsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginsReloadBackendsRequestWithDefaults() *PluginsReloadBackendsRequest {
	var this PluginsReloadBackendsRequest

	return &this
}
