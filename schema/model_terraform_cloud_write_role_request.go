// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TerraformCloudWriteRoleRequest struct for TerraformCloudWriteRoleRequest
type TerraformCloudWriteRoleRequest struct {
	// Maximum time for role. If not set or set to 0, will use system default.
	MaxTtl string `json:"max_ttl,omitempty"`

	// Name of the Terraform Cloud or Enterprise organization
	Organization string `json:"organization,omitempty"`

	// ID of the Terraform Cloud or Enterprise team under organization (e.g., settings/teams/team-xxxxxxxxxxxxx)
	TeamId string `json:"team_id,omitempty"`

	// Default lease for generated credentials. If not set or set to 0, will use system default.
	Ttl string `json:"ttl,omitempty"`

	// ID of the Terraform Cloud or Enterprise user (e.g., user-xxxxxxxxxxxxxxxx)
	UserId string `json:"user_id,omitempty"`
}
