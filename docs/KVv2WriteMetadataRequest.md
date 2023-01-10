# KVv2WriteMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CasRequired** | Pointer to **bool** | If true the key will require the cas parameter to be set on all write requests. If false, the backend’s configuration will be used. | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** | User-provided key-value pairs that are used to describe arbitrary and version-agnostic information about a secret. | [optional] 
**DeleteVersionAfter** | Pointer to **int32** | The length of time before a version is deleted. If not set, the backend&#39;s configured delete_version_after is used. Cannot be greater than the backend&#39;s delete_version_after. A zero duration clears the current setting. A negative duration will cause an error. | [optional] 
**MaxVersions** | Pointer to **int32** | The number of versions to keep. If not set, the backend’s configured max version is used. | [optional] 

## Methods

### NewKVv2WriteMetadataRequest

`func NewKVv2WriteMetadataRequest() *KVv2WriteMetadataRequest`

NewKVv2WriteMetadataRequest instantiates a new KVv2WriteMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKVv2WriteMetadataRequestWithDefaults

`func NewKVv2WriteMetadataRequestWithDefaults() *KVv2WriteMetadataRequest`

NewKVv2WriteMetadataRequestWithDefaults instantiates a new KVv2WriteMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCasRequired

`func (o *KVv2WriteMetadataRequest) GetCasRequired() bool`

GetCasRequired returns the CasRequired field if non-nil, zero value otherwise.

### GetCasRequiredOk

`func (o *KVv2WriteMetadataRequest) GetCasRequiredOk() (*bool, bool)`

GetCasRequiredOk returns a tuple with the CasRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasRequired

`func (o *KVv2WriteMetadataRequest) SetCasRequired(v bool)`

SetCasRequired sets CasRequired field to given value.

### HasCasRequired

`func (o *KVv2WriteMetadataRequest) HasCasRequired() bool`

HasCasRequired returns a boolean if a field has been set.

### GetCustomMetadata

`func (o *KVv2WriteMetadataRequest) GetCustomMetadata() map[string]interface{}`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *KVv2WriteMetadataRequest) GetCustomMetadataOk() (*map[string]interface{}, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *KVv2WriteMetadataRequest) SetCustomMetadata(v map[string]interface{})`

SetCustomMetadata sets CustomMetadata field to given value.

### HasCustomMetadata

`func (o *KVv2WriteMetadataRequest) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.

### GetDeleteVersionAfter

`func (o *KVv2WriteMetadataRequest) GetDeleteVersionAfter() int32`

GetDeleteVersionAfter returns the DeleteVersionAfter field if non-nil, zero value otherwise.

### GetDeleteVersionAfterOk

`func (o *KVv2WriteMetadataRequest) GetDeleteVersionAfterOk() (*int32, bool)`

GetDeleteVersionAfterOk returns a tuple with the DeleteVersionAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteVersionAfter

`func (o *KVv2WriteMetadataRequest) SetDeleteVersionAfter(v int32)`

SetDeleteVersionAfter sets DeleteVersionAfter field to given value.

### HasDeleteVersionAfter

`func (o *KVv2WriteMetadataRequest) HasDeleteVersionAfter() bool`

HasDeleteVersionAfter returns a boolean if a field has been set.

### GetMaxVersions

`func (o *KVv2WriteMetadataRequest) GetMaxVersions() int32`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *KVv2WriteMetadataRequest) GetMaxVersionsOk() (*int32, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *KVv2WriteMetadataRequest) SetMaxVersions(v int32)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *KVv2WriteMetadataRequest) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


