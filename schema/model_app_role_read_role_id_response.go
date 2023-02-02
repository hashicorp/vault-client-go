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

// AppRoleReadRoleIDResponse struct for AppRoleReadRoleIDResponse
type AppRoleReadRoleIDResponse struct {
	// Identifier of the role. Defaults to a UUID.

	RoleId string `json:"role_id"`
}

// NewAppRoleReadRoleIDResponseWithDefaults instantiates a new AppRoleReadRoleIDResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadRoleIDResponseWithDefaults() *AppRoleReadRoleIDResponse {
	var this AppRoleReadRoleIDResponse

	return &this
}

func (o AppRoleReadRoleIDResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["role_id"] = o.RoleId

	return json.Marshal(toSerialize)
}
