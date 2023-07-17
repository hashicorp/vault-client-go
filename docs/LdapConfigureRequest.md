# LdapConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnonymousGroupSearch** | Pointer to **bool** | Use anonymous binds when performing LDAP group searches (if true the initial credentials will still be used for the initial connection test). | [optional] [default to false]
**Binddn** | Pointer to **string** | LDAP DN for searching for the user DN (optional) | [optional] 
**Bindpass** | Pointer to **string** | LDAP password for searching for the user DN (optional) | [optional] 
**CaseSensitiveNames** | Pointer to **bool** | If true, case sensitivity will be used when comparing usernames and groups for matching policies. | [optional] 
**Certificate** | Pointer to **string** | CA certificate to use when verifying LDAP server certificate, must be x509 PEM encoded (optional) | [optional] 
**ClientTlsCert** | Pointer to **string** | Client certificate to provide to the LDAP server, must be x509 PEM encoded (optional) | [optional] 
**ClientTlsKey** | Pointer to **string** | Client certificate key to provide to the LDAP server, must be x509 PEM encoded (optional) | [optional] 
**ConnectionTimeout** | Pointer to **string** | Timeout, in seconds, when attempting to connect to the LDAP server before trying the next URL in the configuration. | [optional] [default to "30s"]
**DenyNullBind** | Pointer to **bool** | Denies an unauthenticated LDAP bind request if the user&#x27;s password is empty; defaults to true | [optional] [default to true]
**DereferenceAliases** | Pointer to **string** | When aliases should be dereferenced on search operations. Accepted values are &#x27;never&#x27;, &#x27;finding&#x27;, &#x27;searching&#x27;, &#x27;always&#x27;. Defaults to &#x27;never&#x27;. | [optional] [default to "never"]
**Discoverdn** | Pointer to **bool** | Use anonymous bind to discover the bind DN of a user (optional) | [optional] 
**Groupattr** | Pointer to **string** | LDAP attribute to follow on objects returned by &lt;groupfilter&gt; in order to enumerate user group membership. Examples: \&quot;cn\&quot; or \&quot;memberOf\&quot;, etc. Default: cn | [optional] [default to "cn"]
**Groupdn** | Pointer to **string** | LDAP search base to use for group membership search (eg: ou&#x3D;Groups,dc&#x3D;example,dc&#x3D;org) | [optional] 
**Groupfilter** | Pointer to **string** | Go template for querying group membership of user (optional) The template can access the following context variables: UserDN, Username Example: (&amp;(objectClass&#x3D;group)(member:1.2.840.113556.1.4.1941:&#x3D;{{.UserDN}})) Default: (|(memberUid&#x3D;{{.Username}})(member&#x3D;{{.UserDN}})(uniqueMember&#x3D;{{.UserDN}})) | [optional] [default to "(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))"]
**InsecureTls** | Pointer to **bool** | Skip LDAP server SSL Certificate verification - VERY insecure (optional) | [optional] 
**Length** | Pointer to **int32** | The desired length of passwords that Vault generates. | [optional] 
**MaxPageSize** | Pointer to **int32** | If set to a value greater than 0, the LDAP backend will use the LDAP server&#x27;s paged search control to request pages of up to the given size. This can be used to avoid hitting the LDAP server&#x27;s maximum result size limit. Otherwise, the LDAP backend will not use the paged search control. | [optional] [default to 0]
**MaxTtl** | Pointer to **string** | The maximum password time-to-live. | [optional] 
**PasswordPolicy** | Pointer to **string** | Password policy to use to generate passwords | [optional] 
**RequestTimeout** | Pointer to **string** | Timeout, in seconds, for the connection when making requests against the server before returning back an error. | [optional] [default to "90s"]
**Schema** | Pointer to **string** | The desired LDAP schema used when modifying user account passwords. | [optional] [default to "openldap"]
**Starttls** | Pointer to **bool** | Issue a StartTLS command after establishing unencrypted connection (optional) | [optional] 
**TlsMaxVersion** | Pointer to **string** | Maximum TLS version to use. Accepted values are &#x27;tls10&#x27;, &#x27;tls11&#x27;, &#x27;tls12&#x27; or &#x27;tls13&#x27;. Defaults to &#x27;tls12&#x27; | [optional] [default to "tls12"]
**TlsMinVersion** | Pointer to **string** | Minimum TLS version to use. Accepted values are &#x27;tls10&#x27;, &#x27;tls11&#x27;, &#x27;tls12&#x27; or &#x27;tls13&#x27;. Defaults to &#x27;tls12&#x27; | [optional] [default to "tls12"]
**Ttl** | Pointer to **string** | The default password time-to-live. | [optional] 
**Upndomain** | Pointer to **string** | Enables userPrincipalDomain login with [username]@UPNDomain (optional) | [optional] 
**Url** | Pointer to **string** | LDAP URL to connect to (default: ldap://127.0.0.1). Multiple URLs can be specified by concatenating them with commas; they will be tried in-order. | [optional] [default to "ldap://127.0.0.1"]
**UsePre111GroupCnBehavior** | Pointer to **bool** | In Vault 1.1.1 a fix for handling group CN values of different cases unfortunately introduced a regression that could cause previously defined groups to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for matching group CNs will be used. This is only needed in some upgrade scenarios for backwards compatibility. It is enabled by default if the config is upgraded but disabled by default on new configurations. | [optional] 
**UseTokenGroups** | Pointer to **bool** | If true, use the Active Directory tokenGroups constructed attribute of the user to find the group memberships. This will find all security groups including nested ones. | [optional] [default to false]
**Userattr** | Pointer to **string** | Attribute used for users (default: cn) | [optional] [default to "cn"]
**Userdn** | Pointer to **string** | LDAP domain to use for users (eg: ou&#x3D;People,dc&#x3D;example,dc&#x3D;org) | [optional] 
**Userfilter** | Pointer to **string** | Go template for LDAP user search filer (optional) The template can access the following context variables: UserAttr, Username Default: ({{.UserAttr}}&#x3D;{{.Username}}) | [optional] [default to "({{.UserAttr}}={{.Username}})"]
**UsernameAsAlias** | Pointer to **bool** | If true, sets the alias name to the username | [optional] [default to false]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


