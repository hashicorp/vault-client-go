// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AzureConfigureAuthRequest struct for AzureConfigureAuthRequest
type AzureConfigureAuthRequest struct {
	// The OAuth2 client id to connection to Azure. This value can also be provided with the AZURE_CLIENT_ID environment variable.
	ClientId string `json:"client_id"`

	// The OAuth2 client secret to connection to Azure. This value can also be provided with the AZURE_CLIENT_SECRET environment variable.
	ClientSecret string `json:"client_secret"`

	// The Azure environment name. If not provided, AzurePublicCloud is used. This value can also be provided with the AZURE_ENVIRONMENT environment variable.
	Environment string `json:"environment"`

	// The resource URL for the vault application in Azure Active Directory. This value can also be provided with the AZURE_AD_RESOURCE environment variable.
	Resource string `json:"resource"`

	// The TTL of the root password in Azure. This can be either a number of seconds or a time formatted duration (ex: 24h, 48ds)
	RootPasswordTtl int32 `json:"root_password_ttl"`

	// The tenant id for the Azure Active Directory. This is sometimes referred to as Directory ID in AD. This value can also be provided with the AZURE_TENANT_ID environment variable.
	TenantId string `json:"tenant_id"`
}

// NewAzureConfigureAuthRequestWithDefaults instantiates a new AzureConfigureAuthRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureConfigureAuthRequestWithDefaults() *AzureConfigureAuthRequest {
	var this AzureConfigureAuthRequest

	return &this
}

func (o AzureConfigureAuthRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["client_id"] = o.ClientId
	toSerialize["client_secret"] = o.ClientSecret
	toSerialize["environment"] = o.Environment
	toSerialize["resource"] = o.Resource
	toSerialize["root_password_ttl"] = o.RootPasswordTtl
	toSerialize["tenant_id"] = o.TenantId

	return json.Marshal(toSerialize)
}
