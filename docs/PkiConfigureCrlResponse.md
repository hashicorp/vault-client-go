# PkiConfigureCrlResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRebuild** | Pointer to **bool** | If set to true, enables automatic rebuilding of the CRL | [optional] 
**AutoRebuildGracePeriod** | Pointer to **string** | The time before the CRL expires to automatically rebuild it, when enabled. Must be shorter than the CRL expiry. Defaults to 12h. | [optional] [default to "12h"]
**CrossClusterRevocation** | Pointer to **bool** | Whether to enable a global, cross-cluster revocation queue. Must be used with auto_rebuild&#x3D;true. | [optional] 
**DeltaRebuildInterval** | Pointer to **string** | The time between delta CRL rebuilds if a new revocation has occurred. Must be shorter than the CRL expiry. Defaults to 15m. | [optional] [default to "15m"]
**Disable** | Pointer to **bool** | If set to true, disables generating the CRL entirely. | [optional] 
**EnableDelta** | Pointer to **bool** | Whether to enable delta CRLs between authoritative CRL rebuilds | [optional] 
**Expiry** | Pointer to **string** | The amount of time the generated CRL should be valid; defaults to 72 hours | [optional] [default to "72h"]
**OcspDisable** | Pointer to **bool** | If set to true, ocsp unauthorized responses will be returned. | [optional] 
**OcspExpiry** | Pointer to **string** | The amount of time an OCSP response will be valid (controls the NextUpdate field); defaults to 12 hours | [optional] [default to "1h"]
**UnifiedCrl** | Pointer to **bool** | If set to true enables global replication of revocation entries, also enabling unified versions of OCSP and CRLs if their respective features are enabled. disable for CRLs and ocsp_disable for OCSP. | [optional] 
**UnifiedCrlOnExistingPaths** | Pointer to **bool** | If set to true, existing CRL and OCSP paths will return the unified CRL instead of a response based on cluster-local data | [optional] 



## Methods


### NewPkiConfigureCrlResponse

`func NewPkiConfigureCrlResponse() *PkiConfigureCrlResponse`

NewPkiConfigureCrlResponse instantiates a new PkiConfigureCrlResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiConfigureCrlResponseWithDefaults

`func NewPkiConfigureCrlResponseWithDefaults() *PkiConfigureCrlResponse`

NewPkiConfigureCrlResponseWithDefaults instantiates a new PkiConfigureCrlResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAutoRebuild

`func (o *PkiConfigureCrlResponse) GetAutoRebuild() bool`

GetAutoRebuild returns the AutoRebuild field if non-nil, zero value otherwise.

### GetAutoRebuildOk

`func (o *PkiConfigureCrlResponse) GetAutoRebuildOk() (*bool, bool)`

GetAutoRebuildOk returns a tuple with the AutoRebuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRebuild

`func (o *PkiConfigureCrlResponse) SetAutoRebuild(v bool)`

SetAutoRebuild sets AutoRebuild field to given value.


### HasAutoRebuild

`func (o *PkiConfigureCrlResponse) HasAutoRebuild() bool`

HasAutoRebuild returns a boolean if a field has been set.




### GetAutoRebuildGracePeriod

`func (o *PkiConfigureCrlResponse) GetAutoRebuildGracePeriod() string`

GetAutoRebuildGracePeriod returns the AutoRebuildGracePeriod field if non-nil, zero value otherwise.

### GetAutoRebuildGracePeriodOk

`func (o *PkiConfigureCrlResponse) GetAutoRebuildGracePeriodOk() (*string, bool)`

GetAutoRebuildGracePeriodOk returns a tuple with the AutoRebuildGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRebuildGracePeriod

`func (o *PkiConfigureCrlResponse) SetAutoRebuildGracePeriod(v string)`

SetAutoRebuildGracePeriod sets AutoRebuildGracePeriod field to given value.


### HasAutoRebuildGracePeriod

`func (o *PkiConfigureCrlResponse) HasAutoRebuildGracePeriod() bool`

HasAutoRebuildGracePeriod returns a boolean if a field has been set.




### GetCrossClusterRevocation

`func (o *PkiConfigureCrlResponse) GetCrossClusterRevocation() bool`

GetCrossClusterRevocation returns the CrossClusterRevocation field if non-nil, zero value otherwise.

### GetCrossClusterRevocationOk

`func (o *PkiConfigureCrlResponse) GetCrossClusterRevocationOk() (*bool, bool)`

GetCrossClusterRevocationOk returns a tuple with the CrossClusterRevocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossClusterRevocation

`func (o *PkiConfigureCrlResponse) SetCrossClusterRevocation(v bool)`

SetCrossClusterRevocation sets CrossClusterRevocation field to given value.


### HasCrossClusterRevocation

`func (o *PkiConfigureCrlResponse) HasCrossClusterRevocation() bool`

HasCrossClusterRevocation returns a boolean if a field has been set.




### GetDeltaRebuildInterval

`func (o *PkiConfigureCrlResponse) GetDeltaRebuildInterval() string`

GetDeltaRebuildInterval returns the DeltaRebuildInterval field if non-nil, zero value otherwise.

### GetDeltaRebuildIntervalOk

`func (o *PkiConfigureCrlResponse) GetDeltaRebuildIntervalOk() (*string, bool)`

GetDeltaRebuildIntervalOk returns a tuple with the DeltaRebuildInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaRebuildInterval

`func (o *PkiConfigureCrlResponse) SetDeltaRebuildInterval(v string)`

SetDeltaRebuildInterval sets DeltaRebuildInterval field to given value.


### HasDeltaRebuildInterval

`func (o *PkiConfigureCrlResponse) HasDeltaRebuildInterval() bool`

HasDeltaRebuildInterval returns a boolean if a field has been set.




### GetDisable

`func (o *PkiConfigureCrlResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *PkiConfigureCrlResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *PkiConfigureCrlResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.


### HasDisable

`func (o *PkiConfigureCrlResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.




### GetEnableDelta

`func (o *PkiConfigureCrlResponse) GetEnableDelta() bool`

GetEnableDelta returns the EnableDelta field if non-nil, zero value otherwise.

### GetEnableDeltaOk

`func (o *PkiConfigureCrlResponse) GetEnableDeltaOk() (*bool, bool)`

GetEnableDeltaOk returns a tuple with the EnableDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDelta

`func (o *PkiConfigureCrlResponse) SetEnableDelta(v bool)`

SetEnableDelta sets EnableDelta field to given value.


### HasEnableDelta

`func (o *PkiConfigureCrlResponse) HasEnableDelta() bool`

HasEnableDelta returns a boolean if a field has been set.




### GetExpiry

`func (o *PkiConfigureCrlResponse) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *PkiConfigureCrlResponse) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *PkiConfigureCrlResponse) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.


### HasExpiry

`func (o *PkiConfigureCrlResponse) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.




### GetOcspDisable

`func (o *PkiConfigureCrlResponse) GetOcspDisable() bool`

GetOcspDisable returns the OcspDisable field if non-nil, zero value otherwise.

### GetOcspDisableOk

`func (o *PkiConfigureCrlResponse) GetOcspDisableOk() (*bool, bool)`

GetOcspDisableOk returns a tuple with the OcspDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspDisable

`func (o *PkiConfigureCrlResponse) SetOcspDisable(v bool)`

SetOcspDisable sets OcspDisable field to given value.


### HasOcspDisable

`func (o *PkiConfigureCrlResponse) HasOcspDisable() bool`

HasOcspDisable returns a boolean if a field has been set.




### GetOcspExpiry

`func (o *PkiConfigureCrlResponse) GetOcspExpiry() string`

GetOcspExpiry returns the OcspExpiry field if non-nil, zero value otherwise.

### GetOcspExpiryOk

`func (o *PkiConfigureCrlResponse) GetOcspExpiryOk() (*string, bool)`

GetOcspExpiryOk returns a tuple with the OcspExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspExpiry

`func (o *PkiConfigureCrlResponse) SetOcspExpiry(v string)`

SetOcspExpiry sets OcspExpiry field to given value.


### HasOcspExpiry

`func (o *PkiConfigureCrlResponse) HasOcspExpiry() bool`

HasOcspExpiry returns a boolean if a field has been set.




### GetUnifiedCrl

`func (o *PkiConfigureCrlResponse) GetUnifiedCrl() bool`

GetUnifiedCrl returns the UnifiedCrl field if non-nil, zero value otherwise.

### GetUnifiedCrlOk

`func (o *PkiConfigureCrlResponse) GetUnifiedCrlOk() (*bool, bool)`

GetUnifiedCrlOk returns a tuple with the UnifiedCrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedCrl

`func (o *PkiConfigureCrlResponse) SetUnifiedCrl(v bool)`

SetUnifiedCrl sets UnifiedCrl field to given value.


### HasUnifiedCrl

`func (o *PkiConfigureCrlResponse) HasUnifiedCrl() bool`

HasUnifiedCrl returns a boolean if a field has been set.




### GetUnifiedCrlOnExistingPaths

`func (o *PkiConfigureCrlResponse) GetUnifiedCrlOnExistingPaths() bool`

GetUnifiedCrlOnExistingPaths returns the UnifiedCrlOnExistingPaths field if non-nil, zero value otherwise.

### GetUnifiedCrlOnExistingPathsOk

`func (o *PkiConfigureCrlResponse) GetUnifiedCrlOnExistingPathsOk() (*bool, bool)`

GetUnifiedCrlOnExistingPathsOk returns a tuple with the UnifiedCrlOnExistingPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedCrlOnExistingPaths

`func (o *PkiConfigureCrlResponse) SetUnifiedCrlOnExistingPaths(v bool)`

SetUnifiedCrlOnExistingPaths sets UnifiedCrlOnExistingPaths field to given value.


### HasUnifiedCrlOnExistingPaths

`func (o *PkiConfigureCrlResponse) HasUnifiedCrlOnExistingPaths() bool`

HasUnifiedCrlOnExistingPaths returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


