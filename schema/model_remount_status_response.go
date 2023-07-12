// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// RemountStatusResponse struct for RemountStatusResponse
type RemountStatusResponse struct {
	MigrationId string `json:"migration_id,omitempty"`

	MigrationInfo map[string]interface{} `json:"migration_info,omitempty"`
}
