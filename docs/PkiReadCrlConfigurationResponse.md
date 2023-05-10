# PkiReadCrlConfigurationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRebuild** | Pointer to **bool** | If set to true, enables automatic rebuilding of the CRL | [optional] 
**AutoRebuildGracePeriod** | Pointer to **string** | The time before the CRL expires to automatically rebuild it, when enabled. Must be shorter than the CRL expiry. Defaults to 12h. | [optional] 
**CrossClusterRevocation** | Pointer to **bool** | Whether to enable a global, cross-cluster revocation queue. Must be used with auto_rebuild&#x3D;true. | [optional] 
**DeltaRebuildInterval** | Pointer to **string** | The time between delta CRL rebuilds if a new revocation has occurred. Must be shorter than the CRL expiry. Defaults to 15m. | [optional] 
**Disable** | Pointer to **bool** | If set to true, disables generating the CRL entirely. | [optional] 
**EnableDelta** | Pointer to **bool** | Whether to enable delta CRLs between authoritative CRL rebuilds | [optional] 
**Expiry** | Pointer to **string** | The amount of time the generated CRL should be valid; defaults to 72 hours | [optional] 
**OcspDisable** | Pointer to **bool** | If set to true, ocsp unauthorized responses will be returned. | [optional] 
**OcspExpiry** | Pointer to **string** | The amount of time an OCSP response will be valid (controls the NextUpdate field); defaults to 12 hours | [optional] 
**UnifiedCrl** | Pointer to **bool** | If set to true enables global replication of revocation entries, also enabling unified versions of OCSP and CRLs if their respective features are enabled. disable for CRLs and ocsp_disable for OCSP. | [optional] 
**UnifiedCrlOnExistingPaths** | Pointer to **bool** | If set to true, existing CRL and OCSP paths will return the unified CRL instead of a response based on cluster-local data | [optional] 



## Methods


### NewPkiReadCrlConfigurationResponse

`func NewPkiReadCrlConfigurationResponse() *PkiReadCrlConfigurationResponse`

NewPkiReadCrlConfigurationResponse instantiates a new PkiReadCrlConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiReadCrlConfigurationResponseWithDefaults

`func NewPkiReadCrlConfigurationResponseWithDefaults() *PkiReadCrlConfigurationResponse`

NewPkiReadCrlConfigurationResponseWithDefaults instantiates a new PkiReadCrlConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAutoRebuild

`func (o *PkiReadCrlConfigurationResponse) GetAutoRebuild() bool`

GetAutoRebuild returns the AutoRebuild field if non-nil, zero value otherwise.

### GetAutoRebuildOk

`func (o *PkiReadCrlConfigurationResponse) GetAutoRebuildOk() (*bool, bool)`

GetAutoRebuildOk returns a tuple with the AutoRebuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRebuild

`func (o *PkiReadCrlConfigurationResponse) SetAutoRebuild(v bool)`

SetAutoRebuild sets AutoRebuild field to given value.


### HasAutoRebuild

`func (o *PkiReadCrlConfigurationResponse) HasAutoRebuild() bool`

HasAutoRebuild returns a boolean if a field has been set.




### GetAutoRebuildGracePeriod

`func (o *PkiReadCrlConfigurationResponse) GetAutoRebuildGracePeriod() string`

GetAutoRebuildGracePeriod returns the AutoRebuildGracePeriod field if non-nil, zero value otherwise.

### GetAutoRebuildGracePeriodOk

`func (o *PkiReadCrlConfigurationResponse) GetAutoRebuildGracePeriodOk() (*string, bool)`

GetAutoRebuildGracePeriodOk returns a tuple with the AutoRebuildGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRebuildGracePeriod

`func (o *PkiReadCrlConfigurationResponse) SetAutoRebuildGracePeriod(v string)`

SetAutoRebuildGracePeriod sets AutoRebuildGracePeriod field to given value.


### HasAutoRebuildGracePeriod

`func (o *PkiReadCrlConfigurationResponse) HasAutoRebuildGracePeriod() bool`

HasAutoRebuildGracePeriod returns a boolean if a field has been set.




### GetCrossClusterRevocation

`func (o *PkiReadCrlConfigurationResponse) GetCrossClusterRevocation() bool`

GetCrossClusterRevocation returns the CrossClusterRevocation field if non-nil, zero value otherwise.

### GetCrossClusterRevocationOk

`func (o *PkiReadCrlConfigurationResponse) GetCrossClusterRevocationOk() (*bool, bool)`

GetCrossClusterRevocationOk returns a tuple with the CrossClusterRevocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossClusterRevocation

`func (o *PkiReadCrlConfigurationResponse) SetCrossClusterRevocation(v bool)`

SetCrossClusterRevocation sets CrossClusterRevocation field to given value.


### HasCrossClusterRevocation

`func (o *PkiReadCrlConfigurationResponse) HasCrossClusterRevocation() bool`

HasCrossClusterRevocation returns a boolean if a field has been set.




### GetDeltaRebuildInterval

`func (o *PkiReadCrlConfigurationResponse) GetDeltaRebuildInterval() string`

GetDeltaRebuildInterval returns the DeltaRebuildInterval field if non-nil, zero value otherwise.

### GetDeltaRebuildIntervalOk

`func (o *PkiReadCrlConfigurationResponse) GetDeltaRebuildIntervalOk() (*string, bool)`

GetDeltaRebuildIntervalOk returns a tuple with the DeltaRebuildInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaRebuildInterval

`func (o *PkiReadCrlConfigurationResponse) SetDeltaRebuildInterval(v string)`

SetDeltaRebuildInterval sets DeltaRebuildInterval field to given value.


### HasDeltaRebuildInterval

`func (o *PkiReadCrlConfigurationResponse) HasDeltaRebuildInterval() bool`

HasDeltaRebuildInterval returns a boolean if a field has been set.




### GetDisable

`func (o *PkiReadCrlConfigurationResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *PkiReadCrlConfigurationResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *PkiReadCrlConfigurationResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.


### HasDisable

`func (o *PkiReadCrlConfigurationResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.




### GetEnableDelta

`func (o *PkiReadCrlConfigurationResponse) GetEnableDelta() bool`

GetEnableDelta returns the EnableDelta field if non-nil, zero value otherwise.

### GetEnableDeltaOk

`func (o *PkiReadCrlConfigurationResponse) GetEnableDeltaOk() (*bool, bool)`

GetEnableDeltaOk returns a tuple with the EnableDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDelta

`func (o *PkiReadCrlConfigurationResponse) SetEnableDelta(v bool)`

SetEnableDelta sets EnableDelta field to given value.


### HasEnableDelta

`func (o *PkiReadCrlConfigurationResponse) HasEnableDelta() bool`

HasEnableDelta returns a boolean if a field has been set.




### GetExpiry

`func (o *PkiReadCrlConfigurationResponse) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *PkiReadCrlConfigurationResponse) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *PkiReadCrlConfigurationResponse) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.


### HasExpiry

`func (o *PkiReadCrlConfigurationResponse) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.




### GetOcspDisable

`func (o *PkiReadCrlConfigurationResponse) GetOcspDisable() bool`

GetOcspDisable returns the OcspDisable field if non-nil, zero value otherwise.

### GetOcspDisableOk

`func (o *PkiReadCrlConfigurationResponse) GetOcspDisableOk() (*bool, bool)`

GetOcspDisableOk returns a tuple with the OcspDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspDisable

`func (o *PkiReadCrlConfigurationResponse) SetOcspDisable(v bool)`

SetOcspDisable sets OcspDisable field to given value.


### HasOcspDisable

`func (o *PkiReadCrlConfigurationResponse) HasOcspDisable() bool`

HasOcspDisable returns a boolean if a field has been set.




### GetOcspExpiry

`func (o *PkiReadCrlConfigurationResponse) GetOcspExpiry() string`

GetOcspExpiry returns the OcspExpiry field if non-nil, zero value otherwise.

### GetOcspExpiryOk

`func (o *PkiReadCrlConfigurationResponse) GetOcspExpiryOk() (*string, bool)`

GetOcspExpiryOk returns a tuple with the OcspExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspExpiry

`func (o *PkiReadCrlConfigurationResponse) SetOcspExpiry(v string)`

SetOcspExpiry sets OcspExpiry field to given value.


### HasOcspExpiry

`func (o *PkiReadCrlConfigurationResponse) HasOcspExpiry() bool`

HasOcspExpiry returns a boolean if a field has been set.




### GetUnifiedCrl

`func (o *PkiReadCrlConfigurationResponse) GetUnifiedCrl() bool`

GetUnifiedCrl returns the UnifiedCrl field if non-nil, zero value otherwise.

### GetUnifiedCrlOk

`func (o *PkiReadCrlConfigurationResponse) GetUnifiedCrlOk() (*bool, bool)`

GetUnifiedCrlOk returns a tuple with the UnifiedCrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedCrl

`func (o *PkiReadCrlConfigurationResponse) SetUnifiedCrl(v bool)`

SetUnifiedCrl sets UnifiedCrl field to given value.


### HasUnifiedCrl

`func (o *PkiReadCrlConfigurationResponse) HasUnifiedCrl() bool`

HasUnifiedCrl returns a boolean if a field has been set.




### GetUnifiedCrlOnExistingPaths

`func (o *PkiReadCrlConfigurationResponse) GetUnifiedCrlOnExistingPaths() bool`

GetUnifiedCrlOnExistingPaths returns the UnifiedCrlOnExistingPaths field if non-nil, zero value otherwise.

### GetUnifiedCrlOnExistingPathsOk

`func (o *PkiReadCrlConfigurationResponse) GetUnifiedCrlOnExistingPathsOk() (*bool, bool)`

GetUnifiedCrlOnExistingPathsOk returns a tuple with the UnifiedCrlOnExistingPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedCrlOnExistingPaths

`func (o *PkiReadCrlConfigurationResponse) SetUnifiedCrlOnExistingPaths(v bool)`

SetUnifiedCrlOnExistingPaths sets UnifiedCrlOnExistingPaths field to given value.


### HasUnifiedCrlOnExistingPaths

`func (o *PkiReadCrlConfigurationResponse) HasUnifiedCrlOnExistingPaths() bool`

HasUnifiedCrlOnExistingPaths returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


