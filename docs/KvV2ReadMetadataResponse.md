# KvV2ReadMetadataResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CasRequired** | Pointer to **bool** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**CurrentVersion** | Pointer to **int64** |  | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** | User-provided key-value pairs that are used to describe arbitrary and version-agnostic information about a secret. | [optional] 
**DeleteVersionAfter** | Pointer to **int32** | The length of time before a version is deleted. | [optional] 
**MaxVersions** | Pointer to **int64** | The number of versions to keep | [optional] 
**OldestVersion** | Pointer to **int64** |  | [optional] 
**UpdatedTime** | Pointer to **time.Time** |  | [optional] 
**Versions** | Pointer to **map[string]interface{}** |  | [optional] 



## Methods


### NewKvV2ReadMetadataResponse

`func NewKvV2ReadMetadataResponse() *KvV2ReadMetadataResponse`

NewKvV2ReadMetadataResponse instantiates a new KvV2ReadMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvV2ReadMetadataResponseWithDefaults

`func NewKvV2ReadMetadataResponseWithDefaults() *KvV2ReadMetadataResponse`

NewKvV2ReadMetadataResponseWithDefaults instantiates a new KvV2ReadMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCasRequired

`func (o *KvV2ReadMetadataResponse) GetCasRequired() bool`

GetCasRequired returns the CasRequired field if non-nil, zero value otherwise.

### GetCasRequiredOk

`func (o *KvV2ReadMetadataResponse) GetCasRequiredOk() (*bool, bool)`

GetCasRequiredOk returns a tuple with the CasRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasRequired

`func (o *KvV2ReadMetadataResponse) SetCasRequired(v bool)`

SetCasRequired sets CasRequired field to given value.


### HasCasRequired

`func (o *KvV2ReadMetadataResponse) HasCasRequired() bool`

HasCasRequired returns a boolean if a field has been set.




### GetCreatedTime

`func (o *KvV2ReadMetadataResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *KvV2ReadMetadataResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *KvV2ReadMetadataResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### HasCreatedTime

`func (o *KvV2ReadMetadataResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.




### GetCurrentVersion

`func (o *KvV2ReadMetadataResponse) GetCurrentVersion() int64`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *KvV2ReadMetadataResponse) GetCurrentVersionOk() (*int64, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *KvV2ReadMetadataResponse) SetCurrentVersion(v int64)`

SetCurrentVersion sets CurrentVersion field to given value.


### HasCurrentVersion

`func (o *KvV2ReadMetadataResponse) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.




### GetCustomMetadata

`func (o *KvV2ReadMetadataResponse) GetCustomMetadata() map[string]interface{}`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *KvV2ReadMetadataResponse) GetCustomMetadataOk() (*map[string]interface{}, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *KvV2ReadMetadataResponse) SetCustomMetadata(v map[string]interface{})`

SetCustomMetadata sets CustomMetadata field to given value.


### HasCustomMetadata

`func (o *KvV2ReadMetadataResponse) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.




### GetDeleteVersionAfter

`func (o *KvV2ReadMetadataResponse) GetDeleteVersionAfter() int32`

GetDeleteVersionAfter returns the DeleteVersionAfter field if non-nil, zero value otherwise.

### GetDeleteVersionAfterOk

`func (o *KvV2ReadMetadataResponse) GetDeleteVersionAfterOk() (*int32, bool)`

GetDeleteVersionAfterOk returns a tuple with the DeleteVersionAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteVersionAfter

`func (o *KvV2ReadMetadataResponse) SetDeleteVersionAfter(v int32)`

SetDeleteVersionAfter sets DeleteVersionAfter field to given value.


### HasDeleteVersionAfter

`func (o *KvV2ReadMetadataResponse) HasDeleteVersionAfter() bool`

HasDeleteVersionAfter returns a boolean if a field has been set.




### GetMaxVersions

`func (o *KvV2ReadMetadataResponse) GetMaxVersions() int64`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *KvV2ReadMetadataResponse) GetMaxVersionsOk() (*int64, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *KvV2ReadMetadataResponse) SetMaxVersions(v int64)`

SetMaxVersions sets MaxVersions field to given value.


### HasMaxVersions

`func (o *KvV2ReadMetadataResponse) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.




### GetOldestVersion

`func (o *KvV2ReadMetadataResponse) GetOldestVersion() int64`

GetOldestVersion returns the OldestVersion field if non-nil, zero value otherwise.

### GetOldestVersionOk

`func (o *KvV2ReadMetadataResponse) GetOldestVersionOk() (*int64, bool)`

GetOldestVersionOk returns a tuple with the OldestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestVersion

`func (o *KvV2ReadMetadataResponse) SetOldestVersion(v int64)`

SetOldestVersion sets OldestVersion field to given value.


### HasOldestVersion

`func (o *KvV2ReadMetadataResponse) HasOldestVersion() bool`

HasOldestVersion returns a boolean if a field has been set.




### GetUpdatedTime

`func (o *KvV2ReadMetadataResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *KvV2ReadMetadataResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *KvV2ReadMetadataResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.


### HasUpdatedTime

`func (o *KvV2ReadMetadataResponse) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.




### GetVersions

`func (o *KvV2ReadMetadataResponse) GetVersions() map[string]interface{}`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *KvV2ReadMetadataResponse) GetVersionsOk() (*map[string]interface{}, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *KvV2ReadMetadataResponse) SetVersions(v map[string]interface{})`

SetVersions sets Versions field to given value.


### HasVersions

`func (o *KvV2ReadMetadataResponse) HasVersions() bool`

HasVersions returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


