// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteBindSecretIDRequest struct for AppRoleWriteBindSecretIDRequest
type AppRoleWriteBindSecretIDRequest struct {
	// Impose secret_id to be presented when logging in using this role.

	BindSecretId bool `json:"bind_secret_id"`
}

// NewAppRoleWriteBindSecretIDRequestWithDefaults instantiates a new AppRoleWriteBindSecretIDRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteBindSecretIDRequestWithDefaults() *AppRoleWriteBindSecretIDRequest {
	var this AppRoleWriteBindSecretIDRequest

	this.BindSecretId = true

	return &this
}

func (o AppRoleWriteBindSecretIDRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["bind_secret_id"] = o.BindSecretId

	return json.Marshal(toSerialize)
}
