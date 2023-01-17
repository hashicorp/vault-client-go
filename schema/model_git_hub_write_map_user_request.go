/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GitHubWriteMapUserRequest struct for GitHubWriteMapUserRequest
type GitHubWriteMapUserRequest struct {
	// Value for users mapping
	Value string `json:"value"`
}

// NewGitHubWriteMapUserRequestWithDefaults instantiates a new GitHubWriteMapUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitHubWriteMapUserRequestWithDefaults() *GitHubWriteMapUserRequest {
	var this GitHubWriteMapUserRequest

	return &this
}

func (o GitHubWriteMapUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["value"] = o.Value

	return json.Marshal(toSerialize)
}
