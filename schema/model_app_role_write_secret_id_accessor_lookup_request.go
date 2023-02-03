// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleWriteSecretIDAccessorLookupRequest struct for AppRoleWriteSecretIDAccessorLookupRequest
type AppRoleWriteSecretIDAccessorLookupRequest struct {
	// Accessor of the SecretID

	SecretIdAccessor string `json:"secret_id_accessor"`
}

// NewAppRoleWriteSecretIDAccessorLookupRequestWithDefaults instantiates a new AppRoleWriteSecretIDAccessorLookupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteSecretIDAccessorLookupRequestWithDefaults() *AppRoleWriteSecretIDAccessorLookupRequest {
	var this AppRoleWriteSecretIDAccessorLookupRequest

	return &this
}

func (o AppRoleWriteSecretIDAccessorLookupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["secret_id_accessor"] = o.SecretIdAccessor

	return json.Marshal(toSerialize)
}
