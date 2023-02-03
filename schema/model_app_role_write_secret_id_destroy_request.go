// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteSecretIDDestroyRequest struct for AppRoleWriteSecretIDDestroyRequest
type AppRoleWriteSecretIDDestroyRequest struct {
	// SecretID attached to the role.
	SecretId string `json:"secret_id"`
}

// NewAppRoleWriteSecretIDDestroyRequestWithDefaults instantiates a new AppRoleWriteSecretIDDestroyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteSecretIDDestroyRequestWithDefaults() *AppRoleWriteSecretIDDestroyRequest {
	var this AppRoleWriteSecretIDDestroyRequest

	return &this
}
