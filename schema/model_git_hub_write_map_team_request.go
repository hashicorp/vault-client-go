// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// GitHubWriteMapTeamRequest struct for GitHubWriteMapTeamRequest
type GitHubWriteMapTeamRequest struct {
	// Value for teams mapping
	Value string `json:"value"`
}

// NewGitHubWriteMapTeamRequestWithDefaults instantiates a new GitHubWriteMapTeamRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitHubWriteMapTeamRequestWithDefaults() *GitHubWriteMapTeamRequest {
	var this GitHubWriteMapTeamRequest

	return &this
}

func (o GitHubWriteMapTeamRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
