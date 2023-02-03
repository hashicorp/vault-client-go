// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadTokenNumUsesResponse struct for AppRoleReadTokenNumUsesResponse
type AppRoleReadTokenNumUsesResponse struct {
	// The maximum number of times a token may be used, a value of zero means unlimited

	TokenNumUses int32 `json:"token_num_uses"`
}

// NewAppRoleReadTokenNumUsesResponseWithDefaults instantiates a new AppRoleReadTokenNumUsesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadTokenNumUsesResponseWithDefaults() *AppRoleReadTokenNumUsesResponse {
	var this AppRoleReadTokenNumUsesResponse

	return &this
}

func (o AppRoleReadTokenNumUsesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token_num_uses"] = o.TokenNumUses

	return json.Marshal(toSerialize)
}
