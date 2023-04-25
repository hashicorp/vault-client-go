// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// TerraformCloudWriteRoleRequest struct for TerraformCloudWriteRoleRequest
type TerraformCloudWriteRoleRequest struct {
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

// NewTerraformCloudWriteRoleRequestWithDefaults instantiates a new TerraformCloudWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerraformCloudWriteRoleRequestWithDefaults() *TerraformCloudWriteRoleRequest {
	var this TerraformCloudWriteRoleRequest

	return &this
}

func (o TerraformCloudWriteRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["max_ttl"] = o.MaxTtl
	toSerialize["organization"] = o.Organization
	toSerialize["team_id"] = o.TeamId
	toSerialize["ttl"] = o.Ttl
	toSerialize["user_id"] = o.UserId

	return json.Marshal(toSerialize)
}
