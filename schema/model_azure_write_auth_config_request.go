// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AzureWriteAuthConfigRequest struct for AzureWriteAuthConfigRequest
type AzureWriteAuthConfigRequest struct {
	// The OAuth2 client id to connection to Azure. This value can also be provided with the AZURE_CLIENT_ID environment variable.
	ClientId string `json:"client_id"`

	// The OAuth2 client secret to connection to Azure. This value can also be provided with the AZURE_CLIENT_SECRET environment variable.
	ClientSecret string `json:"client_secret"`

	// The Azure environment name. If not provided, AzurePublicCloud is used. This value can also be provided with the AZURE_ENVIRONMENT environment variable.
	Environment string `json:"environment"`

	// The resource URL for the vault application in Azure Active Directory. This value can also be provided with the AZURE_AD_RESOURCE environment variable.
	Resource string `json:"resource"`

	// The tenant id for the Azure Active Directory. This is sometimes referred to as Directory ID in AD. This value can also be provided with the AZURE_TENANT_ID environment variable.
	TenantId string `json:"tenant_id"`
}

// NewAzureWriteAuthConfigRequestWithDefaults instantiates a new AzureWriteAuthConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureWriteAuthConfigRequestWithDefaults() *AzureWriteAuthConfigRequest {
	var this AzureWriteAuthConfigRequest

	return &this
}
