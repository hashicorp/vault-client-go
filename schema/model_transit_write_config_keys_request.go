// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TransitWriteConfigKeysRequest struct for TransitWriteConfigKeysRequest
type TransitWriteConfigKeysRequest struct {
	// Whether to allow automatic upserting (creation) of keys on the encrypt endpoint.
	DisableUpsert bool `json:"disable_upsert"`
}

// NewTransitWriteConfigKeysRequestWithDefaults instantiates a new TransitWriteConfigKeysRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitWriteConfigKeysRequestWithDefaults() *TransitWriteConfigKeysRequest {
	var this TransitWriteConfigKeysRequest

	return &this
}

func (o TransitWriteConfigKeysRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["disable_upsert"] = o.DisableUpsert

	return json.Marshal(toSerialize)
}
