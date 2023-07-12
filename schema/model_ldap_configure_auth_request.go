// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// LdapConfigureAuthRequest struct for LdapConfigureAuthRequest
type LdapConfigureAuthRequest struct {
	// Use anonymous binds when performing LDAP group searches (if true the initial credentials will still be used for the initial connection test).
	AnonymousGroupSearch bool `json:"anonymous_group_search,omitempty"`

	// LDAP DN for searching for the user DN (optional)
	Binddn string `json:"binddn,omitempty"`

	// LDAP password for searching for the user DN (optional)
	Bindpass string `json:"bindpass,omitempty"`

	// If true, case sensitivity will be used when comparing usernames and groups for matching policies.
	CaseSensitiveNames bool `json:"case_sensitive_names,omitempty"`

	// CA certificate to use when verifying LDAP server certificate, must be x509 PEM encoded (optional)
	Certificate string `json:"certificate,omitempty"`

	// Client certificate to provide to the LDAP server, must be x509 PEM encoded (optional)
	ClientTlsCert string `json:"client_tls_cert,omitempty"`

	// Client certificate key to provide to the LDAP server, must be x509 PEM encoded (optional)
	ClientTlsKey string `json:"client_tls_key,omitempty"`

	// Timeout, in seconds, when attempting to connect to the LDAP server before trying the next URL in the configuration.
	ConnectionTimeout string `json:"connection_timeout,omitempty"`

	// Denies an unauthenticated LDAP bind request if the user's password is empty; defaults to true
	DenyNullBind bool `json:"deny_null_bind,omitempty"`

	// When aliases should be dereferenced on search operations. Accepted values are 'never', 'finding', 'searching', 'always'. Defaults to 'never'.
	DereferenceAliases string `json:"dereference_aliases,omitempty"`

	// Use anonymous bind to discover the bind DN of a user (optional)
	Discoverdn bool `json:"discoverdn,omitempty"`

	// LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate user group membership. Examples: \"cn\" or \"memberOf\", etc. Default: cn
	Groupattr string `json:"groupattr,omitempty"`

	// LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org)
	Groupdn string `json:"groupdn,omitempty"`

	// Go template for querying group membership of user (optional) The template can access the following context variables: UserDN, Username Example: (&(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}})) Default: (|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))
	Groupfilter string `json:"groupfilter,omitempty"`

	// Skip LDAP server SSL Certificate verification - VERY insecure (optional)
	InsecureTls bool `json:"insecure_tls,omitempty"`

	// If set to a value greater than 0, the LDAP backend will use the LDAP server's paged search control to request pages of up to the given size. This can be used to avoid hitting the LDAP server's maximum result size limit. Otherwise, the LDAP backend will not use the paged search control.
	MaxPageSize int32 `json:"max_page_size,omitempty"`

	// Timeout, in seconds, for the connection when making requests against the server before returning back an error.
	RequestTimeout string `json:"request_timeout,omitempty"`

	// Issue a StartTLS command after establishing unencrypted connection (optional)
	Starttls bool `json:"starttls,omitempty"`

	// Maximum TLS version to use. Accepted values are 'tls10', 'tls11', 'tls12' or 'tls13'. Defaults to 'tls12'
	TlsMaxVersion string `json:"tls_max_version,omitempty"`

	// Minimum TLS version to use. Accepted values are 'tls10', 'tls11', 'tls12' or 'tls13'. Defaults to 'tls12'
	TlsMinVersion string `json:"tls_min_version,omitempty"`

	// Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token.
	TokenBoundCidrs []string `json:"token_bound_cidrs,omitempty"`

	// If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed.
	TokenExplicitMaxTtl string `json:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime of the generated token
	TokenMaxTtl string `json:"token_max_ttl,omitempty"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy bool `json:"token_no_default_policy,omitempty"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses int32 `json:"token_num_uses,omitempty"`

	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \"24h\").
	TokenPeriod string `json:"token_period,omitempty"`

	// Comma-separated list of policies. This will apply to all tokens generated by this auth method, in addition to any configured for specific users/groups.
	TokenPolicies []string `json:"token_policies,omitempty"`

	// The initial ttl of the token to generate
	TokenTtl string `json:"token_ttl,omitempty"`

	// The type of token to generate, service or batch
	TokenType string `json:"token_type,omitempty"`

	// Enables userPrincipalDomain login with [username]@UPNDomain (optional)
	Upndomain string `json:"upndomain,omitempty"`

	// LDAP URL to connect to (default: ldap://127.0.0.1). Multiple URLs can be specified by concatenating them with commas; they will be tried in-order.
	Url string `json:"url,omitempty"`

	// In Vault 1.1.1 a fix for handling group CN values of different cases unfortunately introduced a regression that could cause previously defined groups to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for matching group CNs will be used. This is only needed in some upgrade scenarios for backwards compatibility. It is enabled by default if the config is upgraded but disabled by default on new configurations.
	UsePre111GroupCnBehavior bool `json:"use_pre111_group_cn_behavior,omitempty"`

	// If true, use the Active Directory tokenGroups constructed attribute of the user to find the group memberships. This will find all security groups including nested ones.
	UseTokenGroups bool `json:"use_token_groups,omitempty"`

	// Attribute used for users (default: cn)
	Userattr string `json:"userattr,omitempty"`

	// LDAP domain to use for users (eg: ou=People,dc=example,dc=org)
	Userdn string `json:"userdn,omitempty"`

	// Go template for LDAP user search filer (optional) The template can access the following context variables: UserAttr, Username Default: ({{.UserAttr}}={{.Username}})
	Userfilter string `json:"userfilter,omitempty"`

	// If true, sets the alias name to the username
	UsernameAsAlias bool `json:"username_as_alias,omitempty"`
}
