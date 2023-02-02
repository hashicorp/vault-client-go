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

// GoogleCloudLoginRequest struct for GoogleCloudLoginRequest
type GoogleCloudLoginRequest struct {
	// A signed JWT. This is either a self-signed service account JWT ('iam' roles only) or a GCE identity metadata token ('iam', 'gce' roles).

	Jwt string `json:"jwt"`

	// Name of the role against which the login is being attempted. Required.

	Role string `json:"role"`
}

// NewGoogleCloudLoginRequestWithDefaults instantiates a new GoogleCloudLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudLoginRequestWithDefaults() *GoogleCloudLoginRequest {
	var this GoogleCloudLoginRequest

	return &this
}

func (o GoogleCloudLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["jwt"] = o.Jwt

	toSerialize["role"] = o.Role

	return json.Marshal(toSerialize)
}
