# AuthReadTuningInformationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedManagedKeys** | Pointer to **[]string** |  | [optional] 
**AllowedResponseHeaders** | Pointer to **[]string** |  | [optional] 
**AuditNonHmacRequestKeys** | Pointer to **[]string** |  | [optional] 
**AuditNonHmacResponseKeys** | Pointer to **[]string** |  | [optional] 
**DefaultLeaseTtl** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalEntropyAccess** | Pointer to **bool** |  | [optional] 
**ForceNoCache** | Pointer to **bool** |  | [optional] 
**ListingVisibility** | Pointer to **string** |  | [optional] 
**MaxLeaseTtl** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 
**PassthroughRequestHeaders** | Pointer to **[]string** |  | [optional] 
**PluginVersion** | Pointer to **string** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 
**UserLockoutCounterResetDuration** | Pointer to **int64** |  | [optional] 
**UserLockoutDisable** | Pointer to **bool** |  | [optional] 
**UserLockoutDuration** | Pointer to **int64** |  | [optional] 
**UserLockoutThreshold** | Pointer to **int64** |  | [optional] 



## Methods


### NewAuthReadTuningInformationResponse

`func NewAuthReadTuningInformationResponse() *AuthReadTuningInformationResponse`

NewAuthReadTuningInformationResponse instantiates a new AuthReadTuningInformationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthReadTuningInformationResponseWithDefaults

`func NewAuthReadTuningInformationResponseWithDefaults() *AuthReadTuningInformationResponse`

NewAuthReadTuningInformationResponseWithDefaults instantiates a new AuthReadTuningInformationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAllowedManagedKeys

`func (o *AuthReadTuningInformationResponse) GetAllowedManagedKeys() []string`

GetAllowedManagedKeys returns the AllowedManagedKeys field if non-nil, zero value otherwise.

### GetAllowedManagedKeysOk

`func (o *AuthReadTuningInformationResponse) GetAllowedManagedKeysOk() (*[]string, bool)`

GetAllowedManagedKeysOk returns a tuple with the AllowedManagedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedManagedKeys

`func (o *AuthReadTuningInformationResponse) SetAllowedManagedKeys(v []string)`

SetAllowedManagedKeys sets AllowedManagedKeys field to given value.


### HasAllowedManagedKeys

`func (o *AuthReadTuningInformationResponse) HasAllowedManagedKeys() bool`

HasAllowedManagedKeys returns a boolean if a field has been set.




### GetAllowedResponseHeaders

`func (o *AuthReadTuningInformationResponse) GetAllowedResponseHeaders() []string`

GetAllowedResponseHeaders returns the AllowedResponseHeaders field if non-nil, zero value otherwise.

### GetAllowedResponseHeadersOk

`func (o *AuthReadTuningInformationResponse) GetAllowedResponseHeadersOk() (*[]string, bool)`

GetAllowedResponseHeadersOk returns a tuple with the AllowedResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedResponseHeaders

`func (o *AuthReadTuningInformationResponse) SetAllowedResponseHeaders(v []string)`

SetAllowedResponseHeaders sets AllowedResponseHeaders field to given value.


### HasAllowedResponseHeaders

`func (o *AuthReadTuningInformationResponse) HasAllowedResponseHeaders() bool`

HasAllowedResponseHeaders returns a boolean if a field has been set.




### GetAuditNonHmacRequestKeys

`func (o *AuthReadTuningInformationResponse) GetAuditNonHmacRequestKeys() []string`

GetAuditNonHmacRequestKeys returns the AuditNonHmacRequestKeys field if non-nil, zero value otherwise.

### GetAuditNonHmacRequestKeysOk

`func (o *AuthReadTuningInformationResponse) GetAuditNonHmacRequestKeysOk() (*[]string, bool)`

GetAuditNonHmacRequestKeysOk returns a tuple with the AuditNonHmacRequestKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditNonHmacRequestKeys

`func (o *AuthReadTuningInformationResponse) SetAuditNonHmacRequestKeys(v []string)`

SetAuditNonHmacRequestKeys sets AuditNonHmacRequestKeys field to given value.


### HasAuditNonHmacRequestKeys

`func (o *AuthReadTuningInformationResponse) HasAuditNonHmacRequestKeys() bool`

HasAuditNonHmacRequestKeys returns a boolean if a field has been set.




### GetAuditNonHmacResponseKeys

`func (o *AuthReadTuningInformationResponse) GetAuditNonHmacResponseKeys() []string`

GetAuditNonHmacResponseKeys returns the AuditNonHmacResponseKeys field if non-nil, zero value otherwise.

### GetAuditNonHmacResponseKeysOk

`func (o *AuthReadTuningInformationResponse) GetAuditNonHmacResponseKeysOk() (*[]string, bool)`

GetAuditNonHmacResponseKeysOk returns a tuple with the AuditNonHmacResponseKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditNonHmacResponseKeys

`func (o *AuthReadTuningInformationResponse) SetAuditNonHmacResponseKeys(v []string)`

SetAuditNonHmacResponseKeys sets AuditNonHmacResponseKeys field to given value.


### HasAuditNonHmacResponseKeys

`func (o *AuthReadTuningInformationResponse) HasAuditNonHmacResponseKeys() bool`

HasAuditNonHmacResponseKeys returns a boolean if a field has been set.




### GetDefaultLeaseTtl

`func (o *AuthReadTuningInformationResponse) GetDefaultLeaseTtl() int32`

GetDefaultLeaseTtl returns the DefaultLeaseTtl field if non-nil, zero value otherwise.

### GetDefaultLeaseTtlOk

`func (o *AuthReadTuningInformationResponse) GetDefaultLeaseTtlOk() (*int32, bool)`

GetDefaultLeaseTtlOk returns a tuple with the DefaultLeaseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLeaseTtl

`func (o *AuthReadTuningInformationResponse) SetDefaultLeaseTtl(v int32)`

SetDefaultLeaseTtl sets DefaultLeaseTtl field to given value.


### HasDefaultLeaseTtl

`func (o *AuthReadTuningInformationResponse) HasDefaultLeaseTtl() bool`

HasDefaultLeaseTtl returns a boolean if a field has been set.




### GetDescription

`func (o *AuthReadTuningInformationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthReadTuningInformationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthReadTuningInformationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### HasDescription

`func (o *AuthReadTuningInformationResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.




### GetExternalEntropyAccess

`func (o *AuthReadTuningInformationResponse) GetExternalEntropyAccess() bool`

GetExternalEntropyAccess returns the ExternalEntropyAccess field if non-nil, zero value otherwise.

### GetExternalEntropyAccessOk

`func (o *AuthReadTuningInformationResponse) GetExternalEntropyAccessOk() (*bool, bool)`

GetExternalEntropyAccessOk returns a tuple with the ExternalEntropyAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEntropyAccess

`func (o *AuthReadTuningInformationResponse) SetExternalEntropyAccess(v bool)`

SetExternalEntropyAccess sets ExternalEntropyAccess field to given value.


### HasExternalEntropyAccess

`func (o *AuthReadTuningInformationResponse) HasExternalEntropyAccess() bool`

HasExternalEntropyAccess returns a boolean if a field has been set.




### GetForceNoCache

`func (o *AuthReadTuningInformationResponse) GetForceNoCache() bool`

GetForceNoCache returns the ForceNoCache field if non-nil, zero value otherwise.

### GetForceNoCacheOk

`func (o *AuthReadTuningInformationResponse) GetForceNoCacheOk() (*bool, bool)`

GetForceNoCacheOk returns a tuple with the ForceNoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceNoCache

`func (o *AuthReadTuningInformationResponse) SetForceNoCache(v bool)`

SetForceNoCache sets ForceNoCache field to given value.


### HasForceNoCache

`func (o *AuthReadTuningInformationResponse) HasForceNoCache() bool`

HasForceNoCache returns a boolean if a field has been set.




### GetListingVisibility

`func (o *AuthReadTuningInformationResponse) GetListingVisibility() string`

GetListingVisibility returns the ListingVisibility field if non-nil, zero value otherwise.

### GetListingVisibilityOk

`func (o *AuthReadTuningInformationResponse) GetListingVisibilityOk() (*string, bool)`

GetListingVisibilityOk returns a tuple with the ListingVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingVisibility

`func (o *AuthReadTuningInformationResponse) SetListingVisibility(v string)`

SetListingVisibility sets ListingVisibility field to given value.


### HasListingVisibility

`func (o *AuthReadTuningInformationResponse) HasListingVisibility() bool`

HasListingVisibility returns a boolean if a field has been set.




### GetMaxLeaseTtl

`func (o *AuthReadTuningInformationResponse) GetMaxLeaseTtl() int32`

GetMaxLeaseTtl returns the MaxLeaseTtl field if non-nil, zero value otherwise.

### GetMaxLeaseTtlOk

`func (o *AuthReadTuningInformationResponse) GetMaxLeaseTtlOk() (*int32, bool)`

GetMaxLeaseTtlOk returns a tuple with the MaxLeaseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLeaseTtl

`func (o *AuthReadTuningInformationResponse) SetMaxLeaseTtl(v int32)`

SetMaxLeaseTtl sets MaxLeaseTtl field to given value.


### HasMaxLeaseTtl

`func (o *AuthReadTuningInformationResponse) HasMaxLeaseTtl() bool`

HasMaxLeaseTtl returns a boolean if a field has been set.




### GetOptions

`func (o *AuthReadTuningInformationResponse) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AuthReadTuningInformationResponse) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AuthReadTuningInformationResponse) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### HasOptions

`func (o *AuthReadTuningInformationResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.




### GetPassthroughRequestHeaders

`func (o *AuthReadTuningInformationResponse) GetPassthroughRequestHeaders() []string`

GetPassthroughRequestHeaders returns the PassthroughRequestHeaders field if non-nil, zero value otherwise.

### GetPassthroughRequestHeadersOk

`func (o *AuthReadTuningInformationResponse) GetPassthroughRequestHeadersOk() (*[]string, bool)`

GetPassthroughRequestHeadersOk returns a tuple with the PassthroughRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthroughRequestHeaders

`func (o *AuthReadTuningInformationResponse) SetPassthroughRequestHeaders(v []string)`

SetPassthroughRequestHeaders sets PassthroughRequestHeaders field to given value.


### HasPassthroughRequestHeaders

`func (o *AuthReadTuningInformationResponse) HasPassthroughRequestHeaders() bool`

HasPassthroughRequestHeaders returns a boolean if a field has been set.




### GetPluginVersion

`func (o *AuthReadTuningInformationResponse) GetPluginVersion() string`

GetPluginVersion returns the PluginVersion field if non-nil, zero value otherwise.

### GetPluginVersionOk

`func (o *AuthReadTuningInformationResponse) GetPluginVersionOk() (*string, bool)`

GetPluginVersionOk returns a tuple with the PluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginVersion

`func (o *AuthReadTuningInformationResponse) SetPluginVersion(v string)`

SetPluginVersion sets PluginVersion field to given value.


### HasPluginVersion

`func (o *AuthReadTuningInformationResponse) HasPluginVersion() bool`

HasPluginVersion returns a boolean if a field has been set.




### GetTokenType

`func (o *AuthReadTuningInformationResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AuthReadTuningInformationResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AuthReadTuningInformationResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### HasTokenType

`func (o *AuthReadTuningInformationResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.




### GetUserLockoutCounterResetDuration

`func (o *AuthReadTuningInformationResponse) GetUserLockoutCounterResetDuration() int64`

GetUserLockoutCounterResetDuration returns the UserLockoutCounterResetDuration field if non-nil, zero value otherwise.

### GetUserLockoutCounterResetDurationOk

`func (o *AuthReadTuningInformationResponse) GetUserLockoutCounterResetDurationOk() (*int64, bool)`

GetUserLockoutCounterResetDurationOk returns a tuple with the UserLockoutCounterResetDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutCounterResetDuration

`func (o *AuthReadTuningInformationResponse) SetUserLockoutCounterResetDuration(v int64)`

SetUserLockoutCounterResetDuration sets UserLockoutCounterResetDuration field to given value.


### HasUserLockoutCounterResetDuration

`func (o *AuthReadTuningInformationResponse) HasUserLockoutCounterResetDuration() bool`

HasUserLockoutCounterResetDuration returns a boolean if a field has been set.




### GetUserLockoutDisable

`func (o *AuthReadTuningInformationResponse) GetUserLockoutDisable() bool`

GetUserLockoutDisable returns the UserLockoutDisable field if non-nil, zero value otherwise.

### GetUserLockoutDisableOk

`func (o *AuthReadTuningInformationResponse) GetUserLockoutDisableOk() (*bool, bool)`

GetUserLockoutDisableOk returns a tuple with the UserLockoutDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutDisable

`func (o *AuthReadTuningInformationResponse) SetUserLockoutDisable(v bool)`

SetUserLockoutDisable sets UserLockoutDisable field to given value.


### HasUserLockoutDisable

`func (o *AuthReadTuningInformationResponse) HasUserLockoutDisable() bool`

HasUserLockoutDisable returns a boolean if a field has been set.




### GetUserLockoutDuration

`func (o *AuthReadTuningInformationResponse) GetUserLockoutDuration() int64`

GetUserLockoutDuration returns the UserLockoutDuration field if non-nil, zero value otherwise.

### GetUserLockoutDurationOk

`func (o *AuthReadTuningInformationResponse) GetUserLockoutDurationOk() (*int64, bool)`

GetUserLockoutDurationOk returns a tuple with the UserLockoutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutDuration

`func (o *AuthReadTuningInformationResponse) SetUserLockoutDuration(v int64)`

SetUserLockoutDuration sets UserLockoutDuration field to given value.


### HasUserLockoutDuration

`func (o *AuthReadTuningInformationResponse) HasUserLockoutDuration() bool`

HasUserLockoutDuration returns a boolean if a field has been set.




### GetUserLockoutThreshold

`func (o *AuthReadTuningInformationResponse) GetUserLockoutThreshold() int64`

GetUserLockoutThreshold returns the UserLockoutThreshold field if non-nil, zero value otherwise.

### GetUserLockoutThresholdOk

`func (o *AuthReadTuningInformationResponse) GetUserLockoutThresholdOk() (*int64, bool)`

GetUserLockoutThresholdOk returns a tuple with the UserLockoutThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutThreshold

`func (o *AuthReadTuningInformationResponse) SetUserLockoutThreshold(v int64)`

SetUserLockoutThreshold sets UserLockoutThreshold field to given value.


### HasUserLockoutThreshold

`func (o *AuthReadTuningInformationResponse) HasUserLockoutThreshold() bool`

HasUserLockoutThreshold returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


