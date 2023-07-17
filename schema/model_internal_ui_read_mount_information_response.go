// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// InternalUiReadMountInformationResponse struct for InternalUiReadMountInformationResponse
type InternalUiReadMountInformationResponse struct {
	Accessor string `json:"accessor,omitempty"`

	Config map[string]interface{} `json:"config,omitempty"`

	Description string `json:"description,omitempty"`

	ExternalEntropyAccess bool `json:"external_entropy_access,omitempty"`

	Local bool `json:"local,omitempty"`

	Options map[string]interface{} `json:"options,omitempty"`

	Path string `json:"path,omitempty"`

	PluginVersion string `json:"plugin_version,omitempty"`

	RunningPluginVersion string `json:"running_plugin_version,omitempty"`

	RunningSha256 string `json:"running_sha256,omitempty"`

	SealWrap bool `json:"seal_wrap,omitempty"`

	Type string `json:"type,omitempty"`

	Uuid string `json:"uuid,omitempty"`
}
