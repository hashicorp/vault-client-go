// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AuthReadConfigurationResponse struct for AuthReadConfigurationResponse
type AuthReadConfigurationResponse struct {
	Accessor string `json:"accessor"`

	Config map[string]interface{} `json:"config"`

	DeprecationStatus string `json:"deprecation_status"`

	Description string `json:"description"`

	ExternalEntropyAccess bool `json:"external_entropy_access"`

	Local bool `json:"local"`

	Options map[string]interface{} `json:"options"`

	PluginVersion string `json:"plugin_version"`

	RunningPluginVersion string `json:"running_plugin_version"`

	RunningSha256 string `json:"running_sha256"`

	SealWrap bool `json:"seal_wrap"`

	Type string `json:"type"`

	Uuid string `json:"uuid"`
}

// NewAuthReadConfigurationResponseWithDefaults instantiates a new AuthReadConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthReadConfigurationResponseWithDefaults() *AuthReadConfigurationResponse {
	var this AuthReadConfigurationResponse

	return &this
}

func (o AuthReadConfigurationResponse) MarshalJSON() ([]byte, error) {
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
