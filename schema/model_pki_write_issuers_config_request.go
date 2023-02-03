// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIWriteIssuersConfigRequest struct for PKIWriteIssuersConfigRequest
type PKIWriteIssuersConfigRequest struct {
	// Reference (name or identifier) to the default issuer.

	Default string `json:"default"`

	// Whether the default issuer should automatically follow the latest generated or imported issuer. Defaults to false.

	DefaultFollowsLatestIssuer bool `json:"default_follows_latest_issuer"`
}

// NewPKIWriteIssuersConfigRequestWithDefaults instantiates a new PKIWriteIssuersConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIWriteIssuersConfigRequestWithDefaults() *PKIWriteIssuersConfigRequest {
	var this PKIWriteIssuersConfigRequest

	this.DefaultFollowsLatestIssuer = false

	return &this
}

func (o PKIWriteIssuersConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["default"] = o.Default

	toSerialize["default_follows_latest_issuer"] = o.DefaultFollowsLatestIssuer

	return json.Marshal(toSerialize)
}
