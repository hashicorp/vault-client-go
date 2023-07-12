// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitRestoreKeyRequest struct for TransitRestoreKeyRequest
type TransitRestoreKeyRequest struct {
	// Backed up key data to be restored. This should be the output from the 'backup/' endpoint.
	Backup string `json:"backup,omitempty"`

	// If set and a key by the given name exists, force the restore operation and override the key.
	Force bool `json:"force,omitempty"`

	// If set, this will be the name of the restored key.
	Name string `json:"name,omitempty"`
}
