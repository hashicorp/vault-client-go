// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AzureLoginRequest struct for AzureLoginRequest
type AzureLoginRequest struct {
	// A signed JWT
	Jwt string `json:"jwt,omitempty"`

	// The resource group from the instance.
	ResourceGroupName string `json:"resource_group_name,omitempty"`

	// The fully qualified ID of the resource, includingthe resource name and resource type. Use the format, /subscriptions/{guid}/resourceGroups/{resource-group-name}/{resource-provider-namespace}/{resource-type}/{resource-name}. This value is ignored if vm_name or vmss_name is specified.
	ResourceId string `json:"resource_id,omitempty"`

	// The token role.
	Role string `json:"role,omitempty"`

	// The subscription id for the instance.
	SubscriptionId string `json:"subscription_id,omitempty"`

	// The name of the virtual machine. This value is ignored if vmss_name is specified.
	VmName string `json:"vm_name,omitempty"`

	// The name of the virtual machine scale set the instance is in.
	VmssName string `json:"vmss_name,omitempty"`
}
