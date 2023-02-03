// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// SysWritePluginsCatalogNameRequest struct for SysWritePluginsCatalogNameRequest
type SysWritePluginsCatalogNameRequest struct {
	// The args passed to plugin command.
	Args []string `json:"args"`

	// The command used to start the plugin. The executable defined in this command must exist in vault's plugin directory.
	Command string `json:"command"`

	// The environment variables passed to plugin command. Each entry is of the form \"key=value\".
	Env []string `json:"env"`

	// The SHA256 sum of the executable used in the command field. This should be HEX encoded.
	Sha256 string `json:"sha256"`

	// The type of the plugin, may be auth, secret, or database
	Type string `json:"type"`

	// The semantic version of the plugin to use.
	Version string `json:"version"`
}

// NewSysWritePluginsCatalogNameRequestWithDefaults instantiates a new SysWritePluginsCatalogNameRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSysWritePluginsCatalogNameRequestWithDefaults() *SysWritePluginsCatalogNameRequest {
	var this SysWritePluginsCatalogNameRequest

	return &this
}
