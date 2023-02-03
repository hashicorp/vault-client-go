// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// AppRoleWritePoliciesRequest struct for AppRoleWritePoliciesRequest
type AppRoleWritePoliciesRequest struct {
	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies"`

	// Comma-separated list of policies
	TokenPolicies []string `json:"token_policies"`
}

// NewAppRoleWritePoliciesRequestWithDefaults instantiates a new AppRoleWritePoliciesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWritePoliciesRequestWithDefaults() *AppRoleWritePoliciesRequest {
	var this AppRoleWritePoliciesRequest

	return &this
}
