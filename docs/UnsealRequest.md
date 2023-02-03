# UnsealRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Key** | Pointer to **string** | Specifies a single unseal key share. This is required unless reset is true. | [optional] 
**Reset** | Pointer to **bool** | Specifies if previously-provided unseal keys are discarded and the unseal process is reset. | [optional] 



## Methods


### NewUnsealRequest

`func NewUnsealRequest() *UnsealRequest`

NewUnsealRequest instantiates a new UnsealRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnsealRequestWithDefaults

`func NewUnsealRequestWithDefaults() *UnsealRequest`

NewUnsealRequestWithDefaults instantiates a new UnsealRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetKey

`func (o *UnsealRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UnsealRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UnsealRequest) SetKey(v string)`

SetKey sets Key field to given value.


### HasKey

`func (o *UnsealRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.




### GetReset

`func (o *UnsealRequest) GetReset() bool`

GetReset returns the Reset field if non-nil, zero value otherwise.

### GetResetOk

`func (o *UnsealRequest) GetResetOk() (*bool, bool)`

GetResetOk returns a tuple with the Reset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReset

`func (o *UnsealRequest) SetReset(v bool)`

SetReset sets Reset field to given value.


### HasReset

`func (o *UnsealRequest) HasReset() bool`

HasReset returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


