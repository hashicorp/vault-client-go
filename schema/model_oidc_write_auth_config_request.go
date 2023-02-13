// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OIDCWriteAuthConfigRequest struct for OIDCWriteAuthConfigRequest
type OIDCWriteAuthConfigRequest struct {
	// The value against which to match the 'iss' claim in a JWT. Optional.
	BoundIssuer string `json:"bound_issuer"`

	// The default role to use if none is provided during login. If not set, a role is required during login.
	DefaultRole string `json:"default_role"`

	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
	JwksCaPem string `json:"jwks_ca_pem"`

	// JWKS URL to use to authenticate signatures. Cannot be used with \"oidc_discovery_url\" or \"jwt_validation_pubkeys\".
	JwksUrl string `json:"jwks_url"`

	// A list of supported signing algorithms. Defaults to RS256.
	JwtSupportedAlgs []string `json:"jwt_supported_algs"`

	// A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used with \"jwks_url\" or \"oidc_discovery_url\".
	JwtValidationPubkeys []string `json:"jwt_validation_pubkeys"`

	// Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs.
	NamespaceInState bool `json:"namespace_in_state"`

	// The OAuth Client ID configured with your OIDC provider.
	OidcClientId string `json:"oidc_client_id"`

	// The OAuth Client Secret configured with your OIDC provider.
	OidcClientSecret string `json:"oidc_client_secret"`

	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used.
	OidcDiscoveryCaPem string `json:"oidc_discovery_ca_pem"`

	// OIDC Discovery URL, without any .well-known component (base path). Cannot be used with \"jwks_url\" or \"jwt_validation_pubkeys\".
	OidcDiscoveryUrl string `json:"oidc_discovery_url"`

	// The response mode to be used in the OAuth2 request. Allowed values are 'query' and 'form_post'.
	OidcResponseMode string `json:"oidc_response_mode"`

	// The response types to request. Allowed values are 'code' and 'id_token'. Defaults to 'code'.
	OidcResponseTypes []string `json:"oidc_response_types"`

	// Provider-specific configuration. Optional.
	ProviderConfig map[string]interface{} `json:"provider_config"`
}

// NewOIDCWriteAuthConfigRequestWithDefaults instantiates a new OIDCWriteAuthConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWriteAuthConfigRequestWithDefaults() *OIDCWriteAuthConfigRequest {
	var this OIDCWriteAuthConfigRequest

	return &this
}

func (o OIDCWriteAuthConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["bound_issuer"] = o.BoundIssuer
	toSerialize["default_role"] = o.DefaultRole
	toSerialize["jwks_ca_pem"] = o.JwksCaPem
	toSerialize["jwks_url"] = o.JwksUrl
	toSerialize["jwt_supported_algs"] = o.JwtSupportedAlgs
	toSerialize["jwt_validation_pubkeys"] = o.JwtValidationPubkeys
	toSerialize["namespace_in_state"] = o.NamespaceInState
	toSerialize["oidc_client_id"] = o.OidcClientId
	toSerialize["oidc_client_secret"] = o.OidcClientSecret
	toSerialize["oidc_discovery_ca_pem"] = o.OidcDiscoveryCaPem
	toSerialize["oidc_discovery_url"] = o.OidcDiscoveryUrl
	toSerialize["oidc_response_mode"] = o.OidcResponseMode
	toSerialize["oidc_response_types"] = o.OidcResponseTypes
	toSerialize["provider_config"] = o.ProviderConfig

	return json.Marshal(toSerialize)
}
