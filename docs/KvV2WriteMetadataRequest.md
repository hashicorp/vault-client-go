# KvV2WriteMetadataRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CasRequired** | Pointer to **bool** | If true the key will require the cas parameter to be set on all write requests. If false, the backend’s configuration will be used. | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** | User-provided key-value pairs that are used to describe arbitrary and version-agnostic information about a secret. | [optional] 
**DeleteVersionAfter** | Pointer to **string** | The length of time before a version is deleted. If not set, the backend&#x27;s configured delete_version_after is used. Cannot be greater than the backend&#x27;s delete_version_after. A zero duration clears the current setting. A negative duration will cause an error. | [optional] 
**MaxVersions** | Pointer to **int32** | The number of versions to keep. If not set, the backend’s configured max version is used. | [optional] 



## Methods


### NewKvV2WriteMetadataRequest

`func NewKvV2WriteMetadataRequest() *KvV2WriteMetadataRequest`

NewKvV2WriteMetadataRequest instantiates a new KvV2WriteMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvV2WriteMetadataRequestWithDefaults

`func NewKvV2WriteMetadataRequestWithDefaults() *KvV2WriteMetadataRequest`

NewKvV2WriteMetadataRequestWithDefaults instantiates a new KvV2WriteMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCasRequired

`func (o *KvV2WriteMetadataRequest) GetCasRequired() bool`

GetCasRequired returns the CasRequired field if non-nil, zero value otherwise.

### GetCasRequiredOk

`func (o *KvV2WriteMetadataRequest) GetCasRequiredOk() (*bool, bool)`

GetCasRequiredOk returns a tuple with the CasRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasRequired

`func (o *KvV2WriteMetadataRequest) SetCasRequired(v bool)`

SetCasRequired sets CasRequired field to given value.


### HasCasRequired

`func (o *KvV2WriteMetadataRequest) HasCasRequired() bool`

HasCasRequired returns a boolean if a field has been set.




### GetCustomMetadata

`func (o *KvV2WriteMetadataRequest) GetCustomMetadata() map[string]interface{}`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *KvV2WriteMetadataRequest) GetCustomMetadataOk() (*map[string]interface{}, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *KvV2WriteMetadataRequest) SetCustomMetadata(v map[string]interface{})`

SetCustomMetadata sets CustomMetadata field to given value.


### HasCustomMetadata

`func (o *KvV2WriteMetadataRequest) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.




### GetDeleteVersionAfter

`func (o *KvV2WriteMetadataRequest) GetDeleteVersionAfter() string`

GetDeleteVersionAfter returns the DeleteVersionAfter field if non-nil, zero value otherwise.

### GetDeleteVersionAfterOk

`func (o *KvV2WriteMetadataRequest) GetDeleteVersionAfterOk() (*string, bool)`

GetDeleteVersionAfterOk returns a tuple with the DeleteVersionAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteVersionAfter

`func (o *KvV2WriteMetadataRequest) SetDeleteVersionAfter(v string)`

SetDeleteVersionAfter sets DeleteVersionAfter field to given value.


### HasDeleteVersionAfter

`func (o *KvV2WriteMetadataRequest) HasDeleteVersionAfter() bool`

HasDeleteVersionAfter returns a boolean if a field has been set.




### GetMaxVersions

`func (o *KvV2WriteMetadataRequest) GetMaxVersions() int32`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *KvV2WriteMetadataRequest) GetMaxVersionsOk() (*int32, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *KvV2WriteMetadataRequest) SetMaxVersions(v int32)`

SetMaxVersions sets MaxVersions field to given value.


### HasMaxVersions

`func (o *KvV2WriteMetadataRequest) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


