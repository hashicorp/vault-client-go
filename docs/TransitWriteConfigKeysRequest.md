# TransitWriteConfigKeysRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableUpsert** | Pointer to **bool** | Whether to allow automatic upserting (creation) of keys on the encrypt endpoint. | [optional] 

## Methods

### NewTransitWriteConfigKeysRequest

`func NewTransitWriteConfigKeysRequest() *TransitWriteConfigKeysRequest`

NewTransitWriteConfigKeysRequest instantiates a new TransitWriteConfigKeysRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitWriteConfigKeysRequestWithDefaults

`func NewTransitWriteConfigKeysRequestWithDefaults() *TransitWriteConfigKeysRequest`

NewTransitWriteConfigKeysRequestWithDefaults instantiates a new TransitWriteConfigKeysRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableUpsert

`func (o *TransitWriteConfigKeysRequest) GetDisableUpsert() bool`

GetDisableUpsert returns the DisableUpsert field if non-nil, zero value otherwise.

### GetDisableUpsertOk

`func (o *TransitWriteConfigKeysRequest) GetDisableUpsertOk() (*bool, bool)`

GetDisableUpsertOk returns a tuple with the DisableUpsert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUpsert

`func (o *TransitWriteConfigKeysRequest) SetDisableUpsert(v bool)`

SetDisableUpsert sets DisableUpsert field to given value.

### HasDisableUpsert

`func (o *TransitWriteConfigKeysRequest) HasDisableUpsert() bool`

HasDisableUpsert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


