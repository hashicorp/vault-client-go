// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// MongoDbAtlasWriteRoleRequest struct for MongoDbAtlasWriteRoleRequest
type MongoDbAtlasWriteRoleRequest struct {
	// Access list entry in CIDR notation to be added for the API key. Optional for organization and project keys.
	CidrBlocks []string `json:"cidr_blocks,omitempty"`

	// IP address to be added to the access list for the API key. Optional for organization and project keys.
	IpAddresses []string `json:"ip_addresses,omitempty"`

	// The maximum allowed lifetime of credentials issued using this role.
	MaxTtl string `json:"max_ttl,omitempty"`

	// Organization ID required for an organization API key
	OrganizationId string `json:"organization_id,omitempty"`

	// Project ID the project API key belongs to.
	ProjectId string `json:"project_id,omitempty"`

	// Roles assigned when an organization API Key is assigned to a project API key
	ProjectRoles []string `json:"project_roles,omitempty"`

	// List of roles that the API Key should be granted. A minimum of one role must be provided. Any roles provided must be valid for the assigned Project, required for organization and project keys.
	Roles []string `json:"roles"`

	// Duration in seconds after which the issued credential should expire. Defaults to 0, in which case the value will fallback to the system/mount defaults.
	Ttl string `json:"ttl,omitempty"`
}
