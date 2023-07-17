// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudEditServiceAccountsForRoleRequest struct for GoogleCloudEditServiceAccountsForRoleRequest
type GoogleCloudEditServiceAccountsForRoleRequest struct {
	// Service-account emails or IDs to add.
	Add []string `json:"add,omitempty"`

	// Service-account emails or IDs to remove.
	Remove []string `json:"remove,omitempty"`
}
