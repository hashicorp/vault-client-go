// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// MountsReadConfigurationResponse struct for MountsReadConfigurationResponse
type MountsReadConfigurationResponse struct {
	Accessor string `json:"accessor,omitempty"`

	// Configuration for this mount, such as default_lease_ttl and max_lease_ttl.
	Config map[string]interface{} `json:"config,omitempty"`

	DeprecationStatus string `json:"deprecation_status,omitempty"`

	// User-friendly description for this mount.
	Description string `json:"description,omitempty"`

	ExternalEntropyAccess bool `json:"external_entropy_access,omitempty"`

	// Mark the mount as a local mount, which is not replicated and is unaffected by replication.
	Local bool `json:"local,omitempty"`

	// The options to pass into the backend. Should be a json object with string keys and values.
	Options map[string]interface{} `json:"options,omitempty"`

	// The semantic version of the plugin to use, or image tag if oci_image is provided.
	PluginVersion string `json:"plugin_version,omitempty"`

	RunningPluginVersion string `json:"running_plugin_version,omitempty"`

	RunningSha256 string `json:"running_sha256,omitempty"`

	// Whether to turn on seal wrapping for the mount.
	SealWrap bool `json:"seal_wrap,omitempty"`

	// The type of the backend. Example: \"passthrough\"
	Type string `json:"type,omitempty"`

	Uuid string `json:"uuid,omitempty"`
}
