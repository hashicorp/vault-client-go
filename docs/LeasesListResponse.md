# LeasesListResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counts** | Pointer to **int32** | Number of matching leases per mount | [optional] 
**LeaseCount** | Pointer to **int32** | Number of matching leases | [optional] 



## Methods


### NewLeasesListResponse

`func NewLeasesListResponse() *LeasesListResponse`

NewLeasesListResponse instantiates a new LeasesListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeasesListResponseWithDefaults

`func NewLeasesListResponseWithDefaults() *LeasesListResponse`

NewLeasesListResponseWithDefaults instantiates a new LeasesListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCounts

`func (o *LeasesListResponse) GetCounts() int32`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *LeasesListResponse) GetCountsOk() (*int32, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *LeasesListResponse) SetCounts(v int32)`

SetCounts sets Counts field to given value.


### HasCounts

`func (o *LeasesListResponse) HasCounts() bool`

HasCounts returns a boolean if a field has been set.




### GetLeaseCount

`func (o *LeasesListResponse) GetLeaseCount() int32`

GetLeaseCount returns the LeaseCount field if non-nil, zero value otherwise.

### GetLeaseCountOk

`func (o *LeasesListResponse) GetLeaseCountOk() (*int32, bool)`

GetLeaseCountOk returns a tuple with the LeaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseCount

`func (o *LeasesListResponse) SetLeaseCount(v int32)`

SetLeaseCount sets LeaseCount field to given value.


### HasLeaseCount

`func (o *LeasesListResponse) HasLeaseCount() bool`

HasLeaseCount returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


