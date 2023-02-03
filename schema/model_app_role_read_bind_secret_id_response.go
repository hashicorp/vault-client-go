// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadBindSecretIDResponse struct for AppRoleReadBindSecretIDResponse
type AppRoleReadBindSecretIDResponse struct {
	// Impose secret_id to be presented when logging in using this role. Defaults to 'true'.
	BindSecretId bool `json:"bind_secret_id"`
}

// NewAppRoleReadBindSecretIDResponseWithDefaults instantiates a new AppRoleReadBindSecretIDResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadBindSecretIDResponseWithDefaults() *AppRoleReadBindSecretIDResponse {
	var this AppRoleReadBindSecretIDResponse

	return &this
}

func (o AppRoleReadBindSecretIDResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
