// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AuditingEnableDeviceRequest struct for AuditingEnableDeviceRequest
type AuditingEnableDeviceRequest struct {
	// User-friendly description for this audit backend.
	Description string `json:"description"`

	// Mark the mount as a local mount, which is not replicated and is unaffected by replication.
	Local bool `json:"local"`

	// Configuration options for the audit backend.
	Options map[string]interface{} `json:"options"`

	// The type of the backend. Example: \"mysql\"
	Type string `json:"type"`
}

// NewAuditingEnableDeviceRequestWithDefaults instantiates a new AuditingEnableDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditingEnableDeviceRequestWithDefaults() *AuditingEnableDeviceRequest {
	var this AuditingEnableDeviceRequest

	this.Local = false

	return &this
}

func (o AuditingEnableDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["description"] = o.Description
	toSerialize["local"] = o.Local
	toSerialize["options"] = o.Options
	toSerialize["type"] = o.Type

	return json.Marshal(toSerialize)
}
