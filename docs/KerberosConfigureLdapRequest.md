# KerberosConfigureLdapRequest


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
**MaxPageSize** | Pointer to **int32** | If set to a value greater than 0, the LDAP backend will use the LDAP server&#x27;s paged search control to request pages of up to the given size. This can be used to avoid hitting the LDAP server&#x27;s maximum result size limit. Otherwise, the LDAP backend will not use the paged search control. | [optional] [default to 0]
**RequestTimeout** | Pointer to **string** | Timeout, in seconds, for the connection when making requests against the server before returning back an error. | [optional] [default to "90s"]
**Starttls** | Pointer to **bool** | Issue a StartTLS command after establishing unencrypted connection (optional) | [optional] 
**TlsMaxVersion** | Pointer to **string** | Maximum TLS version to use. Accepted values are &#x27;tls10&#x27;, &#x27;tls11&#x27;, &#x27;tls12&#x27; or &#x27;tls13&#x27;. Defaults to &#x27;tls12&#x27; | [optional] [default to "tls12"]
**TlsMinVersion** | Pointer to **string** | Minimum TLS version to use. Accepted values are &#x27;tls10&#x27;, &#x27;tls11&#x27;, &#x27;tls12&#x27; or &#x27;tls13&#x27;. Defaults to &#x27;tls12&#x27; | [optional] [default to "tls12"]
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **string** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **string** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#x27;default&#x27; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **string** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies. This will apply to all tokens generated by this auth method, in addition to any configured for specific users/groups. | [optional] 
**TokenTtl** | Pointer to **string** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]
**Upndomain** | Pointer to **string** | Enables userPrincipalDomain login with [username]@UPNDomain (optional) | [optional] 
**Url** | Pointer to **string** | LDAP URL to connect to (default: ldap://127.0.0.1). Multiple URLs can be specified by concatenating them with commas; they will be tried in-order. | [optional] [default to "ldap://127.0.0.1"]
**UsePre111GroupCnBehavior** | Pointer to **bool** | In Vault 1.1.1 a fix for handling group CN values of different cases unfortunately introduced a regression that could cause previously defined groups to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for matching group CNs will be used. This is only needed in some upgrade scenarios for backwards compatibility. It is enabled by default if the config is upgraded but disabled by default on new configurations. | [optional] 
**UseTokenGroups** | Pointer to **bool** | If true, use the Active Directory tokenGroups constructed attribute of the user to find the group memberships. This will find all security groups including nested ones. | [optional] [default to false]
**Userattr** | Pointer to **string** | Attribute used for users (default: cn) | [optional] [default to "cn"]
**Userdn** | Pointer to **string** | LDAP domain to use for users (eg: ou&#x3D;People,dc&#x3D;example,dc&#x3D;org) | [optional] 
**Userfilter** | Pointer to **string** | Go template for LDAP user search filer (optional) The template can access the following context variables: UserAttr, Username Default: ({{.UserAttr}}&#x3D;{{.Username}}) | [optional] [default to "({{.UserAttr}}={{.Username}})"]
**UsernameAsAlias** | Pointer to **bool** | If true, sets the alias name to the username | [optional] [default to false]



## Methods


### NewKerberosConfigureLdapRequest

`func NewKerberosConfigureLdapRequest() *KerberosConfigureLdapRequest`

NewKerberosConfigureLdapRequest instantiates a new KerberosConfigureLdapRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosConfigureLdapRequestWithDefaults

`func NewKerberosConfigureLdapRequestWithDefaults() *KerberosConfigureLdapRequest`

NewKerberosConfigureLdapRequestWithDefaults instantiates a new KerberosConfigureLdapRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAnonymousGroupSearch

`func (o *KerberosConfigureLdapRequest) GetAnonymousGroupSearch() bool`

GetAnonymousGroupSearch returns the AnonymousGroupSearch field if non-nil, zero value otherwise.

### GetAnonymousGroupSearchOk

`func (o *KerberosConfigureLdapRequest) GetAnonymousGroupSearchOk() (*bool, bool)`

GetAnonymousGroupSearchOk returns a tuple with the AnonymousGroupSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousGroupSearch

`func (o *KerberosConfigureLdapRequest) SetAnonymousGroupSearch(v bool)`

SetAnonymousGroupSearch sets AnonymousGroupSearch field to given value.


### HasAnonymousGroupSearch

`func (o *KerberosConfigureLdapRequest) HasAnonymousGroupSearch() bool`

HasAnonymousGroupSearch returns a boolean if a field has been set.




### GetBinddn

`func (o *KerberosConfigureLdapRequest) GetBinddn() string`

GetBinddn returns the Binddn field if non-nil, zero value otherwise.

### GetBinddnOk

`func (o *KerberosConfigureLdapRequest) GetBinddnOk() (*string, bool)`

GetBinddnOk returns a tuple with the Binddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinddn

`func (o *KerberosConfigureLdapRequest) SetBinddn(v string)`

SetBinddn sets Binddn field to given value.


### HasBinddn

`func (o *KerberosConfigureLdapRequest) HasBinddn() bool`

HasBinddn returns a boolean if a field has been set.




### GetBindpass

`func (o *KerberosConfigureLdapRequest) GetBindpass() string`

GetBindpass returns the Bindpass field if non-nil, zero value otherwise.

### GetBindpassOk

`func (o *KerberosConfigureLdapRequest) GetBindpassOk() (*string, bool)`

GetBindpassOk returns a tuple with the Bindpass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindpass

`func (o *KerberosConfigureLdapRequest) SetBindpass(v string)`

SetBindpass sets Bindpass field to given value.


### HasBindpass

`func (o *KerberosConfigureLdapRequest) HasBindpass() bool`

HasBindpass returns a boolean if a field has been set.




### GetCaseSensitiveNames

`func (o *KerberosConfigureLdapRequest) GetCaseSensitiveNames() bool`

GetCaseSensitiveNames returns the CaseSensitiveNames field if non-nil, zero value otherwise.

### GetCaseSensitiveNamesOk

`func (o *KerberosConfigureLdapRequest) GetCaseSensitiveNamesOk() (*bool, bool)`

GetCaseSensitiveNamesOk returns a tuple with the CaseSensitiveNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitiveNames

`func (o *KerberosConfigureLdapRequest) SetCaseSensitiveNames(v bool)`

SetCaseSensitiveNames sets CaseSensitiveNames field to given value.


### HasCaseSensitiveNames

`func (o *KerberosConfigureLdapRequest) HasCaseSensitiveNames() bool`

HasCaseSensitiveNames returns a boolean if a field has been set.




### GetCertificate

`func (o *KerberosConfigureLdapRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *KerberosConfigureLdapRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *KerberosConfigureLdapRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *KerberosConfigureLdapRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.




### GetClientTlsCert

`func (o *KerberosConfigureLdapRequest) GetClientTlsCert() string`

GetClientTlsCert returns the ClientTlsCert field if non-nil, zero value otherwise.

### GetClientTlsCertOk

`func (o *KerberosConfigureLdapRequest) GetClientTlsCertOk() (*string, bool)`

GetClientTlsCertOk returns a tuple with the ClientTlsCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsCert

`func (o *KerberosConfigureLdapRequest) SetClientTlsCert(v string)`

SetClientTlsCert sets ClientTlsCert field to given value.


### HasClientTlsCert

`func (o *KerberosConfigureLdapRequest) HasClientTlsCert() bool`

HasClientTlsCert returns a boolean if a field has been set.




### GetClientTlsKey

`func (o *KerberosConfigureLdapRequest) GetClientTlsKey() string`

GetClientTlsKey returns the ClientTlsKey field if non-nil, zero value otherwise.

### GetClientTlsKeyOk

`func (o *KerberosConfigureLdapRequest) GetClientTlsKeyOk() (*string, bool)`

GetClientTlsKeyOk returns a tuple with the ClientTlsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsKey

`func (o *KerberosConfigureLdapRequest) SetClientTlsKey(v string)`

SetClientTlsKey sets ClientTlsKey field to given value.


### HasClientTlsKey

`func (o *KerberosConfigureLdapRequest) HasClientTlsKey() bool`

HasClientTlsKey returns a boolean if a field has been set.




### GetConnectionTimeout

`func (o *KerberosConfigureLdapRequest) GetConnectionTimeout() string`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *KerberosConfigureLdapRequest) GetConnectionTimeoutOk() (*string, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *KerberosConfigureLdapRequest) SetConnectionTimeout(v string)`

SetConnectionTimeout sets ConnectionTimeout field to given value.


### HasConnectionTimeout

`func (o *KerberosConfigureLdapRequest) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.




### GetDenyNullBind

`func (o *KerberosConfigureLdapRequest) GetDenyNullBind() bool`

GetDenyNullBind returns the DenyNullBind field if non-nil, zero value otherwise.

### GetDenyNullBindOk

`func (o *KerberosConfigureLdapRequest) GetDenyNullBindOk() (*bool, bool)`

GetDenyNullBindOk returns a tuple with the DenyNullBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyNullBind

`func (o *KerberosConfigureLdapRequest) SetDenyNullBind(v bool)`

SetDenyNullBind sets DenyNullBind field to given value.


### HasDenyNullBind

`func (o *KerberosConfigureLdapRequest) HasDenyNullBind() bool`

HasDenyNullBind returns a boolean if a field has been set.




### GetDereferenceAliases

`func (o *KerberosConfigureLdapRequest) GetDereferenceAliases() string`

GetDereferenceAliases returns the DereferenceAliases field if non-nil, zero value otherwise.

### GetDereferenceAliasesOk

`func (o *KerberosConfigureLdapRequest) GetDereferenceAliasesOk() (*string, bool)`

GetDereferenceAliasesOk returns a tuple with the DereferenceAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDereferenceAliases

`func (o *KerberosConfigureLdapRequest) SetDereferenceAliases(v string)`

SetDereferenceAliases sets DereferenceAliases field to given value.


### HasDereferenceAliases

`func (o *KerberosConfigureLdapRequest) HasDereferenceAliases() bool`

HasDereferenceAliases returns a boolean if a field has been set.




### GetDiscoverdn

`func (o *KerberosConfigureLdapRequest) GetDiscoverdn() bool`

GetDiscoverdn returns the Discoverdn field if non-nil, zero value otherwise.

### GetDiscoverdnOk

`func (o *KerberosConfigureLdapRequest) GetDiscoverdnOk() (*bool, bool)`

GetDiscoverdnOk returns a tuple with the Discoverdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverdn

`func (o *KerberosConfigureLdapRequest) SetDiscoverdn(v bool)`

SetDiscoverdn sets Discoverdn field to given value.


### HasDiscoverdn

`func (o *KerberosConfigureLdapRequest) HasDiscoverdn() bool`

HasDiscoverdn returns a boolean if a field has been set.




### GetGroupattr

`func (o *KerberosConfigureLdapRequest) GetGroupattr() string`

GetGroupattr returns the Groupattr field if non-nil, zero value otherwise.

### GetGroupattrOk

`func (o *KerberosConfigureLdapRequest) GetGroupattrOk() (*string, bool)`

GetGroupattrOk returns a tuple with the Groupattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupattr

`func (o *KerberosConfigureLdapRequest) SetGroupattr(v string)`

SetGroupattr sets Groupattr field to given value.


### HasGroupattr

`func (o *KerberosConfigureLdapRequest) HasGroupattr() bool`

HasGroupattr returns a boolean if a field has been set.




### GetGroupdn

`func (o *KerberosConfigureLdapRequest) GetGroupdn() string`

GetGroupdn returns the Groupdn field if non-nil, zero value otherwise.

### GetGroupdnOk

`func (o *KerberosConfigureLdapRequest) GetGroupdnOk() (*string, bool)`

GetGroupdnOk returns a tuple with the Groupdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupdn

`func (o *KerberosConfigureLdapRequest) SetGroupdn(v string)`

SetGroupdn sets Groupdn field to given value.


### HasGroupdn

`func (o *KerberosConfigureLdapRequest) HasGroupdn() bool`

HasGroupdn returns a boolean if a field has been set.




### GetGroupfilter

`func (o *KerberosConfigureLdapRequest) GetGroupfilter() string`

GetGroupfilter returns the Groupfilter field if non-nil, zero value otherwise.

### GetGroupfilterOk

`func (o *KerberosConfigureLdapRequest) GetGroupfilterOk() (*string, bool)`

GetGroupfilterOk returns a tuple with the Groupfilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupfilter

`func (o *KerberosConfigureLdapRequest) SetGroupfilter(v string)`

SetGroupfilter sets Groupfilter field to given value.


### HasGroupfilter

`func (o *KerberosConfigureLdapRequest) HasGroupfilter() bool`

HasGroupfilter returns a boolean if a field has been set.




### GetInsecureTls

`func (o *KerberosConfigureLdapRequest) GetInsecureTls() bool`

GetInsecureTls returns the InsecureTls field if non-nil, zero value otherwise.

### GetInsecureTlsOk

`func (o *KerberosConfigureLdapRequest) GetInsecureTlsOk() (*bool, bool)`

GetInsecureTlsOk returns a tuple with the InsecureTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureTls

`func (o *KerberosConfigureLdapRequest) SetInsecureTls(v bool)`

SetInsecureTls sets InsecureTls field to given value.


### HasInsecureTls

`func (o *KerberosConfigureLdapRequest) HasInsecureTls() bool`

HasInsecureTls returns a boolean if a field has been set.




### GetMaxPageSize

`func (o *KerberosConfigureLdapRequest) GetMaxPageSize() int32`

GetMaxPageSize returns the MaxPageSize field if non-nil, zero value otherwise.

### GetMaxPageSizeOk

`func (o *KerberosConfigureLdapRequest) GetMaxPageSizeOk() (*int32, bool)`

GetMaxPageSizeOk returns a tuple with the MaxPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPageSize

`func (o *KerberosConfigureLdapRequest) SetMaxPageSize(v int32)`

SetMaxPageSize sets MaxPageSize field to given value.


### HasMaxPageSize

`func (o *KerberosConfigureLdapRequest) HasMaxPageSize() bool`

HasMaxPageSize returns a boolean if a field has been set.




### GetRequestTimeout

`func (o *KerberosConfigureLdapRequest) GetRequestTimeout() string`

GetRequestTimeout returns the RequestTimeout field if non-nil, zero value otherwise.

### GetRequestTimeoutOk

`func (o *KerberosConfigureLdapRequest) GetRequestTimeoutOk() (*string, bool)`

GetRequestTimeoutOk returns a tuple with the RequestTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimeout

`func (o *KerberosConfigureLdapRequest) SetRequestTimeout(v string)`

SetRequestTimeout sets RequestTimeout field to given value.


### HasRequestTimeout

`func (o *KerberosConfigureLdapRequest) HasRequestTimeout() bool`

HasRequestTimeout returns a boolean if a field has been set.




### GetStarttls

`func (o *KerberosConfigureLdapRequest) GetStarttls() bool`

GetStarttls returns the Starttls field if non-nil, zero value otherwise.

### GetStarttlsOk

`func (o *KerberosConfigureLdapRequest) GetStarttlsOk() (*bool, bool)`

GetStarttlsOk returns a tuple with the Starttls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarttls

`func (o *KerberosConfigureLdapRequest) SetStarttls(v bool)`

SetStarttls sets Starttls field to given value.


### HasStarttls

`func (o *KerberosConfigureLdapRequest) HasStarttls() bool`

HasStarttls returns a boolean if a field has been set.




### GetTlsMaxVersion

`func (o *KerberosConfigureLdapRequest) GetTlsMaxVersion() string`

GetTlsMaxVersion returns the TlsMaxVersion field if non-nil, zero value otherwise.

### GetTlsMaxVersionOk

`func (o *KerberosConfigureLdapRequest) GetTlsMaxVersionOk() (*string, bool)`

GetTlsMaxVersionOk returns a tuple with the TlsMaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMaxVersion

`func (o *KerberosConfigureLdapRequest) SetTlsMaxVersion(v string)`

SetTlsMaxVersion sets TlsMaxVersion field to given value.


### HasTlsMaxVersion

`func (o *KerberosConfigureLdapRequest) HasTlsMaxVersion() bool`

HasTlsMaxVersion returns a boolean if a field has been set.




### GetTlsMinVersion

`func (o *KerberosConfigureLdapRequest) GetTlsMinVersion() string`

GetTlsMinVersion returns the TlsMinVersion field if non-nil, zero value otherwise.

### GetTlsMinVersionOk

`func (o *KerberosConfigureLdapRequest) GetTlsMinVersionOk() (*string, bool)`

GetTlsMinVersionOk returns a tuple with the TlsMinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMinVersion

`func (o *KerberosConfigureLdapRequest) SetTlsMinVersion(v string)`

SetTlsMinVersion sets TlsMinVersion field to given value.


### HasTlsMinVersion

`func (o *KerberosConfigureLdapRequest) HasTlsMinVersion() bool`

HasTlsMinVersion returns a boolean if a field has been set.




### GetTokenBoundCidrs

`func (o *KerberosConfigureLdapRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *KerberosConfigureLdapRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *KerberosConfigureLdapRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.


### HasTokenBoundCidrs

`func (o *KerberosConfigureLdapRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.




### GetTokenExplicitMaxTtl

`func (o *KerberosConfigureLdapRequest) GetTokenExplicitMaxTtl() string`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *KerberosConfigureLdapRequest) GetTokenExplicitMaxTtlOk() (*string, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *KerberosConfigureLdapRequest) SetTokenExplicitMaxTtl(v string)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.


### HasTokenExplicitMaxTtl

`func (o *KerberosConfigureLdapRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.




### GetTokenMaxTtl

`func (o *KerberosConfigureLdapRequest) GetTokenMaxTtl() string`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *KerberosConfigureLdapRequest) GetTokenMaxTtlOk() (*string, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *KerberosConfigureLdapRequest) SetTokenMaxTtl(v string)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.


### HasTokenMaxTtl

`func (o *KerberosConfigureLdapRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.




### GetTokenNoDefaultPolicy

`func (o *KerberosConfigureLdapRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *KerberosConfigureLdapRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *KerberosConfigureLdapRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.


### HasTokenNoDefaultPolicy

`func (o *KerberosConfigureLdapRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.




### GetTokenNumUses

`func (o *KerberosConfigureLdapRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *KerberosConfigureLdapRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *KerberosConfigureLdapRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.


### HasTokenNumUses

`func (o *KerberosConfigureLdapRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.




### GetTokenPeriod

`func (o *KerberosConfigureLdapRequest) GetTokenPeriod() string`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *KerberosConfigureLdapRequest) GetTokenPeriodOk() (*string, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *KerberosConfigureLdapRequest) SetTokenPeriod(v string)`

SetTokenPeriod sets TokenPeriod field to given value.


### HasTokenPeriod

`func (o *KerberosConfigureLdapRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.




### GetTokenPolicies

`func (o *KerberosConfigureLdapRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *KerberosConfigureLdapRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *KerberosConfigureLdapRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.


### HasTokenPolicies

`func (o *KerberosConfigureLdapRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.




### GetTokenTtl

`func (o *KerberosConfigureLdapRequest) GetTokenTtl() string`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *KerberosConfigureLdapRequest) GetTokenTtlOk() (*string, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *KerberosConfigureLdapRequest) SetTokenTtl(v string)`

SetTokenTtl sets TokenTtl field to given value.


### HasTokenTtl

`func (o *KerberosConfigureLdapRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.




### GetTokenType

`func (o *KerberosConfigureLdapRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *KerberosConfigureLdapRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *KerberosConfigureLdapRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### HasTokenType

`func (o *KerberosConfigureLdapRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.




### GetUpndomain

`func (o *KerberosConfigureLdapRequest) GetUpndomain() string`

GetUpndomain returns the Upndomain field if non-nil, zero value otherwise.

### GetUpndomainOk

`func (o *KerberosConfigureLdapRequest) GetUpndomainOk() (*string, bool)`

GetUpndomainOk returns a tuple with the Upndomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpndomain

`func (o *KerberosConfigureLdapRequest) SetUpndomain(v string)`

SetUpndomain sets Upndomain field to given value.


### HasUpndomain

`func (o *KerberosConfigureLdapRequest) HasUpndomain() bool`

HasUpndomain returns a boolean if a field has been set.




### GetUrl

`func (o *KerberosConfigureLdapRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *KerberosConfigureLdapRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *KerberosConfigureLdapRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### HasUrl

`func (o *KerberosConfigureLdapRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.




### GetUsePre111GroupCnBehavior

`func (o *KerberosConfigureLdapRequest) GetUsePre111GroupCnBehavior() bool`

GetUsePre111GroupCnBehavior returns the UsePre111GroupCnBehavior field if non-nil, zero value otherwise.

### GetUsePre111GroupCnBehaviorOk

`func (o *KerberosConfigureLdapRequest) GetUsePre111GroupCnBehaviorOk() (*bool, bool)`

GetUsePre111GroupCnBehaviorOk returns a tuple with the UsePre111GroupCnBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePre111GroupCnBehavior

`func (o *KerberosConfigureLdapRequest) SetUsePre111GroupCnBehavior(v bool)`

SetUsePre111GroupCnBehavior sets UsePre111GroupCnBehavior field to given value.


### HasUsePre111GroupCnBehavior

`func (o *KerberosConfigureLdapRequest) HasUsePre111GroupCnBehavior() bool`

HasUsePre111GroupCnBehavior returns a boolean if a field has been set.




### GetUseTokenGroups

`func (o *KerberosConfigureLdapRequest) GetUseTokenGroups() bool`

GetUseTokenGroups returns the UseTokenGroups field if non-nil, zero value otherwise.

### GetUseTokenGroupsOk

`func (o *KerberosConfigureLdapRequest) GetUseTokenGroupsOk() (*bool, bool)`

GetUseTokenGroupsOk returns a tuple with the UseTokenGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTokenGroups

`func (o *KerberosConfigureLdapRequest) SetUseTokenGroups(v bool)`

SetUseTokenGroups sets UseTokenGroups field to given value.


### HasUseTokenGroups

`func (o *KerberosConfigureLdapRequest) HasUseTokenGroups() bool`

HasUseTokenGroups returns a boolean if a field has been set.




### GetUserattr

`func (o *KerberosConfigureLdapRequest) GetUserattr() string`

GetUserattr returns the Userattr field if non-nil, zero value otherwise.

### GetUserattrOk

`func (o *KerberosConfigureLdapRequest) GetUserattrOk() (*string, bool)`

GetUserattrOk returns a tuple with the Userattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserattr

`func (o *KerberosConfigureLdapRequest) SetUserattr(v string)`

SetUserattr sets Userattr field to given value.


### HasUserattr

`func (o *KerberosConfigureLdapRequest) HasUserattr() bool`

HasUserattr returns a boolean if a field has been set.




### GetUserdn

`func (o *KerberosConfigureLdapRequest) GetUserdn() string`

GetUserdn returns the Userdn field if non-nil, zero value otherwise.

### GetUserdnOk

`func (o *KerberosConfigureLdapRequest) GetUserdnOk() (*string, bool)`

GetUserdnOk returns a tuple with the Userdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdn

`func (o *KerberosConfigureLdapRequest) SetUserdn(v string)`

SetUserdn sets Userdn field to given value.


### HasUserdn

`func (o *KerberosConfigureLdapRequest) HasUserdn() bool`

HasUserdn returns a boolean if a field has been set.




### GetUserfilter

`func (o *KerberosConfigureLdapRequest) GetUserfilter() string`

GetUserfilter returns the Userfilter field if non-nil, zero value otherwise.

### GetUserfilterOk

`func (o *KerberosConfigureLdapRequest) GetUserfilterOk() (*string, bool)`

GetUserfilterOk returns a tuple with the Userfilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserfilter

`func (o *KerberosConfigureLdapRequest) SetUserfilter(v string)`

SetUserfilter sets Userfilter field to given value.


### HasUserfilter

`func (o *KerberosConfigureLdapRequest) HasUserfilter() bool`

HasUserfilter returns a boolean if a field has been set.




### GetUsernameAsAlias

`func (o *KerberosConfigureLdapRequest) GetUsernameAsAlias() bool`

GetUsernameAsAlias returns the UsernameAsAlias field if non-nil, zero value otherwise.

### GetUsernameAsAliasOk

`func (o *KerberosConfigureLdapRequest) GetUsernameAsAliasOk() (*bool, bool)`

GetUsernameAsAliasOk returns a tuple with the UsernameAsAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameAsAlias

`func (o *KerberosConfigureLdapRequest) SetUsernameAsAlias(v bool)`

SetUsernameAsAlias sets UsernameAsAlias field to given value.


### HasUsernameAsAlias

`func (o *KerberosConfigureLdapRequest) HasUsernameAsAlias() bool`

HasUsernameAsAlias returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


