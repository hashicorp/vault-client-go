// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PkiPatchRoleResponse struct for PkiPatchRoleResponse
type PkiPatchRoleResponse struct {
	// If set, clients can request certificates for any domain, regardless of allowed_domains restrictions. See the documentation for more information.
	AllowAnyName bool `json:"allow_any_name"`

	// If set, clients can request certificates for the base domains themselves, e.g. \"example.com\" of domains listed in allowed_domains. This is a separate option as in some cases this can be considered a security threat. See the documentation for more information.
	AllowBareDomains bool `json:"allow_bare_domains"`

	// If set, domains specified in allowed_domains can include shell-style glob patterns, e.g. \"ftp*.example.com\". See the documentation for more information.
	AllowGlobDomains bool `json:"allow_glob_domains"`

	// If set, IP Subject Alternative Names are allowed. Any valid IP is accepted and No authorization checking is performed.
	AllowIpSans bool `json:"allow_ip_sans"`

	// Whether to allow \"localhost\" and \"localdomain\" as a valid common name in a request, independent of allowed_domains value.
	AllowLocalhost bool `json:"allow_localhost"`

	// If set, clients can request certificates for subdomains of domains listed in allowed_domains, including wildcard subdomains. See the documentation for more information.
	AllowSubdomains bool `json:"allow_subdomains"`

	// Whether to allow \"localhost\" and \"localdomain\" as a valid common name in a request, independent of allowed_domains value.
	AllowTokenDisplayname bool `json:"allow_token_displayname"`

	// If set, allows certificates with wildcards in the common name to be issued, conforming to RFC 6125's Section 6.4.3; e.g., \"*.example.net\" or \"b*z.example.net\". See the documentation for more information.
	AllowWildcardCertificates bool `json:"allow_wildcard_certificates"`

	// Specifies the domains this role is allowed to issue certificates for. This is used with the allow_bare_domains, allow_subdomains, and allow_glob_domains to determine matches for the common name, DNS-typed SAN entries, and Email-typed SAN entries of certificates. See the documentation for more information. This parameter accepts a comma-separated string or list of domains.
	AllowedDomains []string `json:"allowed_domains"`

	// If set, Allowed domains can be specified using identity template policies. Non-templated domains are also permitted.
	AllowedDomainsTemplate bool `json:"allowed_domains_template"`

	// If set, an array of allowed other names to put in SANs. These values support globbing and must be in the format <oid>;<type>:<value>. Currently only \"utf8\" is a valid type. All values, including globbing values, must use this syntax, with the exception being a single \"*\" which allows any OID and any value (but type must still be utf8).
	AllowedOtherSans []string `json:"allowed_other_sans"`

	// If set, an array of allowed serial numbers to put in Subject. These values support globbing.
	AllowedSerialNumbers []string `json:"allowed_serial_numbers"`

	// If set, an array of allowed URIs for URI Subject Alternative Names. Any valid URI is accepted, these values support globbing.
	AllowedUriSans []string `json:"allowed_uri_sans"`

	// If set, Allowed URI SANs can be specified using identity template policies. Non-templated URI SANs are also permitted.
	AllowedUriSansTemplate bool `json:"allowed_uri_sans_template"`

	// If set, an array of allowed user-ids to put in user system login name specified here: https://www.rfc-editor.org/rfc/rfc1274#section-9.3.1
	AllowedUserIds []string `json:"allowed_user_ids"`

	// Mark Basic Constraints valid when issuing non-CA certificates.
	BasicConstraintsValidForNonCa bool `json:"basic_constraints_valid_for_non_ca"`

	// If set, certificates are flagged for client auth use. Defaults to true. See also RFC 5280 Section 4.2.1.12.
	ClientFlag bool `json:"client_flag"`

	// List of allowed validations to run against the Common Name field. Values can include 'email' to validate the CN is a email address, 'hostname' to validate the CN is a valid hostname (potentially including wildcards). When multiple validations are specified, these take OR semantics (either email OR hostname are allowed). The special value 'disabled' allows disabling all CN name validations, allowing for arbitrary non-Hostname, non-Email address CNs.
	CnValidations []string `json:"cn_validations"`

	// If set, certificates are flagged for code signing use. Defaults to false. See also RFC 5280 Section 4.2.1.12.
	CodeSigningFlag bool `json:"code_signing_flag"`

	// If set, Country will be set to this value in certificates issued by this role.
	Country []string `json:"country"`

	// If set, certificates are flagged for email protection use. Defaults to false. See also RFC 5280 Section 4.2.1.12.
	EmailProtectionFlag bool `json:"email_protection_flag"`

	// If set, only valid host names are allowed for CN and DNS SANs, and the host part of email addresses. Defaults to true.
	EnforceHostnames bool `json:"enforce_hostnames"`

	// A comma-separated string or list of extended key usages. Valid values can be found at https://golang.org/pkg/crypto/x509/#ExtKeyUsage -- simply drop the \"ExtKeyUsage\" part of the name. To remove all key usages from being set, set this value to an empty list. See also RFC 5280 Section 4.2.1.12.
	ExtKeyUsage []string `json:"ext_key_usage"`

	// A comma-separated string or list of extended key usage oids.
	ExtKeyUsageOids []string `json:"ext_key_usage_oids"`

	// If set, certificates issued/signed against this role will have Vault leases attached to them. Defaults to \"false\". Certificates can be added to the CRL by \"vault revoke <lease_id>\" when certificates are associated with leases. It can also be done using the \"pki/revoke\" endpoint. However, when lease generation is disabled, invoking \"pki/revoke\" would be the only way to add the certificates to the CRL. When large number of certificates are generated with long lifetimes, it is recommended that lease generation be disabled, as large amount of leases adversely affect the startup time of Vault.
	GenerateLease bool `json:"generate_lease"`

	// Reference to the issuer used to sign requests serviced by this role.
	IssuerRef string `json:"issuer_ref"`

	// The number of bits to use. Allowed values are 0 (universal default); with rsa key_type: 2048 (default), 3072, or 4096; with ec key_type: 224, 256 (default), 384, or 521; ignored with ed25519.
	KeyBits int32 `json:"key_bits"`

	// The type of key to use; defaults to RSA. \"rsa\" \"ec\", \"ed25519\" and \"any\" are the only valid values.
	KeyType string `json:"key_type"`

	// A comma-separated string or list of key usages (not extended key usages). Valid values can be found at https://golang.org/pkg/crypto/x509/#KeyUsage -- simply drop the \"KeyUsage\" part of the name. To remove all key usages from being set, set this value to an empty list. See also RFC 5280 Section 4.2.1.3.
	KeyUsage []string `json:"key_usage"`

	// If set, Locality will be set to this value in certificates issued by this role.
	Locality []string `json:"locality"`

	// The maximum allowed lease duration. If not set, defaults to the system maximum lease TTL.
	MaxTtl int32 `json:"max_ttl"`

	// If set, certificates issued/signed against this role will not be stored in the storage backend. This can improve performance when issuing large numbers of certificates. However, certificates issued in this way cannot be enumerated or revoked, so this option is recommended only for certificates that are non-sensitive, or extremely short-lived. This option implies a value of \"false\" for \"generate_lease\".
	NoStore bool `json:"no_store"`

	// Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ.
	NotAfter string `json:"not_after"`

	// The duration before now which the certificate needs to be backdated by.
	NotBeforeDuration int32 `json:"not_before_duration"`

	// If set, O (Organization) will be set to this value in certificates issued by this role.
	Organization []string `json:"organization"`

	// If set, OU (OrganizationalUnit) will be set to this value in certificates issued by this role.
	Ou []string `json:"ou"`

	// A comma-separated string or list of policy OIDs, or a JSON list of qualified policy information, which must include an oid, and may include a notice and/or cps url, using the form [{\"oid\"=\"1.3.6.1.4.1.7.8\",\"notice\"=\"I am a user Notice\"}, {\"oid\"=\"1.3.6.1.4.1.44947.1.2.4 \",\"cps\"=\"https://example.com\"}].
	PolicyIdentifiers []string `json:"policy_identifiers"`

	// If set, Postal Code will be set to this value in certificates issued by this role.
	PostalCode []string `json:"postal_code"`

	// If set, Province will be set to this value in certificates issued by this role.
	Province []string `json:"province"`

	// If set to false, makes the 'common_name' field optional while generating a certificate.
	RequireCn bool `json:"require_cn"`

	// If set, certificates are flagged for server auth use. Defaults to true. See also RFC 5280 Section 4.2.1.12.
	ServerFlag bool `json:"server_flag"`

	// The number of bits to use in the signature algorithm; accepts 256 for SHA-2-256, 384 for SHA-2-384, and 512 for SHA-2-512. Defaults to 0 to automatically detect based on key length (SHA-2-256 for RSA keys, and matching the curve size for NIST P-Curves).
	SignatureBits int32 `json:"signature_bits"`

	// If set, Street Address will be set to this value in certificates issued by this role.
	StreetAddress []string `json:"street_address"`

	// The lease duration (validity period of the certificate) if no specific lease duration is requested. The lease duration controls the expiration of certificates issued by this backend. Defaults to the system default value or the value of max_ttl, whichever is shorter.
	Ttl int32 `json:"ttl"`

	// If set, when used with a signing profile, the common name in the CSR will be used. This does *not* include any requested Subject Alternative Names; use use_csr_sans for that. Defaults to true.
	UseCsrCommonName bool `json:"use_csr_common_name"`

	// If set, when used with a signing profile, the SANs in the CSR will be used. This does *not* include the Common Name (cn); use use_csr_common_name for that. Defaults to true.
	UseCsrSans bool `json:"use_csr_sans"`

	// Whether or not to use PSS signatures when using a RSA key-type issuer. Defaults to false.
	UsePss bool `json:"use_pss"`
}

// NewPkiPatchRoleResponseWithDefaults instantiates a new PkiPatchRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkiPatchRoleResponseWithDefaults() *PkiPatchRoleResponse {
	var this PkiPatchRoleResponse

	this.ServerFlag = true

	return &this
}

func (o PkiPatchRoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["allow_any_name"] = o.AllowAnyName
	toSerialize["allow_bare_domains"] = o.AllowBareDomains
	toSerialize["allow_glob_domains"] = o.AllowGlobDomains
	toSerialize["allow_ip_sans"] = o.AllowIpSans
	toSerialize["allow_localhost"] = o.AllowLocalhost
	toSerialize["allow_subdomains"] = o.AllowSubdomains
	toSerialize["allow_token_displayname"] = o.AllowTokenDisplayname
	toSerialize["allow_wildcard_certificates"] = o.AllowWildcardCertificates
	toSerialize["allowed_domains"] = o.AllowedDomains
	toSerialize["allowed_domains_template"] = o.AllowedDomainsTemplate
	toSerialize["allowed_other_sans"] = o.AllowedOtherSans
	toSerialize["allowed_serial_numbers"] = o.AllowedSerialNumbers
	toSerialize["allowed_uri_sans"] = o.AllowedUriSans
	toSerialize["allowed_uri_sans_template"] = o.AllowedUriSansTemplate
	toSerialize["allowed_user_ids"] = o.AllowedUserIds
	toSerialize["basic_constraints_valid_for_non_ca"] = o.BasicConstraintsValidForNonCa
	toSerialize["client_flag"] = o.ClientFlag
	toSerialize["cn_validations"] = o.CnValidations
	toSerialize["code_signing_flag"] = o.CodeSigningFlag
	toSerialize["country"] = o.Country
	toSerialize["email_protection_flag"] = o.EmailProtectionFlag
	toSerialize["enforce_hostnames"] = o.EnforceHostnames
	toSerialize["ext_key_usage"] = o.ExtKeyUsage
	toSerialize["ext_key_usage_oids"] = o.ExtKeyUsageOids
	toSerialize["generate_lease"] = o.GenerateLease
	toSerialize["issuer_ref"] = o.IssuerRef
	toSerialize["key_bits"] = o.KeyBits
	toSerialize["key_type"] = o.KeyType
	toSerialize["key_usage"] = o.KeyUsage
	toSerialize["locality"] = o.Locality
	toSerialize["max_ttl"] = o.MaxTtl
	toSerialize["no_store"] = o.NoStore
	toSerialize["not_after"] = o.NotAfter
	toSerialize["not_before_duration"] = o.NotBeforeDuration
	toSerialize["organization"] = o.Organization
	toSerialize["ou"] = o.Ou
	toSerialize["policy_identifiers"] = o.PolicyIdentifiers
	toSerialize["postal_code"] = o.PostalCode
	toSerialize["province"] = o.Province
	toSerialize["require_cn"] = o.RequireCn
	toSerialize["server_flag"] = o.ServerFlag
	toSerialize["signature_bits"] = o.SignatureBits
	toSerialize["street_address"] = o.StreetAddress
	toSerialize["ttl"] = o.Ttl
	toSerialize["use_csr_common_name"] = o.UseCsrCommonName
	toSerialize["use_csr_sans"] = o.UseCsrSans
	toSerialize["use_pss"] = o.UsePss

	return json.Marshal(toSerialize)
}
