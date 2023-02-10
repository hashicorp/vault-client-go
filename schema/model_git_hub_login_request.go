// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GitHubLoginRequest struct for GitHubLoginRequest
type GitHubLoginRequest struct {
	// GitHub personal API token
	Token string `json:"token"`
}

// NewGitHubLoginRequestWithDefaults instantiates a new GitHubLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitHubLoginRequestWithDefaults() *GitHubLoginRequest {
	var this GitHubLoginRequest

	return &this
}

func (o GitHubLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token"] = o.Token

	return json.Marshal(toSerialize)
}
