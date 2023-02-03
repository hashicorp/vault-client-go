// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// OIDCWriteProviderTokenRequest struct for OIDCWriteProviderTokenRequest
type OIDCWriteProviderTokenRequest struct {
	// The ID of the requesting client.
	ClientId string `json:"client_id"`

	// The secret of the requesting client.
	ClientSecret string `json:"client_secret"`

	// The authorization code received from the provider's authorization endpoint.
	Code string `json:"code"`

	// The code verifier associated with the authorization code.
	CodeVerifier string `json:"code_verifier"`

	// The authorization grant type. The following grant types are supported: 'authorization_code'.
	GrantType string `json:"grant_type"`

	// The callback location where the authentication response was sent.
	RedirectUri string `json:"redirect_uri"`
}

// NewOIDCWriteProviderTokenRequestWithDefaults instantiates a new OIDCWriteProviderTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWriteProviderTokenRequestWithDefaults() *OIDCWriteProviderTokenRequest {
	var this OIDCWriteProviderTokenRequest

	return &this
}
