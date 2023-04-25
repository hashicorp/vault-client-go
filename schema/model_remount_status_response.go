// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RemountStatusResponse struct for RemountStatusResponse
type RemountStatusResponse struct {
	MigrationId string `json:"migration_id"`

	MigrationInfo map[string]interface{} `json:"migration_info"`
}

// NewRemountStatusResponseWithDefaults instantiates a new RemountStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemountStatusResponseWithDefaults() *RemountStatusResponse {
	var this RemountStatusResponse

	return &this
}

func (o RemountStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["migration_id"] = o.MigrationId
	toSerialize["migration_info"] = o.MigrationInfo

	return json.Marshal(toSerialize)
}
