// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitTrimKeyRequest struct for TransitTrimKeyRequest
type TransitTrimKeyRequest struct {
	// The minimum available version for the key ring. All versions before this version will be permanently deleted. This value can at most be equal to the lesser of 'min_decryption_version' and 'min_encryption_version'. This is not allowed to be set when either 'min_encryption_version' or 'min_decryption_version' is set to zero.
	MinAvailableVersion int32 `json:"min_available_version"`
}

// NewTransitTrimKeyRequestWithDefaults instantiates a new TransitTrimKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitTrimKeyRequestWithDefaults() *TransitTrimKeyRequest {
	var this TransitTrimKeyRequest

	return &this
}

func (o TransitTrimKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
