// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// PkiConfigureAcmeRequest struct for PkiConfigureAcmeRequest
type PkiConfigureAcmeRequest struct {
	// whether the ExtKeyUsage field from a role is used, defaults to false meaning that certificate will be signed with ServerAuth.
	AllowRoleExtKeyUsage bool `json:"allow_role_ext_key_usage,omitempty"`

	// which issuers are allowed for use with ACME; by default, this will only be the primary (default) issuer
	AllowedIssuers []string `json:"allowed_issuers,omitempty"`

	// which roles are allowed for use with ACME; by default via '*', these will be all roles including sign-verbatim; when concrete role names are specified, any default_directory_policy role must be included to allow usage of the default acme directories under /pki/acme/directory and /pki/issuer/:issuer_id/acme/directory.
	AllowedRoles []string `json:"allowed_roles,omitempty"`

	// the policy to be used for non-role-qualified ACME requests; by default ACME issuance will be otherwise unrestricted, equivalent to the sign-verbatim endpoint; one may also specify a role to use as this policy, as \"role:<role_name>\", the specified role must be allowed by allowed_roles
	DefaultDirectoryPolicy string `json:"default_directory_policy,omitempty"`

	// DNS resolver to use for domain resolution on this mount. Defaults to using the default system resolver. Must be in the format <host>:<port>, with both parts mandatory.
	DnsResolver string `json:"dns_resolver,omitempty"`

	// Specify the policy to use for external account binding behaviour, 'not-required', 'new-account-required' or 'always-required'
	EabPolicy string `json:"eab_policy,omitempty"`

	// whether ACME is enabled, defaults to false meaning that clusters will by default not get ACME support
	Enabled bool `json:"enabled,omitempty"`
}
