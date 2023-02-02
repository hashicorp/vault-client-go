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

// GoogleCloudWriteRoleLabelsRequest struct for GoogleCloudWriteRoleLabelsRequest
type GoogleCloudWriteRoleLabelsRequest struct {
	// BoundLabels to add (in $key:$value)

	Add []string `json:"add"`

	// Label key values to remove

	Remove []string `json:"remove"`
}

// NewGoogleCloudWriteRoleLabelsRequestWithDefaults instantiates a new GoogleCloudWriteRoleLabelsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudWriteRoleLabelsRequestWithDefaults() *GoogleCloudWriteRoleLabelsRequest {
	var this GoogleCloudWriteRoleLabelsRequest

	return &this
}

func (o GoogleCloudWriteRoleLabelsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["add"] = o.Add

	toSerialize["remove"] = o.Remove

	return json.Marshal(toSerialize)
}
