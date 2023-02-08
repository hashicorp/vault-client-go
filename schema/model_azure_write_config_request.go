// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AzureWriteConfigRequest struct for AzureWriteConfigRequest
type AzureWriteConfigRequest struct {
	// The OAuth2 client id to connect to Azure. This value can also be provided with the AZURE_CLIENT_ID environment variable.
	ClientId string `json:"client_id"`

	// The OAuth2 client secret to connect to Azure. This value can also be provided with the AZURE_CLIENT_SECRET environment variable.
	ClientSecret string `json:"client_secret"`

	// The Azure environment name. If not provided, AzurePublicCloud is used. This value can also be provided with the AZURE_ENVIRONMENT environment variable.
	Environment string `json:"environment"`

	// Name of the password policy to use to generate passwords for dynamic credentials.
	PasswordPolicy string `json:"password_policy"`

	// The TTL of the root password in Azure. This can be either a number of seconds or a time formatted duration (ex: 24h, 48ds)
	RootPasswordTtl int32 `json:"root_password_ttl"`

	// The subscription id for the Azure Active Directory. This value can also be provided with the AZURE_SUBSCRIPTION_ID environment variable.
	SubscriptionId string `json:"subscription_id"`

	// The tenant id for the Azure Active Directory. This value can also be provided with the AZURE_TENANT_ID environment variable.
	TenantId string `json:"tenant_id"`
}

// NewAzureWriteConfigRequestWithDefaults instantiates a new AzureWriteConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureWriteConfigRequestWithDefaults() *AzureWriteConfigRequest {
	var this AzureWriteConfigRequest

	return &this
}

func (o AzureWriteConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}
