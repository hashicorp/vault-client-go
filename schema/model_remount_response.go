// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RemountResponse struct for RemountResponse
type RemountResponse struct {
	MigrationId string `json:"migration_id"`
}

// NewRemountResponseWithDefaults instantiates a new RemountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemountResponseWithDefaults() *RemountResponse {
	var this RemountResponse

	return &this
}

func (o RemountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["migration_id"] = o.MigrationId

	return json.Marshal(toSerialize)
}
