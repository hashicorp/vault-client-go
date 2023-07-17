// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudEditLabelsForRoleRequest struct for GoogleCloudEditLabelsForRoleRequest
type GoogleCloudEditLabelsForRoleRequest struct {
	// BoundLabels to add (in $key:$value)
	Add []string `json:"add,omitempty"`

	// Label key values to remove
	Remove []string `json:"remove,omitempty"`
}
