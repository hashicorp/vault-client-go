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
	return json.Marshal(o)
}
