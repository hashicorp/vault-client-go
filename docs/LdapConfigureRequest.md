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



## Methods


### NewLdapConfigureRequest

`func NewLdapConfigureRequest() *LdapConfigureRequest`

NewLdapConfigureRequest instantiates a new LdapConfigureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapConfigureRequestWithDefaults

`func NewLdapConfigureRequestWithDefaults() *LdapConfigureRequest`

NewLdapConfigureRequestWithDefaults instantiates a new LdapConfigureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAnonymousGroupSearch

`func (o *LdapConfigureRequest) GetAnonymousGroupSearch() bool`

GetAnonymousGroupSearch returns the AnonymousGroupSearch field if non-nil, zero value otherwise.

### GetAnonymousGroupSearchOk

`func (o *LdapConfigureRequest) GetAnonymousGroupSearchOk() (*bool, bool)`

GetAnonymousGroupSearchOk returns a tuple with the AnonymousGroupSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousGroupSearch

`func (o *LdapConfigureRequest) SetAnonymousGroupSearch(v bool)`

SetAnonymousGroupSearch sets AnonymousGroupSearch field to given value.


### HasAnonymousGroupSearch

`func (o *LdapConfigureRequest) HasAnonymousGroupSearch() bool`

HasAnonymousGroupSearch returns a boolean if a field has been set.




### GetBinddn

`func (o *LdapConfigureRequest) GetBinddn() string`

GetBinddn returns the Binddn field if non-nil, zero value otherwise.

### GetBinddnOk

`func (o *LdapConfigureRequest) GetBinddnOk() (*string, bool)`

GetBinddnOk returns a tuple with the Binddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinddn

`func (o *LdapConfigureRequest) SetBinddn(v string)`

SetBinddn sets Binddn field to given value.


### HasBinddn

`func (o *LdapConfigureRequest) HasBinddn() bool`

HasBinddn returns a boolean if a field has been set.




### GetBindpass

`func (o *LdapConfigureRequest) GetBindpass() string`

GetBindpass returns the Bindpass field if non-nil, zero value otherwise.

### GetBindpassOk

`func (o *LdapConfigureRequest) GetBindpassOk() (*string, bool)`

GetBindpassOk returns a tuple with the Bindpass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindpass

`func (o *LdapConfigureRequest) SetBindpass(v string)`

SetBindpass sets Bindpass field to given value.


### HasBindpass

`func (o *LdapConfigureRequest) HasBindpass() bool`

HasBindpass returns a boolean if a field has been set.




### GetCaseSensitiveNames

`func (o *LdapConfigureRequest) GetCaseSensitiveNames() bool`

GetCaseSensitiveNames returns the CaseSensitiveNames field if non-nil, zero value otherwise.

### GetCaseSensitiveNamesOk

`func (o *LdapConfigureRequest) GetCaseSensitiveNamesOk() (*bool, bool)`

GetCaseSensitiveNamesOk returns a tuple with the CaseSensitiveNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitiveNames

`func (o *LdapConfigureRequest) SetCaseSensitiveNames(v bool)`

SetCaseSensitiveNames sets CaseSensitiveNames field to given value.


### HasCaseSensitiveNames

`func (o *LdapConfigureRequest) HasCaseSensitiveNames() bool`

HasCaseSensitiveNames returns a boolean if a field has been set.




### GetCertificate

`func (o *LdapConfigureRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *LdapConfigureRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *LdapConfigureRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *LdapConfigureRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.




### GetClientTlsCert

`func (o *LdapConfigureRequest) GetClientTlsCert() string`

GetClientTlsCert returns the ClientTlsCert field if non-nil, zero value otherwise.

### GetClientTlsCertOk

`func (o *LdapConfigureRequest) GetClientTlsCertOk() (*string, bool)`

GetClientTlsCertOk returns a tuple with the ClientTlsCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsCert

`func (o *LdapConfigureRequest) SetClientTlsCert(v string)`

SetClientTlsCert sets ClientTlsCert field to given value.


### HasClientTlsCert

`func (o *LdapConfigureRequest) HasClientTlsCert() bool`

HasClientTlsCert returns a boolean if a field has been set.




### GetClientTlsKey

`func (o *LdapConfigureRequest) GetClientTlsKey() string`

GetClientTlsKey returns the ClientTlsKey field if non-nil, zero value otherwise.

### GetClientTlsKeyOk

`func (o *LdapConfigureRequest) GetClientTlsKeyOk() (*string, bool)`

GetClientTlsKeyOk returns a tuple with the ClientTlsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsKey

`func (o *LdapConfigureRequest) SetClientTlsKey(v string)`

SetClientTlsKey sets ClientTlsKey field to given value.


### HasClientTlsKey

`func (o *LdapConfigureRequest) HasClientTlsKey() bool`

HasClientTlsKey returns a boolean if a field has been set.




### GetConnectionTimeout

`func (o *LdapConfigureRequest) GetConnectionTimeout() string`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *LdapConfigureRequest) GetConnectionTimeoutOk() (*string, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *LdapConfigureRequest) SetConnectionTimeout(v string)`

SetConnectionTimeout sets ConnectionTimeout field to given value.


### HasConnectionTimeout

`func (o *LdapConfigureRequest) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.




### GetDenyNullBind

`func (o *LdapConfigureRequest) GetDenyNullBind() bool`

GetDenyNullBind returns the DenyNullBind field if non-nil, zero value otherwise.

### GetDenyNullBindOk

`func (o *LdapConfigureRequest) GetDenyNullBindOk() (*bool, bool)`

GetDenyNullBindOk returns a tuple with the DenyNullBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyNullBind

`func (o *LdapConfigureRequest) SetDenyNullBind(v bool)`

SetDenyNullBind sets DenyNullBind field to given value.


### HasDenyNullBind

`func (o *LdapConfigureRequest) HasDenyNullBind() bool`

HasDenyNullBind returns a boolean if a field has been set.




### GetDereferenceAliases

`func (o *LdapConfigureRequest) GetDereferenceAliases() string`

GetDereferenceAliases returns the DereferenceAliases field if non-nil, zero value otherwise.

### GetDereferenceAliasesOk

`func (o *LdapConfigureRequest) GetDereferenceAliasesOk() (*string, bool)`

GetDereferenceAliasesOk returns a tuple with the DereferenceAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDereferenceAliases

`func (o *LdapConfigureRequest) SetDereferenceAliases(v string)`

SetDereferenceAliases sets DereferenceAliases field to given value.


### HasDereferenceAliases

`func (o *LdapConfigureRequest) HasDereferenceAliases() bool`

HasDereferenceAliases returns a boolean if a field has been set.




### GetDiscoverdn

`func (o *LdapConfigureRequest) GetDiscoverdn() bool`

GetDiscoverdn returns the Discoverdn field if non-nil, zero value otherwise.

### GetDiscoverdnOk

`func (o *LdapConfigureRequest) GetDiscoverdnOk() (*bool, bool)`

GetDiscoverdnOk returns a tuple with the Discoverdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverdn

`func (o *LdapConfigureRequest) SetDiscoverdn(v bool)`

SetDiscoverdn sets Discoverdn field to given value.


### HasDiscoverdn

`func (o *LdapConfigureRequest) HasDiscoverdn() bool`

HasDiscoverdn returns a boolean if a field has been set.




### GetGroupattr

`func (o *LdapConfigureRequest) GetGroupattr() string`

GetGroupattr returns the Groupattr field if non-nil, zero value otherwise.

### GetGroupattrOk

`func (o *LdapConfigureRequest) GetGroupattrOk() (*string, bool)`

GetGroupattrOk returns a tuple with the Groupattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupattr

`func (o *LdapConfigureRequest) SetGroupattr(v string)`

SetGroupattr sets Groupattr field to given value.


### HasGroupattr

`func (o *LdapConfigureRequest) HasGroupattr() bool`

HasGroupattr returns a boolean if a field has been set.




### GetGroupdn

`func (o *LdapConfigureRequest) GetGroupdn() string`

GetGroupdn returns the Groupdn field if non-nil, zero value otherwise.

### GetGroupdnOk

`func (o *LdapConfigureRequest) GetGroupdnOk() (*string, bool)`

GetGroupdnOk returns a tuple with the Groupdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupdn

`func (o *LdapConfigureRequest) SetGroupdn(v string)`

SetGroupdn sets Groupdn field to given value.


### HasGroupdn

`func (o *LdapConfigureRequest) HasGroupdn() bool`

HasGroupdn returns a boolean if a field has been set.




### GetGroupfilter

`func (o *LdapConfigureRequest) GetGroupfilter() string`

GetGroupfilter returns the Groupfilter field if non-nil, zero value otherwise.

### GetGroupfilterOk

`func (o *LdapConfigureRequest) GetGroupfilterOk() (*string, bool)`

GetGroupfilterOk returns a tuple with the Groupfilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupfilter

`func (o *LdapConfigureRequest) SetGroupfilter(v string)`

SetGroupfilter sets Groupfilter field to given value.


### HasGroupfilter

`func (o *LdapConfigureRequest) HasGroupfilter() bool`

HasGroupfilter returns a boolean if a field has been set.




### GetInsecureTls

`func (o *LdapConfigureRequest) GetInsecureTls() bool`

GetInsecureTls returns the InsecureTls field if non-nil, zero value otherwise.

### GetInsecureTlsOk

`func (o *LdapConfigureRequest) GetInsecureTlsOk() (*bool, bool)`

GetInsecureTlsOk returns a tuple with the InsecureTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureTls

`func (o *LdapConfigureRequest) SetInsecureTls(v bool)`

SetInsecureTls sets InsecureTls field to given value.


### HasInsecureTls

`func (o *LdapConfigureRequest) HasInsecureTls() bool`

HasInsecureTls returns a boolean if a field has been set.




### GetLength

`func (o *LdapConfigureRequest) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *LdapConfigureRequest) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *LdapConfigureRequest) SetLength(v int32)`

SetLength sets Length field to given value.


### HasLength

`func (o *LdapConfigureRequest) HasLength() bool`

HasLength returns a boolean if a field has been set.




### GetMaxPageSize

`func (o *LdapConfigureRequest) GetMaxPageSize() int32`

GetMaxPageSize returns the MaxPageSize field if non-nil, zero value otherwise.

### GetMaxPageSizeOk

`func (o *LdapConfigureRequest) GetMaxPageSizeOk() (*int32, bool)`

GetMaxPageSizeOk returns a tuple with the MaxPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPageSize

`func (o *LdapConfigureRequest) SetMaxPageSize(v int32)`

SetMaxPageSize sets MaxPageSize field to given value.


### HasMaxPageSize

`func (o *LdapConfigureRequest) HasMaxPageSize() bool`

HasMaxPageSize returns a boolean if a field has been set.




### GetMaxTtl

`func (o *LdapConfigureRequest) GetMaxTtl() string`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *LdapConfigureRequest) GetMaxTtlOk() (*string, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *LdapConfigureRequest) SetMaxTtl(v string)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *LdapConfigureRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetPasswordPolicy

`func (o *LdapConfigureRequest) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *LdapConfigureRequest) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *LdapConfigureRequest) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.


### HasPasswordPolicy

`func (o *LdapConfigureRequest) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.




### GetRequestTimeout

`func (o *LdapConfigureRequest) GetRequestTimeout() string`

GetRequestTimeout returns the RequestTimeout field if non-nil, zero value otherwise.

### GetRequestTimeoutOk

`func (o *LdapConfigureRequest) GetRequestTimeoutOk() (*string, bool)`

GetRequestTimeoutOk returns a tuple with the RequestTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimeout

`func (o *LdapConfigureRequest) SetRequestTimeout(v string)`

SetRequestTimeout sets RequestTimeout field to given value.


### HasRequestTimeout

`func (o *LdapConfigureRequest) HasRequestTimeout() bool`

HasRequestTimeout returns a boolean if a field has been set.




### GetSchema

`func (o *LdapConfigureRequest) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *LdapConfigureRequest) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *LdapConfigureRequest) SetSchema(v string)`

SetSchema sets Schema field to given value.


### HasSchema

`func (o *LdapConfigureRequest) HasSchema() bool`

HasSchema returns a boolean if a field has been set.




### GetStarttls

`func (o *LdapConfigureRequest) GetStarttls() bool`

GetStarttls returns the Starttls field if non-nil, zero value otherwise.

### GetStarttlsOk

`func (o *LdapConfigureRequest) GetStarttlsOk() (*bool, bool)`

GetStarttlsOk returns a tuple with the Starttls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarttls

`func (o *LdapConfigureRequest) SetStarttls(v bool)`

SetStarttls sets Starttls field to given value.


### HasStarttls

`func (o *LdapConfigureRequest) HasStarttls() bool`

HasStarttls returns a boolean if a field has been set.




### GetTlsMaxVersion

`func (o *LdapConfigureRequest) GetTlsMaxVersion() string`

GetTlsMaxVersion returns the TlsMaxVersion field if non-nil, zero value otherwise.

### GetTlsMaxVersionOk

`func (o *LdapConfigureRequest) GetTlsMaxVersionOk() (*string, bool)`

GetTlsMaxVersionOk returns a tuple with the TlsMaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMaxVersion

`func (o *LdapConfigureRequest) SetTlsMaxVersion(v string)`

SetTlsMaxVersion sets TlsMaxVersion field to given value.


### HasTlsMaxVersion

`func (o *LdapConfigureRequest) HasTlsMaxVersion() bool`

HasTlsMaxVersion returns a boolean if a field has been set.




### GetTlsMinVersion

`func (o *LdapConfigureRequest) GetTlsMinVersion() string`

GetTlsMinVersion returns the TlsMinVersion field if non-nil, zero value otherwise.

### GetTlsMinVersionOk

`func (o *LdapConfigureRequest) GetTlsMinVersionOk() (*string, bool)`

GetTlsMinVersionOk returns a tuple with the TlsMinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMinVersion

`func (o *LdapConfigureRequest) SetTlsMinVersion(v string)`

SetTlsMinVersion sets TlsMinVersion field to given value.


### HasTlsMinVersion

`func (o *LdapConfigureRequest) HasTlsMinVersion() bool`

HasTlsMinVersion returns a boolean if a field has been set.




### GetTtl

`func (o *LdapConfigureRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *LdapConfigureRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *LdapConfigureRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *LdapConfigureRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.




### GetUpndomain

`func (o *LdapConfigureRequest) GetUpndomain() string`

GetUpndomain returns the Upndomain field if non-nil, zero value otherwise.

### GetUpndomainOk

`func (o *LdapConfigureRequest) GetUpndomainOk() (*string, bool)`

GetUpndomainOk returns a tuple with the Upndomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpndomain

`func (o *LdapConfigureRequest) SetUpndomain(v string)`

SetUpndomain sets Upndomain field to given value.


### HasUpndomain

`func (o *LdapConfigureRequest) HasUpndomain() bool`

HasUpndomain returns a boolean if a field has been set.




### GetUrl

`func (o *LdapConfigureRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LdapConfigureRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LdapConfigureRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### HasUrl

`func (o *LdapConfigureRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.




### GetUsePre111GroupCnBehavior

`func (o *LdapConfigureRequest) GetUsePre111GroupCnBehavior() bool`

GetUsePre111GroupCnBehavior returns the UsePre111GroupCnBehavior field if non-nil, zero value otherwise.

### GetUsePre111GroupCnBehaviorOk

`func (o *LdapConfigureRequest) GetUsePre111GroupCnBehaviorOk() (*bool, bool)`

GetUsePre111GroupCnBehaviorOk returns a tuple with the UsePre111GroupCnBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePre111GroupCnBehavior

`func (o *LdapConfigureRequest) SetUsePre111GroupCnBehavior(v bool)`

SetUsePre111GroupCnBehavior sets UsePre111GroupCnBehavior field to given value.


### HasUsePre111GroupCnBehavior

`func (o *LdapConfigureRequest) HasUsePre111GroupCnBehavior() bool`

HasUsePre111GroupCnBehavior returns a boolean if a field has been set.




### GetUseTokenGroups

`func (o *LdapConfigureRequest) GetUseTokenGroups() bool`

GetUseTokenGroups returns the UseTokenGroups field if non-nil, zero value otherwise.

### GetUseTokenGroupsOk

`func (o *LdapConfigureRequest) GetUseTokenGroupsOk() (*bool, bool)`

GetUseTokenGroupsOk returns a tuple with the UseTokenGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTokenGroups

`func (o *LdapConfigureRequest) SetUseTokenGroups(v bool)`

SetUseTokenGroups sets UseTokenGroups field to given value.


### HasUseTokenGroups

`func (o *LdapConfigureRequest) HasUseTokenGroups() bool`

HasUseTokenGroups returns a boolean if a field has been set.




### GetUserattr

`func (o *LdapConfigureRequest) GetUserattr() string`

GetUserattr returns the Userattr field if non-nil, zero value otherwise.

### GetUserattrOk

`func (o *LdapConfigureRequest) GetUserattrOk() (*string, bool)`

GetUserattrOk returns a tuple with the Userattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserattr

`func (o *LdapConfigureRequest) SetUserattr(v string)`

SetUserattr sets Userattr field to given value.


### HasUserattr

`func (o *LdapConfigureRequest) HasUserattr() bool`

HasUserattr returns a boolean if a field has been set.




### GetUserdn

`func (o *LdapConfigureRequest) GetUserdn() string`

GetUserdn returns the Userdn field if non-nil, zero value otherwise.

### GetUserdnOk

`func (o *LdapConfigureRequest) GetUserdnOk() (*string, bool)`

GetUserdnOk returns a tuple with the Userdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdn

`func (o *LdapConfigureRequest) SetUserdn(v string)`

SetUserdn sets Userdn field to given value.


### HasUserdn

`func (o *LdapConfigureRequest) HasUserdn() bool`

HasUserdn returns a boolean if a field has been set.




### GetUserfilter

`func (o *LdapConfigureRequest) GetUserfilter() string`

GetUserfilter returns the Userfilter field if non-nil, zero value otherwise.

### GetUserfilterOk

`func (o *LdapConfigureRequest) GetUserfilterOk() (*string, bool)`

GetUserfilterOk returns a tuple with the Userfilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserfilter

`func (o *LdapConfigureRequest) SetUserfilter(v string)`

SetUserfilter sets Userfilter field to given value.


### HasUserfilter

`func (o *LdapConfigureRequest) HasUserfilter() bool`

HasUserfilter returns a boolean if a field has been set.




### GetUsernameAsAlias

`func (o *LdapConfigureRequest) GetUsernameAsAlias() bool`

GetUsernameAsAlias returns the UsernameAsAlias field if non-nil, zero value otherwise.

### GetUsernameAsAliasOk

`func (o *LdapConfigureRequest) GetUsernameAsAliasOk() (*bool, bool)`

GetUsernameAsAliasOk returns a tuple with the UsernameAsAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameAsAlias

`func (o *LdapConfigureRequest) SetUsernameAsAlias(v bool)`

SetUsernameAsAlias sets UsernameAsAlias field to given value.


### HasUsernameAsAlias

`func (o *LdapConfigureRequest) HasUsernameAsAlias() bool`

HasUsernameAsAlias returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


