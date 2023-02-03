// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// NomadWriteRoleRequest struct for NomadWriteRoleRequest
type NomadWriteRoleRequest struct {
	// Boolean value describing if the token should be global or not. Defaults to false.

	Global bool `json:"global"`

	// Comma-separated string or list of policies as previously created in Nomad. Required for 'client' token.

	Policies []string `json:"policies"`

	// Which type of token to create: 'client' or 'management'. If a 'management' token, the \"policies\" parameter is not required. Defaults to 'client'.

	Type string `json:"type"`
}

// NewNomadWriteRoleRequestWithDefaults instantiates a new NomadWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNomadWriteRoleRequestWithDefaults() *NomadWriteRoleRequest {
	var this NomadWriteRoleRequest

	this.Type = "client"

	return &this
}

func (o NomadWriteRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["global"] = o.Global

	toSerialize["policies"] = o.Policies

	toSerialize["type"] = o.Type

	return json.Marshal(toSerialize)
}
