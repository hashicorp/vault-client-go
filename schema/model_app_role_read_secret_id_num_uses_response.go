// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleReadSecretIDNumUsesResponse struct for AppRoleReadSecretIDNumUsesResponse
type AppRoleReadSecretIDNumUsesResponse struct {
	// Number of times a secret ID can access the role, after which the SecretID will expire. Defaults to 0 meaning that the secret ID is of unlimited use.
	SecretIdNumUses int32 `json:"secret_id_num_uses"`
}

// NewAppRoleReadSecretIDNumUsesResponseWithDefaults instantiates a new AppRoleReadSecretIDNumUsesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadSecretIDNumUsesResponseWithDefaults() *AppRoleReadSecretIDNumUsesResponse {
	var this AppRoleReadSecretIDNumUsesResponse

	return &this
}
