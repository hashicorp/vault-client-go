# KvV2WriteResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**DeletionTime** | Pointer to **string** |  | [optional] 
**Destroyed** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 



## Methods


### NewKvV2WriteResponse

`func NewKvV2WriteResponse() *KvV2WriteResponse`

NewKvV2WriteResponse instantiates a new KvV2WriteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvV2WriteResponseWithDefaults

`func NewKvV2WriteResponseWithDefaults() *KvV2WriteResponse`

NewKvV2WriteResponseWithDefaults instantiates a new KvV2WriteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCreatedTime

`func (o *KvV2WriteResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *KvV2WriteResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *KvV2WriteResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### HasCreatedTime

`func (o *KvV2WriteResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.




### GetCustomMetadata

`func (o *KvV2WriteResponse) GetCustomMetadata() map[string]interface{}`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *KvV2WriteResponse) GetCustomMetadataOk() (*map[string]interface{}, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *KvV2WriteResponse) SetCustomMetadata(v map[string]interface{})`

SetCustomMetadata sets CustomMetadata field to given value.


### HasCustomMetadata

`func (o *KvV2WriteResponse) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.




### GetDeletionTime

`func (o *KvV2WriteResponse) GetDeletionTime() string`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *KvV2WriteResponse) GetDeletionTimeOk() (*string, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *KvV2WriteResponse) SetDeletionTime(v string)`

SetDeletionTime sets DeletionTime field to given value.


### HasDeletionTime

`func (o *KvV2WriteResponse) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.




### GetDestroyed

`func (o *KvV2WriteResponse) GetDestroyed() bool`

GetDestroyed returns the Destroyed field if non-nil, zero value otherwise.

### GetDestroyedOk

`func (o *KvV2WriteResponse) GetDestroyedOk() (*bool, bool)`

GetDestroyedOk returns a tuple with the Destroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyed

`func (o *KvV2WriteResponse) SetDestroyed(v bool)`

SetDestroyed sets Destroyed field to given value.


### HasDestroyed

`func (o *KvV2WriteResponse) HasDestroyed() bool`

HasDestroyed returns a boolean if a field has been set.




### GetVersion

`func (o *KvV2WriteResponse) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KvV2WriteResponse) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KvV2WriteResponse) SetVersion(v int64)`

SetVersion sets Version field to given value.


### HasVersion

`func (o *KvV2WriteResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


