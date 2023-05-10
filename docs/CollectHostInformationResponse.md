# CollectHostInformationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CpuTimes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Disk** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Host** | Pointer to **map[string]interface{}** |  | [optional] 
**Memory** | Pointer to **map[string]interface{}** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 



## Methods


### NewCollectHostInformationResponse

`func NewCollectHostInformationResponse() *CollectHostInformationResponse`

NewCollectHostInformationResponse instantiates a new CollectHostInformationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectHostInformationResponseWithDefaults

`func NewCollectHostInformationResponseWithDefaults() *CollectHostInformationResponse`

NewCollectHostInformationResponseWithDefaults instantiates a new CollectHostInformationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCpu

`func (o *CollectHostInformationResponse) GetCpu() []map[string]interface{}`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *CollectHostInformationResponse) GetCpuOk() (*[]map[string]interface{}, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *CollectHostInformationResponse) SetCpu(v []map[string]interface{})`

SetCpu sets Cpu field to given value.


### HasCpu

`func (o *CollectHostInformationResponse) HasCpu() bool`

HasCpu returns a boolean if a field has been set.




### GetCpuTimes

`func (o *CollectHostInformationResponse) GetCpuTimes() []map[string]interface{}`

GetCpuTimes returns the CpuTimes field if non-nil, zero value otherwise.

### GetCpuTimesOk

`func (o *CollectHostInformationResponse) GetCpuTimesOk() (*[]map[string]interface{}, bool)`

GetCpuTimesOk returns a tuple with the CpuTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTimes

`func (o *CollectHostInformationResponse) SetCpuTimes(v []map[string]interface{})`

SetCpuTimes sets CpuTimes field to given value.


### HasCpuTimes

`func (o *CollectHostInformationResponse) HasCpuTimes() bool`

HasCpuTimes returns a boolean if a field has been set.




### GetDisk

`func (o *CollectHostInformationResponse) GetDisk() []map[string]interface{}`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *CollectHostInformationResponse) GetDiskOk() (*[]map[string]interface{}, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *CollectHostInformationResponse) SetDisk(v []map[string]interface{})`

SetDisk sets Disk field to given value.


### HasDisk

`func (o *CollectHostInformationResponse) HasDisk() bool`

HasDisk returns a boolean if a field has been set.




### GetHost

`func (o *CollectHostInformationResponse) GetHost() map[string]interface{}`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CollectHostInformationResponse) GetHostOk() (*map[string]interface{}, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CollectHostInformationResponse) SetHost(v map[string]interface{})`

SetHost sets Host field to given value.


### HasHost

`func (o *CollectHostInformationResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.




### GetMemory

`func (o *CollectHostInformationResponse) GetMemory() map[string]interface{}`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CollectHostInformationResponse) GetMemoryOk() (*map[string]interface{}, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CollectHostInformationResponse) SetMemory(v map[string]interface{})`

SetMemory sets Memory field to given value.


### HasMemory

`func (o *CollectHostInformationResponse) HasMemory() bool`

HasMemory returns a boolean if a field has been set.




### GetTimestamp

`func (o *CollectHostInformationResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CollectHostInformationResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CollectHostInformationResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### HasTimestamp

`func (o *CollectHostInformationResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


