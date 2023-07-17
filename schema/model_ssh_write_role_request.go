// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// SshWriteRoleRequest struct for SshWriteRoleRequest
type SshWriteRoleRequest struct {
	// [Not applicable for OTP type] [Optional for CA type] When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512, default, or the empty string.
	AlgorithmSigner string `json:"algorithm_signer,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] If set, host certificates that are requested are allowed to use the base domains listed in \"allowed_domains\", e.g. \"example.com\". This is a separate option as in some cases this can be considered a security threat.
	AllowBareDomains bool `json:"allow_bare_domains,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] If set, certificates are allowed to be signed for use as a 'host'.
	AllowHostCertificates bool `json:"allow_host_certificates,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] If set, host certificates that are requested are allowed to use subdomains of those listed in \"allowed_domains\".
	AllowSubdomains bool `json:"allow_subdomains,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] If set, certificates are allowed to be signed for use as a 'user'.
	AllowUserCertificates bool `json:"allow_user_certificates,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] If true, users can override the key ID for a signed certificate with the \"key_id\" field. When false, the key ID will always be the token display name. The key ID is logged by the SSH server and can be useful for auditing.
	AllowUserKeyIds bool `json:"allow_user_key_ids,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] A comma-separated list of critical options that certificates can have when signed. To allow any critical options, set this to an empty string.
	AllowedCriticalOptions string `json:"allowed_critical_options,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] If this option is not specified, client can request for a signed certificate for any valid host. If only certain domains are allowed, then this list enforces it.
	AllowedDomains string `json:"allowed_domains,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] If set, Allowed domains can be specified using identity template policies. Non-templated domains are also permitted.
	AllowedDomainsTemplate bool `json:"allowed_domains_template,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] A comma-separated list of extensions that certificates can have when signed. An empty list means that no extension overrides are allowed by an end-user; explicitly specify '*' to allow any extensions to be set.
	AllowedExtensions string `json:"allowed_extensions,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] If set, allows the enforcement of key types and minimum key sizes to be signed.
	AllowedUserKeyLengths map[string]interface{} `json:"allowed_user_key_lengths,omitempty"`

	// [Optional for all types] [Works differently for CA type] If this option is not specified, or is '*', client can request a credential for any valid user at the remote host, including the admin user. If only certain usernames are to be allowed, then this list enforces it. If this field is set, then credentials can only be created for default_user and usernames present in this list. Setting this option will enable all the users with access to this role to fetch credentials for all other usernames in this list. Use with caution. N.B.: with the CA type, an empty list means that no users are allowed; explicitly specify '*' to allow any user.
	AllowedUsers string `json:"allowed_users,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] If set, Allowed users can be specified using identity template policies. Non-templated users are also permitted.
	AllowedUsersTemplate bool `json:"allowed_users_template,omitempty"`

	// [Optional for OTP type] [Not applicable for CA type] Comma separated list of CIDR blocks for which the role is applicable for. CIDR blocks can belong to more than one role.
	CidrList string `json:"cidr_list,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] Critical options certificates should have if none are provided when signing. This field takes in key value pairs in JSON format. Note that these are not restricted by \"allowed_critical_options\". Defaults to none.
	DefaultCriticalOptions map[string]interface{} `json:"default_critical_options,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] Extensions certificates should have if none are provided when signing. This field takes in key value pairs in JSON format. Note that these are not restricted by \"allowed_extensions\". Defaults to none.
	DefaultExtensions map[string]interface{} `json:"default_extensions,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] If set, Default extension values can be specified using identity template policies. Non-templated extension values are also permitted.
	DefaultExtensionsTemplate bool `json:"default_extensions_template,omitempty"`

	// [Required for OTP type] [Optional for CA type] Default username for which a credential will be generated. When the endpoint 'creds/' is used without a username, this value will be used as default username.
	DefaultUser string `json:"default_user,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] If set, Default user can be specified using identity template policies. Non-templated users are also permitted.
	DefaultUserTemplate bool `json:"default_user_template,omitempty"`

	// [Optional for OTP type] [Not applicable for CA type] Comma separated list of CIDR blocks. IP addresses belonging to these blocks are not accepted by the role. This is particularly useful when big CIDR blocks are being used by the role and certain parts of it needs to be kept out.
	ExcludeCidrList string `json:"exclude_cidr_list,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] When supplied, this value specifies a custom format for the key id of a signed certificate. The following variables are available for use: '{{token_display_name}}' - The display name of the token used to make the request. '{{role_name}}' - The name of the role signing the request. '{{public_key_hash}}' - A SHA256 checksum of the public key that is being signed.
	KeyIdFormat string `json:"key_id_format,omitempty"`

	// [Required for all types] Type of key used to login to hosts. It can be either 'otp' or 'ca'. 'otp' type requires agent to be installed in remote hosts.
	KeyType string `json:"key_type,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] The maximum allowed lease duration
	MaxTtl string `json:"max_ttl,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] The duration that the SSH certificate should be backdated by at issuance.
	NotBeforeDuration string `json:"not_before_duration,omitempty"`

	// [Optional for OTP type] [Not applicable for CA type] Port number for SSH connection. Default is '22'. Port number does not play any role in creation of OTP. For 'otp' type, this is just a way to inform client about the port number to use. Port number will be returned to client by Vault server along with OTP.
	Port int32 `json:"port,omitempty"`

	// [Not applicable for OTP type] [Optional for CA type] The lease duration if no specific lease duration is requested. The lease duration controls the expiration of certificates issued by this backend. Defaults to the value of max_ttl.
	Ttl string `json:"ttl,omitempty"`
}
