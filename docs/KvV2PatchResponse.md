# KvV2PatchResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**DeletionTime** | Pointer to **string** |  | [optional] 
**Destroyed** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 



## Methods


### NewKvV2PatchResponse

`func NewKvV2PatchResponse() *KvV2PatchResponse`

NewKvV2PatchResponse instantiates a new KvV2PatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvV2PatchResponseWithDefaults

`func NewKvV2PatchResponseWithDefaults() *KvV2PatchResponse`

NewKvV2PatchResponseWithDefaults instantiates a new KvV2PatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCreatedTime

`func (o *KvV2PatchResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *KvV2PatchResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *KvV2PatchResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### HasCreatedTime

`func (o *KvV2PatchResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.




### GetCustomMetadata

`func (o *KvV2PatchResponse) GetCustomMetadata() map[string]interface{}`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *KvV2PatchResponse) GetCustomMetadataOk() (*map[string]interface{}, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *KvV2PatchResponse) SetCustomMetadata(v map[string]interface{})`

SetCustomMetadata sets CustomMetadata field to given value.


### HasCustomMetadata

`func (o *KvV2PatchResponse) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.




### GetDeletionTime

`func (o *KvV2PatchResponse) GetDeletionTime() string`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *KvV2PatchResponse) GetDeletionTimeOk() (*string, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *KvV2PatchResponse) SetDeletionTime(v string)`

SetDeletionTime sets DeletionTime field to given value.


### HasDeletionTime

`func (o *KvV2PatchResponse) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.




### GetDestroyed

`func (o *KvV2PatchResponse) GetDestroyed() bool`

GetDestroyed returns the Destroyed field if non-nil, zero value otherwise.

### GetDestroyedOk

`func (o *KvV2PatchResponse) GetDestroyedOk() (*bool, bool)`

GetDestroyedOk returns a tuple with the Destroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyed

`func (o *KvV2PatchResponse) SetDestroyed(v bool)`

SetDestroyed sets Destroyed field to given value.


### HasDestroyed

`func (o *KvV2PatchResponse) HasDestroyed() bool`

HasDestroyed returns a boolean if a field has been set.




### GetVersion

`func (o *KvV2PatchResponse) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KvV2PatchResponse) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KvV2PatchResponse) SetVersion(v int64)`

SetVersion sets Version field to given value.


### HasVersion

`func (o *KvV2PatchResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


