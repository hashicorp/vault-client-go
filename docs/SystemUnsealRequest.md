# SystemUnsealRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Specifies a single unseal key share. This is required unless reset is true. | [optional] 
**Reset** | Pointer to **bool** | Specifies if previously-provided unseal keys are discarded and the unseal process is reset. | [optional] 

## Methods

### NewSystemUnsealRequest

`func NewSystemUnsealRequest() *SystemUnsealRequest`

NewSystemUnsealRequest instantiates a new SystemUnsealRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemUnsealRequestWithDefaults

`func NewSystemUnsealRequestWithDefaults() *SystemUnsealRequest`

NewSystemUnsealRequestWithDefaults instantiates a new SystemUnsealRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SystemUnsealRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SystemUnsealRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SystemUnsealRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SystemUnsealRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetReset

`func (o *SystemUnsealRequest) GetReset() bool`

GetReset returns the Reset field if non-nil, zero value otherwise.

### GetResetOk

`func (o *SystemUnsealRequest) GetResetOk() (*bool, bool)`

GetResetOk returns a tuple with the Reset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReset

`func (o *SystemUnsealRequest) SetReset(v bool)`

SetReset sets Reset field to given value.

### HasReset

`func (o *SystemUnsealRequest) HasReset() bool`

HasReset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


