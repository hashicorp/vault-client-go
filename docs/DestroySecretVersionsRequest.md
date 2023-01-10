# DestroySecretVersionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | Pointer to **[]int32** | The versions to destroy. Their data will be permanently deleted. | [optional] 

## Methods

### NewDestroySecretVersionsRequest

`func NewDestroySecretVersionsRequest() *DestroySecretVersionsRequest`

NewDestroySecretVersionsRequest instantiates a new DestroySecretVersionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestroySecretVersionsRequestWithDefaults

`func NewDestroySecretVersionsRequestWithDefaults() *DestroySecretVersionsRequest`

NewDestroySecretVersionsRequestWithDefaults instantiates a new DestroySecretVersionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *DestroySecretVersionsRequest) GetVersions() []int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *DestroySecretVersionsRequest) GetVersionsOk() (*[]int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *DestroySecretVersionsRequest) SetVersions(v []int32)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *DestroySecretVersionsRequest) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


