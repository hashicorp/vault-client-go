// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AWSConfigWriteClientRequest struct for AWSConfigWriteClientRequest
type AWSConfigWriteClientRequest struct {
	// AWS Access Key ID for the account used to make AWS API requests.
	AccessKey string `json:"access_key"`

	// List of additional headers that are allowed to be in AWS STS request headers
	AllowedStsHeaderValues []string `json:"allowed_sts_header_values"`

	// URL to override the default generated endpoint for making AWS EC2 API calls.
	Endpoint string `json:"endpoint"`

	// URL to override the default generated endpoint for making AWS IAM API calls.
	IamEndpoint string `json:"iam_endpoint"`

	// Value to require in the X-Vault-AWS-IAM-Server-ID request header
	IamServerIdHeaderValue string `json:"iam_server_id_header_value"`

	// Maximum number of retries for recoverable exceptions of AWS APIs
	MaxRetries int32 `json:"max_retries"`

	// AWS Secret Access Key for the account used to make AWS API requests.
	SecretKey string `json:"secret_key"`

	// URL to override the default generated endpoint for making AWS STS API calls.
	StsEndpoint string `json:"sts_endpoint"`

	// The region ID for the sts_endpoint, if set.
	StsRegion string `json:"sts_region"`
}

// NewAWSConfigWriteClientRequestWithDefaults instantiates a new AWSConfigWriteClientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSConfigWriteClientRequestWithDefaults() *AWSConfigWriteClientRequest {
	var this AWSConfigWriteClientRequest

	this.AccessKey = ""
	this.Endpoint = ""
	this.IamEndpoint = ""
	this.IamServerIdHeaderValue = ""
	this.MaxRetries = -1
	this.SecretKey = ""
	this.StsEndpoint = ""
	this.StsRegion = ""

	return &this
}

func (o AWSConfigWriteClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
