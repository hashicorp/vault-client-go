# OpenLDAPWriteConfigRequest

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
**DenyNullBind** | Pointer to **bool** | Denies an unauthenticated LDAP bind request if the user&#39;s password is empty; defaults to true | [optional] [default to true]
**Discoverdn** | Pointer to **bool** | Use anonymous bind to discover the bind DN of a user (optional) | [optional] 
**Groupattr** | Pointer to **string** | LDAP attribute to follow on objects returned by &lt;groupfilter&gt; in order to enumerate user group membership. Examples: \&quot;cn\&quot; or \&quot;memberOf\&quot;, etc. Default: cn | [optional] [default to "cn"]
**Groupdn** | Pointer to **string** | LDAP search base to use for group membership search (eg: ou&#x3D;Groups,dc&#x3D;example,dc&#x3D;org) | [optional] 
**Groupfilter** | Pointer to **string** | Go template for querying group membership of user (optional) The template can access the following context variables: UserDN, Username Example: (&amp;(objectClass&#x3D;group)(member:1.2.840.113556.1.4.1941:&#x3D;{{.UserDN}})) Default: (|(memberUid&#x3D;{{.Username}})(member&#x3D;{{.UserDN}})(uniqueMember&#x3D;{{.UserDN}})) | [optional] [default to "(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))"]
**InsecureTls** | Pointer to **bool** | Skip LDAP server SSL Certificate verification - VERY insecure (optional) | [optional] 
**Length** | Pointer to **int32** | The desired length of passwords that Vault generates. | [optional] 
**MaxTtl** | Pointer to **int32** | The maximum password time-to-live. | [optional] 
**PasswordPolicy** | Pointer to **string** | Password policy to use to generate passwords | [optional] 
**RequestTimeout** | Pointer to **int32** | Timeout, in seconds, for the connection when making requests against the server before returning back an error. | [optional] 
**Schema** | Pointer to **string** | The desired LDAP schema used when modifying user account passwords. | [optional] [default to "openldap"]
**Starttls** | Pointer to **bool** | Issue a StartTLS command after establishing unencrypted connection (optional) | [optional] 
**TlsMaxVersion** | Pointer to **string** | Maximum TLS version to use. Accepted values are &#39;tls10&#39;, &#39;tls11&#39;, &#39;tls12&#39; or &#39;tls13&#39;. Defaults to &#39;tls12&#39; | [optional] [default to "tls12"]
**TlsMinVersion** | Pointer to **string** | Minimum TLS version to use. Accepted values are &#39;tls10&#39;, &#39;tls11&#39;, &#39;tls12&#39; or &#39;tls13&#39;. Defaults to &#39;tls12&#39; | [optional] [default to "tls12"]
**Ttl** | Pointer to **int32** | The default password time-to-live. | [optional] 
**Upndomain** | Pointer to **string** | Enables userPrincipalDomain login with [username]@UPNDomain (optional) | [optional] 
**Url** | Pointer to **string** | LDAP URL to connect to (default: ldap://127.0.0.1). Multiple URLs can be specified by concatenating them with commas; they will be tried in-order. | [optional] [default to "ldap://127.0.0.1"]
**UsePre111GroupCnBehavior** | Pointer to **bool** | In Vault 1.1.1 a fix for handling group CN values of different cases unfortunately introduced a regression that could cause previously defined groups to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for matching group CNs will be used. This is only needed in some upgrade scenarios for backwards compatibility. It is enabled by default if the config is upgraded but disabled by default on new configurations. | [optional] 
**UseTokenGroups** | Pointer to **bool** | If true, use the Active Directory tokenGroups constructed attribute of the user to find the group memberships. This will find all security groups including nested ones. | [optional] [default to false]
**Userattr** | Pointer to **string** | Attribute used for users (default: cn) | [optional] [default to "cn"]
**Userdn** | Pointer to **string** | LDAP domain to use for users (eg: ou&#x3D;People,dc&#x3D;example,dc&#x3D;org) | [optional] 
**Userfilter** | Pointer to **string** | Go template for LDAP user search filer (optional) The template can access the following context variables: UserAttr, Username Default: ({{.UserAttr}}&#x3D;{{.Username}}) | [optional] [default to "({{.UserAttr}}={{.Username}})"]
**UsernameAsAlias** | Pointer to **bool** | If true, sets the alias name to the username | [optional] [default to false]

## Methods

### NewOpenLDAPWriteConfigRequest

`func NewOpenLDAPWriteConfigRequest() *OpenLDAPWriteConfigRequest`

NewOpenLDAPWriteConfigRequest instantiates a new OpenLDAPWriteConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenLDAPWriteConfigRequestWithDefaults

`func NewOpenLDAPWriteConfigRequestWithDefaults() *OpenLDAPWriteConfigRequest`

NewOpenLDAPWriteConfigRequestWithDefaults instantiates a new OpenLDAPWriteConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymousGroupSearch

`func (o *OpenLDAPWriteConfigRequest) GetAnonymousGroupSearch() bool`

GetAnonymousGroupSearch returns the AnonymousGroupSearch field if non-nil, zero value otherwise.

### GetAnonymousGroupSearchOk

`func (o *OpenLDAPWriteConfigRequest) GetAnonymousGroupSearchOk() (*bool, bool)`

GetAnonymousGroupSearchOk returns a tuple with the AnonymousGroupSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousGroupSearch

`func (o *OpenLDAPWriteConfigRequest) SetAnonymousGroupSearch(v bool)`

SetAnonymousGroupSearch sets AnonymousGroupSearch field to given value.

### HasAnonymousGroupSearch

`func (o *OpenLDAPWriteConfigRequest) HasAnonymousGroupSearch() bool`

HasAnonymousGroupSearch returns a boolean if a field has been set.

### GetBinddn

`func (o *OpenLDAPWriteConfigRequest) GetBinddn() string`

GetBinddn returns the Binddn field if non-nil, zero value otherwise.

### GetBinddnOk

`func (o *OpenLDAPWriteConfigRequest) GetBinddnOk() (*string, bool)`

GetBinddnOk returns a tuple with the Binddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinddn

`func (o *OpenLDAPWriteConfigRequest) SetBinddn(v string)`

SetBinddn sets Binddn field to given value.

### HasBinddn

`func (o *OpenLDAPWriteConfigRequest) HasBinddn() bool`

HasBinddn returns a boolean if a field has been set.

### GetBindpass

`func (o *OpenLDAPWriteConfigRequest) GetBindpass() string`

GetBindpass returns the Bindpass field if non-nil, zero value otherwise.

### GetBindpassOk

`func (o *OpenLDAPWriteConfigRequest) GetBindpassOk() (*string, bool)`

GetBindpassOk returns a tuple with the Bindpass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindpass

`func (o *OpenLDAPWriteConfigRequest) SetBindpass(v string)`

SetBindpass sets Bindpass field to given value.

### HasBindpass

`func (o *OpenLDAPWriteConfigRequest) HasBindpass() bool`

HasBindpass returns a boolean if a field has been set.

### GetCaseSensitiveNames

`func (o *OpenLDAPWriteConfigRequest) GetCaseSensitiveNames() bool`

GetCaseSensitiveNames returns the CaseSensitiveNames field if non-nil, zero value otherwise.

### GetCaseSensitiveNamesOk

`func (o *OpenLDAPWriteConfigRequest) GetCaseSensitiveNamesOk() (*bool, bool)`

GetCaseSensitiveNamesOk returns a tuple with the CaseSensitiveNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitiveNames

`func (o *OpenLDAPWriteConfigRequest) SetCaseSensitiveNames(v bool)`

SetCaseSensitiveNames sets CaseSensitiveNames field to given value.

### HasCaseSensitiveNames

`func (o *OpenLDAPWriteConfigRequest) HasCaseSensitiveNames() bool`

HasCaseSensitiveNames returns a boolean if a field has been set.

### GetCertificate

`func (o *OpenLDAPWriteConfigRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *OpenLDAPWriteConfigRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *OpenLDAPWriteConfigRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *OpenLDAPWriteConfigRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetClientTlsCert

`func (o *OpenLDAPWriteConfigRequest) GetClientTlsCert() string`

GetClientTlsCert returns the ClientTlsCert field if non-nil, zero value otherwise.

### GetClientTlsCertOk

`func (o *OpenLDAPWriteConfigRequest) GetClientTlsCertOk() (*string, bool)`

GetClientTlsCertOk returns a tuple with the ClientTlsCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsCert

`func (o *OpenLDAPWriteConfigRequest) SetClientTlsCert(v string)`

SetClientTlsCert sets ClientTlsCert field to given value.

### HasClientTlsCert

`func (o *OpenLDAPWriteConfigRequest) HasClientTlsCert() bool`

HasClientTlsCert returns a boolean if a field has been set.

### GetClientTlsKey

`func (o *OpenLDAPWriteConfigRequest) GetClientTlsKey() string`

GetClientTlsKey returns the ClientTlsKey field if non-nil, zero value otherwise.

### GetClientTlsKeyOk

`func (o *OpenLDAPWriteConfigRequest) GetClientTlsKeyOk() (*string, bool)`

GetClientTlsKeyOk returns a tuple with the ClientTlsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsKey

`func (o *OpenLDAPWriteConfigRequest) SetClientTlsKey(v string)`

SetClientTlsKey sets ClientTlsKey field to given value.

### HasClientTlsKey

`func (o *OpenLDAPWriteConfigRequest) HasClientTlsKey() bool`

HasClientTlsKey returns a boolean if a field has been set.

### GetDenyNullBind

`func (o *OpenLDAPWriteConfigRequest) GetDenyNullBind() bool`

GetDenyNullBind returns the DenyNullBind field if non-nil, zero value otherwise.

### GetDenyNullBindOk

`func (o *OpenLDAPWriteConfigRequest) GetDenyNullBindOk() (*bool, bool)`

GetDenyNullBindOk returns a tuple with the DenyNullBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyNullBind

`func (o *OpenLDAPWriteConfigRequest) SetDenyNullBind(v bool)`

SetDenyNullBind sets DenyNullBind field to given value.

### HasDenyNullBind

`func (o *OpenLDAPWriteConfigRequest) HasDenyNullBind() bool`

HasDenyNullBind returns a boolean if a field has been set.

### GetDiscoverdn

`func (o *OpenLDAPWriteConfigRequest) GetDiscoverdn() bool`

GetDiscoverdn returns the Discoverdn field if non-nil, zero value otherwise.

### GetDiscoverdnOk

`func (o *OpenLDAPWriteConfigRequest) GetDiscoverdnOk() (*bool, bool)`

GetDiscoverdnOk returns a tuple with the Discoverdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverdn

`func (o *OpenLDAPWriteConfigRequest) SetDiscoverdn(v bool)`

SetDiscoverdn sets Discoverdn field to given value.

### HasDiscoverdn

`func (o *OpenLDAPWriteConfigRequest) HasDiscoverdn() bool`

HasDiscoverdn returns a boolean if a field has been set.

### GetGroupattr

`func (o *OpenLDAPWriteConfigRequest) GetGroupattr() string`

GetGroupattr returns the Groupattr field if non-nil, zero value otherwise.

### GetGroupattrOk

`func (o *OpenLDAPWriteConfigRequest) GetGroupattrOk() (*string, bool)`

GetGroupattrOk returns a tuple with the Groupattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupattr

`func (o *OpenLDAPWriteConfigRequest) SetGroupattr(v string)`

SetGroupattr sets Groupattr field to given value.

### HasGroupattr

`func (o *OpenLDAPWriteConfigRequest) HasGroupattr() bool`

HasGroupattr returns a boolean if a field has been set.

### GetGroupdn

`func (o *OpenLDAPWriteConfigRequest) GetGroupdn() string`

GetGroupdn returns the Groupdn field if non-nil, zero value otherwise.

### GetGroupdnOk

`func (o *OpenLDAPWriteConfigRequest) GetGroupdnOk() (*string, bool)`

GetGroupdnOk returns a tuple with the Groupdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupdn

`func (o *OpenLDAPWriteConfigRequest) SetGroupdn(v string)`

SetGroupdn sets Groupdn field to given value.

### HasGroupdn

`func (o *OpenLDAPWriteConfigRequest) HasGroupdn() bool`

HasGroupdn returns a boolean if a field has been set.

### GetGroupfilter

`func (o *OpenLDAPWriteConfigRequest) GetGroupfilter() string`

GetGroupfilter returns the Groupfilter field if non-nil, zero value otherwise.

### GetGroupfilterOk

`func (o *OpenLDAPWriteConfigRequest) GetGroupfilterOk() (*string, bool)`

GetGroupfilterOk returns a tuple with the Groupfilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupfilter

`func (o *OpenLDAPWriteConfigRequest) SetGroupfilter(v string)`

SetGroupfilter sets Groupfilter field to given value.

### HasGroupfilter

`func (o *OpenLDAPWriteConfigRequest) HasGroupfilter() bool`

HasGroupfilter returns a boolean if a field has been set.

### GetInsecureTls

`func (o *OpenLDAPWriteConfigRequest) GetInsecureTls() bool`

GetInsecureTls returns the InsecureTls field if non-nil, zero value otherwise.

### GetInsecureTlsOk

`func (o *OpenLDAPWriteConfigRequest) GetInsecureTlsOk() (*bool, bool)`

GetInsecureTlsOk returns a tuple with the InsecureTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureTls

`func (o *OpenLDAPWriteConfigRequest) SetInsecureTls(v bool)`

SetInsecureTls sets InsecureTls field to given value.

### HasInsecureTls

`func (o *OpenLDAPWriteConfigRequest) HasInsecureTls() bool`

HasInsecureTls returns a boolean if a field has been set.

### GetLength

`func (o *OpenLDAPWriteConfigRequest) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *OpenLDAPWriteConfigRequest) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *OpenLDAPWriteConfigRequest) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *OpenLDAPWriteConfigRequest) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetMaxTtl

`func (o *OpenLDAPWriteConfigRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *OpenLDAPWriteConfigRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *OpenLDAPWriteConfigRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *OpenLDAPWriteConfigRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetPasswordPolicy

`func (o *OpenLDAPWriteConfigRequest) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *OpenLDAPWriteConfigRequest) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *OpenLDAPWriteConfigRequest) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *OpenLDAPWriteConfigRequest) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### GetRequestTimeout

`func (o *OpenLDAPWriteConfigRequest) GetRequestTimeout() int32`

GetRequestTimeout returns the RequestTimeout field if non-nil, zero value otherwise.

### GetRequestTimeoutOk

`func (o *OpenLDAPWriteConfigRequest) GetRequestTimeoutOk() (*int32, bool)`

GetRequestTimeoutOk returns a tuple with the RequestTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimeout

`func (o *OpenLDAPWriteConfigRequest) SetRequestTimeout(v int32)`

SetRequestTimeout sets RequestTimeout field to given value.

### HasRequestTimeout

`func (o *OpenLDAPWriteConfigRequest) HasRequestTimeout() bool`

HasRequestTimeout returns a boolean if a field has been set.

### GetSchema

`func (o *OpenLDAPWriteConfigRequest) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OpenLDAPWriteConfigRequest) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OpenLDAPWriteConfigRequest) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OpenLDAPWriteConfigRequest) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetStarttls

`func (o *OpenLDAPWriteConfigRequest) GetStarttls() bool`

GetStarttls returns the Starttls field if non-nil, zero value otherwise.

### GetStarttlsOk

`func (o *OpenLDAPWriteConfigRequest) GetStarttlsOk() (*bool, bool)`

GetStarttlsOk returns a tuple with the Starttls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarttls

`func (o *OpenLDAPWriteConfigRequest) SetStarttls(v bool)`

SetStarttls sets Starttls field to given value.

### HasStarttls

`func (o *OpenLDAPWriteConfigRequest) HasStarttls() bool`

HasStarttls returns a boolean if a field has been set.

### GetTlsMaxVersion

`func (o *OpenLDAPWriteConfigRequest) GetTlsMaxVersion() string`

GetTlsMaxVersion returns the TlsMaxVersion field if non-nil, zero value otherwise.

### GetTlsMaxVersionOk

`func (o *OpenLDAPWriteConfigRequest) GetTlsMaxVersionOk() (*string, bool)`

GetTlsMaxVersionOk returns a tuple with the TlsMaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMaxVersion

`func (o *OpenLDAPWriteConfigRequest) SetTlsMaxVersion(v string)`

SetTlsMaxVersion sets TlsMaxVersion field to given value.

### HasTlsMaxVersion

`func (o *OpenLDAPWriteConfigRequest) HasTlsMaxVersion() bool`

HasTlsMaxVersion returns a boolean if a field has been set.

### GetTlsMinVersion

`func (o *OpenLDAPWriteConfigRequest) GetTlsMinVersion() string`

GetTlsMinVersion returns the TlsMinVersion field if non-nil, zero value otherwise.

### GetTlsMinVersionOk

`func (o *OpenLDAPWriteConfigRequest) GetTlsMinVersionOk() (*string, bool)`

GetTlsMinVersionOk returns a tuple with the TlsMinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMinVersion

`func (o *OpenLDAPWriteConfigRequest) SetTlsMinVersion(v string)`

SetTlsMinVersion sets TlsMinVersion field to given value.

### HasTlsMinVersion

`func (o *OpenLDAPWriteConfigRequest) HasTlsMinVersion() bool`

HasTlsMinVersion returns a boolean if a field has been set.

### GetTtl

`func (o *OpenLDAPWriteConfigRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *OpenLDAPWriteConfigRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *OpenLDAPWriteConfigRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *OpenLDAPWriteConfigRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUpndomain

`func (o *OpenLDAPWriteConfigRequest) GetUpndomain() string`

GetUpndomain returns the Upndomain field if non-nil, zero value otherwise.

### GetUpndomainOk

`func (o *OpenLDAPWriteConfigRequest) GetUpndomainOk() (*string, bool)`

GetUpndomainOk returns a tuple with the Upndomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpndomain

`func (o *OpenLDAPWriteConfigRequest) SetUpndomain(v string)`

SetUpndomain sets Upndomain field to given value.

### HasUpndomain

`func (o *OpenLDAPWriteConfigRequest) HasUpndomain() bool`

HasUpndomain returns a boolean if a field has been set.

### GetUrl

`func (o *OpenLDAPWriteConfigRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OpenLDAPWriteConfigRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OpenLDAPWriteConfigRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OpenLDAPWriteConfigRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsePre111GroupCnBehavior

`func (o *OpenLDAPWriteConfigRequest) GetUsePre111GroupCnBehavior() bool`

GetUsePre111GroupCnBehavior returns the UsePre111GroupCnBehavior field if non-nil, zero value otherwise.

### GetUsePre111GroupCnBehaviorOk

`func (o *OpenLDAPWriteConfigRequest) GetUsePre111GroupCnBehaviorOk() (*bool, bool)`

GetUsePre111GroupCnBehaviorOk returns a tuple with the UsePre111GroupCnBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePre111GroupCnBehavior

`func (o *OpenLDAPWriteConfigRequest) SetUsePre111GroupCnBehavior(v bool)`

SetUsePre111GroupCnBehavior sets UsePre111GroupCnBehavior field to given value.

### HasUsePre111GroupCnBehavior

`func (o *OpenLDAPWriteConfigRequest) HasUsePre111GroupCnBehavior() bool`

HasUsePre111GroupCnBehavior returns a boolean if a field has been set.

### GetUseTokenGroups

`func (o *OpenLDAPWriteConfigRequest) GetUseTokenGroups() bool`

GetUseTokenGroups returns the UseTokenGroups field if non-nil, zero value otherwise.

### GetUseTokenGroupsOk

`func (o *OpenLDAPWriteConfigRequest) GetUseTokenGroupsOk() (*bool, bool)`

GetUseTokenGroupsOk returns a tuple with the UseTokenGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTokenGroups

`func (o *OpenLDAPWriteConfigRequest) SetUseTokenGroups(v bool)`

SetUseTokenGroups sets UseTokenGroups field to given value.

### HasUseTokenGroups

`func (o *OpenLDAPWriteConfigRequest) HasUseTokenGroups() bool`

HasUseTokenGroups returns a boolean if a field has been set.

### GetUserattr

`func (o *OpenLDAPWriteConfigRequest) GetUserattr() string`

GetUserattr returns the Userattr field if non-nil, zero value otherwise.

### GetUserattrOk

`func (o *OpenLDAPWriteConfigRequest) GetUserattrOk() (*string, bool)`

GetUserattrOk returns a tuple with the Userattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserattr

`func (o *OpenLDAPWriteConfigRequest) SetUserattr(v string)`

SetUserattr sets Userattr field to given value.

### HasUserattr

`func (o *OpenLDAPWriteConfigRequest) HasUserattr() bool`

HasUserattr returns a boolean if a field has been set.

### GetUserdn

`func (o *OpenLDAPWriteConfigRequest) GetUserdn() string`

GetUserdn returns the Userdn field if non-nil, zero value otherwise.

### GetUserdnOk

`func (o *OpenLDAPWriteConfigRequest) GetUserdnOk() (*string, bool)`

GetUserdnOk returns a tuple with the Userdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdn

`func (o *OpenLDAPWriteConfigRequest) SetUserdn(v string)`

SetUserdn sets Userdn field to given value.

### HasUserdn

`func (o *OpenLDAPWriteConfigRequest) HasUserdn() bool`

HasUserdn returns a boolean if a field has been set.

### GetUserfilter

`func (o *OpenLDAPWriteConfigRequest) GetUserfilter() string`

GetUserfilter returns the Userfilter field if non-nil, zero value otherwise.

### GetUserfilterOk

`func (o *OpenLDAPWriteConfigRequest) GetUserfilterOk() (*string, bool)`

GetUserfilterOk returns a tuple with the Userfilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserfilter

`func (o *OpenLDAPWriteConfigRequest) SetUserfilter(v string)`

SetUserfilter sets Userfilter field to given value.

### HasUserfilter

`func (o *OpenLDAPWriteConfigRequest) HasUserfilter() bool`

HasUserfilter returns a boolean if a field has been set.

### GetUsernameAsAlias

`func (o *OpenLDAPWriteConfigRequest) GetUsernameAsAlias() bool`

GetUsernameAsAlias returns the UsernameAsAlias field if non-nil, zero value otherwise.

### GetUsernameAsAliasOk

`func (o *OpenLDAPWriteConfigRequest) GetUsernameAsAliasOk() (*bool, bool)`

GetUsernameAsAliasOk returns a tuple with the UsernameAsAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameAsAlias

`func (o *OpenLDAPWriteConfigRequest) SetUsernameAsAlias(v bool)`

SetUsernameAsAlias sets UsernameAsAlias field to given value.

### HasUsernameAsAlias

`func (o *OpenLDAPWriteConfigRequest) HasUsernameAsAlias() bool`

HasUsernameAsAlias returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

