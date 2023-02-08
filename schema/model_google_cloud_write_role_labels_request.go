// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GoogleCloudWriteRoleLabelsRequest struct for GoogleCloudWriteRoleLabelsRequest
type GoogleCloudWriteRoleLabelsRequest struct {
	// BoundLabels to add (in $key:$value)
	Add []string `json:"add"`

	// Label key values to remove
	Remove []string `json:"remove"`
}

// NewGoogleCloudWriteRoleLabelsRequestWithDefaults instantiates a new GoogleCloudWriteRoleLabelsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudWriteRoleLabelsRequestWithDefaults() *GoogleCloudWriteRoleLabelsRequest {
	var this GoogleCloudWriteRoleLabelsRequest

	return &this
}

func (o GoogleCloudWriteRoleLabelsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
