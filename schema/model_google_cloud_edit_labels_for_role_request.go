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

// NewGoogleCloudEditLabelsForRoleRequestWithDefaults instantiates a new GoogleCloudEditLabelsForRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudEditLabelsForRoleRequestWithDefaults() *GoogleCloudEditLabelsForRoleRequest {
	var this GoogleCloudEditLabelsForRoleRequest

	return &this
}
