// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PluginsCatalogRegisterPluginWithTypeRequest struct for PluginsCatalogRegisterPluginWithTypeRequest
type PluginsCatalogRegisterPluginWithTypeRequest struct {
	// The args passed to plugin command.
	Args []string `json:"args"`

	// The command used to start the plugin. The executable defined in this command must exist in vault's plugin directory.
	Command string `json:"command"`

	// The environment variables passed to plugin command. Each entry is of the form \"key=value\".
	Env []string `json:"env"`

	// The SHA256 sum of the executable used in the command field. This should be HEX encoded.
	Sha256 string `json:"sha256"`

	// The semantic version of the plugin to use.
	Version string `json:"version"`
}

// NewPluginsCatalogRegisterPluginWithTypeRequestWithDefaults instantiates a new PluginsCatalogRegisterPluginWithTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginsCatalogRegisterPluginWithTypeRequestWithDefaults() *PluginsCatalogRegisterPluginWithTypeRequest {
	var this PluginsCatalogRegisterPluginWithTypeRequest

	return &this
}

func (o PluginsCatalogRegisterPluginWithTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["args"] = o.Args
	toSerialize["command"] = o.Command
	toSerialize["env"] = o.Env
	toSerialize["sha256"] = o.Sha256
	toSerialize["version"] = o.Version

	return json.Marshal(toSerialize)
}
