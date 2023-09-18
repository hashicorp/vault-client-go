// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PluginsCatalogRegisterPluginRequest struct for PluginsCatalogRegisterPluginRequest
type PluginsCatalogRegisterPluginRequest struct {
	// The args passed to plugin command.
	Args []string `json:"args,omitempty"`

	// The command used to start the plugin. The executable defined in this command must exist in vault's plugin directory.
	Command string `json:"command,omitempty"`

	// The environment variables passed to plugin command. Each entry is of the form \"key=value\".
	Env []string `json:"env,omitempty"`

	// The name of the OCI image to be run, without the tag or SHA256. Must already be present on the machine.
	OciImage string `json:"oci_image,omitempty"`

	// The SHA256 sum of the executable or container to be run. This should be HEX encoded.
	Sha256 string `json:"sha256,omitempty"`

	// The semantic version of the plugin to use, or image tag if oci_image is provided.
	Version string `json:"version,omitempty"`
}
