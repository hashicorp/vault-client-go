# CertCertsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedCommonNames** | Pointer to **[]string** | A comma-separated list of names. At least one must exist in the Common Name. Supports globbing. | [optional] 
**AllowedDnsSans** | Pointer to **[]string** | A comma-separated list of DNS names. At least one must exist in the SANs. Supports globbing. | [optional] 
**AllowedEmailSans** | Pointer to **[]string** | A comma-separated list of Email Addresses. At least one must exist in the SANs. Supports globbing. | [optional] 
**AllowedMetadataExtensions** | Pointer to **[]string** | A comma-separated string or array of oid extensions. Upon successfull authentication, these extensions will be added as metadata if they are present in the certificate. The metadata key will be the string consisting of the oid numbers separated by a dash (-) instead of a dot (.) to allow usage in ACL templates. | [optional] 
**AllowedNames** | Pointer to **[]string** | A comma-separated list of names. At least one must exist in either the Common Name or SANs. Supports globbing. This parameter is deprecated, please use allowed_common_names, allowed_dns_sans, allowed_email_sans, allowed_uri_sans. | [optional] 
**AllowedOrganizationalUnits** | Pointer to **[]string** | A comma-separated list of Organizational Units names. At least one must exist in the OU field. | [optional] 
**AllowedUriSans** | Pointer to **[]string** | A comma-separated list of URIs. At least one must exist in the SANs. Supports globbing. | [optional] 
**BoundCidrs** | Pointer to **[]string** | Use \&quot;token_bound_cidrs\&quot; instead. If this and \&quot;token_bound_cidrs\&quot; are both specified, only \&quot;token_bound_cidrs\&quot; will be used. | [optional] 
**Certificate** | Pointer to **string** | The public certificate that should be trusted. Must be x509 PEM encoded. | [optional] 
**DisplayName** | Pointer to **string** | The display name to use for clients using this certificate. | [optional] 
**Lease** | Pointer to **int32** | Use \&quot;token_ttl\&quot; instead. If this and \&quot;token_ttl\&quot; are both specified, only \&quot;token_ttl\&quot; will be used. | [optional] 
**MaxTtl** | Pointer to **int32** | Use \&quot;token_max_ttl\&quot; instead. If this and \&quot;token_max_ttl\&quot; are both specified, only \&quot;token_max_ttl\&quot; will be used. | [optional] 
**OcspCaCertificates** | Pointer to **string** | Any additional CA certificates needed to communicate with OCSP servers | [optional] 
**OcspEnabled** | Pointer to **bool** | Whether to attempt OCSP verification of certificates at login | [optional] 
**OcspFailOpen** | Pointer to **bool** | If set to true, if an OCSP revocation cannot be made successfully, login will proceed rather than failing. If false, failing to get an OCSP status fails the request. | [optional] [default to false]
**OcspQueryAllServers** | Pointer to **bool** | If set to true, rather than accepting the first successful OCSP response, query all servers and consider the certificate valid only if all servers agree. | [optional] [default to false]
**OcspServersOverride** | Pointer to **[]string** | A comma-separated list of OCSP server addresses. If unset, the OCSP server is determined from the AuthorityInformationAccess extension on the certificate being inspected. | [optional] 
**Period** | Pointer to **int32** | Use \&quot;token_period\&quot; instead. If this and \&quot;token_period\&quot; are both specified, only \&quot;token_period\&quot; will be used. | [optional] 
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
**RequiredExtensions** | Pointer to **[]string** | A comma-separated string or array of extensions formatted as \&quot;oid:value\&quot;. Expects the extension value to be some type of ASN1 encoded string. All values much match. Supports globbing on \&quot;value\&quot;. | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **int32** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **int32** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#39;default&#39; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **int32** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies | [optional] 
**TokenTtl** | Pointer to **int32** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]
**Ttl** | Pointer to **int32** | Use \&quot;token_ttl\&quot; instead. If this and \&quot;token_ttl\&quot; are both specified, only \&quot;token_ttl\&quot; will be used. | [optional] 

## Methods

### NewCertCertsRequest

`func NewCertCertsRequest() *CertCertsRequest`

NewCertCertsRequest instantiates a new CertCertsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertCertsRequestWithDefaults

`func NewCertCertsRequestWithDefaults() *CertCertsRequest`

NewCertCertsRequestWithDefaults instantiates a new CertCertsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedCommonNames

`func (o *CertCertsRequest) GetAllowedCommonNames() []string`

GetAllowedCommonNames returns the AllowedCommonNames field if non-nil, zero value otherwise.

### GetAllowedCommonNamesOk

`func (o *CertCertsRequest) GetAllowedCommonNamesOk() (*[]string, bool)`

GetAllowedCommonNamesOk returns a tuple with the AllowedCommonNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCommonNames

`func (o *CertCertsRequest) SetAllowedCommonNames(v []string)`

SetAllowedCommonNames sets AllowedCommonNames field to given value.

### HasAllowedCommonNames

`func (o *CertCertsRequest) HasAllowedCommonNames() bool`

HasAllowedCommonNames returns a boolean if a field has been set.

### GetAllowedDnsSans

`func (o *CertCertsRequest) GetAllowedDnsSans() []string`

GetAllowedDnsSans returns the AllowedDnsSans field if non-nil, zero value otherwise.

### GetAllowedDnsSansOk

`func (o *CertCertsRequest) GetAllowedDnsSansOk() (*[]string, bool)`

GetAllowedDnsSansOk returns a tuple with the AllowedDnsSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDnsSans

`func (o *CertCertsRequest) SetAllowedDnsSans(v []string)`

SetAllowedDnsSans sets AllowedDnsSans field to given value.

### HasAllowedDnsSans

`func (o *CertCertsRequest) HasAllowedDnsSans() bool`

HasAllowedDnsSans returns a boolean if a field has been set.

### GetAllowedEmailSans

`func (o *CertCertsRequest) GetAllowedEmailSans() []string`

GetAllowedEmailSans returns the AllowedEmailSans field if non-nil, zero value otherwise.

### GetAllowedEmailSansOk

`func (o *CertCertsRequest) GetAllowedEmailSansOk() (*[]string, bool)`

GetAllowedEmailSansOk returns a tuple with the AllowedEmailSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEmailSans

`func (o *CertCertsRequest) SetAllowedEmailSans(v []string)`

SetAllowedEmailSans sets AllowedEmailSans field to given value.

### HasAllowedEmailSans

`func (o *CertCertsRequest) HasAllowedEmailSans() bool`

HasAllowedEmailSans returns a boolean if a field has been set.

### GetAllowedMetadataExtensions

`func (o *CertCertsRequest) GetAllowedMetadataExtensions() []string`

GetAllowedMetadataExtensions returns the AllowedMetadataExtensions field if non-nil, zero value otherwise.

### GetAllowedMetadataExtensionsOk

`func (o *CertCertsRequest) GetAllowedMetadataExtensionsOk() (*[]string, bool)`

GetAllowedMetadataExtensionsOk returns a tuple with the AllowedMetadataExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMetadataExtensions

`func (o *CertCertsRequest) SetAllowedMetadataExtensions(v []string)`

SetAllowedMetadataExtensions sets AllowedMetadataExtensions field to given value.

### HasAllowedMetadataExtensions

`func (o *CertCertsRequest) HasAllowedMetadataExtensions() bool`

HasAllowedMetadataExtensions returns a boolean if a field has been set.

### GetAllowedNames

`func (o *CertCertsRequest) GetAllowedNames() []string`

GetAllowedNames returns the AllowedNames field if non-nil, zero value otherwise.

### GetAllowedNamesOk

`func (o *CertCertsRequest) GetAllowedNamesOk() (*[]string, bool)`

GetAllowedNamesOk returns a tuple with the AllowedNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNames

`func (o *CertCertsRequest) SetAllowedNames(v []string)`

SetAllowedNames sets AllowedNames field to given value.

### HasAllowedNames

`func (o *CertCertsRequest) HasAllowedNames() bool`

HasAllowedNames returns a boolean if a field has been set.

### GetAllowedOrganizationalUnits

`func (o *CertCertsRequest) GetAllowedOrganizationalUnits() []string`

GetAllowedOrganizationalUnits returns the AllowedOrganizationalUnits field if non-nil, zero value otherwise.

### GetAllowedOrganizationalUnitsOk

`func (o *CertCertsRequest) GetAllowedOrganizationalUnitsOk() (*[]string, bool)`

GetAllowedOrganizationalUnitsOk returns a tuple with the AllowedOrganizationalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrganizationalUnits

`func (o *CertCertsRequest) SetAllowedOrganizationalUnits(v []string)`

SetAllowedOrganizationalUnits sets AllowedOrganizationalUnits field to given value.

### HasAllowedOrganizationalUnits

`func (o *CertCertsRequest) HasAllowedOrganizationalUnits() bool`

HasAllowedOrganizationalUnits returns a boolean if a field has been set.

### GetAllowedUriSans

`func (o *CertCertsRequest) GetAllowedUriSans() []string`

GetAllowedUriSans returns the AllowedUriSans field if non-nil, zero value otherwise.

### GetAllowedUriSansOk

`func (o *CertCertsRequest) GetAllowedUriSansOk() (*[]string, bool)`

GetAllowedUriSansOk returns a tuple with the AllowedUriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUriSans

`func (o *CertCertsRequest) SetAllowedUriSans(v []string)`

SetAllowedUriSans sets AllowedUriSans field to given value.

### HasAllowedUriSans

`func (o *CertCertsRequest) HasAllowedUriSans() bool`

HasAllowedUriSans returns a boolean if a field has been set.

### GetBoundCidrs

`func (o *CertCertsRequest) GetBoundCidrs() []string`

GetBoundCidrs returns the BoundCidrs field if non-nil, zero value otherwise.

### GetBoundCidrsOk

`func (o *CertCertsRequest) GetBoundCidrsOk() (*[]string, bool)`

GetBoundCidrsOk returns a tuple with the BoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCidrs

`func (o *CertCertsRequest) SetBoundCidrs(v []string)`

SetBoundCidrs sets BoundCidrs field to given value.

### HasBoundCidrs

`func (o *CertCertsRequest) HasBoundCidrs() bool`

HasBoundCidrs returns a boolean if a field has been set.

### GetCertificate

`func (o *CertCertsRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertCertsRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertCertsRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertCertsRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetDisplayName

`func (o *CertCertsRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CertCertsRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CertCertsRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CertCertsRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLease

`func (o *CertCertsRequest) GetLease() int32`

GetLease returns the Lease field if non-nil, zero value otherwise.

### GetLeaseOk

`func (o *CertCertsRequest) GetLeaseOk() (*int32, bool)`

GetLeaseOk returns a tuple with the Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease

`func (o *CertCertsRequest) SetLease(v int32)`

SetLease sets Lease field to given value.

### HasLease

`func (o *CertCertsRequest) HasLease() bool`

HasLease returns a boolean if a field has been set.

### GetMaxTtl

`func (o *CertCertsRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *CertCertsRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *CertCertsRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *CertCertsRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetOcspCaCertificates

`func (o *CertCertsRequest) GetOcspCaCertificates() string`

GetOcspCaCertificates returns the OcspCaCertificates field if non-nil, zero value otherwise.

### GetOcspCaCertificatesOk

`func (o *CertCertsRequest) GetOcspCaCertificatesOk() (*string, bool)`

GetOcspCaCertificatesOk returns a tuple with the OcspCaCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspCaCertificates

`func (o *CertCertsRequest) SetOcspCaCertificates(v string)`

SetOcspCaCertificates sets OcspCaCertificates field to given value.

### HasOcspCaCertificates

`func (o *CertCertsRequest) HasOcspCaCertificates() bool`

HasOcspCaCertificates returns a boolean if a field has been set.

### GetOcspEnabled

`func (o *CertCertsRequest) GetOcspEnabled() bool`

GetOcspEnabled returns the OcspEnabled field if non-nil, zero value otherwise.

### GetOcspEnabledOk

`func (o *CertCertsRequest) GetOcspEnabledOk() (*bool, bool)`

GetOcspEnabledOk returns a tuple with the OcspEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspEnabled

`func (o *CertCertsRequest) SetOcspEnabled(v bool)`

SetOcspEnabled sets OcspEnabled field to given value.

### HasOcspEnabled

`func (o *CertCertsRequest) HasOcspEnabled() bool`

HasOcspEnabled returns a boolean if a field has been set.

### GetOcspFailOpen

`func (o *CertCertsRequest) GetOcspFailOpen() bool`

GetOcspFailOpen returns the OcspFailOpen field if non-nil, zero value otherwise.

### GetOcspFailOpenOk

`func (o *CertCertsRequest) GetOcspFailOpenOk() (*bool, bool)`

GetOcspFailOpenOk returns a tuple with the OcspFailOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspFailOpen

`func (o *CertCertsRequest) SetOcspFailOpen(v bool)`

SetOcspFailOpen sets OcspFailOpen field to given value.

### HasOcspFailOpen

`func (o *CertCertsRequest) HasOcspFailOpen() bool`

HasOcspFailOpen returns a boolean if a field has been set.

### GetOcspQueryAllServers

`func (o *CertCertsRequest) GetOcspQueryAllServers() bool`

GetOcspQueryAllServers returns the OcspQueryAllServers field if non-nil, zero value otherwise.

### GetOcspQueryAllServersOk

`func (o *CertCertsRequest) GetOcspQueryAllServersOk() (*bool, bool)`

GetOcspQueryAllServersOk returns a tuple with the OcspQueryAllServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspQueryAllServers

`func (o *CertCertsRequest) SetOcspQueryAllServers(v bool)`

SetOcspQueryAllServers sets OcspQueryAllServers field to given value.

### HasOcspQueryAllServers

`func (o *CertCertsRequest) HasOcspQueryAllServers() bool`

HasOcspQueryAllServers returns a boolean if a field has been set.

### GetOcspServersOverride

`func (o *CertCertsRequest) GetOcspServersOverride() []string`

GetOcspServersOverride returns the OcspServersOverride field if non-nil, zero value otherwise.

### GetOcspServersOverrideOk

`func (o *CertCertsRequest) GetOcspServersOverrideOk() (*[]string, bool)`

GetOcspServersOverrideOk returns a tuple with the OcspServersOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspServersOverride

`func (o *CertCertsRequest) SetOcspServersOverride(v []string)`

SetOcspServersOverride sets OcspServersOverride field to given value.

### HasOcspServersOverride

`func (o *CertCertsRequest) HasOcspServersOverride() bool`

HasOcspServersOverride returns a boolean if a field has been set.

### GetPeriod

`func (o *CertCertsRequest) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CertCertsRequest) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CertCertsRequest) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CertCertsRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPolicies

`func (o *CertCertsRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CertCertsRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CertCertsRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *CertCertsRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetRequiredExtensions

`func (o *CertCertsRequest) GetRequiredExtensions() []string`

GetRequiredExtensions returns the RequiredExtensions field if non-nil, zero value otherwise.

### GetRequiredExtensionsOk

`func (o *CertCertsRequest) GetRequiredExtensionsOk() (*[]string, bool)`

GetRequiredExtensionsOk returns a tuple with the RequiredExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredExtensions

`func (o *CertCertsRequest) SetRequiredExtensions(v []string)`

SetRequiredExtensions sets RequiredExtensions field to given value.

### HasRequiredExtensions

`func (o *CertCertsRequest) HasRequiredExtensions() bool`

HasRequiredExtensions returns a boolean if a field has been set.

### GetTokenBoundCidrs

`func (o *CertCertsRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *CertCertsRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *CertCertsRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.

### HasTokenBoundCidrs

`func (o *CertCertsRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.

### GetTokenExplicitMaxTtl

`func (o *CertCertsRequest) GetTokenExplicitMaxTtl() int32`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *CertCertsRequest) GetTokenExplicitMaxTtlOk() (*int32, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *CertCertsRequest) SetTokenExplicitMaxTtl(v int32)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.

### HasTokenExplicitMaxTtl

`func (o *CertCertsRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.

### GetTokenMaxTtl

`func (o *CertCertsRequest) GetTokenMaxTtl() int32`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *CertCertsRequest) GetTokenMaxTtlOk() (*int32, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *CertCertsRequest) SetTokenMaxTtl(v int32)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.

### HasTokenMaxTtl

`func (o *CertCertsRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.

### GetTokenNoDefaultPolicy

`func (o *CertCertsRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *CertCertsRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *CertCertsRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.

### HasTokenNoDefaultPolicy

`func (o *CertCertsRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.

### GetTokenNumUses

`func (o *CertCertsRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *CertCertsRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *CertCertsRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.

### HasTokenNumUses

`func (o *CertCertsRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.

### GetTokenPeriod

`func (o *CertCertsRequest) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *CertCertsRequest) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *CertCertsRequest) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.

### HasTokenPeriod

`func (o *CertCertsRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.

### GetTokenPolicies

`func (o *CertCertsRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *CertCertsRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *CertCertsRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.

### HasTokenPolicies

`func (o *CertCertsRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.

### GetTokenTtl

`func (o *CertCertsRequest) GetTokenTtl() int32`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *CertCertsRequest) GetTokenTtlOk() (*int32, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *CertCertsRequest) SetTokenTtl(v int32)`

SetTokenTtl sets TokenTtl field to given value.

### HasTokenTtl

`func (o *CertCertsRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.

### GetTokenType

`func (o *CertCertsRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *CertCertsRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *CertCertsRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *CertCertsRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetTtl

`func (o *CertCertsRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CertCertsRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CertCertsRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CertCertsRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


