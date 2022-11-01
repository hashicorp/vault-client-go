# PkiConfigCrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRebuild** | Pointer to **bool** | If set to true, enables automatic rebuilding of the CRL | [optional] 
**AutoRebuildGracePeriod** | Pointer to **string** | The time before the CRL expires to automatically rebuild it, when enabled. Must be shorter than the CRL expiry. Defaults to 12h. | [optional] [default to "12h"]
**DeltaRebuildInterval** | Pointer to **string** | The time between delta CRL rebuilds if a new revocation has occurred. Must be shorter than the CRL expiry. Defaults to 15m. | [optional] [default to "15m"]
**Disable** | Pointer to **bool** | If set to true, disables generating the CRL entirely. | [optional] 
**EnableDelta** | Pointer to **bool** | Whether to enable delta CRLs between authoritative CRL rebuilds | [optional] 
**Expiry** | Pointer to **string** | The amount of time the generated CRL should be valid; defaults to 72 hours | [optional] [default to "72h"]
**OcspDisable** | Pointer to **bool** | If set to true, ocsp unauthorized responses will be returned. | [optional] 
**OcspExpiry** | Pointer to **string** | The amount of time an OCSP response will be valid (controls the NextUpdate field); defaults to 12 hours | [optional] [default to "1h"]

## Methods

### NewPkiConfigCrlRequest

`func NewPkiConfigCrlRequest() *PkiConfigCrlRequest`

NewPkiConfigCrlRequest instantiates a new PkiConfigCrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiConfigCrlRequestWithDefaults

`func NewPkiConfigCrlRequestWithDefaults() *PkiConfigCrlRequest`

NewPkiConfigCrlRequestWithDefaults instantiates a new PkiConfigCrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRebuild

`func (o *PkiConfigCrlRequest) GetAutoRebuild() bool`

GetAutoRebuild returns the AutoRebuild field if non-nil, zero value otherwise.

### GetAutoRebuildOk

`func (o *PkiConfigCrlRequest) GetAutoRebuildOk() (*bool, bool)`

GetAutoRebuildOk returns a tuple with the AutoRebuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRebuild

`func (o *PkiConfigCrlRequest) SetAutoRebuild(v bool)`

SetAutoRebuild sets AutoRebuild field to given value.

### HasAutoRebuild

`func (o *PkiConfigCrlRequest) HasAutoRebuild() bool`

HasAutoRebuild returns a boolean if a field has been set.

### GetAutoRebuildGracePeriod

`func (o *PkiConfigCrlRequest) GetAutoRebuildGracePeriod() string`

GetAutoRebuildGracePeriod returns the AutoRebuildGracePeriod field if non-nil, zero value otherwise.

### GetAutoRebuildGracePeriodOk

`func (o *PkiConfigCrlRequest) GetAutoRebuildGracePeriodOk() (*string, bool)`

GetAutoRebuildGracePeriodOk returns a tuple with the AutoRebuildGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRebuildGracePeriod

`func (o *PkiConfigCrlRequest) SetAutoRebuildGracePeriod(v string)`

SetAutoRebuildGracePeriod sets AutoRebuildGracePeriod field to given value.

### HasAutoRebuildGracePeriod

`func (o *PkiConfigCrlRequest) HasAutoRebuildGracePeriod() bool`

HasAutoRebuildGracePeriod returns a boolean if a field has been set.

### GetDeltaRebuildInterval

`func (o *PkiConfigCrlRequest) GetDeltaRebuildInterval() string`

GetDeltaRebuildInterval returns the DeltaRebuildInterval field if non-nil, zero value otherwise.

### GetDeltaRebuildIntervalOk

`func (o *PkiConfigCrlRequest) GetDeltaRebuildIntervalOk() (*string, bool)`

GetDeltaRebuildIntervalOk returns a tuple with the DeltaRebuildInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaRebuildInterval

`func (o *PkiConfigCrlRequest) SetDeltaRebuildInterval(v string)`

SetDeltaRebuildInterval sets DeltaRebuildInterval field to given value.

### HasDeltaRebuildInterval

`func (o *PkiConfigCrlRequest) HasDeltaRebuildInterval() bool`

HasDeltaRebuildInterval returns a boolean if a field has been set.

### GetDisable

`func (o *PkiConfigCrlRequest) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *PkiConfigCrlRequest) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *PkiConfigCrlRequest) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *PkiConfigCrlRequest) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEnableDelta

`func (o *PkiConfigCrlRequest) GetEnableDelta() bool`

GetEnableDelta returns the EnableDelta field if non-nil, zero value otherwise.

### GetEnableDeltaOk

`func (o *PkiConfigCrlRequest) GetEnableDeltaOk() (*bool, bool)`

GetEnableDeltaOk returns a tuple with the EnableDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDelta

`func (o *PkiConfigCrlRequest) SetEnableDelta(v bool)`

SetEnableDelta sets EnableDelta field to given value.

### HasEnableDelta

`func (o *PkiConfigCrlRequest) HasEnableDelta() bool`

HasEnableDelta returns a boolean if a field has been set.

### GetExpiry

`func (o *PkiConfigCrlRequest) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *PkiConfigCrlRequest) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *PkiConfigCrlRequest) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *PkiConfigCrlRequest) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetOcspDisable

`func (o *PkiConfigCrlRequest) GetOcspDisable() bool`

GetOcspDisable returns the OcspDisable field if non-nil, zero value otherwise.

### GetOcspDisableOk

`func (o *PkiConfigCrlRequest) GetOcspDisableOk() (*bool, bool)`

GetOcspDisableOk returns a tuple with the OcspDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspDisable

`func (o *PkiConfigCrlRequest) SetOcspDisable(v bool)`

SetOcspDisable sets OcspDisable field to given value.

### HasOcspDisable

`func (o *PkiConfigCrlRequest) HasOcspDisable() bool`

HasOcspDisable returns a boolean if a field has been set.

### GetOcspExpiry

`func (o *PkiConfigCrlRequest) GetOcspExpiry() string`

GetOcspExpiry returns the OcspExpiry field if non-nil, zero value otherwise.

### GetOcspExpiryOk

`func (o *PkiConfigCrlRequest) GetOcspExpiryOk() (*string, bool)`

GetOcspExpiryOk returns a tuple with the OcspExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspExpiry

`func (o *PkiConfigCrlRequest) SetOcspExpiry(v string)`

SetOcspExpiry sets OcspExpiry field to given value.

### HasOcspExpiry

`func (o *PkiConfigCrlRequest) HasOcspExpiry() bool`

HasOcspExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


