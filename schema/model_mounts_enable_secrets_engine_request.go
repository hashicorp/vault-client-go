// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// MountsEnableSecretsEngineRequest struct for MountsEnableSecretsEngineRequest
type MountsEnableSecretsEngineRequest struct {
	// Configuration for this mount, such as default_lease_ttl and max_lease_ttl.
	Config map[string]interface{} `json:"config,omitempty"`

	// User-friendly description for this mount.
	Description string `json:"description,omitempty"`

	// Whether to give the mount access to Vault's external entropy.
	ExternalEntropyAccess bool `json:"external_entropy_access,omitempty"`

	// Mark the mount as a local mount, which is not replicated and is unaffected by replication.
	Local bool `json:"local,omitempty"`

	// The options to pass into the backend. Should be a json object with string keys and values.
	Options map[string]interface{} `json:"options,omitempty"`

	// Name of the plugin to mount based from the name registered in the plugin catalog.
	PluginName string `json:"plugin_name,omitempty"`

	// The semantic version of the plugin to use, or image tag if oci_image is provided.
	PluginVersion string `json:"plugin_version,omitempty"`

	// Whether to turn on seal wrapping for the mount.
	SealWrap bool `json:"seal_wrap,omitempty"`

	// The type of the backend. Example: \"passthrough\"
	Type string `json:"type,omitempty"`
}
