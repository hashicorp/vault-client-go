// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadTokenTTLResponse struct for AppRoleReadTokenTTLResponse
type AppRoleReadTokenTTLResponse struct {
	// The initial ttl of the token to generate
	TokenTtl int32 `json:"token_ttl"`
}

// NewAppRoleReadTokenTTLResponseWithDefaults instantiates a new AppRoleReadTokenTTLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadTokenTTLResponseWithDefaults() *AppRoleReadTokenTTLResponse {
	var this AppRoleReadTokenTTLResponse

	return &this
}

func (o AppRoleReadTokenTTLResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
