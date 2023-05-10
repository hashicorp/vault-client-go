// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GithubLoginRequest struct for GithubLoginRequest
type GithubLoginRequest struct {
	// GitHub personal API token
	Token string `json:"token"`
}

// NewGithubLoginRequestWithDefaults instantiates a new GithubLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubLoginRequestWithDefaults() *GithubLoginRequest {
	var this GithubLoginRequest

	return &this
}

func (o GithubLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token"] = o.Token

	return json.Marshal(toSerialize)
}
