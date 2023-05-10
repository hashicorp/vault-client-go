// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AuthReadConfigurationResponse struct for AuthReadConfigurationResponse
type AuthReadConfigurationResponse struct {
	Accessor string `json:"accessor,omitempty"`

	Config map[string]interface{} `json:"config,omitempty"`

	DeprecationStatus string `json:"deprecation_status,omitempty"`

	Description string `json:"description,omitempty"`

	ExternalEntropyAccess bool `json:"external_entropy_access,omitempty"`

	Local bool `json:"local,omitempty"`

	Options map[string]interface{} `json:"options,omitempty"`

	PluginVersion string `json:"plugin_version,omitempty"`

	RunningPluginVersion string `json:"running_plugin_version,omitempty"`

	RunningSha256 string `json:"running_sha256,omitempty"`

	SealWrap bool `json:"seal_wrap,omitempty"`

	Type string `json:"type,omitempty"`

	Uuid string `json:"uuid,omitempty"`
}

// NewAuthReadConfigurationResponseWithDefaults instantiates a new AuthReadConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthReadConfigurationResponseWithDefaults() *AuthReadConfigurationResponse {
	var this AuthReadConfigurationResponse

	return &this
}
