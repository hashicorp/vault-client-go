// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AuditingEnableDeviceRequest struct for AuditingEnableDeviceRequest
type AuditingEnableDeviceRequest struct {
	// User-friendly description for this audit backend.
	Description string `json:"description,omitempty"`

	// Mark the mount as a local mount, which is not replicated and is unaffected by replication.
	Local bool `json:"local,omitempty"`

	// Configuration options for the audit backend.
	Options map[string]interface{} `json:"options,omitempty"`

	// The type of the backend. Example: \"mysql\"
	Type string `json:"type,omitempty"`
}
