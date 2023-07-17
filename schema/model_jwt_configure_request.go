// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// JwtConfigureRequest struct for JwtConfigureRequest
type JwtConfigureRequest struct {
	// The value against which to match the 'iss' claim in a JWT. Optional.
	BoundIssuer string `json:"bound_issuer,omitempty"`

	// The default role to use if none is provided during login. If not set, a role is required during login.
	DefaultRole string `json:"default_role,omitempty"`

	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
	JwksCaPem string `json:"jwks_ca_pem,omitempty"`

	// JWKS URL to use to authenticate signatures. Cannot be used with \"oidc_discovery_url\" or \"jwt_validation_pubkeys\".
	JwksUrl string `json:"jwks_url,omitempty"`

	// A list of supported signing algorithms. Defaults to RS256.
	JwtSupportedAlgs []string `json:"jwt_supported_algs,omitempty"`

	// A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used with \"jwks_url\" or \"oidc_discovery_url\".
	JwtValidationPubkeys []string `json:"jwt_validation_pubkeys,omitempty"`

	// Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs.
	NamespaceInState bool `json:"namespace_in_state,omitempty"`

	// The OAuth Client ID configured with your OIDC provider.
	OidcClientId string `json:"oidc_client_id,omitempty"`

	// The OAuth Client Secret configured with your OIDC provider.
	OidcClientSecret string `json:"oidc_client_secret,omitempty"`

	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used.
	OidcDiscoveryCaPem string `json:"oidc_discovery_ca_pem,omitempty"`

	// OIDC Discovery URL, without any .well-known component (base path). Cannot be used with \"jwks_url\" or \"jwt_validation_pubkeys\".
	OidcDiscoveryUrl string `json:"oidc_discovery_url,omitempty"`

	// The response mode to be used in the OAuth2 request. Allowed values are 'query' and 'form_post'.
	OidcResponseMode string `json:"oidc_response_mode,omitempty"`

	// The response types to request. Allowed values are 'code' and 'id_token'. Defaults to 'code'.
	OidcResponseTypes []string `json:"oidc_response_types,omitempty"`

	// Provider-specific configuration. Optional.
	ProviderConfig map[string]interface{} `json:"provider_config,omitempty"`
}
