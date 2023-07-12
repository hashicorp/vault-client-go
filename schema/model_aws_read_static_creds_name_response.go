// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AwsReadStaticCredsNameResponse struct for AwsReadStaticCredsNameResponse
type AwsReadStaticCredsNameResponse struct {
	// The access key of the AWS Credential
	AccessKey string `json:"access_key,omitempty"`

	// The secret key of the AWS Credential
	SecretKey string `json:"secret_key,omitempty"`
}
