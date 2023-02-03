// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteMountRequest struct for WriteMountRequest
type WriteMountRequest struct {
	// Configuration for this mount, such as default_lease_ttl and max_lease_ttl.

	Config map[string]interface{} `json:"config"`

	// User-friendly description for this mount.

	Description string `json:"description"`

	// Whether to give the mount access to Vault's external entropy.

	ExternalEntropyAccess bool `json:"external_entropy_access"`

	// Mark the mount as a local mount, which is not replicated and is unaffected by replication.

	Local bool `json:"local"`

	// The options to pass into the backend. Should be a json object with string keys and values.

	Options map[string]interface{} `json:"options"`

	// Name of the plugin to mount based from the name registered in the plugin catalog.

	PluginName string `json:"plugin_name"`

	// The semantic version of the plugin to use.

	PluginVersion string `json:"plugin_version"`

	// Whether to turn on seal wrapping for the mount.

	SealWrap bool `json:"seal_wrap"`

	// The type of the backend. Example: \"passthrough\"

	Type string `json:"type"`
}

// NewWriteMountRequestWithDefaults instantiates a new WriteMountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteMountRequestWithDefaults() *WriteMountRequest {
	var this WriteMountRequest

	this.ExternalEntropyAccess = false

	this.Local = false

	this.SealWrap = false

	return &this
}

func (o WriteMountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["config"] = o.Config

	toSerialize["description"] = o.Description

	toSerialize["external_entropy_access"] = o.ExternalEntropyAccess

	toSerialize["local"] = o.Local

	toSerialize["options"] = o.Options

	toSerialize["plugin_name"] = o.PluginName

	toSerialize["plugin_version"] = o.PluginVersion

	toSerialize["seal_wrap"] = o.SealWrap

	toSerialize["type"] = o.Type

	return json.Marshal(toSerialize)
}
