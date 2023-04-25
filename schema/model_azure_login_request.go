// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AzureLoginRequest struct for AzureLoginRequest
type AzureLoginRequest struct {
	// A signed JWT
	Jwt string `json:"jwt"`

	// The resource group from the instance.
	ResourceGroupName string `json:"resource_group_name"`

	// The fully qualified ID of the resource, includingthe resource name and resource type. Use the format, /subscriptions/{guid}/resourceGroups/{resource-group-name}/{resource-provider-namespace}/{resource-type}/{resource-name}. This value is ignored if vm_name or vmss_name is specified.
	ResourceId string `json:"resource_id"`

	// The token role.
	Role string `json:"role"`

	// The subscription id for the instance.
	SubscriptionId string `json:"subscription_id"`

	// The name of the virtual machine. This value is ignored if vmss_name is specified.
	VmName string `json:"vm_name"`

	// The name of the virtual machine scale set the instance is in.
	VmssName string `json:"vmss_name"`
}

// NewAzureLoginRequestWithDefaults instantiates a new AzureLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureLoginRequestWithDefaults() *AzureLoginRequest {
	var this AzureLoginRequest

	return &this
}

func (o AzureLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["jwt"] = o.Jwt
	toSerialize["resource_group_name"] = o.ResourceGroupName
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["role"] = o.Role
	toSerialize["subscription_id"] = o.SubscriptionId
	toSerialize["vm_name"] = o.VmName
	toSerialize["vmss_name"] = o.VmssName

	return json.Marshal(toSerialize)
}
