// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LDAPWriteConfigRequest struct for LDAPWriteConfigRequest
type LDAPWriteConfigRequest struct {
	// Use anonymous binds when performing LDAP group searches (if true the initial credentials will still be used for the initial connection test).

	AnonymousGroupSearch bool `json:"anonymous_group_search"`

	// LDAP DN for searching for the user DN (optional)

	Binddn string `json:"binddn"`

	// LDAP password for searching for the user DN (optional)

	Bindpass string `json:"bindpass"`

	// If true, case sensitivity will be used when comparing usernames and groups for matching policies.

	CaseSensitiveNames bool `json:"case_sensitive_names"`

	// CA certificate to use when verifying LDAP server certificate, must be x509 PEM encoded (optional)

	Certificate string `json:"certificate"`

	// Client certificate to provide to the LDAP server, must be x509 PEM encoded (optional)

	ClientTlsCert string `json:"client_tls_cert"`

	// Client certificate key to provide to the LDAP server, must be x509 PEM encoded (optional)

	ClientTlsKey string `json:"client_tls_key"`

	// Denies an unauthenticated LDAP bind request if the user's password is empty; defaults to true

	DenyNullBind bool `json:"deny_null_bind"`

	// Use anonymous bind to discover the bind DN of a user (optional)

	Discoverdn bool `json:"discoverdn"`

	// LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate user group membership. Examples: \"cn\" or \"memberOf\", etc. Default: cn

	Groupattr string `json:"groupattr"`

	// LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org)

	Groupdn string `json:"groupdn"`

	// Go template for querying group membership of user (optional) The template can access the following context variables: UserDN, Username Example: (&(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}})) Default: (|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))

	Groupfilter string `json:"groupfilter"`

	// Skip LDAP server SSL Certificate verification - VERY insecure (optional)

	InsecureTls bool `json:"insecure_tls"`

	// The desired length of passwords that Vault generates.

	// Deprecated

	Length int32 `json:"length"`

	// The maximum password time-to-live.

	MaxTtl int32 `json:"max_ttl"`

	// Password policy to use to generate passwords

	PasswordPolicy string `json:"password_policy"`

	// Timeout, in seconds, for the connection when making requests against the server before returning back an error.

	RequestTimeout int32 `json:"request_timeout"`

	// The desired LDAP schema used when modifying user account passwords.

	Schema string `json:"schema"`

	// Issue a StartTLS command after establishing unencrypted connection (optional)

	Starttls bool `json:"starttls"`

	// Maximum TLS version to use. Accepted values are 'tls10', 'tls11', 'tls12' or 'tls13'. Defaults to 'tls12'

	TlsMaxVersion string `json:"tls_max_version"`

	// Minimum TLS version to use. Accepted values are 'tls10', 'tls11', 'tls12' or 'tls13'. Defaults to 'tls12'

	TlsMinVersion string `json:"tls_min_version"`

	// The default password time-to-live.

	Ttl int32 `json:"ttl"`

	// Enables userPrincipalDomain login with [username]@UPNDomain (optional)

	Upndomain string `json:"upndomain"`

	// LDAP URL to connect to (default: ldap://127.0.0.1). Multiple URLs can be specified by concatenating them with commas; they will be tried in-order.

	Url string `json:"url"`

	// In Vault 1.1.1 a fix for handling group CN values of different cases unfortunately introduced a regression that could cause previously defined groups to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for matching group CNs will be used. This is only needed in some upgrade scenarios for backwards compatibility. It is enabled by default if the config is upgraded but disabled by default on new configurations.

	UsePre111GroupCnBehavior bool `json:"use_pre111_group_cn_behavior"`

	// If true, use the Active Directory tokenGroups constructed attribute of the user to find the group memberships. This will find all security groups including nested ones.

	UseTokenGroups bool `json:"use_token_groups"`

	// Attribute used for users (default: cn)

	Userattr string `json:"userattr"`

	// LDAP domain to use for users (eg: ou=People,dc=example,dc=org)

	Userdn string `json:"userdn"`

	// Go template for LDAP user search filer (optional) The template can access the following context variables: UserAttr, Username Default: ({{.UserAttr}}={{.Username}})

	Userfilter string `json:"userfilter"`

	// If true, sets the alias name to the username

	UsernameAsAlias bool `json:"username_as_alias"`
}

// NewLDAPWriteConfigRequestWithDefaults instantiates a new LDAPWriteConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPWriteConfigRequestWithDefaults() *LDAPWriteConfigRequest {
	var this LDAPWriteConfigRequest

	this.AnonymousGroupSearch = false

	this.DenyNullBind = true

	this.Groupattr = "cn"

	this.Groupfilter = "(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))"

	this.Schema = "openldap"

	this.TlsMaxVersion = "tls12"

	this.TlsMinVersion = "tls12"

	this.Url = "ldap://127.0.0.1"

	this.UseTokenGroups = false

	this.Userattr = "cn"

	this.Userfilter = "({{.UserAttr}}={{.Username}})"

	this.UsernameAsAlias = false

	return &this
}

func (o LDAPWriteConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["anonymous_group_search"] = o.AnonymousGroupSearch

	toSerialize["binddn"] = o.Binddn

	toSerialize["bindpass"] = o.Bindpass

	toSerialize["case_sensitive_names"] = o.CaseSensitiveNames

	toSerialize["certificate"] = o.Certificate

	toSerialize["client_tls_cert"] = o.ClientTlsCert

	toSerialize["client_tls_key"] = o.ClientTlsKey

	toSerialize["deny_null_bind"] = o.DenyNullBind

	toSerialize["discoverdn"] = o.Discoverdn

	toSerialize["groupattr"] = o.Groupattr

	toSerialize["groupdn"] = o.Groupdn

	toSerialize["groupfilter"] = o.Groupfilter

	toSerialize["insecure_tls"] = o.InsecureTls

	toSerialize["length"] = o.Length

	toSerialize["max_ttl"] = o.MaxTtl

	toSerialize["password_policy"] = o.PasswordPolicy

	toSerialize["request_timeout"] = o.RequestTimeout

	toSerialize["schema"] = o.Schema

	toSerialize["starttls"] = o.Starttls

	toSerialize["tls_max_version"] = o.TlsMaxVersion

	toSerialize["tls_min_version"] = o.TlsMinVersion

	toSerialize["ttl"] = o.Ttl

	toSerialize["upndomain"] = o.Upndomain

	toSerialize["url"] = o.Url

	toSerialize["use_pre111_group_cn_behavior"] = o.UsePre111GroupCnBehavior

	toSerialize["use_token_groups"] = o.UseTokenGroups

	toSerialize["userattr"] = o.Userattr

	toSerialize["userdn"] = o.Userdn

	toSerialize["userfilter"] = o.Userfilter

	toSerialize["username_as_alias"] = o.UsernameAsAlias

	return json.Marshal(toSerialize)
}
