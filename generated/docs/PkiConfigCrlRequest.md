# PkiConfigCrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disable** | Pointer to **bool** | If set to true, disables generating the CRL entirely. | [optional] 
**Expiry** | Pointer to **string** | The amount of time the generated CRL should be valid; defaults to 72 hours | [optional] [default to "72h"]

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


