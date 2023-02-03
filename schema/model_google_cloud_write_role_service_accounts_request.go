// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// GoogleCloudWriteRoleServiceAccountsRequest struct for GoogleCloudWriteRoleServiceAccountsRequest
type GoogleCloudWriteRoleServiceAccountsRequest struct {
	// Service-account emails or IDs to add.
	Add []string `json:"add"`

	// Service-account emails or IDs to remove.
	Remove []string `json:"remove"`
}

// NewGoogleCloudWriteRoleServiceAccountsRequestWithDefaults instantiates a new GoogleCloudWriteRoleServiceAccountsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudWriteRoleServiceAccountsRequestWithDefaults() *GoogleCloudWriteRoleServiceAccountsRequest {
	var this GoogleCloudWriteRoleServiceAccountsRequest

	return &this
}
