# MountsReadTuningInformationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedManagedKeys** | Pointer to **[]string** |  | [optional] 
**AllowedResponseHeaders** | Pointer to **[]string** | A list of headers to whitelist and allow a plugin to set on responses. | [optional] 
**AuditNonHmacRequestKeys** | Pointer to **[]string** |  | [optional] 
**AuditNonHmacResponseKeys** | Pointer to **[]string** |  | [optional] 
**DefaultLeaseTtl** | Pointer to **int32** | The default lease TTL for this mount. | [optional] 
**Description** | Pointer to **string** | User-friendly description for this credential backend. | [optional] 
**ExternalEntropyAccess** | Pointer to **bool** |  | [optional] 
**ForceNoCache** | Pointer to **bool** |  | [optional] 
**ListingVisibility** | Pointer to **string** |  | [optional] 
**MaxLeaseTtl** | Pointer to **int32** | The max lease TTL for this mount. | [optional] 
**Options** | Pointer to **map[string]interface{}** | The options to pass into the backend. Should be a json object with string keys and values. | [optional] 
**PassthroughRequestHeaders** | Pointer to **[]string** |  | [optional] 
**PluginVersion** | Pointer to **string** | The semantic version of the plugin to use. | [optional] 
**TokenType** | Pointer to **string** | The type of token to issue (service or batch). | [optional] 
**UserLockoutCounterResetDuration** | Pointer to **int64** |  | [optional] 
**UserLockoutDisable** | Pointer to **bool** |  | [optional] 
**UserLockoutDuration** | Pointer to **int64** |  | [optional] 
**UserLockoutThreshold** | Pointer to **int64** |  | [optional] 



## Methods


### NewMountsReadTuningInformationResponse

`func NewMountsReadTuningInformationResponse() *MountsReadTuningInformationResponse`

NewMountsReadTuningInformationResponse instantiates a new MountsReadTuningInformationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountsReadTuningInformationResponseWithDefaults

`func NewMountsReadTuningInformationResponseWithDefaults() *MountsReadTuningInformationResponse`

NewMountsReadTuningInformationResponseWithDefaults instantiates a new MountsReadTuningInformationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAllowedManagedKeys

`func (o *MountsReadTuningInformationResponse) GetAllowedManagedKeys() []string`

GetAllowedManagedKeys returns the AllowedManagedKeys field if non-nil, zero value otherwise.

### GetAllowedManagedKeysOk

`func (o *MountsReadTuningInformationResponse) GetAllowedManagedKeysOk() (*[]string, bool)`

GetAllowedManagedKeysOk returns a tuple with the AllowedManagedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedManagedKeys

`func (o *MountsReadTuningInformationResponse) SetAllowedManagedKeys(v []string)`

SetAllowedManagedKeys sets AllowedManagedKeys field to given value.


### HasAllowedManagedKeys

`func (o *MountsReadTuningInformationResponse) HasAllowedManagedKeys() bool`

HasAllowedManagedKeys returns a boolean if a field has been set.




### GetAllowedResponseHeaders

`func (o *MountsReadTuningInformationResponse) GetAllowedResponseHeaders() []string`

GetAllowedResponseHeaders returns the AllowedResponseHeaders field if non-nil, zero value otherwise.

### GetAllowedResponseHeadersOk

`func (o *MountsReadTuningInformationResponse) GetAllowedResponseHeadersOk() (*[]string, bool)`

GetAllowedResponseHeadersOk returns a tuple with the AllowedResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedResponseHeaders

`func (o *MountsReadTuningInformationResponse) SetAllowedResponseHeaders(v []string)`

SetAllowedResponseHeaders sets AllowedResponseHeaders field to given value.


### HasAllowedResponseHeaders

`func (o *MountsReadTuningInformationResponse) HasAllowedResponseHeaders() bool`

HasAllowedResponseHeaders returns a boolean if a field has been set.




### GetAuditNonHmacRequestKeys

`func (o *MountsReadTuningInformationResponse) GetAuditNonHmacRequestKeys() []string`

GetAuditNonHmacRequestKeys returns the AuditNonHmacRequestKeys field if non-nil, zero value otherwise.

### GetAuditNonHmacRequestKeysOk

`func (o *MountsReadTuningInformationResponse) GetAuditNonHmacRequestKeysOk() (*[]string, bool)`

GetAuditNonHmacRequestKeysOk returns a tuple with the AuditNonHmacRequestKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditNonHmacRequestKeys

`func (o *MountsReadTuningInformationResponse) SetAuditNonHmacRequestKeys(v []string)`

SetAuditNonHmacRequestKeys sets AuditNonHmacRequestKeys field to given value.


### HasAuditNonHmacRequestKeys

`func (o *MountsReadTuningInformationResponse) HasAuditNonHmacRequestKeys() bool`

HasAuditNonHmacRequestKeys returns a boolean if a field has been set.




### GetAuditNonHmacResponseKeys

`func (o *MountsReadTuningInformationResponse) GetAuditNonHmacResponseKeys() []string`

GetAuditNonHmacResponseKeys returns the AuditNonHmacResponseKeys field if non-nil, zero value otherwise.

### GetAuditNonHmacResponseKeysOk

`func (o *MountsReadTuningInformationResponse) GetAuditNonHmacResponseKeysOk() (*[]string, bool)`

GetAuditNonHmacResponseKeysOk returns a tuple with the AuditNonHmacResponseKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditNonHmacResponseKeys

`func (o *MountsReadTuningInformationResponse) SetAuditNonHmacResponseKeys(v []string)`

SetAuditNonHmacResponseKeys sets AuditNonHmacResponseKeys field to given value.


### HasAuditNonHmacResponseKeys

`func (o *MountsReadTuningInformationResponse) HasAuditNonHmacResponseKeys() bool`

HasAuditNonHmacResponseKeys returns a boolean if a field has been set.




### GetDefaultLeaseTtl

`func (o *MountsReadTuningInformationResponse) GetDefaultLeaseTtl() int32`

GetDefaultLeaseTtl returns the DefaultLeaseTtl field if non-nil, zero value otherwise.

### GetDefaultLeaseTtlOk

`func (o *MountsReadTuningInformationResponse) GetDefaultLeaseTtlOk() (*int32, bool)`

GetDefaultLeaseTtlOk returns a tuple with the DefaultLeaseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLeaseTtl

`func (o *MountsReadTuningInformationResponse) SetDefaultLeaseTtl(v int32)`

SetDefaultLeaseTtl sets DefaultLeaseTtl field to given value.


### HasDefaultLeaseTtl

`func (o *MountsReadTuningInformationResponse) HasDefaultLeaseTtl() bool`

HasDefaultLeaseTtl returns a boolean if a field has been set.




### GetDescription

`func (o *MountsReadTuningInformationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MountsReadTuningInformationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MountsReadTuningInformationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### HasDescription

`func (o *MountsReadTuningInformationResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.




### GetExternalEntropyAccess

`func (o *MountsReadTuningInformationResponse) GetExternalEntropyAccess() bool`

GetExternalEntropyAccess returns the ExternalEntropyAccess field if non-nil, zero value otherwise.

### GetExternalEntropyAccessOk

`func (o *MountsReadTuningInformationResponse) GetExternalEntropyAccessOk() (*bool, bool)`

GetExternalEntropyAccessOk returns a tuple with the ExternalEntropyAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEntropyAccess

`func (o *MountsReadTuningInformationResponse) SetExternalEntropyAccess(v bool)`

SetExternalEntropyAccess sets ExternalEntropyAccess field to given value.


### HasExternalEntropyAccess

`func (o *MountsReadTuningInformationResponse) HasExternalEntropyAccess() bool`

HasExternalEntropyAccess returns a boolean if a field has been set.




### GetForceNoCache

`func (o *MountsReadTuningInformationResponse) GetForceNoCache() bool`

GetForceNoCache returns the ForceNoCache field if non-nil, zero value otherwise.

### GetForceNoCacheOk

`func (o *MountsReadTuningInformationResponse) GetForceNoCacheOk() (*bool, bool)`

GetForceNoCacheOk returns a tuple with the ForceNoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceNoCache

`func (o *MountsReadTuningInformationResponse) SetForceNoCache(v bool)`

SetForceNoCache sets ForceNoCache field to given value.


### HasForceNoCache

`func (o *MountsReadTuningInformationResponse) HasForceNoCache() bool`

HasForceNoCache returns a boolean if a field has been set.




### GetListingVisibility

`func (o *MountsReadTuningInformationResponse) GetListingVisibility() string`

GetListingVisibility returns the ListingVisibility field if non-nil, zero value otherwise.

### GetListingVisibilityOk

`func (o *MountsReadTuningInformationResponse) GetListingVisibilityOk() (*string, bool)`

GetListingVisibilityOk returns a tuple with the ListingVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingVisibility

`func (o *MountsReadTuningInformationResponse) SetListingVisibility(v string)`

SetListingVisibility sets ListingVisibility field to given value.


### HasListingVisibility

`func (o *MountsReadTuningInformationResponse) HasListingVisibility() bool`

HasListingVisibility returns a boolean if a field has been set.




### GetMaxLeaseTtl

`func (o *MountsReadTuningInformationResponse) GetMaxLeaseTtl() int32`

GetMaxLeaseTtl returns the MaxLeaseTtl field if non-nil, zero value otherwise.

### GetMaxLeaseTtlOk

`func (o *MountsReadTuningInformationResponse) GetMaxLeaseTtlOk() (*int32, bool)`

GetMaxLeaseTtlOk returns a tuple with the MaxLeaseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLeaseTtl

`func (o *MountsReadTuningInformationResponse) SetMaxLeaseTtl(v int32)`

SetMaxLeaseTtl sets MaxLeaseTtl field to given value.


### HasMaxLeaseTtl

`func (o *MountsReadTuningInformationResponse) HasMaxLeaseTtl() bool`

HasMaxLeaseTtl returns a boolean if a field has been set.




### GetOptions

`func (o *MountsReadTuningInformationResponse) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MountsReadTuningInformationResponse) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MountsReadTuningInformationResponse) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### HasOptions

`func (o *MountsReadTuningInformationResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.




### GetPassthroughRequestHeaders

`func (o *MountsReadTuningInformationResponse) GetPassthroughRequestHeaders() []string`

GetPassthroughRequestHeaders returns the PassthroughRequestHeaders field if non-nil, zero value otherwise.

### GetPassthroughRequestHeadersOk

`func (o *MountsReadTuningInformationResponse) GetPassthroughRequestHeadersOk() (*[]string, bool)`

GetPassthroughRequestHeadersOk returns a tuple with the PassthroughRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthroughRequestHeaders

`func (o *MountsReadTuningInformationResponse) SetPassthroughRequestHeaders(v []string)`

SetPassthroughRequestHeaders sets PassthroughRequestHeaders field to given value.


### HasPassthroughRequestHeaders

`func (o *MountsReadTuningInformationResponse) HasPassthroughRequestHeaders() bool`

HasPassthroughRequestHeaders returns a boolean if a field has been set.




### GetPluginVersion

`func (o *MountsReadTuningInformationResponse) GetPluginVersion() string`

GetPluginVersion returns the PluginVersion field if non-nil, zero value otherwise.

### GetPluginVersionOk

`func (o *MountsReadTuningInformationResponse) GetPluginVersionOk() (*string, bool)`

GetPluginVersionOk returns a tuple with the PluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginVersion

`func (o *MountsReadTuningInformationResponse) SetPluginVersion(v string)`

SetPluginVersion sets PluginVersion field to given value.


### HasPluginVersion

`func (o *MountsReadTuningInformationResponse) HasPluginVersion() bool`

HasPluginVersion returns a boolean if a field has been set.




### GetTokenType

`func (o *MountsReadTuningInformationResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *MountsReadTuningInformationResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *MountsReadTuningInformationResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### HasTokenType

`func (o *MountsReadTuningInformationResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.




### GetUserLockoutCounterResetDuration

`func (o *MountsReadTuningInformationResponse) GetUserLockoutCounterResetDuration() int64`

GetUserLockoutCounterResetDuration returns the UserLockoutCounterResetDuration field if non-nil, zero value otherwise.

### GetUserLockoutCounterResetDurationOk

`func (o *MountsReadTuningInformationResponse) GetUserLockoutCounterResetDurationOk() (*int64, bool)`

GetUserLockoutCounterResetDurationOk returns a tuple with the UserLockoutCounterResetDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutCounterResetDuration

`func (o *MountsReadTuningInformationResponse) SetUserLockoutCounterResetDuration(v int64)`

SetUserLockoutCounterResetDuration sets UserLockoutCounterResetDuration field to given value.


### HasUserLockoutCounterResetDuration

`func (o *MountsReadTuningInformationResponse) HasUserLockoutCounterResetDuration() bool`

HasUserLockoutCounterResetDuration returns a boolean if a field has been set.




### GetUserLockoutDisable

`func (o *MountsReadTuningInformationResponse) GetUserLockoutDisable() bool`

GetUserLockoutDisable returns the UserLockoutDisable field if non-nil, zero value otherwise.

### GetUserLockoutDisableOk

`func (o *MountsReadTuningInformationResponse) GetUserLockoutDisableOk() (*bool, bool)`

GetUserLockoutDisableOk returns a tuple with the UserLockoutDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutDisable

`func (o *MountsReadTuningInformationResponse) SetUserLockoutDisable(v bool)`

SetUserLockoutDisable sets UserLockoutDisable field to given value.


### HasUserLockoutDisable

`func (o *MountsReadTuningInformationResponse) HasUserLockoutDisable() bool`

HasUserLockoutDisable returns a boolean if a field has been set.




### GetUserLockoutDuration

`func (o *MountsReadTuningInformationResponse) GetUserLockoutDuration() int64`

GetUserLockoutDuration returns the UserLockoutDuration field if non-nil, zero value otherwise.

### GetUserLockoutDurationOk

`func (o *MountsReadTuningInformationResponse) GetUserLockoutDurationOk() (*int64, bool)`

GetUserLockoutDurationOk returns a tuple with the UserLockoutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutDuration

`func (o *MountsReadTuningInformationResponse) SetUserLockoutDuration(v int64)`

SetUserLockoutDuration sets UserLockoutDuration field to given value.


### HasUserLockoutDuration

`func (o *MountsReadTuningInformationResponse) HasUserLockoutDuration() bool`

HasUserLockoutDuration returns a boolean if a field has been set.




### GetUserLockoutThreshold

`func (o *MountsReadTuningInformationResponse) GetUserLockoutThreshold() int64`

GetUserLockoutThreshold returns the UserLockoutThreshold field if non-nil, zero value otherwise.

### GetUserLockoutThresholdOk

`func (o *MountsReadTuningInformationResponse) GetUserLockoutThresholdOk() (*int64, bool)`

GetUserLockoutThresholdOk returns a tuple with the UserLockoutThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutThreshold

`func (o *MountsReadTuningInformationResponse) SetUserLockoutThreshold(v int64)`

SetUserLockoutThreshold sets UserLockoutThreshold field to given value.


### HasUserLockoutThreshold

`func (o *MountsReadTuningInformationResponse) HasUserLockoutThreshold() bool`

HasUserLockoutThreshold returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


