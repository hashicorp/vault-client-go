// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AWSWriteSecurityTokenServiceRequest struct for AWSWriteSecurityTokenServiceRequest
type AWSWriteSecurityTokenServiceRequest struct {
	// ARN of role to assume when credential_type is assumed_role
	RoleArn string `json:"role_arn"`

	// Session name to use when assuming role. Max chars: 64
	RoleSessionName string `json:"role_session_name"`

	// Lifetime of the returned credentials in seconds
	Ttl int32 `json:"ttl"`
}

// NewAWSWriteSecurityTokenServiceRequestWithDefaults instantiates a new AWSWriteSecurityTokenServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSWriteSecurityTokenServiceRequestWithDefaults() *AWSWriteSecurityTokenServiceRequest {
	var this AWSWriteSecurityTokenServiceRequest

	this.Ttl = 3600

	return &this
}

func (o AWSWriteSecurityTokenServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
