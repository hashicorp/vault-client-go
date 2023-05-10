# KvV2WriteRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | The contents of the data map will be stored and returned on read. | [optional] 
**Options** | Pointer to **map[string]interface{}** | Options for writing a KV entry. Set the \&quot;cas\&quot; value to use a Check-And-Set operation. If not set the write will be allowed. If set to 0 a write will only be allowed if the key doesn’t exist. If the index is non-zero the write will only be allowed if the key’s current version matches the version specified in the cas parameter. | [optional] 
**Version** | Pointer to **int32** | If provided during a read, the value at the version number will be returned | [optional] 



## Methods


### NewKvV2WriteRequest

`func NewKvV2WriteRequest() *KvV2WriteRequest`

NewKvV2WriteRequest instantiates a new KvV2WriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvV2WriteRequestWithDefaults

`func NewKvV2WriteRequestWithDefaults() *KvV2WriteRequest`

NewKvV2WriteRequestWithDefaults instantiates a new KvV2WriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetData

`func (o *KvV2WriteRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KvV2WriteRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KvV2WriteRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### HasData

`func (o *KvV2WriteRequest) HasData() bool`

HasData returns a boolean if a field has been set.




### GetOptions

`func (o *KvV2WriteRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *KvV2WriteRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *KvV2WriteRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.


### HasOptions

`func (o *KvV2WriteRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.




### GetVersion

`func (o *KvV2WriteRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KvV2WriteRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KvV2WriteRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.


### HasVersion

`func (o *KvV2WriteRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


