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

// AWSWriteCredentialsRequest struct for AWSWriteCredentialsRequest
type AWSWriteCredentialsRequest struct {
	// Name of the role

	Name string `json:"name"`

	// ARN of role to assume when credential_type is assumed_role

	RoleArn string `json:"role_arn"`

	// Session name to use when assuming role. Max chars: 64

	RoleSessionName string `json:"role_session_name"`

	// Lifetime of the returned credentials in seconds

	Ttl int32 `json:"ttl"`
}

// NewAWSWriteCredentialsRequestWithDefaults instantiates a new AWSWriteCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSWriteCredentialsRequestWithDefaults() *AWSWriteCredentialsRequest {
	var this AWSWriteCredentialsRequest

	this.Ttl = 3600

	return &this
}

func (o AWSWriteCredentialsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["name"] = o.Name

	toSerialize["role_arn"] = o.RoleArn

	toSerialize["role_session_name"] = o.RoleSessionName

	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
