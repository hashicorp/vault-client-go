// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// CentrifyWriteConfigRequest struct for CentrifyWriteConfigRequest
type CentrifyWriteConfigRequest struct {
	// OAuth2 App ID
	AppId string `json:"app_id"`

	// OAuth2 Client ID
	ClientId string `json:"client_id"`

	// OAuth2 Client Secret
	ClientSecret string `json:"client_secret"`

	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies"`

	// OAuth2 App Scope
	Scope string `json:"scope"`

	// Service URL (https://<tenant>.my.centrify.com)
	ServiceUrl string `json:"service_url"`

	// Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token.
	TokenBoundCidrs []string `json:"token_bound_cidrs"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy bool `json:"token_no_default_policy"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses int32 `json:"token_num_uses"`

	// Comma-separated list of policies
	TokenPolicies []string `json:"token_policies"`

	// The initial ttl of the token to generate
	TokenTtl int32 `json:"token_ttl"`

	// The type of token to generate, service or batch
	TokenType string `json:"token_type"`
}

// NewCentrifyWriteConfigRequestWithDefaults instantiates a new CentrifyWriteConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentrifyWriteConfigRequestWithDefaults() *CentrifyWriteConfigRequest {
	var this CentrifyWriteConfigRequest

	this.AppId = "vault_io_integration"
	this.Scope = "vault_io_integration"
	this.TokenType = "default-service"

	return &this
}
