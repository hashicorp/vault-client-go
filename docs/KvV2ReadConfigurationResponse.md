# KvV2ReadConfigurationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CasRequired** | Pointer to **bool** | If true, the backend will require the cas parameter to be set for each write | [optional] 
**DeleteVersionAfter** | Pointer to **int32** | The length of time before a version is deleted. | [optional] 
**MaxVersions** | Pointer to **int32** | The number of versions to keep for each key. | [optional] 



## Methods


### NewKvV2ReadConfigurationResponse

`func NewKvV2ReadConfigurationResponse() *KvV2ReadConfigurationResponse`

NewKvV2ReadConfigurationResponse instantiates a new KvV2ReadConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvV2ReadConfigurationResponseWithDefaults

`func NewKvV2ReadConfigurationResponseWithDefaults() *KvV2ReadConfigurationResponse`

NewKvV2ReadConfigurationResponseWithDefaults instantiates a new KvV2ReadConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCasRequired

`func (o *KvV2ReadConfigurationResponse) GetCasRequired() bool`

GetCasRequired returns the CasRequired field if non-nil, zero value otherwise.

### GetCasRequiredOk

`func (o *KvV2ReadConfigurationResponse) GetCasRequiredOk() (*bool, bool)`

GetCasRequiredOk returns a tuple with the CasRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasRequired

`func (o *KvV2ReadConfigurationResponse) SetCasRequired(v bool)`

SetCasRequired sets CasRequired field to given value.


### HasCasRequired

`func (o *KvV2ReadConfigurationResponse) HasCasRequired() bool`

HasCasRequired returns a boolean if a field has been set.




### GetDeleteVersionAfter

`func (o *KvV2ReadConfigurationResponse) GetDeleteVersionAfter() int32`

GetDeleteVersionAfter returns the DeleteVersionAfter field if non-nil, zero value otherwise.

### GetDeleteVersionAfterOk

`func (o *KvV2ReadConfigurationResponse) GetDeleteVersionAfterOk() (*int32, bool)`

GetDeleteVersionAfterOk returns a tuple with the DeleteVersionAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteVersionAfter

`func (o *KvV2ReadConfigurationResponse) SetDeleteVersionAfter(v int32)`

SetDeleteVersionAfter sets DeleteVersionAfter field to given value.


### HasDeleteVersionAfter

`func (o *KvV2ReadConfigurationResponse) HasDeleteVersionAfter() bool`

HasDeleteVersionAfter returns a boolean if a field has been set.




### GetMaxVersions

`func (o *KvV2ReadConfigurationResponse) GetMaxVersions() int32`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *KvV2ReadConfigurationResponse) GetMaxVersionsOk() (*int32, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *KvV2ReadConfigurationResponse) SetMaxVersions(v int32)`

SetMaxVersions sets MaxVersions field to given value.


### HasMaxVersions

`func (o *KvV2ReadConfigurationResponse) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


