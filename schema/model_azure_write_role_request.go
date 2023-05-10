// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AzureWriteRoleRequest struct for AzureWriteRoleRequest
type AzureWriteRoleRequest struct {
	// Application Object ID to use for static service principal credentials.
	ApplicationObjectId string `json:"application_object_id,omitempty"`

	// JSON list of Azure groups to add the service principal to.
	AzureGroups string `json:"azure_groups,omitempty"`

	// JSON list of Azure roles to assign.
	AzureRoles string `json:"azure_roles,omitempty"`

	// Maximum time a service principal. If not set or set to 0, will use system default.
	MaxTtl int32 `json:"max_ttl,omitempty"`

	// Indicates whether new application objects should be permanently deleted. If not set, objects will not be permanently deleted.
	PermanentlyDelete bool `json:"permanently_delete,omitempty"`

	// Persist the app between generated credentials. Useful if the app needs to maintain owner ship of resources it creates
	PersistApp bool `json:"persist_app,omitempty"`

	// Default lease for generated credentials. If not set or set to 0, will use system default.
	Ttl int32 `json:"ttl,omitempty"`
}

// NewAzureWriteRoleRequestWithDefaults instantiates a new AzureWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureWriteRoleRequestWithDefaults() *AzureWriteRoleRequest {
	var this AzureWriteRoleRequest

	this.PermanentlyDelete = false
	this.PersistApp = false

	return &this
}
