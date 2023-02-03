// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteSecretIDAccessorDestroyRequest struct for AppRoleWriteSecretIDAccessorDestroyRequest
type AppRoleWriteSecretIDAccessorDestroyRequest struct {
	// Accessor of the SecretID
	SecretIdAccessor string `json:"secret_id_accessor"`
}

// NewAppRoleWriteSecretIDAccessorDestroyRequestWithDefaults instantiates a new AppRoleWriteSecretIDAccessorDestroyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteSecretIDAccessorDestroyRequestWithDefaults() *AppRoleWriteSecretIDAccessorDestroyRequest {
	var this AppRoleWriteSecretIDAccessorDestroyRequest

	return &this
}

func (o AppRoleWriteSecretIDAccessorDestroyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
