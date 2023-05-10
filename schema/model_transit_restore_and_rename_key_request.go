// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitRestoreAndRenameKeyRequest struct for TransitRestoreAndRenameKeyRequest
type TransitRestoreAndRenameKeyRequest struct {
	// Backed up key data to be restored. This should be the output from the 'backup/' endpoint.
	Backup string `json:"backup"`

	// If set and a key by the given name exists, force the restore operation and override the key.
	Force bool `json:"force"`
}

// NewTransitRestoreAndRenameKeyRequestWithDefaults instantiates a new TransitRestoreAndRenameKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitRestoreAndRenameKeyRequestWithDefaults() *TransitRestoreAndRenameKeyRequest {
	var this TransitRestoreAndRenameKeyRequest

	this.Force = false

	return &this
}

func (o TransitRestoreAndRenameKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["backup"] = o.Backup
	toSerialize["force"] = o.Force

	return json.Marshal(toSerialize)
}
