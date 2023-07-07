# EncryptionKeyConfigureRotationRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether automatic rotation is enabled. | [optional] 
**Interval** | Pointer to **string** | How long after installation of an active key term that the key will be automatically rotated. | [optional] 
**MaxOperations** | Pointer to **int64** | The number of encryption operations performed before the barrier key is automatically rotated. | [optional] 



## Methods


### NewEncryptionKeyConfigureRotationRequest

`func NewEncryptionKeyConfigureRotationRequest() *EncryptionKeyConfigureRotationRequest`

NewEncryptionKeyConfigureRotationRequest instantiates a new EncryptionKeyConfigureRotationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionKeyConfigureRotationRequestWithDefaults

`func NewEncryptionKeyConfigureRotationRequestWithDefaults() *EncryptionKeyConfigureRotationRequest`

NewEncryptionKeyConfigureRotationRequestWithDefaults instantiates a new EncryptionKeyConfigureRotationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetEnabled

`func (o *EncryptionKeyConfigureRotationRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EncryptionKeyConfigureRotationRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EncryptionKeyConfigureRotationRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### HasEnabled

`func (o *EncryptionKeyConfigureRotationRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.




### GetInterval

`func (o *EncryptionKeyConfigureRotationRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *EncryptionKeyConfigureRotationRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *EncryptionKeyConfigureRotationRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.


### HasInterval

`func (o *EncryptionKeyConfigureRotationRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.




### GetMaxOperations

`func (o *EncryptionKeyConfigureRotationRequest) GetMaxOperations() int64`

GetMaxOperations returns the MaxOperations field if non-nil, zero value otherwise.

### GetMaxOperationsOk

`func (o *EncryptionKeyConfigureRotationRequest) GetMaxOperationsOk() (*int64, bool)`

GetMaxOperationsOk returns a tuple with the MaxOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOperations

`func (o *EncryptionKeyConfigureRotationRequest) SetMaxOperations(v int64)`

SetMaxOperations sets MaxOperations field to given value.


### HasMaxOperations

`func (o *EncryptionKeyConfigureRotationRequest) HasMaxOperations() bool`

HasMaxOperations returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


