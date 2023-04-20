// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GoogleCloudEditServiceAccountsForRoleRequest struct for GoogleCloudEditServiceAccountsForRoleRequest
type GoogleCloudEditServiceAccountsForRoleRequest struct {
	// Service-account emails or IDs to add.
	Add []string `json:"add"`

	// Service-account emails or IDs to remove.
	Remove []string `json:"remove"`
}

// NewGoogleCloudEditServiceAccountsForRoleRequestWithDefaults instantiates a new GoogleCloudEditServiceAccountsForRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudEditServiceAccountsForRoleRequestWithDefaults() *GoogleCloudEditServiceAccountsForRoleRequest {
	var this GoogleCloudEditServiceAccountsForRoleRequest

	return &this
}

func (o GoogleCloudEditServiceAccountsForRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["add"] = o.Add
	toSerialize["remove"] = o.Remove

	return json.Marshal(toSerialize)
}
