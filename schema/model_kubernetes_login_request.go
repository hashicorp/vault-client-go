// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KubernetesLoginRequest struct for KubernetesLoginRequest
type KubernetesLoginRequest struct {
	// A signed JWT for authenticating a service account. This field is required.

	Jwt string `json:"jwt"`

	// Name of the role against which the login is being attempted. This field is required

	Role string `json:"role"`
}

// NewKubernetesLoginRequestWithDefaults instantiates a new KubernetesLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesLoginRequestWithDefaults() *KubernetesLoginRequest {
	var this KubernetesLoginRequest

	return &this
}

func (o KubernetesLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["jwt"] = o.Jwt

	toSerialize["role"] = o.Role

	return json.Marshal(toSerialize)
}
