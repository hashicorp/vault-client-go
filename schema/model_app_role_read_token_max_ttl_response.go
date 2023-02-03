// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadTokenMaxTTLResponse struct for AppRoleReadTokenMaxTTLResponse
type AppRoleReadTokenMaxTTLResponse struct {
	// The maximum lifetime of the generated token

	TokenMaxTtl int32 `json:"token_max_ttl"`
}

// NewAppRoleReadTokenMaxTTLResponseWithDefaults instantiates a new AppRoleReadTokenMaxTTLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadTokenMaxTTLResponseWithDefaults() *AppRoleReadTokenMaxTTLResponse {
	var this AppRoleReadTokenMaxTTLResponse

	return &this
}

func (o AppRoleReadTokenMaxTTLResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token_max_ttl"] = o.TokenMaxTtl

	return json.Marshal(toSerialize)
}
