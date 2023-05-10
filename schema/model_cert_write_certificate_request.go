// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// CertWriteCertificateRequest struct for CertWriteCertificateRequest
type CertWriteCertificateRequest struct {
	// A comma-separated list of names. At least one must exist in the Common Name. Supports globbing.
	AllowedCommonNames []string `json:"allowed_common_names"`

	// A comma-separated list of DNS names. At least one must exist in the SANs. Supports globbing.
	AllowedDnsSans []string `json:"allowed_dns_sans"`

	// A comma-separated list of Email Addresses. At least one must exist in the SANs. Supports globbing.
	AllowedEmailSans []string `json:"allowed_email_sans"`

	// A comma-separated string or array of oid extensions. Upon successful authentication, these extensions will be added as metadata if they are present in the certificate. The metadata key will be the string consisting of the oid numbers separated by a dash (-) instead of a dot (.) to allow usage in ACL templates.
	AllowedMetadataExtensions []string `json:"allowed_metadata_extensions"`

	// A comma-separated list of names. At least one must exist in either the Common Name or SANs. Supports globbing. This parameter is deprecated, please use allowed_common_names, allowed_dns_sans, allowed_email_sans, allowed_uri_sans.
	AllowedNames []string `json:"allowed_names"`

	// A comma-separated list of Organizational Units names. At least one must exist in the OU field.
	AllowedOrganizationalUnits []string `json:"allowed_organizational_units"`

	// A comma-separated list of URIs. At least one must exist in the SANs. Supports globbing.
	AllowedUriSans []string `json:"allowed_uri_sans"`

	// Use \"token_bound_cidrs\" instead. If this and \"token_bound_cidrs\" are both specified, only \"token_bound_cidrs\" will be used.
	// Deprecated
	BoundCidrs []string `json:"bound_cidrs"`

	// The public certificate that should be trusted. Must be x509 PEM encoded.
	Certificate string `json:"certificate"`

	// The display name to use for clients using this certificate.
	DisplayName string `json:"display_name"`

	// Use \"token_ttl\" instead. If this and \"token_ttl\" are both specified, only \"token_ttl\" will be used.
	// Deprecated
	Lease int32 `json:"lease"`

	// Use \"token_max_ttl\" instead. If this and \"token_max_ttl\" are both specified, only \"token_max_ttl\" will be used.
	// Deprecated
	MaxTtl int32 `json:"max_ttl"`

	// Any additional CA certificates needed to communicate with OCSP servers
	OcspCaCertificates string `json:"ocsp_ca_certificates"`

	// Whether to attempt OCSP verification of certificates at login
	OcspEnabled bool `json:"ocsp_enabled"`

	// If set to true, if an OCSP revocation cannot be made successfully, login will proceed rather than failing. If false, failing to get an OCSP status fails the request.
	OcspFailOpen bool `json:"ocsp_fail_open"`

	// If set to true, rather than accepting the first successful OCSP response, query all servers and consider the certificate valid only if all servers agree.
	OcspQueryAllServers bool `json:"ocsp_query_all_servers"`

	// A comma-separated list of OCSP server addresses. If unset, the OCSP server is determined from the AuthorityInformationAccess extension on the certificate being inspected.
	OcspServersOverride []string `json:"ocsp_servers_override"`

	// Use \"token_period\" instead. If this and \"token_period\" are both specified, only \"token_period\" will be used.
	// Deprecated
	Period int32 `json:"period"`

	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies"`

	// A comma-separated string or array of extensions formatted as \"oid:value\". Expects the extension value to be some type of ASN1 encoded string. All values much match. Supports globbing on \"value\".
	RequiredExtensions []string `json:"required_extensions"`

	// Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token.
	TokenBoundCidrs []string `json:"token_bound_cidrs"`

	// If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed.
	TokenExplicitMaxTtl int32 `json:"token_explicit_max_ttl"`

	// The maximum lifetime of the generated token
	TokenMaxTtl int32 `json:"token_max_ttl"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy bool `json:"token_no_default_policy"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses int32 `json:"token_num_uses"`

	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \"24h\").
	TokenPeriod int32 `json:"token_period"`

	// Comma-separated list of policies
	TokenPolicies []string `json:"token_policies"`

	// The initial ttl of the token to generate
	TokenTtl int32 `json:"token_ttl"`

	// The type of token to generate, service or batch
	TokenType string `json:"token_type"`

	// Use \"token_ttl\" instead. If this and \"token_ttl\" are both specified, only \"token_ttl\" will be used.
	// Deprecated
	Ttl int32 `json:"ttl"`
}

// NewCertWriteCertificateRequestWithDefaults instantiates a new CertWriteCertificateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertWriteCertificateRequestWithDefaults() *CertWriteCertificateRequest {
	var this CertWriteCertificateRequest

	this.OcspFailOpen = false
	this.OcspQueryAllServers = false
	this.TokenType = "default-service"

	return &this
}

func (o CertWriteCertificateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["allowed_common_names"] = o.AllowedCommonNames
	toSerialize["allowed_dns_sans"] = o.AllowedDnsSans
	toSerialize["allowed_email_sans"] = o.AllowedEmailSans
	toSerialize["allowed_metadata_extensions"] = o.AllowedMetadataExtensions
	toSerialize["allowed_names"] = o.AllowedNames
	toSerialize["allowed_organizational_units"] = o.AllowedOrganizationalUnits
	toSerialize["allowed_uri_sans"] = o.AllowedUriSans
	toSerialize["bound_cidrs"] = o.BoundCidrs
	toSerialize["certificate"] = o.Certificate
	toSerialize["display_name"] = o.DisplayName
	toSerialize["lease"] = o.Lease
	toSerialize["max_ttl"] = o.MaxTtl
	toSerialize["ocsp_ca_certificates"] = o.OcspCaCertificates
	toSerialize["ocsp_enabled"] = o.OcspEnabled
	toSerialize["ocsp_fail_open"] = o.OcspFailOpen
	toSerialize["ocsp_query_all_servers"] = o.OcspQueryAllServers
	toSerialize["ocsp_servers_override"] = o.OcspServersOverride
	toSerialize["period"] = o.Period
	toSerialize["policies"] = o.Policies
	toSerialize["required_extensions"] = o.RequiredExtensions
	toSerialize["token_bound_cidrs"] = o.TokenBoundCidrs
	toSerialize["token_explicit_max_ttl"] = o.TokenExplicitMaxTtl
	toSerialize["token_max_ttl"] = o.TokenMaxTtl
	toSerialize["token_no_default_policy"] = o.TokenNoDefaultPolicy
	toSerialize["token_num_uses"] = o.TokenNumUses
	toSerialize["token_period"] = o.TokenPeriod
	toSerialize["token_policies"] = o.TokenPolicies
	toSerialize["token_ttl"] = o.TokenTtl
	toSerialize["token_type"] = o.TokenType
	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
