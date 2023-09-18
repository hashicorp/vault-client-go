// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TransitRestoreAndRenameKeyRequest struct for TransitRestoreAndRenameKeyRequest
type TransitRestoreAndRenameKeyRequest struct {
	// Backed up key data to be restored. This should be the output from the 'backup/' endpoint.
	Backup string `json:"backup,omitempty"`

	// If set and a key by the given name exists, force the restore operation and override the key.
	Force bool `json:"force,omitempty"`
}
