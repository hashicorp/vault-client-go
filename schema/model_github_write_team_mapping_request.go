// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GithubWriteTeamMappingRequest struct for GithubWriteTeamMappingRequest
type GithubWriteTeamMappingRequest struct {
	// Value for teams mapping
	Value string `json:"value"`
}

// NewGithubWriteTeamMappingRequestWithDefaults instantiates a new GithubWriteTeamMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubWriteTeamMappingRequestWithDefaults() *GithubWriteTeamMappingRequest {
	var this GithubWriteTeamMappingRequest

	return &this
}

func (o GithubWriteTeamMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["value"] = o.Value

	return json.Marshal(toSerialize)
}
