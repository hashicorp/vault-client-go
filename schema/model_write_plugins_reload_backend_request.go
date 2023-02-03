// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// WritePluginsReloadBackendRequest struct for WritePluginsReloadBackendRequest
type WritePluginsReloadBackendRequest struct {
	// The mount paths of the plugin backends to reload.
	Mounts []string `json:"mounts"`

	// The name of the plugin to reload, as registered in the plugin catalog.
	Plugin string `json:"plugin"`

	Scope string `json:"scope"`
}

// NewWritePluginsReloadBackendRequestWithDefaults instantiates a new WritePluginsReloadBackendRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritePluginsReloadBackendRequestWithDefaults() *WritePluginsReloadBackendRequest {
	var this WritePluginsReloadBackendRequest

	return &this
}
