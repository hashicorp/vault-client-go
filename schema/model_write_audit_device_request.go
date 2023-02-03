// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// WriteAuditDeviceRequest struct for WriteAuditDeviceRequest
type WriteAuditDeviceRequest struct {
	// User-friendly description for this audit backend.

	Description string `json:"description"`

	// Mark the mount as a local mount, which is not replicated and is unaffected by replication.

	Local bool `json:"local"`

	// Configuration options for the audit backend.

	Options map[string]interface{} `json:"options"`

	// The type of the backend. Example: \"mysql\"

	Type string `json:"type"`
}

// NewWriteAuditDeviceRequestWithDefaults instantiates a new WriteAuditDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteAuditDeviceRequestWithDefaults() *WriteAuditDeviceRequest {
	var this WriteAuditDeviceRequest

	this.Local = false

	return &this
}

func (o WriteAuditDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["description"] = o.Description

	toSerialize["local"] = o.Local

	toSerialize["options"] = o.Options

	toSerialize["type"] = o.Type

	return json.Marshal(toSerialize)
}
