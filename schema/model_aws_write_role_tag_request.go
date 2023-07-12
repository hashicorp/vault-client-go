// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AwsWriteRoleTagRequest struct for AwsWriteRoleTagRequest
type AwsWriteRoleTagRequest struct {
	// If set, allows migration of the underlying instance where the client resides. This keys off of pendingTime in the metadata document, so essentially, this disables the client nonce check whenever the instance is migrated to a new host and pendingTime is newer than the previously-remembered time. Use with caution.
	AllowInstanceMigration bool `json:"allow_instance_migration,omitempty"`

	// If set, only allows a single token to be granted per instance ID. In order to perform a fresh login, the entry in access list for the instance ID needs to be cleared using the 'auth/aws-ec2/identity-accesslist/<instance_id>' endpoint.
	DisallowReauthentication bool `json:"disallow_reauthentication,omitempty"`

	// Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID.
	InstanceId string `json:"instance_id,omitempty"`

	// If set, specifies the maximum allowed token lifetime.
	MaxTtl string `json:"max_ttl,omitempty"`

	// Policies to be associated with the tag. If set, must be a subset of the role's policies. If set, but set to an empty value, only the 'default' policy will be given to issued tokens.
	Policies []string `json:"policies,omitempty"`
}
