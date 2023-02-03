// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteSecretIDLookupRequest struct for AppRoleWriteSecretIDLookupRequest
type AppRoleWriteSecretIDLookupRequest struct {
	// SecretID attached to the role.

	SecretId string `json:"secret_id"`
}

// NewAppRoleWriteSecretIDLookupRequestWithDefaults instantiates a new AppRoleWriteSecretIDLookupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteSecretIDLookupRequestWithDefaults() *AppRoleWriteSecretIDLookupRequest {
	var this AppRoleWriteSecretIDLookupRequest

	return &this
}

func (o AppRoleWriteSecretIDLookupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["secret_id"] = o.SecretId

	return json.Marshal(toSerialize)
}
