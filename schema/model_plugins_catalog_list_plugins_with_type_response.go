// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PluginsCatalogListPluginsWithTypeResponse struct for PluginsCatalogListPluginsWithTypeResponse
type PluginsCatalogListPluginsWithTypeResponse struct {
	// List of plugin names in the catalog
	Keys []string `json:"keys,omitempty"`
}

// NewPluginsCatalogListPluginsWithTypeResponseWithDefaults instantiates a new PluginsCatalogListPluginsWithTypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginsCatalogListPluginsWithTypeResponseWithDefaults() *PluginsCatalogListPluginsWithTypeResponse {
	var this PluginsCatalogListPluginsWithTypeResponse

	return &this
}
