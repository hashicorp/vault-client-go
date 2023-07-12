// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AzureConfigureAuthRequest struct for AzureConfigureAuthRequest
type AzureConfigureAuthRequest struct {
	// The OAuth2 client id to connection to Azure. This value can also be provided with the AZURE_CLIENT_ID environment variable.
	ClientId string `json:"client_id,omitempty"`

	// The OAuth2 client secret to connection to Azure. This value can also be provided with the AZURE_CLIENT_SECRET environment variable.
	ClientSecret string `json:"client_secret,omitempty"`

	// The Azure environment name. If not provided, AzurePublicCloud is used. This value can also be provided with the AZURE_ENVIRONMENT environment variable.
	Environment string `json:"environment,omitempty"`

	// The resource URL for the vault application in Azure Active Directory. This value can also be provided with the AZURE_AD_RESOURCE environment variable.
	Resource string `json:"resource,omitempty"`

	// The TTL of the root password in Azure. This can be either a number of seconds or a time formatted duration (ex: 24h, 48ds)
	RootPasswordTtl string `json:"root_password_ttl,omitempty"`

	// The tenant id for the Azure Active Directory. This is sometimes referred to as Directory ID in AD. This value can also be provided with the AZURE_TENANT_ID environment variable.
	TenantId string `json:"tenant_id,omitempty"`
}
