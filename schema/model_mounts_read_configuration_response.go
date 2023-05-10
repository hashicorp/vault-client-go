// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// MountsReadConfigurationResponse struct for MountsReadConfigurationResponse
type MountsReadConfigurationResponse struct {
	Accessor string `json:"accessor"`

	// Configuration for this mount, such as default_lease_ttl and max_lease_ttl.
	Config map[string]interface{} `json:"config"`

	DeprecationStatus string `json:"deprecation_status"`

	// User-friendly description for this mount.
	Description string `json:"description"`

	ExternalEntropyAccess bool `json:"external_entropy_access"`

	// Mark the mount as a local mount, which is not replicated and is unaffected by replication.
	Local bool `json:"local"`

	// The options to pass into the backend. Should be a json object with string keys and values.
	Options map[string]interface{} `json:"options"`

	// The semantic version of the plugin to use.
	PluginVersion string `json:"plugin_version"`

	RunningPluginVersion string `json:"running_plugin_version"`

	RunningSha256 string `json:"running_sha256"`

	// Whether to turn on seal wrapping for the mount.
	SealWrap bool `json:"seal_wrap"`

	// The type of the backend. Example: \"passthrough\"
	Type string `json:"type"`

	Uuid string `json:"uuid"`
}

// NewMountsReadConfigurationResponseWithDefaults instantiates a new MountsReadConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMountsReadConfigurationResponseWithDefaults() *MountsReadConfigurationResponse {
	var this MountsReadConfigurationResponse

	this.Local = false
	this.SealWrap = false

	return &this
}

func (o MountsReadConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["accessor"] = o.Accessor
	toSerialize["config"] = o.Config
	toSerialize["deprecation_status"] = o.DeprecationStatus
	toSerialize["description"] = o.Description
	toSerialize["external_entropy_access"] = o.ExternalEntropyAccess
	toSerialize["local"] = o.Local
	toSerialize["options"] = o.Options
	toSerialize["plugin_version"] = o.PluginVersion
	toSerialize["running_plugin_version"] = o.RunningPluginVersion
	toSerialize["running_sha256"] = o.RunningSha256
	toSerialize["seal_wrap"] = o.SealWrap
	toSerialize["type"] = o.Type
	toSerialize["uuid"] = o.Uuid

	return json.Marshal(toSerialize)
}
