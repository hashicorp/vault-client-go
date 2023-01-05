# KVv2UndeleteVersionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | Pointer to **[]int32** | The versions to unarchive. The versions will be restored and their data will be returned on normal get requests. | [optional] 

## Methods

### NewKVv2UndeleteVersionsRequest

`func NewKVv2UndeleteVersionsRequest() *KVv2UndeleteVersionsRequest`

NewKVv2UndeleteVersionsRequest instantiates a new KVv2UndeleteVersionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKVv2UndeleteVersionsRequestWithDefaults

`func NewKVv2UndeleteVersionsRequestWithDefaults() *KVv2UndeleteVersionsRequest`

NewKVv2UndeleteVersionsRequestWithDefaults instantiates a new KVv2UndeleteVersionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *KVv2UndeleteVersionsRequest) GetVersions() []int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *KVv2UndeleteVersionsRequest) GetVersionsOk() (*[]int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *KVv2UndeleteVersionsRequest) SetVersions(v []int32)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *KVv2UndeleteVersionsRequest) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


