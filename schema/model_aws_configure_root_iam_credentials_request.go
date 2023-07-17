// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AwsConfigureRootIamCredentialsRequest struct for AwsConfigureRootIamCredentialsRequest
type AwsConfigureRootIamCredentialsRequest struct {
	// Access key with permission to create new keys.
	AccessKey string `json:"access_key,omitempty"`

	// Endpoint to custom IAM server URL
	IamEndpoint string `json:"iam_endpoint,omitempty"`

	// Maximum number of retries for recoverable exceptions of AWS APIs
	MaxRetries int32 `json:"max_retries,omitempty"`

	// Region for API calls.
	Region string `json:"region,omitempty"`

	// Secret key with permission to create new keys.
	SecretKey string `json:"secret_key,omitempty"`

	// Endpoint to custom STS server URL
	StsEndpoint string `json:"sts_endpoint,omitempty"`

	// Template to generate custom IAM usernames
	UsernameTemplate string `json:"username_template,omitempty"`
}
