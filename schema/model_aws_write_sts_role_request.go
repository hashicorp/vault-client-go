// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AwsWriteStsRoleRequest struct for AwsWriteStsRoleRequest
type AwsWriteStsRoleRequest struct {
	// AWS ARN for STS role to be assumed when interacting with the account specified. The Vault server must have permissions to assume this role.
	StsRole string `json:"sts_role,omitempty"`
}
