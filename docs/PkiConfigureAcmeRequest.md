# PkiConfigureAcmeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedIssuers** | Pointer to **[]string** | which issuers are allowed for use with ACME; by default, this will only be the primary (default) issuer | [optional] [default to ["*"]]
**AllowedRoles** | Pointer to **[]string** | which roles are allowed for use with ACME; by default via &#x27;*&#x27;, these will be all roles including sign-verbatim; when concrete role names are specified, any default_directory_policy role must be included to allow usage of the default acme directories under /pki/acme/directory and /pki/issuer/:issuer_id/acme/directory. | [optional] [default to ["*"]]
**DefaultDirectoryPolicy** | Pointer to **string** | the policy to be used for non-role-qualified ACME requests; by default ACME issuance will be otherwise unrestricted, equivalent to the sign-verbatim endpoint; one may also specify a role to use as this policy, as \&quot;role:&lt;role_name&gt;\&quot;, the specified role must be allowed by allowed_roles | [optional] [default to "sign-verbatim"]
**DnsResolver** | Pointer to **string** | DNS resolver to use for domain resolution on this mount. Defaults to using the default system resolver. Must be in the format &lt;host&gt;:&lt;port&gt;, with both parts mandatory. | [optional] [default to ""]
**EabPolicy** | Pointer to **string** | Specify the policy to use for external account binding behaviour, &#x27;not-required&#x27;, &#x27;new-account-required&#x27; or &#x27;always-required&#x27; | [optional] [default to "always-required"]
**Enabled** | Pointer to **bool** | whether ACME is enabled, defaults to false meaning that clusters will by default not get ACME support | [optional] [default to false]



## Methods


### NewPkiConfigureAcmeRequest

`func NewPkiConfigureAcmeRequest() *PkiConfigureAcmeRequest`

NewPkiConfigureAcmeRequest instantiates a new PkiConfigureAcmeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiConfigureAcmeRequestWithDefaults

`func NewPkiConfigureAcmeRequestWithDefaults() *PkiConfigureAcmeRequest`

NewPkiConfigureAcmeRequestWithDefaults instantiates a new PkiConfigureAcmeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAllowedIssuers

`func (o *PkiConfigureAcmeRequest) GetAllowedIssuers() []string`

GetAllowedIssuers returns the AllowedIssuers field if non-nil, zero value otherwise.

### GetAllowedIssuersOk

`func (o *PkiConfigureAcmeRequest) GetAllowedIssuersOk() (*[]string, bool)`

GetAllowedIssuersOk returns a tuple with the AllowedIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIssuers

`func (o *PkiConfigureAcmeRequest) SetAllowedIssuers(v []string)`

SetAllowedIssuers sets AllowedIssuers field to given value.


### HasAllowedIssuers

`func (o *PkiConfigureAcmeRequest) HasAllowedIssuers() bool`

HasAllowedIssuers returns a boolean if a field has been set.




### GetAllowedRoles

`func (o *PkiConfigureAcmeRequest) GetAllowedRoles() []string`

GetAllowedRoles returns the AllowedRoles field if non-nil, zero value otherwise.

### GetAllowedRolesOk

`func (o *PkiConfigureAcmeRequest) GetAllowedRolesOk() (*[]string, bool)`

GetAllowedRolesOk returns a tuple with the AllowedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRoles

`func (o *PkiConfigureAcmeRequest) SetAllowedRoles(v []string)`

SetAllowedRoles sets AllowedRoles field to given value.


### HasAllowedRoles

`func (o *PkiConfigureAcmeRequest) HasAllowedRoles() bool`

HasAllowedRoles returns a boolean if a field has been set.




### GetDefaultDirectoryPolicy

`func (o *PkiConfigureAcmeRequest) GetDefaultDirectoryPolicy() string`

GetDefaultDirectoryPolicy returns the DefaultDirectoryPolicy field if non-nil, zero value otherwise.

### GetDefaultDirectoryPolicyOk

`func (o *PkiConfigureAcmeRequest) GetDefaultDirectoryPolicyOk() (*string, bool)`

GetDefaultDirectoryPolicyOk returns a tuple with the DefaultDirectoryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDirectoryPolicy

`func (o *PkiConfigureAcmeRequest) SetDefaultDirectoryPolicy(v string)`

SetDefaultDirectoryPolicy sets DefaultDirectoryPolicy field to given value.


### HasDefaultDirectoryPolicy

`func (o *PkiConfigureAcmeRequest) HasDefaultDirectoryPolicy() bool`

HasDefaultDirectoryPolicy returns a boolean if a field has been set.




### GetDnsResolver

`func (o *PkiConfigureAcmeRequest) GetDnsResolver() string`

GetDnsResolver returns the DnsResolver field if non-nil, zero value otherwise.

### GetDnsResolverOk

`func (o *PkiConfigureAcmeRequest) GetDnsResolverOk() (*string, bool)`

GetDnsResolverOk returns a tuple with the DnsResolver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsResolver

`func (o *PkiConfigureAcmeRequest) SetDnsResolver(v string)`

SetDnsResolver sets DnsResolver field to given value.


### HasDnsResolver

`func (o *PkiConfigureAcmeRequest) HasDnsResolver() bool`

HasDnsResolver returns a boolean if a field has been set.




### GetEabPolicy

`func (o *PkiConfigureAcmeRequest) GetEabPolicy() string`

GetEabPolicy returns the EabPolicy field if non-nil, zero value otherwise.

### GetEabPolicyOk

`func (o *PkiConfigureAcmeRequest) GetEabPolicyOk() (*string, bool)`

GetEabPolicyOk returns a tuple with the EabPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabPolicy

`func (o *PkiConfigureAcmeRequest) SetEabPolicy(v string)`

SetEabPolicy sets EabPolicy field to given value.


### HasEabPolicy

`func (o *PkiConfigureAcmeRequest) HasEabPolicy() bool`

HasEabPolicy returns a boolean if a field has been set.




### GetEnabled

`func (o *PkiConfigureAcmeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PkiConfigureAcmeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PkiConfigureAcmeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### HasEnabled

`func (o *PkiConfigureAcmeRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


