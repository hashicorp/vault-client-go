// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AwsConfigureClientRequest struct for AwsConfigureClientRequest
type AwsConfigureClientRequest struct {
	// AWS Access Key ID for the account used to make AWS API requests.
	AccessKey string `json:"access_key,omitempty"`

	// List of additional headers that are allowed to be in AWS STS request headers
	AllowedStsHeaderValues []string `json:"allowed_sts_header_values,omitempty"`

	// URL to override the default generated endpoint for making AWS EC2 API calls.
	Endpoint string `json:"endpoint,omitempty"`

	// URL to override the default generated endpoint for making AWS IAM API calls.
	IamEndpoint string `json:"iam_endpoint,omitempty"`

	// Value to require in the X-Vault-AWS-IAM-Server-ID request header
	IamServerIdHeaderValue string `json:"iam_server_id_header_value,omitempty"`

	// Maximum number of retries for recoverable exceptions of AWS APIs
	MaxRetries int32 `json:"max_retries,omitempty"`

	// AWS Secret Access Key for the account used to make AWS API requests.
	SecretKey string `json:"secret_key,omitempty"`

	// URL to override the default generated endpoint for making AWS STS API calls.
	StsEndpoint string `json:"sts_endpoint,omitempty"`

	// The region ID for the sts_endpoint, if set.
	StsRegion string `json:"sts_region,omitempty"`

	// Uses the STS region from client requests for making AWS STS API calls.
	UseStsRegionFromClient bool `json:"use_sts_region_from_client,omitempty"`
}
