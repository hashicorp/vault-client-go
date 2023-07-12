// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TerraformCloudConfigureRequest struct for TerraformCloudConfigureRequest
type TerraformCloudConfigureRequest struct {
	// The address to access Terraform Cloud or Enterprise. Default is \"https://app.terraform.io\".
	Address string `json:"address,omitempty"`

	// The base path for the Terraform Cloud or Enterprise API. Default is \"/api/v2/\".
	BasePath string `json:"base_path,omitempty"`

	// The token to access Terraform Cloud
	Token string `json:"token"`
}
