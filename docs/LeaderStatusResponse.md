# LeaderStatusResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveTime** | Pointer to **time.Time** |  | [optional] 
**HaEnabled** | Pointer to **bool** |  | [optional] 
**IsSelf** | Pointer to **bool** |  | [optional] 
**LastWal** | Pointer to **int64** |  | [optional] 
**LeaderAddress** | Pointer to **string** |  | [optional] 
**LeaderClusterAddress** | Pointer to **string** |  | [optional] 
**PerformanceStandby** | Pointer to **bool** |  | [optional] 
**PerformanceStandbyLastRemoteWal** | Pointer to **int64** |  | [optional] 
**RaftAppliedIndex** | Pointer to **int64** |  | [optional] 
**RaftCommittedIndex** | Pointer to **int64** |  | [optional] 



## Methods


### NewLeaderStatusResponse

`func NewLeaderStatusResponse() *LeaderStatusResponse`

NewLeaderStatusResponse instantiates a new LeaderStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderStatusResponseWithDefaults

`func NewLeaderStatusResponseWithDefaults() *LeaderStatusResponse`

NewLeaderStatusResponseWithDefaults instantiates a new LeaderStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetActiveTime

`func (o *LeaderStatusResponse) GetActiveTime() time.Time`

GetActiveTime returns the ActiveTime field if non-nil, zero value otherwise.

### GetActiveTimeOk

`func (o *LeaderStatusResponse) GetActiveTimeOk() (*time.Time, bool)`

GetActiveTimeOk returns a tuple with the ActiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTime

`func (o *LeaderStatusResponse) SetActiveTime(v time.Time)`

SetActiveTime sets ActiveTime field to given value.


### HasActiveTime

`func (o *LeaderStatusResponse) HasActiveTime() bool`

HasActiveTime returns a boolean if a field has been set.




### GetHaEnabled

`func (o *LeaderStatusResponse) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *LeaderStatusResponse) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *LeaderStatusResponse) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.


### HasHaEnabled

`func (o *LeaderStatusResponse) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.




### GetIsSelf

`func (o *LeaderStatusResponse) GetIsSelf() bool`

GetIsSelf returns the IsSelf field if non-nil, zero value otherwise.

### GetIsSelfOk

`func (o *LeaderStatusResponse) GetIsSelfOk() (*bool, bool)`

GetIsSelfOk returns a tuple with the IsSelf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelf

`func (o *LeaderStatusResponse) SetIsSelf(v bool)`

SetIsSelf sets IsSelf field to given value.


### HasIsSelf

`func (o *LeaderStatusResponse) HasIsSelf() bool`

HasIsSelf returns a boolean if a field has been set.




### GetLastWal

`func (o *LeaderStatusResponse) GetLastWal() int64`

GetLastWal returns the LastWal field if non-nil, zero value otherwise.

### GetLastWalOk

`func (o *LeaderStatusResponse) GetLastWalOk() (*int64, bool)`

GetLastWalOk returns a tuple with the LastWal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWal

`func (o *LeaderStatusResponse) SetLastWal(v int64)`

SetLastWal sets LastWal field to given value.


### HasLastWal

`func (o *LeaderStatusResponse) HasLastWal() bool`

HasLastWal returns a boolean if a field has been set.




### GetLeaderAddress

`func (o *LeaderStatusResponse) GetLeaderAddress() string`

GetLeaderAddress returns the LeaderAddress field if non-nil, zero value otherwise.

### GetLeaderAddressOk

`func (o *LeaderStatusResponse) GetLeaderAddressOk() (*string, bool)`

GetLeaderAddressOk returns a tuple with the LeaderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderAddress

`func (o *LeaderStatusResponse) SetLeaderAddress(v string)`

SetLeaderAddress sets LeaderAddress field to given value.


### HasLeaderAddress

`func (o *LeaderStatusResponse) HasLeaderAddress() bool`

HasLeaderAddress returns a boolean if a field has been set.




### GetLeaderClusterAddress

`func (o *LeaderStatusResponse) GetLeaderClusterAddress() string`

GetLeaderClusterAddress returns the LeaderClusterAddress field if non-nil, zero value otherwise.

### GetLeaderClusterAddressOk

`func (o *LeaderStatusResponse) GetLeaderClusterAddressOk() (*string, bool)`

GetLeaderClusterAddressOk returns a tuple with the LeaderClusterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderClusterAddress

`func (o *LeaderStatusResponse) SetLeaderClusterAddress(v string)`

SetLeaderClusterAddress sets LeaderClusterAddress field to given value.


### HasLeaderClusterAddress

`func (o *LeaderStatusResponse) HasLeaderClusterAddress() bool`

HasLeaderClusterAddress returns a boolean if a field has been set.




### GetPerformanceStandby

`func (o *LeaderStatusResponse) GetPerformanceStandby() bool`

GetPerformanceStandby returns the PerformanceStandby field if non-nil, zero value otherwise.

### GetPerformanceStandbyOk

`func (o *LeaderStatusResponse) GetPerformanceStandbyOk() (*bool, bool)`

GetPerformanceStandbyOk returns a tuple with the PerformanceStandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceStandby

`func (o *LeaderStatusResponse) SetPerformanceStandby(v bool)`

SetPerformanceStandby sets PerformanceStandby field to given value.


### HasPerformanceStandby

`func (o *LeaderStatusResponse) HasPerformanceStandby() bool`

HasPerformanceStandby returns a boolean if a field has been set.




### GetPerformanceStandbyLastRemoteWal

`func (o *LeaderStatusResponse) GetPerformanceStandbyLastRemoteWal() int64`

GetPerformanceStandbyLastRemoteWal returns the PerformanceStandbyLastRemoteWal field if non-nil, zero value otherwise.

### GetPerformanceStandbyLastRemoteWalOk

`func (o *LeaderStatusResponse) GetPerformanceStandbyLastRemoteWalOk() (*int64, bool)`

GetPerformanceStandbyLastRemoteWalOk returns a tuple with the PerformanceStandbyLastRemoteWal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceStandbyLastRemoteWal

`func (o *LeaderStatusResponse) SetPerformanceStandbyLastRemoteWal(v int64)`

SetPerformanceStandbyLastRemoteWal sets PerformanceStandbyLastRemoteWal field to given value.


### HasPerformanceStandbyLastRemoteWal

`func (o *LeaderStatusResponse) HasPerformanceStandbyLastRemoteWal() bool`

HasPerformanceStandbyLastRemoteWal returns a boolean if a field has been set.




### GetRaftAppliedIndex

`func (o *LeaderStatusResponse) GetRaftAppliedIndex() int64`

GetRaftAppliedIndex returns the RaftAppliedIndex field if non-nil, zero value otherwise.

### GetRaftAppliedIndexOk

`func (o *LeaderStatusResponse) GetRaftAppliedIndexOk() (*int64, bool)`

GetRaftAppliedIndexOk returns a tuple with the RaftAppliedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaftAppliedIndex

`func (o *LeaderStatusResponse) SetRaftAppliedIndex(v int64)`

SetRaftAppliedIndex sets RaftAppliedIndex field to given value.


### HasRaftAppliedIndex

`func (o *LeaderStatusResponse) HasRaftAppliedIndex() bool`

HasRaftAppliedIndex returns a boolean if a field has been set.




### GetRaftCommittedIndex

`func (o *LeaderStatusResponse) GetRaftCommittedIndex() int64`

GetRaftCommittedIndex returns the RaftCommittedIndex field if non-nil, zero value otherwise.

### GetRaftCommittedIndexOk

`func (o *LeaderStatusResponse) GetRaftCommittedIndexOk() (*int64, bool)`

GetRaftCommittedIndexOk returns a tuple with the RaftCommittedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaftCommittedIndex

`func (o *LeaderStatusResponse) SetRaftCommittedIndex(v int64)`

SetRaftCommittedIndex sets RaftCommittedIndex field to given value.


### HasRaftCommittedIndex

`func (o *LeaderStatusResponse) HasRaftCommittedIndex() bool`

HasRaftCommittedIndex returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


