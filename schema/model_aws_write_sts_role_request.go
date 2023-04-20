// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AwsWriteStsRoleRequest struct for AwsWriteStsRoleRequest
type AwsWriteStsRoleRequest struct {
	// AWS ARN for STS role to be assumed when interacting with the account specified. The Vault server must have permissions to assume this role.
	StsRole string `json:"sts_role"`
}

// NewAwsWriteStsRoleRequestWithDefaults instantiates a new AwsWriteStsRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsWriteStsRoleRequestWithDefaults() *AwsWriteStsRoleRequest {
	var this AwsWriteStsRoleRequest

	return &this
}

func (o AwsWriteStsRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["sts_role"] = o.StsRole

	return json.Marshal(toSerialize)
}
