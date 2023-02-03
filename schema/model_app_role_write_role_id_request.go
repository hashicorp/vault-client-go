// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteRoleIDRequest struct for AppRoleWriteRoleIDRequest
type AppRoleWriteRoleIDRequest struct {
	// Identifier of the role. Defaults to a UUID.
	RoleId string `json:"role_id"`
}

// NewAppRoleWriteRoleIDRequestWithDefaults instantiates a new AppRoleWriteRoleIDRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteRoleIDRequestWithDefaults() *AppRoleWriteRoleIDRequest {
	var this AppRoleWriteRoleIDRequest

	return &this
}

func (o AppRoleWriteRoleIDRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
