// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AWSConfigWriteRootIAMCredentialsRequest struct for AWSConfigWriteRootIAMCredentialsRequest
type AWSConfigWriteRootIAMCredentialsRequest struct {
	// Access key with permission to create new keys.
	AccessKey string `json:"access_key"`

	// Endpoint to custom IAM server URL
	IamEndpoint string `json:"iam_endpoint"`

	// Maximum number of retries for recoverable exceptions of AWS APIs
	MaxRetries int32 `json:"max_retries"`

	// Region for API calls.
	Region string `json:"region"`

	// Secret key with permission to create new keys.
	SecretKey string `json:"secret_key"`

	// Endpoint to custom STS server URL
	StsEndpoint string `json:"sts_endpoint"`

	// Template to generate custom IAM usernames
	UsernameTemplate string `json:"username_template"`
}

// NewAWSConfigWriteRootIAMCredentialsRequestWithDefaults instantiates a new AWSConfigWriteRootIAMCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSConfigWriteRootIAMCredentialsRequestWithDefaults() *AWSConfigWriteRootIAMCredentialsRequest {
	var this AWSConfigWriteRootIAMCredentialsRequest

	this.MaxRetries = -1

	return &this
}

func (o AWSConfigWriteRootIAMCredentialsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
