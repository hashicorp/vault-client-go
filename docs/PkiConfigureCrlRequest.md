# PkiConfigureCrlRequest


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
**UnifiedCrl** | Pointer to **bool** | If set to true enables global replication of revocation entries, also enabling unified versions of OCSP and CRLs if their respective features are enabled. disable for CRLs and ocsp_disable for OCSP. | [optional] [default to false]
**UnifiedCrlOnExistingPaths** | Pointer to **bool** | If set to true, existing CRL and OCSP paths will return the unified CRL instead of a response based on cluster-local data | [optional] [default to false]



## Methods


### NewPkiConfigureCrlRequest

`func NewPkiConfigureCrlRequest() *PkiConfigureCrlRequest`

NewPkiConfigureCrlRequest instantiates a new PkiConfigureCrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiConfigureCrlRequestWithDefaults

`func NewPkiConfigureCrlRequestWithDefaults() *PkiConfigureCrlRequest`

NewPkiConfigureCrlRequestWithDefaults instantiates a new PkiConfigureCrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAutoRebuild

`func (o *PkiConfigureCrlRequest) GetAutoRebuild() bool`

GetAutoRebuild returns the AutoRebuild field if non-nil, zero value otherwise.

### GetAutoRebuildOk

`func (o *PkiConfigureCrlRequest) GetAutoRebuildOk() (*bool, bool)`

GetAutoRebuildOk returns a tuple with the AutoRebuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRebuild

`func (o *PkiConfigureCrlRequest) SetAutoRebuild(v bool)`

SetAutoRebuild sets AutoRebuild field to given value.


### HasAutoRebuild

`func (o *PkiConfigureCrlRequest) HasAutoRebuild() bool`

HasAutoRebuild returns a boolean if a field has been set.




### GetAutoRebuildGracePeriod

`func (o *PkiConfigureCrlRequest) GetAutoRebuildGracePeriod() string`

GetAutoRebuildGracePeriod returns the AutoRebuildGracePeriod field if non-nil, zero value otherwise.

### GetAutoRebuildGracePeriodOk

`func (o *PkiConfigureCrlRequest) GetAutoRebuildGracePeriodOk() (*string, bool)`

GetAutoRebuildGracePeriodOk returns a tuple with the AutoRebuildGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRebuildGracePeriod

`func (o *PkiConfigureCrlRequest) SetAutoRebuildGracePeriod(v string)`

SetAutoRebuildGracePeriod sets AutoRebuildGracePeriod field to given value.


### HasAutoRebuildGracePeriod

`func (o *PkiConfigureCrlRequest) HasAutoRebuildGracePeriod() bool`

HasAutoRebuildGracePeriod returns a boolean if a field has been set.




### GetCrossClusterRevocation

`func (o *PkiConfigureCrlRequest) GetCrossClusterRevocation() bool`

GetCrossClusterRevocation returns the CrossClusterRevocation field if non-nil, zero value otherwise.

### GetCrossClusterRevocationOk

`func (o *PkiConfigureCrlRequest) GetCrossClusterRevocationOk() (*bool, bool)`

GetCrossClusterRevocationOk returns a tuple with the CrossClusterRevocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossClusterRevocation

`func (o *PkiConfigureCrlRequest) SetCrossClusterRevocation(v bool)`

SetCrossClusterRevocation sets CrossClusterRevocation field to given value.


### HasCrossClusterRevocation

`func (o *PkiConfigureCrlRequest) HasCrossClusterRevocation() bool`

HasCrossClusterRevocation returns a boolean if a field has been set.




### GetDeltaRebuildInterval

`func (o *PkiConfigureCrlRequest) GetDeltaRebuildInterval() string`

GetDeltaRebuildInterval returns the DeltaRebuildInterval field if non-nil, zero value otherwise.

### GetDeltaRebuildIntervalOk

`func (o *PkiConfigureCrlRequest) GetDeltaRebuildIntervalOk() (*string, bool)`

GetDeltaRebuildIntervalOk returns a tuple with the DeltaRebuildInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaRebuildInterval

`func (o *PkiConfigureCrlRequest) SetDeltaRebuildInterval(v string)`

SetDeltaRebuildInterval sets DeltaRebuildInterval field to given value.


### HasDeltaRebuildInterval

`func (o *PkiConfigureCrlRequest) HasDeltaRebuildInterval() bool`

HasDeltaRebuildInterval returns a boolean if a field has been set.




### GetDisable

`func (o *PkiConfigureCrlRequest) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *PkiConfigureCrlRequest) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *PkiConfigureCrlRequest) SetDisable(v bool)`

SetDisable sets Disable field to given value.


### HasDisable

`func (o *PkiConfigureCrlRequest) HasDisable() bool`

HasDisable returns a boolean if a field has been set.




### GetEnableDelta

`func (o *PkiConfigureCrlRequest) GetEnableDelta() bool`

GetEnableDelta returns the EnableDelta field if non-nil, zero value otherwise.

### GetEnableDeltaOk

`func (o *PkiConfigureCrlRequest) GetEnableDeltaOk() (*bool, bool)`

GetEnableDeltaOk returns a tuple with the EnableDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDelta

`func (o *PkiConfigureCrlRequest) SetEnableDelta(v bool)`

SetEnableDelta sets EnableDelta field to given value.


### HasEnableDelta

`func (o *PkiConfigureCrlRequest) HasEnableDelta() bool`

HasEnableDelta returns a boolean if a field has been set.




### GetExpiry

`func (o *PkiConfigureCrlRequest) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *PkiConfigureCrlRequest) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *PkiConfigureCrlRequest) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.


### HasExpiry

`func (o *PkiConfigureCrlRequest) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.




### GetOcspDisable

`func (o *PkiConfigureCrlRequest) GetOcspDisable() bool`

GetOcspDisable returns the OcspDisable field if non-nil, zero value otherwise.

### GetOcspDisableOk

`func (o *PkiConfigureCrlRequest) GetOcspDisableOk() (*bool, bool)`

GetOcspDisableOk returns a tuple with the OcspDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspDisable

`func (o *PkiConfigureCrlRequest) SetOcspDisable(v bool)`

SetOcspDisable sets OcspDisable field to given value.


### HasOcspDisable

`func (o *PkiConfigureCrlRequest) HasOcspDisable() bool`

HasOcspDisable returns a boolean if a field has been set.




### GetOcspExpiry

`func (o *PkiConfigureCrlRequest) GetOcspExpiry() string`

GetOcspExpiry returns the OcspExpiry field if non-nil, zero value otherwise.

### GetOcspExpiryOk

`func (o *PkiConfigureCrlRequest) GetOcspExpiryOk() (*string, bool)`

GetOcspExpiryOk returns a tuple with the OcspExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspExpiry

`func (o *PkiConfigureCrlRequest) SetOcspExpiry(v string)`

SetOcspExpiry sets OcspExpiry field to given value.


### HasOcspExpiry

`func (o *PkiConfigureCrlRequest) HasOcspExpiry() bool`

HasOcspExpiry returns a boolean if a field has been set.




### GetUnifiedCrl

`func (o *PkiConfigureCrlRequest) GetUnifiedCrl() bool`

GetUnifiedCrl returns the UnifiedCrl field if non-nil, zero value otherwise.

### GetUnifiedCrlOk

`func (o *PkiConfigureCrlRequest) GetUnifiedCrlOk() (*bool, bool)`

GetUnifiedCrlOk returns a tuple with the UnifiedCrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedCrl

`func (o *PkiConfigureCrlRequest) SetUnifiedCrl(v bool)`

SetUnifiedCrl sets UnifiedCrl field to given value.


### HasUnifiedCrl

`func (o *PkiConfigureCrlRequest) HasUnifiedCrl() bool`

HasUnifiedCrl returns a boolean if a field has been set.




### GetUnifiedCrlOnExistingPaths

`func (o *PkiConfigureCrlRequest) GetUnifiedCrlOnExistingPaths() bool`

GetUnifiedCrlOnExistingPaths returns the UnifiedCrlOnExistingPaths field if non-nil, zero value otherwise.

### GetUnifiedCrlOnExistingPathsOk

`func (o *PkiConfigureCrlRequest) GetUnifiedCrlOnExistingPathsOk() (*bool, bool)`

GetUnifiedCrlOnExistingPathsOk returns a tuple with the UnifiedCrlOnExistingPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnifiedCrlOnExistingPaths

`func (o *PkiConfigureCrlRequest) SetUnifiedCrlOnExistingPaths(v bool)`

SetUnifiedCrlOnExistingPaths sets UnifiedCrlOnExistingPaths field to given value.


### HasUnifiedCrlOnExistingPaths

`func (o *PkiConfigureCrlRequest) HasUnifiedCrlOnExistingPaths() bool`

HasUnifiedCrlOnExistingPaths returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


