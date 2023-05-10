// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitRotateKeyRequest struct for TransitRotateKeyRequest
type TransitRotateKeyRequest struct {
	// The UUID of the managed key to use for the new version of this transit key
	ManagedKeyId string `json:"managed_key_id"`

	// The name of the managed key to use for the new version of this transit key
	ManagedKeyName string `json:"managed_key_name"`
}

// NewTransitRotateKeyRequestWithDefaults instantiates a new TransitRotateKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitRotateKeyRequestWithDefaults() *TransitRotateKeyRequest {
	var this TransitRotateKeyRequest

	return &this
}

func (o TransitRotateKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["managed_key_id"] = o.ManagedKeyId
	toSerialize["managed_key_name"] = o.ManagedKeyName

	return json.Marshal(toSerialize)
}
