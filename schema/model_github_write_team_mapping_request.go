// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GithubWriteTeamMappingRequest struct for GithubWriteTeamMappingRequest
type GithubWriteTeamMappingRequest struct {
	// Value for teams mapping
	Value string `json:"value,omitempty"`
}

// NewGithubWriteTeamMappingRequestWithDefaults instantiates a new GithubWriteTeamMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubWriteTeamMappingRequestWithDefaults() *GithubWriteTeamMappingRequest {
	var this GithubWriteTeamMappingRequest

	return &this
}
