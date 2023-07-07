# SshWriteRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgorithmSigner** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512, default, or the empty string. | [optional] 
**AllowBareDomains** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, host certificates that are requested are allowed to use the base domains listed in \&quot;allowed_domains\&quot;, e.g. \&quot;example.com\&quot;. This is a separate option as in some cases this can be considered a security threat. | [optional] 
**AllowHostCertificates** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, certificates are allowed to be signed for use as a &#x27;host&#x27;. | [optional] [default to false]
**AllowSubdomains** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, host certificates that are requested are allowed to use subdomains of those listed in \&quot;allowed_domains\&quot;. | [optional] 
**AllowUserCertificates** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, certificates are allowed to be signed for use as a &#x27;user&#x27;. | [optional] [default to false]
**AllowUserKeyIds** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If true, users can override the key ID for a signed certificate with the \&quot;key_id\&quot; field. When false, the key ID will always be the token display name. The key ID is logged by the SSH server and can be useful for auditing. | [optional] 
**AllowedCriticalOptions** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] A comma-separated list of critical options that certificates can have when signed. To allow any critical options, set this to an empty string. | [optional] 
**AllowedDomains** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] If this option is not specified, client can request for a signed certificate for any valid host. If only certain domains are allowed, then this list enforces it. | [optional] 
**AllowedDomainsTemplate** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, Allowed domains can be specified using identity template policies. Non-templated domains are also permitted. | [optional] [default to false]
**AllowedExtensions** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] A comma-separated list of extensions that certificates can have when signed. An empty list means that no extension overrides are allowed by an end-user; explicitly specify &#x27;*&#x27; to allow any extensions to be set. | [optional] 
**AllowedUserKeyLengths** | Pointer to **map[string]interface{}** | [Not applicable for OTP type] [Optional for CA type] If set, allows the enforcement of key types and minimum key sizes to be signed. | [optional] 
**AllowedUsers** | Pointer to **string** | [Optional for all types] [Works differently for CA type] If this option is not specified, or is &#x27;*&#x27;, client can request a credential for any valid user at the remote host, including the admin user. If only certain usernames are to be allowed, then this list enforces it. If this field is set, then credentials can only be created for default_user and usernames present in this list. Setting this option will enable all the users with access to this role to fetch credentials for all other usernames in this list. Use with caution. N.B.: with the CA type, an empty list means that no users are allowed; explicitly specify &#x27;*&#x27; to allow any user. | [optional] 
**AllowedUsersTemplate** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, Allowed users can be specified using identity template policies. Non-templated users are also permitted. | [optional] [default to false]
**CidrList** | Pointer to **string** | [Optional for OTP type] [Not applicable for CA type] Comma separated list of CIDR blocks for which the role is applicable for. CIDR blocks can belong to more than one role. | [optional] 
**DefaultCriticalOptions** | Pointer to **map[string]interface{}** | [Not applicable for OTP type] [Optional for CA type] Critical options certificates should have if none are provided when signing. This field takes in key value pairs in JSON format. Note that these are not restricted by \&quot;allowed_critical_options\&quot;. Defaults to none. | [optional] 
**DefaultExtensions** | Pointer to **map[string]interface{}** | [Not applicable for OTP type] [Optional for CA type] Extensions certificates should have if none are provided when signing. This field takes in key value pairs in JSON format. Note that these are not restricted by \&quot;allowed_extensions\&quot;. Defaults to none. | [optional] 
**DefaultExtensionsTemplate** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, Default extension values can be specified using identity template policies. Non-templated extension values are also permitted. | [optional] [default to false]
**DefaultUser** | Pointer to **string** | [Required for OTP type] [Optional for CA type] Default username for which a credential will be generated. When the endpoint &#x27;creds/&#x27; is used without a username, this value will be used as default username. | [optional] 
**DefaultUserTemplate** | Pointer to **bool** | [Not applicable for OTP type] [Optional for CA type] If set, Default user can be specified using identity template policies. Non-templated users are also permitted. | [optional] [default to false]
**ExcludeCidrList** | Pointer to **string** | [Optional for OTP type] [Not applicable for CA type] Comma separated list of CIDR blocks. IP addresses belonging to these blocks are not accepted by the role. This is particularly useful when big CIDR blocks are being used by the role and certain parts of it needs to be kept out. | [optional] 
**KeyIdFormat** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] When supplied, this value specifies a custom format for the key id of a signed certificate. The following variables are available for use: &#x27;{{token_display_name}}&#x27; - The display name of the token used to make the request. &#x27;{{role_name}}&#x27; - The name of the role signing the request. &#x27;{{public_key_hash}}&#x27; - A SHA256 checksum of the public key that is being signed. | [optional] 
**KeyType** | Pointer to **string** | [Required for all types] Type of key used to login to hosts. It can be either &#x27;otp&#x27; or &#x27;ca&#x27;. &#x27;otp&#x27; type requires agent to be installed in remote hosts. | [optional] 
**MaxTtl** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] The maximum allowed lease duration | [optional] 
**NotBeforeDuration** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] The duration that the SSH certificate should be backdated by at issuance. | [optional] [default to "30"]
**Port** | Pointer to **int32** | [Optional for OTP type] [Not applicable for CA type] Port number for SSH connection. Default is &#x27;22&#x27;. Port number does not play any role in creation of OTP. For &#x27;otp&#x27; type, this is just a way to inform client about the port number to use. Port number will be returned to client by Vault server along with OTP. | [optional] 
**Ttl** | Pointer to **string** | [Not applicable for OTP type] [Optional for CA type] The lease duration if no specific lease duration is requested. The lease duration controls the expiration of certificates issued by this backend. Defaults to the value of max_ttl. | [optional] 



## Methods


### NewSshWriteRoleRequest

`func NewSshWriteRoleRequest() *SshWriteRoleRequest`

NewSshWriteRoleRequest instantiates a new SshWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshWriteRoleRequestWithDefaults

`func NewSshWriteRoleRequestWithDefaults() *SshWriteRoleRequest`

NewSshWriteRoleRequestWithDefaults instantiates a new SshWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAlgorithmSigner

`func (o *SshWriteRoleRequest) GetAlgorithmSigner() string`

GetAlgorithmSigner returns the AlgorithmSigner field if non-nil, zero value otherwise.

### GetAlgorithmSignerOk

`func (o *SshWriteRoleRequest) GetAlgorithmSignerOk() (*string, bool)`

GetAlgorithmSignerOk returns a tuple with the AlgorithmSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithmSigner

`func (o *SshWriteRoleRequest) SetAlgorithmSigner(v string)`

SetAlgorithmSigner sets AlgorithmSigner field to given value.


### HasAlgorithmSigner

`func (o *SshWriteRoleRequest) HasAlgorithmSigner() bool`

HasAlgorithmSigner returns a boolean if a field has been set.




### GetAllowBareDomains

`func (o *SshWriteRoleRequest) GetAllowBareDomains() bool`

GetAllowBareDomains returns the AllowBareDomains field if non-nil, zero value otherwise.

### GetAllowBareDomainsOk

`func (o *SshWriteRoleRequest) GetAllowBareDomainsOk() (*bool, bool)`

GetAllowBareDomainsOk returns a tuple with the AllowBareDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBareDomains

`func (o *SshWriteRoleRequest) SetAllowBareDomains(v bool)`

SetAllowBareDomains sets AllowBareDomains field to given value.


### HasAllowBareDomains

`func (o *SshWriteRoleRequest) HasAllowBareDomains() bool`

HasAllowBareDomains returns a boolean if a field has been set.




### GetAllowHostCertificates

`func (o *SshWriteRoleRequest) GetAllowHostCertificates() bool`

GetAllowHostCertificates returns the AllowHostCertificates field if non-nil, zero value otherwise.

### GetAllowHostCertificatesOk

`func (o *SshWriteRoleRequest) GetAllowHostCertificatesOk() (*bool, bool)`

GetAllowHostCertificatesOk returns a tuple with the AllowHostCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHostCertificates

`func (o *SshWriteRoleRequest) SetAllowHostCertificates(v bool)`

SetAllowHostCertificates sets AllowHostCertificates field to given value.


### HasAllowHostCertificates

`func (o *SshWriteRoleRequest) HasAllowHostCertificates() bool`

HasAllowHostCertificates returns a boolean if a field has been set.




### GetAllowSubdomains

`func (o *SshWriteRoleRequest) GetAllowSubdomains() bool`

GetAllowSubdomains returns the AllowSubdomains field if non-nil, zero value otherwise.

### GetAllowSubdomainsOk

`func (o *SshWriteRoleRequest) GetAllowSubdomainsOk() (*bool, bool)`

GetAllowSubdomainsOk returns a tuple with the AllowSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSubdomains

`func (o *SshWriteRoleRequest) SetAllowSubdomains(v bool)`

SetAllowSubdomains sets AllowSubdomains field to given value.


### HasAllowSubdomains

`func (o *SshWriteRoleRequest) HasAllowSubdomains() bool`

HasAllowSubdomains returns a boolean if a field has been set.




### GetAllowUserCertificates

`func (o *SshWriteRoleRequest) GetAllowUserCertificates() bool`

GetAllowUserCertificates returns the AllowUserCertificates field if non-nil, zero value otherwise.

### GetAllowUserCertificatesOk

`func (o *SshWriteRoleRequest) GetAllowUserCertificatesOk() (*bool, bool)`

GetAllowUserCertificatesOk returns a tuple with the AllowUserCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserCertificates

`func (o *SshWriteRoleRequest) SetAllowUserCertificates(v bool)`

SetAllowUserCertificates sets AllowUserCertificates field to given value.


### HasAllowUserCertificates

`func (o *SshWriteRoleRequest) HasAllowUserCertificates() bool`

HasAllowUserCertificates returns a boolean if a field has been set.




### GetAllowUserKeyIds

`func (o *SshWriteRoleRequest) GetAllowUserKeyIds() bool`

GetAllowUserKeyIds returns the AllowUserKeyIds field if non-nil, zero value otherwise.

### GetAllowUserKeyIdsOk

`func (o *SshWriteRoleRequest) GetAllowUserKeyIdsOk() (*bool, bool)`

GetAllowUserKeyIdsOk returns a tuple with the AllowUserKeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserKeyIds

`func (o *SshWriteRoleRequest) SetAllowUserKeyIds(v bool)`

SetAllowUserKeyIds sets AllowUserKeyIds field to given value.


### HasAllowUserKeyIds

`func (o *SshWriteRoleRequest) HasAllowUserKeyIds() bool`

HasAllowUserKeyIds returns a boolean if a field has been set.




### GetAllowedCriticalOptions

`func (o *SshWriteRoleRequest) GetAllowedCriticalOptions() string`

GetAllowedCriticalOptions returns the AllowedCriticalOptions field if non-nil, zero value otherwise.

### GetAllowedCriticalOptionsOk

`func (o *SshWriteRoleRequest) GetAllowedCriticalOptionsOk() (*string, bool)`

GetAllowedCriticalOptionsOk returns a tuple with the AllowedCriticalOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCriticalOptions

`func (o *SshWriteRoleRequest) SetAllowedCriticalOptions(v string)`

SetAllowedCriticalOptions sets AllowedCriticalOptions field to given value.


### HasAllowedCriticalOptions

`func (o *SshWriteRoleRequest) HasAllowedCriticalOptions() bool`

HasAllowedCriticalOptions returns a boolean if a field has been set.




### GetAllowedDomains

`func (o *SshWriteRoleRequest) GetAllowedDomains() string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *SshWriteRoleRequest) GetAllowedDomainsOk() (*string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *SshWriteRoleRequest) SetAllowedDomains(v string)`

SetAllowedDomains sets AllowedDomains field to given value.


### HasAllowedDomains

`func (o *SshWriteRoleRequest) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.




### GetAllowedDomainsTemplate

`func (o *SshWriteRoleRequest) GetAllowedDomainsTemplate() bool`

GetAllowedDomainsTemplate returns the AllowedDomainsTemplate field if non-nil, zero value otherwise.

### GetAllowedDomainsTemplateOk

`func (o *SshWriteRoleRequest) GetAllowedDomainsTemplateOk() (*bool, bool)`

GetAllowedDomainsTemplateOk returns a tuple with the AllowedDomainsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomainsTemplate

`func (o *SshWriteRoleRequest) SetAllowedDomainsTemplate(v bool)`

SetAllowedDomainsTemplate sets AllowedDomainsTemplate field to given value.


### HasAllowedDomainsTemplate

`func (o *SshWriteRoleRequest) HasAllowedDomainsTemplate() bool`

HasAllowedDomainsTemplate returns a boolean if a field has been set.




### GetAllowedExtensions

`func (o *SshWriteRoleRequest) GetAllowedExtensions() string`

GetAllowedExtensions returns the AllowedExtensions field if non-nil, zero value otherwise.

### GetAllowedExtensionsOk

`func (o *SshWriteRoleRequest) GetAllowedExtensionsOk() (*string, bool)`

GetAllowedExtensionsOk returns a tuple with the AllowedExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedExtensions

`func (o *SshWriteRoleRequest) SetAllowedExtensions(v string)`

SetAllowedExtensions sets AllowedExtensions field to given value.


### HasAllowedExtensions

`func (o *SshWriteRoleRequest) HasAllowedExtensions() bool`

HasAllowedExtensions returns a boolean if a field has been set.




### GetAllowedUserKeyLengths

`func (o *SshWriteRoleRequest) GetAllowedUserKeyLengths() map[string]interface{}`

GetAllowedUserKeyLengths returns the AllowedUserKeyLengths field if non-nil, zero value otherwise.

### GetAllowedUserKeyLengthsOk

`func (o *SshWriteRoleRequest) GetAllowedUserKeyLengthsOk() (*map[string]interface{}, bool)`

GetAllowedUserKeyLengthsOk returns a tuple with the AllowedUserKeyLengths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUserKeyLengths

`func (o *SshWriteRoleRequest) SetAllowedUserKeyLengths(v map[string]interface{})`

SetAllowedUserKeyLengths sets AllowedUserKeyLengths field to given value.


### HasAllowedUserKeyLengths

`func (o *SshWriteRoleRequest) HasAllowedUserKeyLengths() bool`

HasAllowedUserKeyLengths returns a boolean if a field has been set.




### GetAllowedUsers

`func (o *SshWriteRoleRequest) GetAllowedUsers() string`

GetAllowedUsers returns the AllowedUsers field if non-nil, zero value otherwise.

### GetAllowedUsersOk

`func (o *SshWriteRoleRequest) GetAllowedUsersOk() (*string, bool)`

GetAllowedUsersOk returns a tuple with the AllowedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsers

`func (o *SshWriteRoleRequest) SetAllowedUsers(v string)`

SetAllowedUsers sets AllowedUsers field to given value.


### HasAllowedUsers

`func (o *SshWriteRoleRequest) HasAllowedUsers() bool`

HasAllowedUsers returns a boolean if a field has been set.




### GetAllowedUsersTemplate

`func (o *SshWriteRoleRequest) GetAllowedUsersTemplate() bool`

GetAllowedUsersTemplate returns the AllowedUsersTemplate field if non-nil, zero value otherwise.

### GetAllowedUsersTemplateOk

`func (o *SshWriteRoleRequest) GetAllowedUsersTemplateOk() (*bool, bool)`

GetAllowedUsersTemplateOk returns a tuple with the AllowedUsersTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsersTemplate

`func (o *SshWriteRoleRequest) SetAllowedUsersTemplate(v bool)`

SetAllowedUsersTemplate sets AllowedUsersTemplate field to given value.


### HasAllowedUsersTemplate

`func (o *SshWriteRoleRequest) HasAllowedUsersTemplate() bool`

HasAllowedUsersTemplate returns a boolean if a field has been set.




### GetCidrList

`func (o *SshWriteRoleRequest) GetCidrList() string`

GetCidrList returns the CidrList field if non-nil, zero value otherwise.

### GetCidrListOk

`func (o *SshWriteRoleRequest) GetCidrListOk() (*string, bool)`

GetCidrListOk returns a tuple with the CidrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrList

`func (o *SshWriteRoleRequest) SetCidrList(v string)`

SetCidrList sets CidrList field to given value.


### HasCidrList

`func (o *SshWriteRoleRequest) HasCidrList() bool`

HasCidrList returns a boolean if a field has been set.




### GetDefaultCriticalOptions

`func (o *SshWriteRoleRequest) GetDefaultCriticalOptions() map[string]interface{}`

GetDefaultCriticalOptions returns the DefaultCriticalOptions field if non-nil, zero value otherwise.

### GetDefaultCriticalOptionsOk

`func (o *SshWriteRoleRequest) GetDefaultCriticalOptionsOk() (*map[string]interface{}, bool)`

GetDefaultCriticalOptionsOk returns a tuple with the DefaultCriticalOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCriticalOptions

`func (o *SshWriteRoleRequest) SetDefaultCriticalOptions(v map[string]interface{})`

SetDefaultCriticalOptions sets DefaultCriticalOptions field to given value.


### HasDefaultCriticalOptions

`func (o *SshWriteRoleRequest) HasDefaultCriticalOptions() bool`

HasDefaultCriticalOptions returns a boolean if a field has been set.




### GetDefaultExtensions

`func (o *SshWriteRoleRequest) GetDefaultExtensions() map[string]interface{}`

GetDefaultExtensions returns the DefaultExtensions field if non-nil, zero value otherwise.

### GetDefaultExtensionsOk

`func (o *SshWriteRoleRequest) GetDefaultExtensionsOk() (*map[string]interface{}, bool)`

GetDefaultExtensionsOk returns a tuple with the DefaultExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExtensions

`func (o *SshWriteRoleRequest) SetDefaultExtensions(v map[string]interface{})`

SetDefaultExtensions sets DefaultExtensions field to given value.


### HasDefaultExtensions

`func (o *SshWriteRoleRequest) HasDefaultExtensions() bool`

HasDefaultExtensions returns a boolean if a field has been set.




### GetDefaultExtensionsTemplate

`func (o *SshWriteRoleRequest) GetDefaultExtensionsTemplate() bool`

GetDefaultExtensionsTemplate returns the DefaultExtensionsTemplate field if non-nil, zero value otherwise.

### GetDefaultExtensionsTemplateOk

`func (o *SshWriteRoleRequest) GetDefaultExtensionsTemplateOk() (*bool, bool)`

GetDefaultExtensionsTemplateOk returns a tuple with the DefaultExtensionsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExtensionsTemplate

`func (o *SshWriteRoleRequest) SetDefaultExtensionsTemplate(v bool)`

SetDefaultExtensionsTemplate sets DefaultExtensionsTemplate field to given value.


### HasDefaultExtensionsTemplate

`func (o *SshWriteRoleRequest) HasDefaultExtensionsTemplate() bool`

HasDefaultExtensionsTemplate returns a boolean if a field has been set.




### GetDefaultUser

`func (o *SshWriteRoleRequest) GetDefaultUser() string`

GetDefaultUser returns the DefaultUser field if non-nil, zero value otherwise.

### GetDefaultUserOk

`func (o *SshWriteRoleRequest) GetDefaultUserOk() (*string, bool)`

GetDefaultUserOk returns a tuple with the DefaultUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUser

`func (o *SshWriteRoleRequest) SetDefaultUser(v string)`

SetDefaultUser sets DefaultUser field to given value.


### HasDefaultUser

`func (o *SshWriteRoleRequest) HasDefaultUser() bool`

HasDefaultUser returns a boolean if a field has been set.




### GetDefaultUserTemplate

`func (o *SshWriteRoleRequest) GetDefaultUserTemplate() bool`

GetDefaultUserTemplate returns the DefaultUserTemplate field if non-nil, zero value otherwise.

### GetDefaultUserTemplateOk

`func (o *SshWriteRoleRequest) GetDefaultUserTemplateOk() (*bool, bool)`

GetDefaultUserTemplateOk returns a tuple with the DefaultUserTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserTemplate

`func (o *SshWriteRoleRequest) SetDefaultUserTemplate(v bool)`

SetDefaultUserTemplate sets DefaultUserTemplate field to given value.


### HasDefaultUserTemplate

`func (o *SshWriteRoleRequest) HasDefaultUserTemplate() bool`

HasDefaultUserTemplate returns a boolean if a field has been set.




### GetExcludeCidrList

`func (o *SshWriteRoleRequest) GetExcludeCidrList() string`

GetExcludeCidrList returns the ExcludeCidrList field if non-nil, zero value otherwise.

### GetExcludeCidrListOk

`func (o *SshWriteRoleRequest) GetExcludeCidrListOk() (*string, bool)`

GetExcludeCidrListOk returns a tuple with the ExcludeCidrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCidrList

`func (o *SshWriteRoleRequest) SetExcludeCidrList(v string)`

SetExcludeCidrList sets ExcludeCidrList field to given value.


### HasExcludeCidrList

`func (o *SshWriteRoleRequest) HasExcludeCidrList() bool`

HasExcludeCidrList returns a boolean if a field has been set.




### GetKeyIdFormat

`func (o *SshWriteRoleRequest) GetKeyIdFormat() string`

GetKeyIdFormat returns the KeyIdFormat field if non-nil, zero value otherwise.

### GetKeyIdFormatOk

`func (o *SshWriteRoleRequest) GetKeyIdFormatOk() (*string, bool)`

GetKeyIdFormatOk returns a tuple with the KeyIdFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyIdFormat

`func (o *SshWriteRoleRequest) SetKeyIdFormat(v string)`

SetKeyIdFormat sets KeyIdFormat field to given value.


### HasKeyIdFormat

`func (o *SshWriteRoleRequest) HasKeyIdFormat() bool`

HasKeyIdFormat returns a boolean if a field has been set.




### GetKeyType

`func (o *SshWriteRoleRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *SshWriteRoleRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *SshWriteRoleRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### HasKeyType

`func (o *SshWriteRoleRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.




### GetMaxTtl

`func (o *SshWriteRoleRequest) GetMaxTtl() string`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *SshWriteRoleRequest) GetMaxTtlOk() (*string, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *SshWriteRoleRequest) SetMaxTtl(v string)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *SshWriteRoleRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetNotBeforeDuration

`func (o *SshWriteRoleRequest) GetNotBeforeDuration() string`

GetNotBeforeDuration returns the NotBeforeDuration field if non-nil, zero value otherwise.

### GetNotBeforeDurationOk

`func (o *SshWriteRoleRequest) GetNotBeforeDurationOk() (*string, bool)`

GetNotBeforeDurationOk returns a tuple with the NotBeforeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforeDuration

`func (o *SshWriteRoleRequest) SetNotBeforeDuration(v string)`

SetNotBeforeDuration sets NotBeforeDuration field to given value.


### HasNotBeforeDuration

`func (o *SshWriteRoleRequest) HasNotBeforeDuration() bool`

HasNotBeforeDuration returns a boolean if a field has been set.




### GetPort

`func (o *SshWriteRoleRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SshWriteRoleRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SshWriteRoleRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### HasPort

`func (o *SshWriteRoleRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.




### GetTtl

`func (o *SshWriteRoleRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *SshWriteRoleRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *SshWriteRoleRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *SshWriteRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


