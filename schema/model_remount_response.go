// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RemountResponse struct for RemountResponse
type RemountResponse struct {
	MigrationId string `json:"migration_id,omitempty"`
}

// NewRemountResponseWithDefaults instantiates a new RemountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemountResponseWithDefaults() *RemountResponse {
	var this RemountResponse

	return &this
}
