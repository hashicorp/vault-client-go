# CertWriteCertificateRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedCommonNames** | Pointer to **[]string** | A comma-separated list of names. At least one must exist in the Common Name. Supports globbing. | [optional] 
**AllowedDnsSans** | Pointer to **[]string** | A comma-separated list of DNS names. At least one must exist in the SANs. Supports globbing. | [optional] 
**AllowedEmailSans** | Pointer to **[]string** | A comma-separated list of Email Addresses. At least one must exist in the SANs. Supports globbing. | [optional] 
**AllowedMetadataExtensions** | Pointer to **[]string** | A comma-separated string or array of oid extensions. Upon successful authentication, these extensions will be added as metadata if they are present in the certificate. The metadata key will be the string consisting of the oid numbers separated by a dash (-) instead of a dot (.) to allow usage in ACL templates. | [optional] 
**AllowedNames** | Pointer to **[]string** | A comma-separated list of names. At least one must exist in either the Common Name or SANs. Supports globbing. This parameter is deprecated, please use allowed_common_names, allowed_dns_sans, allowed_email_sans, allowed_uri_sans. | [optional] 
**AllowedOrganizationalUnits** | Pointer to **[]string** | A comma-separated list of Organizational Units names. At least one must exist in the OU field. | [optional] 
**AllowedUriSans** | Pointer to **[]string** | A comma-separated list of URIs. At least one must exist in the SANs. Supports globbing. | [optional] 
**BoundCidrs** | Pointer to **[]string** | Use \&quot;token_bound_cidrs\&quot; instead. If this and \&quot;token_bound_cidrs\&quot; are both specified, only \&quot;token_bound_cidrs\&quot; will be used. | [optional] 
**Certificate** | Pointer to **string** | The public certificate that should be trusted. Must be x509 PEM encoded. | [optional] 
**DisplayName** | Pointer to **string** | The display name to use for clients using this certificate. | [optional] 
**Lease** | Pointer to **int32** | Use \&quot;token_ttl\&quot; instead. If this and \&quot;token_ttl\&quot; are both specified, only \&quot;token_ttl\&quot; will be used. | [optional] 
**MaxTtl** | Pointer to **string** | Use \&quot;token_max_ttl\&quot; instead. If this and \&quot;token_max_ttl\&quot; are both specified, only \&quot;token_max_ttl\&quot; will be used. | [optional] 
**OcspCaCertificates** | Pointer to **string** | Any additional CA certificates needed to communicate with OCSP servers | [optional] 
**OcspEnabled** | Pointer to **bool** | Whether to attempt OCSP verification of certificates at login | [optional] 
**OcspFailOpen** | Pointer to **bool** | If set to true, if an OCSP revocation cannot be made successfully, login will proceed rather than failing. If false, failing to get an OCSP status fails the request. | [optional] [default to false]
**OcspQueryAllServers** | Pointer to **bool** | If set to true, rather than accepting the first successful OCSP response, query all servers and consider the certificate valid only if all servers agree. | [optional] [default to false]
**OcspServersOverride** | Pointer to **[]string** | A comma-separated list of OCSP server addresses. If unset, the OCSP server is determined from the AuthorityInformationAccess extension on the certificate being inspected. | [optional] 
**Period** | Pointer to **string** | Use \&quot;token_period\&quot; instead. If this and \&quot;token_period\&quot; are both specified, only \&quot;token_period\&quot; will be used. | [optional] 
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
**RequiredExtensions** | Pointer to **[]string** | A comma-separated string or array of extensions formatted as \&quot;oid:value\&quot;. Expects the extension value to be some type of ASN1 encoded string. All values much match. Supports globbing on \&quot;value\&quot;. | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **string** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **string** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#x27;default&#x27; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **string** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies | [optional] 
**TokenTtl** | Pointer to **string** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]
**Ttl** | Pointer to **string** | Use \&quot;token_ttl\&quot; instead. If this and \&quot;token_ttl\&quot; are both specified, only \&quot;token_ttl\&quot; will be used. | [optional] 



## Methods


### NewCertWriteCertificateRequest

`func NewCertWriteCertificateRequest() *CertWriteCertificateRequest`

NewCertWriteCertificateRequest instantiates a new CertWriteCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertWriteCertificateRequestWithDefaults

`func NewCertWriteCertificateRequestWithDefaults() *CertWriteCertificateRequest`

NewCertWriteCertificateRequestWithDefaults instantiates a new CertWriteCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAllowedCommonNames

`func (o *CertWriteCertificateRequest) GetAllowedCommonNames() []string`

GetAllowedCommonNames returns the AllowedCommonNames field if non-nil, zero value otherwise.

### GetAllowedCommonNamesOk

`func (o *CertWriteCertificateRequest) GetAllowedCommonNamesOk() (*[]string, bool)`

GetAllowedCommonNamesOk returns a tuple with the AllowedCommonNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCommonNames

`func (o *CertWriteCertificateRequest) SetAllowedCommonNames(v []string)`

SetAllowedCommonNames sets AllowedCommonNames field to given value.


### HasAllowedCommonNames

`func (o *CertWriteCertificateRequest) HasAllowedCommonNames() bool`

HasAllowedCommonNames returns a boolean if a field has been set.




### GetAllowedDnsSans

`func (o *CertWriteCertificateRequest) GetAllowedDnsSans() []string`

GetAllowedDnsSans returns the AllowedDnsSans field if non-nil, zero value otherwise.

### GetAllowedDnsSansOk

`func (o *CertWriteCertificateRequest) GetAllowedDnsSansOk() (*[]string, bool)`

GetAllowedDnsSansOk returns a tuple with the AllowedDnsSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDnsSans

`func (o *CertWriteCertificateRequest) SetAllowedDnsSans(v []string)`

SetAllowedDnsSans sets AllowedDnsSans field to given value.


### HasAllowedDnsSans

`func (o *CertWriteCertificateRequest) HasAllowedDnsSans() bool`

HasAllowedDnsSans returns a boolean if a field has been set.




### GetAllowedEmailSans

`func (o *CertWriteCertificateRequest) GetAllowedEmailSans() []string`

GetAllowedEmailSans returns the AllowedEmailSans field if non-nil, zero value otherwise.

### GetAllowedEmailSansOk

`func (o *CertWriteCertificateRequest) GetAllowedEmailSansOk() (*[]string, bool)`

GetAllowedEmailSansOk returns a tuple with the AllowedEmailSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEmailSans

`func (o *CertWriteCertificateRequest) SetAllowedEmailSans(v []string)`

SetAllowedEmailSans sets AllowedEmailSans field to given value.


### HasAllowedEmailSans

`func (o *CertWriteCertificateRequest) HasAllowedEmailSans() bool`

HasAllowedEmailSans returns a boolean if a field has been set.




### GetAllowedMetadataExtensions

`func (o *CertWriteCertificateRequest) GetAllowedMetadataExtensions() []string`

GetAllowedMetadataExtensions returns the AllowedMetadataExtensions field if non-nil, zero value otherwise.

### GetAllowedMetadataExtensionsOk

`func (o *CertWriteCertificateRequest) GetAllowedMetadataExtensionsOk() (*[]string, bool)`

GetAllowedMetadataExtensionsOk returns a tuple with the AllowedMetadataExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMetadataExtensions

`func (o *CertWriteCertificateRequest) SetAllowedMetadataExtensions(v []string)`

SetAllowedMetadataExtensions sets AllowedMetadataExtensions field to given value.


### HasAllowedMetadataExtensions

`func (o *CertWriteCertificateRequest) HasAllowedMetadataExtensions() bool`

HasAllowedMetadataExtensions returns a boolean if a field has been set.




### GetAllowedNames

`func (o *CertWriteCertificateRequest) GetAllowedNames() []string`

GetAllowedNames returns the AllowedNames field if non-nil, zero value otherwise.

### GetAllowedNamesOk

`func (o *CertWriteCertificateRequest) GetAllowedNamesOk() (*[]string, bool)`

GetAllowedNamesOk returns a tuple with the AllowedNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNames

`func (o *CertWriteCertificateRequest) SetAllowedNames(v []string)`

SetAllowedNames sets AllowedNames field to given value.


### HasAllowedNames

`func (o *CertWriteCertificateRequest) HasAllowedNames() bool`

HasAllowedNames returns a boolean if a field has been set.




### GetAllowedOrganizationalUnits

`func (o *CertWriteCertificateRequest) GetAllowedOrganizationalUnits() []string`

GetAllowedOrganizationalUnits returns the AllowedOrganizationalUnits field if non-nil, zero value otherwise.

### GetAllowedOrganizationalUnitsOk

`func (o *CertWriteCertificateRequest) GetAllowedOrganizationalUnitsOk() (*[]string, bool)`

GetAllowedOrganizationalUnitsOk returns a tuple with the AllowedOrganizationalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrganizationalUnits

`func (o *CertWriteCertificateRequest) SetAllowedOrganizationalUnits(v []string)`

SetAllowedOrganizationalUnits sets AllowedOrganizationalUnits field to given value.


### HasAllowedOrganizationalUnits

`func (o *CertWriteCertificateRequest) HasAllowedOrganizationalUnits() bool`

HasAllowedOrganizationalUnits returns a boolean if a field has been set.




### GetAllowedUriSans

`func (o *CertWriteCertificateRequest) GetAllowedUriSans() []string`

GetAllowedUriSans returns the AllowedUriSans field if non-nil, zero value otherwise.

### GetAllowedUriSansOk

`func (o *CertWriteCertificateRequest) GetAllowedUriSansOk() (*[]string, bool)`

GetAllowedUriSansOk returns a tuple with the AllowedUriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUriSans

`func (o *CertWriteCertificateRequest) SetAllowedUriSans(v []string)`

SetAllowedUriSans sets AllowedUriSans field to given value.


### HasAllowedUriSans

`func (o *CertWriteCertificateRequest) HasAllowedUriSans() bool`

HasAllowedUriSans returns a boolean if a field has been set.




### GetBoundCidrs

`func (o *CertWriteCertificateRequest) GetBoundCidrs() []string`

GetBoundCidrs returns the BoundCidrs field if non-nil, zero value otherwise.

### GetBoundCidrsOk

`func (o *CertWriteCertificateRequest) GetBoundCidrsOk() (*[]string, bool)`

GetBoundCidrsOk returns a tuple with the BoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCidrs

`func (o *CertWriteCertificateRequest) SetBoundCidrs(v []string)`

SetBoundCidrs sets BoundCidrs field to given value.


### HasBoundCidrs

`func (o *CertWriteCertificateRequest) HasBoundCidrs() bool`

HasBoundCidrs returns a boolean if a field has been set.




### GetCertificate

`func (o *CertWriteCertificateRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertWriteCertificateRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertWriteCertificateRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### HasCertificate

`func (o *CertWriteCertificateRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.




### GetDisplayName

`func (o *CertWriteCertificateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CertWriteCertificateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CertWriteCertificateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### HasDisplayName

`func (o *CertWriteCertificateRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.




### GetLease

`func (o *CertWriteCertificateRequest) GetLease() int32`

GetLease returns the Lease field if non-nil, zero value otherwise.

### GetLeaseOk

`func (o *CertWriteCertificateRequest) GetLeaseOk() (*int32, bool)`

GetLeaseOk returns a tuple with the Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease

`func (o *CertWriteCertificateRequest) SetLease(v int32)`

SetLease sets Lease field to given value.


### HasLease

`func (o *CertWriteCertificateRequest) HasLease() bool`

HasLease returns a boolean if a field has been set.




### GetMaxTtl

`func (o *CertWriteCertificateRequest) GetMaxTtl() string`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *CertWriteCertificateRequest) GetMaxTtlOk() (*string, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *CertWriteCertificateRequest) SetMaxTtl(v string)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *CertWriteCertificateRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetOcspCaCertificates

`func (o *CertWriteCertificateRequest) GetOcspCaCertificates() string`

GetOcspCaCertificates returns the OcspCaCertificates field if non-nil, zero value otherwise.

### GetOcspCaCertificatesOk

`func (o *CertWriteCertificateRequest) GetOcspCaCertificatesOk() (*string, bool)`

GetOcspCaCertificatesOk returns a tuple with the OcspCaCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspCaCertificates

`func (o *CertWriteCertificateRequest) SetOcspCaCertificates(v string)`

SetOcspCaCertificates sets OcspCaCertificates field to given value.


### HasOcspCaCertificates

`func (o *CertWriteCertificateRequest) HasOcspCaCertificates() bool`

HasOcspCaCertificates returns a boolean if a field has been set.




### GetOcspEnabled

`func (o *CertWriteCertificateRequest) GetOcspEnabled() bool`

GetOcspEnabled returns the OcspEnabled field if non-nil, zero value otherwise.

### GetOcspEnabledOk

`func (o *CertWriteCertificateRequest) GetOcspEnabledOk() (*bool, bool)`

GetOcspEnabledOk returns a tuple with the OcspEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspEnabled

`func (o *CertWriteCertificateRequest) SetOcspEnabled(v bool)`

SetOcspEnabled sets OcspEnabled field to given value.


### HasOcspEnabled

`func (o *CertWriteCertificateRequest) HasOcspEnabled() bool`

HasOcspEnabled returns a boolean if a field has been set.




### GetOcspFailOpen

`func (o *CertWriteCertificateRequest) GetOcspFailOpen() bool`

GetOcspFailOpen returns the OcspFailOpen field if non-nil, zero value otherwise.

### GetOcspFailOpenOk

`func (o *CertWriteCertificateRequest) GetOcspFailOpenOk() (*bool, bool)`

GetOcspFailOpenOk returns a tuple with the OcspFailOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspFailOpen

`func (o *CertWriteCertificateRequest) SetOcspFailOpen(v bool)`

SetOcspFailOpen sets OcspFailOpen field to given value.


### HasOcspFailOpen

`func (o *CertWriteCertificateRequest) HasOcspFailOpen() bool`

HasOcspFailOpen returns a boolean if a field has been set.




### GetOcspQueryAllServers

`func (o *CertWriteCertificateRequest) GetOcspQueryAllServers() bool`

GetOcspQueryAllServers returns the OcspQueryAllServers field if non-nil, zero value otherwise.

### GetOcspQueryAllServersOk

`func (o *CertWriteCertificateRequest) GetOcspQueryAllServersOk() (*bool, bool)`

GetOcspQueryAllServersOk returns a tuple with the OcspQueryAllServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspQueryAllServers

`func (o *CertWriteCertificateRequest) SetOcspQueryAllServers(v bool)`

SetOcspQueryAllServers sets OcspQueryAllServers field to given value.


### HasOcspQueryAllServers

`func (o *CertWriteCertificateRequest) HasOcspQueryAllServers() bool`

HasOcspQueryAllServers returns a boolean if a field has been set.




### GetOcspServersOverride

`func (o *CertWriteCertificateRequest) GetOcspServersOverride() []string`

GetOcspServersOverride returns the OcspServersOverride field if non-nil, zero value otherwise.

### GetOcspServersOverrideOk

`func (o *CertWriteCertificateRequest) GetOcspServersOverrideOk() (*[]string, bool)`

GetOcspServersOverrideOk returns a tuple with the OcspServersOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspServersOverride

`func (o *CertWriteCertificateRequest) SetOcspServersOverride(v []string)`

SetOcspServersOverride sets OcspServersOverride field to given value.


### HasOcspServersOverride

`func (o *CertWriteCertificateRequest) HasOcspServersOverride() bool`

HasOcspServersOverride returns a boolean if a field has been set.




### GetPeriod

`func (o *CertWriteCertificateRequest) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CertWriteCertificateRequest) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CertWriteCertificateRequest) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### HasPeriod

`func (o *CertWriteCertificateRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.




### GetPolicies

`func (o *CertWriteCertificateRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CertWriteCertificateRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CertWriteCertificateRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *CertWriteCertificateRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.




### GetRequiredExtensions

`func (o *CertWriteCertificateRequest) GetRequiredExtensions() []string`

GetRequiredExtensions returns the RequiredExtensions field if non-nil, zero value otherwise.

### GetRequiredExtensionsOk

`func (o *CertWriteCertificateRequest) GetRequiredExtensionsOk() (*[]string, bool)`

GetRequiredExtensionsOk returns a tuple with the RequiredExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredExtensions

`func (o *CertWriteCertificateRequest) SetRequiredExtensions(v []string)`

SetRequiredExtensions sets RequiredExtensions field to given value.


### HasRequiredExtensions

`func (o *CertWriteCertificateRequest) HasRequiredExtensions() bool`

HasRequiredExtensions returns a boolean if a field has been set.




### GetTokenBoundCidrs

`func (o *CertWriteCertificateRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *CertWriteCertificateRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *CertWriteCertificateRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.


### HasTokenBoundCidrs

`func (o *CertWriteCertificateRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.




### GetTokenExplicitMaxTtl

`func (o *CertWriteCertificateRequest) GetTokenExplicitMaxTtl() string`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *CertWriteCertificateRequest) GetTokenExplicitMaxTtlOk() (*string, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *CertWriteCertificateRequest) SetTokenExplicitMaxTtl(v string)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.


### HasTokenExplicitMaxTtl

`func (o *CertWriteCertificateRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.




### GetTokenMaxTtl

`func (o *CertWriteCertificateRequest) GetTokenMaxTtl() string`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *CertWriteCertificateRequest) GetTokenMaxTtlOk() (*string, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *CertWriteCertificateRequest) SetTokenMaxTtl(v string)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.


### HasTokenMaxTtl

`func (o *CertWriteCertificateRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.




### GetTokenNoDefaultPolicy

`func (o *CertWriteCertificateRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *CertWriteCertificateRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *CertWriteCertificateRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.


### HasTokenNoDefaultPolicy

`func (o *CertWriteCertificateRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.




### GetTokenNumUses

`func (o *CertWriteCertificateRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *CertWriteCertificateRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *CertWriteCertificateRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.


### HasTokenNumUses

`func (o *CertWriteCertificateRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.




### GetTokenPeriod

`func (o *CertWriteCertificateRequest) GetTokenPeriod() string`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *CertWriteCertificateRequest) GetTokenPeriodOk() (*string, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *CertWriteCertificateRequest) SetTokenPeriod(v string)`

SetTokenPeriod sets TokenPeriod field to given value.


### HasTokenPeriod

`func (o *CertWriteCertificateRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.




### GetTokenPolicies

`func (o *CertWriteCertificateRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *CertWriteCertificateRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *CertWriteCertificateRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.


### HasTokenPolicies

`func (o *CertWriteCertificateRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.




### GetTokenTtl

`func (o *CertWriteCertificateRequest) GetTokenTtl() string`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *CertWriteCertificateRequest) GetTokenTtlOk() (*string, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *CertWriteCertificateRequest) SetTokenTtl(v string)`

SetTokenTtl sets TokenTtl field to given value.


### HasTokenTtl

`func (o *CertWriteCertificateRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.




### GetTokenType

`func (o *CertWriteCertificateRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *CertWriteCertificateRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *CertWriteCertificateRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### HasTokenType

`func (o *CertWriteCertificateRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.




### GetTtl

`func (o *CertWriteCertificateRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CertWriteCertificateRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CertWriteCertificateRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *CertWriteCertificateRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


