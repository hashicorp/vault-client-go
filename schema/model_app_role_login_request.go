/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleLoginRequest struct for AppRoleLoginRequest
type AppRoleLoginRequest struct {
	// Unique identifier of the Role. Required to be supplied when the 'bind_secret_id' constraint is set.
	RoleId string `json:"role_id"`
	// SecretID belong to the App role
	SecretId string `json:"secret_id"`
}

// NewAppRoleLoginRequestWithDefaults instantiates a new AppRoleLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleLoginRequestWithDefaults() *AppRoleLoginRequest {
	var this AppRoleLoginRequest

	this.SecretId = ""

	return &this
}

func (o AppRoleLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["role_id"] = o.RoleId
	toSerialize["secret_id"] = o.SecretId

	return json.Marshal(toSerialize)
}
