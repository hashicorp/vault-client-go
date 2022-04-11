# SshRolesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminUser** | Pointer to **string** | [Required for Dynamic type] [Not applicable for OTP type] [Not applicable for CA type] Admin user at remote host. The shared key being registered should be for this user and should have root privileges. Everytime a dynamic credential is being generated for other users, Vault uses this admin username to login to remote host and install the generated credential for the other user. | [optional] 
**AlgorithmSigner** | Pointer to **string** | When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512, default, or the empty string. | [optional] 
**AllowBareDomains** | Pointer to **bool** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, host certificates that are requested are allowed to use the base domains listed in \&quot;allowed_domains\&quot;, e.g. \&quot;example.com\&quot;. This is a separate option as in some cases this can be considered a security threat. | [optional] 
**AllowHostCertificates** | Pointer to **bool** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, certificates are allowed to be signed for use as a &#39;host&#39;. | [optional] [default to false]
**AllowSubdomains** | Pointer to **bool** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, host certificates that are requested are allowed to use subdomains of those listed in \&quot;allowed_domains\&quot;. | [optional] 
**AllowUserCertificates** | Pointer to **bool** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, certificates are allowed to be signed for use as a &#39;user&#39;. | [optional] [default to false]
**AllowUserKeyIds** | Pointer to **bool** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If true, users can override the key ID for a signed certificate with the \&quot;key_id\&quot; field. When false, the key ID will always be the token display name. The key ID is logged by the SSH server and can be useful for auditing. | [optional] 
**AllowedCriticalOptions** | Pointer to **string** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] A comma-separated list of critical options that certificates can have when signed. To allow any critical options, set this to an empty string. | [optional] 
**AllowedDomains** | Pointer to **string** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If this option is not specified, client can request for a signed certificate for any valid host. If only certain domains are allowed, then this list enforces it. | [optional] 
**AllowedExtensions** | Pointer to **string** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] A comma-separated list of extensions that certificates can have when signed. An empty list means that no extension overrides are allowed by an end-user; explicitly specify &#39;*&#39; to allow any extensions to be set. | [optional] 
**AllowedUserKeyLengths** | Pointer to **map[string]interface{}** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, allows the enforcement of key types and minimum key sizes to be signed. | [optional] 
**AllowedUsers** | Pointer to **string** | [Optional for all types] [Works differently for CA type] If this option is not specified, or is &#39;*&#39;, client can request a credential for any valid user at the remote host, including the admin user. If only certain usernames are to be allowed, then this list enforces it. If this field is set, then credentials can only be created for default_user and usernames present in this list. Setting this option will enable all the users with access to this role to fetch credentials for all other usernames in this list. Use with caution. N.B.: with the CA type, an empty list means that no users are allowed; explicitly specify &#39;*&#39; to allow any user. | [optional] 
**AllowedUsersTemplate** | Pointer to **bool** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, Allowed users can be specified using identity template policies. Non-templated users are also permitted. | [optional] [default to false]
**CidrList** | Pointer to **string** | [Optional for Dynamic type] [Optional for OTP type] [Not applicable for CA type] Comma separated list of CIDR blocks for which the role is applicable for. CIDR blocks can belong to more than one role. | [optional] 
**DefaultCriticalOptions** | Pointer to **map[string]interface{}** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] Critical options certificates should have if none are provided when signing. This field takes in key value pairs in JSON format. Note that these are not restricted by \&quot;allowed_critical_options\&quot;. Defaults to none. | [optional] 
**DefaultExtensions** | Pointer to **map[string]interface{}** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] Extensions certificates should have if none are provided when signing. This field takes in key value pairs in JSON format. Note that these are not restricted by \&quot;allowed_extensions\&quot;. Defaults to none. | [optional] 
**DefaultExtensionsTemplate** | Pointer to **bool** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] If set, Default extension values can be specified using identity template policies. Non-templated extension values are also permitted. | [optional] [default to false]
**DefaultUser** | Pointer to **string** | [Required for Dynamic type] [Required for OTP type] [Optional for CA type] Default username for which a credential will be generated. When the endpoint &#39;creds/&#39; is used without a username, this value will be used as default username. | [optional] 
**ExcludeCidrList** | Pointer to **string** | [Optional for Dynamic type] [Optional for OTP type] [Not applicable for CA type] Comma separated list of CIDR blocks. IP addresses belonging to these blocks are not accepted by the role. This is particularly useful when big CIDR blocks are being used by the role and certain parts of it needs to be kept out. | [optional] 
**InstallScript** | Pointer to **string** | [Optional for Dynamic type] [Not-applicable for OTP type] [Not applicable for CA type] Script used to install and uninstall public keys in the target machine. The inbuilt default install script will be for Linux hosts. For sample script, refer the project documentation website. | [optional] 
**Key** | Pointer to **string** | [Required for Dynamic type] [Not applicable for OTP type] [Not applicable for CA type] Name of the registered key in Vault. Before creating the role, use the &#39;keys/&#39; endpoint to create a named key. | [optional] 
**KeyBits** | Pointer to **int32** | [Optional for Dynamic type] [Not applicable for OTP type] [Not applicable for CA type] Length of the RSA dynamic key in bits. It is 1024 by default or it can be 2048. | [optional] 
**KeyIdFormat** | Pointer to **string** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] When supplied, this value specifies a custom format for the key id of a signed certificate. The following variables are available for use: &#39;{{token_display_name}}&#39; - The display name of the token used to make the request. &#39;{{role_name}}&#39; - The name of the role signing the request. &#39;{{public_key_hash}}&#39; - A SHA256 checksum of the public key that is being signed. | [optional] 
**KeyOptionSpecs** | Pointer to **string** | [Optional for Dynamic type] [Not applicable for OTP type] [Not applicable for CA type] Comma separated option specifications which will be prefixed to RSA key in authorized_keys file. Options should be valid and comply with authorized_keys file format and should not contain spaces. | [optional] 
**KeyType** | Pointer to **string** | [Required for all types] Type of key used to login to hosts. It can be either &#39;otp&#39;, &#39;dynamic&#39; or &#39;ca&#39;. &#39;otp&#39; type requires agent to be installed in remote hosts. | [optional] 
**MaxTtl** | Pointer to **int32** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] The maximum allowed lease duration | [optional] 
**Port** | Pointer to **int32** | [Optional for Dynamic type] [Optional for OTP type] [Not applicable for CA type] Port number for SSH connection. Default is &#39;22&#39;. Port number does not play any role in creation of OTP. For &#39;otp&#39; type, this is just a way to inform client about the port number to use. Port number will be returned to client by Vault server along with OTP. | [optional] 
**Ttl** | Pointer to **int32** | [Not applicable for Dynamic type] [Not applicable for OTP type] [Optional for CA type] The lease duration if no specific lease duration is requested. The lease duration controls the expiration of certificates issued by this backend. Defaults to the value of max_ttl. | [optional] 

## Methods

### NewSshRolesRequest

`func NewSshRolesRequest() *SshRolesRequest`

NewSshRolesRequest instantiates a new SshRolesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshRolesRequestWithDefaults

`func NewSshRolesRequestWithDefaults() *SshRolesRequest`

NewSshRolesRequestWithDefaults instantiates a new SshRolesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminUser

`func (o *SshRolesRequest) GetAdminUser() string`

GetAdminUser returns the AdminUser field if non-nil, zero value otherwise.

### GetAdminUserOk

`func (o *SshRolesRequest) GetAdminUserOk() (*string, bool)`

GetAdminUserOk returns a tuple with the AdminUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUser

`func (o *SshRolesRequest) SetAdminUser(v string)`

SetAdminUser sets AdminUser field to given value.

### HasAdminUser

`func (o *SshRolesRequest) HasAdminUser() bool`

HasAdminUser returns a boolean if a field has been set.

### GetAlgorithmSigner

`func (o *SshRolesRequest) GetAlgorithmSigner() string`

GetAlgorithmSigner returns the AlgorithmSigner field if non-nil, zero value otherwise.

### GetAlgorithmSignerOk

`func (o *SshRolesRequest) GetAlgorithmSignerOk() (*string, bool)`

GetAlgorithmSignerOk returns a tuple with the AlgorithmSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithmSigner

`func (o *SshRolesRequest) SetAlgorithmSigner(v string)`

SetAlgorithmSigner sets AlgorithmSigner field to given value.

### HasAlgorithmSigner

`func (o *SshRolesRequest) HasAlgorithmSigner() bool`

HasAlgorithmSigner returns a boolean if a field has been set.

### GetAllowBareDomains

`func (o *SshRolesRequest) GetAllowBareDomains() bool`

GetAllowBareDomains returns the AllowBareDomains field if non-nil, zero value otherwise.

### GetAllowBareDomainsOk

`func (o *SshRolesRequest) GetAllowBareDomainsOk() (*bool, bool)`

GetAllowBareDomainsOk returns a tuple with the AllowBareDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBareDomains

`func (o *SshRolesRequest) SetAllowBareDomains(v bool)`

SetAllowBareDomains sets AllowBareDomains field to given value.

### HasAllowBareDomains

`func (o *SshRolesRequest) HasAllowBareDomains() bool`

HasAllowBareDomains returns a boolean if a field has been set.

### GetAllowHostCertificates

`func (o *SshRolesRequest) GetAllowHostCertificates() bool`

GetAllowHostCertificates returns the AllowHostCertificates field if non-nil, zero value otherwise.

### GetAllowHostCertificatesOk

`func (o *SshRolesRequest) GetAllowHostCertificatesOk() (*bool, bool)`

GetAllowHostCertificatesOk returns a tuple with the AllowHostCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHostCertificates

`func (o *SshRolesRequest) SetAllowHostCertificates(v bool)`

SetAllowHostCertificates sets AllowHostCertificates field to given value.

### HasAllowHostCertificates

`func (o *SshRolesRequest) HasAllowHostCertificates() bool`

HasAllowHostCertificates returns a boolean if a field has been set.

### GetAllowSubdomains

`func (o *SshRolesRequest) GetAllowSubdomains() bool`

GetAllowSubdomains returns the AllowSubdomains field if non-nil, zero value otherwise.

### GetAllowSubdomainsOk

`func (o *SshRolesRequest) GetAllowSubdomainsOk() (*bool, bool)`

GetAllowSubdomainsOk returns a tuple with the AllowSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSubdomains

`func (o *SshRolesRequest) SetAllowSubdomains(v bool)`

SetAllowSubdomains sets AllowSubdomains field to given value.

### HasAllowSubdomains

`func (o *SshRolesRequest) HasAllowSubdomains() bool`

HasAllowSubdomains returns a boolean if a field has been set.

### GetAllowUserCertificates

`func (o *SshRolesRequest) GetAllowUserCertificates() bool`

GetAllowUserCertificates returns the AllowUserCertificates field if non-nil, zero value otherwise.

### GetAllowUserCertificatesOk

`func (o *SshRolesRequest) GetAllowUserCertificatesOk() (*bool, bool)`

GetAllowUserCertificatesOk returns a tuple with the AllowUserCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserCertificates

`func (o *SshRolesRequest) SetAllowUserCertificates(v bool)`

SetAllowUserCertificates sets AllowUserCertificates field to given value.

### HasAllowUserCertificates

`func (o *SshRolesRequest) HasAllowUserCertificates() bool`

HasAllowUserCertificates returns a boolean if a field has been set.

### GetAllowUserKeyIds

`func (o *SshRolesRequest) GetAllowUserKeyIds() bool`

GetAllowUserKeyIds returns the AllowUserKeyIds field if non-nil, zero value otherwise.

### GetAllowUserKeyIdsOk

`func (o *SshRolesRequest) GetAllowUserKeyIdsOk() (*bool, bool)`

GetAllowUserKeyIdsOk returns a tuple with the AllowUserKeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserKeyIds

`func (o *SshRolesRequest) SetAllowUserKeyIds(v bool)`

SetAllowUserKeyIds sets AllowUserKeyIds field to given value.

### HasAllowUserKeyIds

`func (o *SshRolesRequest) HasAllowUserKeyIds() bool`

HasAllowUserKeyIds returns a boolean if a field has been set.

### GetAllowedCriticalOptions

`func (o *SshRolesRequest) GetAllowedCriticalOptions() string`

GetAllowedCriticalOptions returns the AllowedCriticalOptions field if non-nil, zero value otherwise.

### GetAllowedCriticalOptionsOk

`func (o *SshRolesRequest) GetAllowedCriticalOptionsOk() (*string, bool)`

GetAllowedCriticalOptionsOk returns a tuple with the AllowedCriticalOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCriticalOptions

`func (o *SshRolesRequest) SetAllowedCriticalOptions(v string)`

SetAllowedCriticalOptions sets AllowedCriticalOptions field to given value.

### HasAllowedCriticalOptions

`func (o *SshRolesRequest) HasAllowedCriticalOptions() bool`

HasAllowedCriticalOptions returns a boolean if a field has been set.

### GetAllowedDomains

`func (o *SshRolesRequest) GetAllowedDomains() string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *SshRolesRequest) GetAllowedDomainsOk() (*string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *SshRolesRequest) SetAllowedDomains(v string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *SshRolesRequest) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### GetAllowedExtensions

`func (o *SshRolesRequest) GetAllowedExtensions() string`

GetAllowedExtensions returns the AllowedExtensions field if non-nil, zero value otherwise.

### GetAllowedExtensionsOk

`func (o *SshRolesRequest) GetAllowedExtensionsOk() (*string, bool)`

GetAllowedExtensionsOk returns a tuple with the AllowedExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedExtensions

`func (o *SshRolesRequest) SetAllowedExtensions(v string)`

SetAllowedExtensions sets AllowedExtensions field to given value.

### HasAllowedExtensions

`func (o *SshRolesRequest) HasAllowedExtensions() bool`

HasAllowedExtensions returns a boolean if a field has been set.

### GetAllowedUserKeyLengths

`func (o *SshRolesRequest) GetAllowedUserKeyLengths() map[string]interface{}`

GetAllowedUserKeyLengths returns the AllowedUserKeyLengths field if non-nil, zero value otherwise.

### GetAllowedUserKeyLengthsOk

`func (o *SshRolesRequest) GetAllowedUserKeyLengthsOk() (*map[string]interface{}, bool)`

GetAllowedUserKeyLengthsOk returns a tuple with the AllowedUserKeyLengths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUserKeyLengths

`func (o *SshRolesRequest) SetAllowedUserKeyLengths(v map[string]interface{})`

SetAllowedUserKeyLengths sets AllowedUserKeyLengths field to given value.

### HasAllowedUserKeyLengths

`func (o *SshRolesRequest) HasAllowedUserKeyLengths() bool`

HasAllowedUserKeyLengths returns a boolean if a field has been set.

### GetAllowedUsers

`func (o *SshRolesRequest) GetAllowedUsers() string`

GetAllowedUsers returns the AllowedUsers field if non-nil, zero value otherwise.

### GetAllowedUsersOk

`func (o *SshRolesRequest) GetAllowedUsersOk() (*string, bool)`

GetAllowedUsersOk returns a tuple with the AllowedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsers

`func (o *SshRolesRequest) SetAllowedUsers(v string)`

SetAllowedUsers sets AllowedUsers field to given value.

### HasAllowedUsers

`func (o *SshRolesRequest) HasAllowedUsers() bool`

HasAllowedUsers returns a boolean if a field has been set.

### GetAllowedUsersTemplate

`func (o *SshRolesRequest) GetAllowedUsersTemplate() bool`

GetAllowedUsersTemplate returns the AllowedUsersTemplate field if non-nil, zero value otherwise.

### GetAllowedUsersTemplateOk

`func (o *SshRolesRequest) GetAllowedUsersTemplateOk() (*bool, bool)`

GetAllowedUsersTemplateOk returns a tuple with the AllowedUsersTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsersTemplate

`func (o *SshRolesRequest) SetAllowedUsersTemplate(v bool)`

SetAllowedUsersTemplate sets AllowedUsersTemplate field to given value.

### HasAllowedUsersTemplate

`func (o *SshRolesRequest) HasAllowedUsersTemplate() bool`

HasAllowedUsersTemplate returns a boolean if a field has been set.

### GetCidrList

`func (o *SshRolesRequest) GetCidrList() string`

GetCidrList returns the CidrList field if non-nil, zero value otherwise.

### GetCidrListOk

`func (o *SshRolesRequest) GetCidrListOk() (*string, bool)`

GetCidrListOk returns a tuple with the CidrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrList

`func (o *SshRolesRequest) SetCidrList(v string)`

SetCidrList sets CidrList field to given value.

### HasCidrList

`func (o *SshRolesRequest) HasCidrList() bool`

HasCidrList returns a boolean if a field has been set.

### GetDefaultCriticalOptions

`func (o *SshRolesRequest) GetDefaultCriticalOptions() map[string]interface{}`

GetDefaultCriticalOptions returns the DefaultCriticalOptions field if non-nil, zero value otherwise.

### GetDefaultCriticalOptionsOk

`func (o *SshRolesRequest) GetDefaultCriticalOptionsOk() (*map[string]interface{}, bool)`

GetDefaultCriticalOptionsOk returns a tuple with the DefaultCriticalOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCriticalOptions

`func (o *SshRolesRequest) SetDefaultCriticalOptions(v map[string]interface{})`

SetDefaultCriticalOptions sets DefaultCriticalOptions field to given value.

### HasDefaultCriticalOptions

`func (o *SshRolesRequest) HasDefaultCriticalOptions() bool`

HasDefaultCriticalOptions returns a boolean if a field has been set.

### GetDefaultExtensions

`func (o *SshRolesRequest) GetDefaultExtensions() map[string]interface{}`

GetDefaultExtensions returns the DefaultExtensions field if non-nil, zero value otherwise.

### GetDefaultExtensionsOk

`func (o *SshRolesRequest) GetDefaultExtensionsOk() (*map[string]interface{}, bool)`

GetDefaultExtensionsOk returns a tuple with the DefaultExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExtensions

`func (o *SshRolesRequest) SetDefaultExtensions(v map[string]interface{})`

SetDefaultExtensions sets DefaultExtensions field to given value.

### HasDefaultExtensions

`func (o *SshRolesRequest) HasDefaultExtensions() bool`

HasDefaultExtensions returns a boolean if a field has been set.

### GetDefaultExtensionsTemplate

`func (o *SshRolesRequest) GetDefaultExtensionsTemplate() bool`

GetDefaultExtensionsTemplate returns the DefaultExtensionsTemplate field if non-nil, zero value otherwise.

### GetDefaultExtensionsTemplateOk

`func (o *SshRolesRequest) GetDefaultExtensionsTemplateOk() (*bool, bool)`

GetDefaultExtensionsTemplateOk returns a tuple with the DefaultExtensionsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExtensionsTemplate

`func (o *SshRolesRequest) SetDefaultExtensionsTemplate(v bool)`

SetDefaultExtensionsTemplate sets DefaultExtensionsTemplate field to given value.

### HasDefaultExtensionsTemplate

`func (o *SshRolesRequest) HasDefaultExtensionsTemplate() bool`

HasDefaultExtensionsTemplate returns a boolean if a field has been set.

### GetDefaultUser

`func (o *SshRolesRequest) GetDefaultUser() string`

GetDefaultUser returns the DefaultUser field if non-nil, zero value otherwise.

### GetDefaultUserOk

`func (o *SshRolesRequest) GetDefaultUserOk() (*string, bool)`

GetDefaultUserOk returns a tuple with the DefaultUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUser

`func (o *SshRolesRequest) SetDefaultUser(v string)`

SetDefaultUser sets DefaultUser field to given value.

### HasDefaultUser

`func (o *SshRolesRequest) HasDefaultUser() bool`

HasDefaultUser returns a boolean if a field has been set.

### GetExcludeCidrList

`func (o *SshRolesRequest) GetExcludeCidrList() string`

GetExcludeCidrList returns the ExcludeCidrList field if non-nil, zero value otherwise.

### GetExcludeCidrListOk

`func (o *SshRolesRequest) GetExcludeCidrListOk() (*string, bool)`

GetExcludeCidrListOk returns a tuple with the ExcludeCidrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCidrList

`func (o *SshRolesRequest) SetExcludeCidrList(v string)`

SetExcludeCidrList sets ExcludeCidrList field to given value.

### HasExcludeCidrList

`func (o *SshRolesRequest) HasExcludeCidrList() bool`

HasExcludeCidrList returns a boolean if a field has been set.

### GetInstallScript

`func (o *SshRolesRequest) GetInstallScript() string`

GetInstallScript returns the InstallScript field if non-nil, zero value otherwise.

### GetInstallScriptOk

`func (o *SshRolesRequest) GetInstallScriptOk() (*string, bool)`

GetInstallScriptOk returns a tuple with the InstallScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallScript

`func (o *SshRolesRequest) SetInstallScript(v string)`

SetInstallScript sets InstallScript field to given value.

### HasInstallScript

`func (o *SshRolesRequest) HasInstallScript() bool`

HasInstallScript returns a boolean if a field has been set.

### GetKey

`func (o *SshRolesRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SshRolesRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SshRolesRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SshRolesRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKeyBits

`func (o *SshRolesRequest) GetKeyBits() int32`

GetKeyBits returns the KeyBits field if non-nil, zero value otherwise.

### GetKeyBitsOk

`func (o *SshRolesRequest) GetKeyBitsOk() (*int32, bool)`

GetKeyBitsOk returns a tuple with the KeyBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyBits

`func (o *SshRolesRequest) SetKeyBits(v int32)`

SetKeyBits sets KeyBits field to given value.

### HasKeyBits

`func (o *SshRolesRequest) HasKeyBits() bool`

HasKeyBits returns a boolean if a field has been set.

### GetKeyIdFormat

`func (o *SshRolesRequest) GetKeyIdFormat() string`

GetKeyIdFormat returns the KeyIdFormat field if non-nil, zero value otherwise.

### GetKeyIdFormatOk

`func (o *SshRolesRequest) GetKeyIdFormatOk() (*string, bool)`

GetKeyIdFormatOk returns a tuple with the KeyIdFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyIdFormat

`func (o *SshRolesRequest) SetKeyIdFormat(v string)`

SetKeyIdFormat sets KeyIdFormat field to given value.

### HasKeyIdFormat

`func (o *SshRolesRequest) HasKeyIdFormat() bool`

HasKeyIdFormat returns a boolean if a field has been set.

### GetKeyOptionSpecs

`func (o *SshRolesRequest) GetKeyOptionSpecs() string`

GetKeyOptionSpecs returns the KeyOptionSpecs field if non-nil, zero value otherwise.

### GetKeyOptionSpecsOk

`func (o *SshRolesRequest) GetKeyOptionSpecsOk() (*string, bool)`

GetKeyOptionSpecsOk returns a tuple with the KeyOptionSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyOptionSpecs

`func (o *SshRolesRequest) SetKeyOptionSpecs(v string)`

SetKeyOptionSpecs sets KeyOptionSpecs field to given value.

### HasKeyOptionSpecs

`func (o *SshRolesRequest) HasKeyOptionSpecs() bool`

HasKeyOptionSpecs returns a boolean if a field has been set.

### GetKeyType

`func (o *SshRolesRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *SshRolesRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *SshRolesRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *SshRolesRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetMaxTtl

`func (o *SshRolesRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *SshRolesRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *SshRolesRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *SshRolesRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetPort

`func (o *SshRolesRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SshRolesRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SshRolesRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SshRolesRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTtl

`func (o *SshRolesRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *SshRolesRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *SshRolesRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *SshRolesRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


