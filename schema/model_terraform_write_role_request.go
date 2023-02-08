// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TerraformWriteRoleRequest struct for TerraformWriteRoleRequest
type TerraformWriteRoleRequest struct {
	// Maximum time for role. If not set or set to 0, will use system default.
	MaxTtl int32 `json:"max_ttl"`

	// Name of the Terraform Cloud or Enterprise organization
	Organization string `json:"organization"`

	// ID of the Terraform Cloud or Enterprise team under organization (e.g., settings/teams/team-xxxxxxxxxxxxx)
	TeamId string `json:"team_id"`

	// Default lease for generated credentials. If not set or set to 0, will use system default.
	Ttl int32 `json:"ttl"`

	// ID of the Terraform Cloud or Enterprise user (e.g., user-xxxxxxxxxxxxxxxx)
	UserId string `json:"user_id"`
}

// NewTerraformWriteRoleRequestWithDefaults instantiates a new TerraformWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerraformWriteRoleRequestWithDefaults() *TerraformWriteRoleRequest {
	var this TerraformWriteRoleRequest

	return &this
}

func (o TerraformWriteRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
