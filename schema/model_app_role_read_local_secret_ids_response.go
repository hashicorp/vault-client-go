// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleReadLocalSecretIDsResponse struct for AppRoleReadLocalSecretIDsResponse
type AppRoleReadLocalSecretIDsResponse struct {
	// If true, the secret identifiers generated using this role will be cluster local. This can only be set during role creation and once set, it can't be reset later
	LocalSecretIds bool `json:"local_secret_ids"`
}

// NewAppRoleReadLocalSecretIDsResponseWithDefaults instantiates a new AppRoleReadLocalSecretIDsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadLocalSecretIDsResponseWithDefaults() *AppRoleReadLocalSecretIDsResponse {
	var this AppRoleReadLocalSecretIDsResponse

	return &this
}
