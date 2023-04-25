// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GithubWriteUserMappingRequest struct for GithubWriteUserMappingRequest
type GithubWriteUserMappingRequest struct {
	// Value for users mapping
	Value string `json:"value"`
}

// NewGithubWriteUserMappingRequestWithDefaults instantiates a new GithubWriteUserMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubWriteUserMappingRequestWithDefaults() *GithubWriteUserMappingRequest {
	var this GithubWriteUserMappingRequest

	return &this
}

func (o GithubWriteUserMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["value"] = o.Value

	return json.Marshal(toSerialize)
}
