// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PluginsCatalogReadPluginConfigurationResponse struct for PluginsCatalogReadPluginConfigurationResponse
type PluginsCatalogReadPluginConfigurationResponse struct {
	// The args passed to plugin command.
	Args []string `json:"args"`

	Builtin bool `json:"builtin"`

	// The command used to start the plugin. The executable defined in this command must exist in vault's plugin directory.
	Command string `json:"command"`

	DeprecationStatus string `json:"deprecation_status"`

	// The name of the plugin
	Name string `json:"name"`

	// The SHA256 sum of the executable used in the command field. This should be HEX encoded.
	Sha256 string `json:"sha256"`

	// The semantic version of the plugin to use.
	Version string `json:"version"`
}

// NewPluginsCatalogReadPluginConfigurationResponseWithDefaults instantiates a new PluginsCatalogReadPluginConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginsCatalogReadPluginConfigurationResponseWithDefaults() *PluginsCatalogReadPluginConfigurationResponse {
	var this PluginsCatalogReadPluginConfigurationResponse

	return &this
}

func (o PluginsCatalogReadPluginConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["args"] = o.Args
	toSerialize["builtin"] = o.Builtin
	toSerialize["command"] = o.Command
	toSerialize["deprecation_status"] = o.DeprecationStatus
	toSerialize["name"] = o.Name
	toSerialize["sha256"] = o.Sha256
	toSerialize["version"] = o.Version

	return json.Marshal(toSerialize)
}
