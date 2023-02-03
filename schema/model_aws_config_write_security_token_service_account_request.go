// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AWSConfigWriteSecurityTokenServiceAccountRequest struct for AWSConfigWriteSecurityTokenServiceAccountRequest
type AWSConfigWriteSecurityTokenServiceAccountRequest struct {
	// AWS ARN for STS role to be assumed when interacting with the account specified. The Vault server must have permissions to assume this role.
	StsRole string `json:"sts_role"`
}

// NewAWSConfigWriteSecurityTokenServiceAccountRequestWithDefaults instantiates a new AWSConfigWriteSecurityTokenServiceAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSConfigWriteSecurityTokenServiceAccountRequestWithDefaults() *AWSConfigWriteSecurityTokenServiceAccountRequest {
	var this AWSConfigWriteSecurityTokenServiceAccountRequest

	return &this
}
