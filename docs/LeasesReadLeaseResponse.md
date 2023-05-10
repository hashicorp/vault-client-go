# LeasesReadLeaseResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpireTime** | Pointer to **time.Time** | Optional lease expiry time | [optional] 
**Id** | Pointer to **string** | Lease id | [optional] 
**IssueTime** | Pointer to **time.Time** | Timestamp for the lease&#x27;s issue time | [optional] 
**LastRenewal** | Pointer to **time.Time** | Optional Timestamp of the last time the lease was renewed | [optional] 
**Renewable** | Pointer to **bool** | True if the lease is able to be renewed | [optional] 
**Ttl** | Pointer to **int32** | Time to Live set for the lease, returns 0 if unset | [optional] 



## Methods


### NewLeasesReadLeaseResponse

`func NewLeasesReadLeaseResponse() *LeasesReadLeaseResponse`

NewLeasesReadLeaseResponse instantiates a new LeasesReadLeaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeasesReadLeaseResponseWithDefaults

`func NewLeasesReadLeaseResponseWithDefaults() *LeasesReadLeaseResponse`

NewLeasesReadLeaseResponseWithDefaults instantiates a new LeasesReadLeaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetExpireTime

`func (o *LeasesReadLeaseResponse) GetExpireTime() time.Time`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *LeasesReadLeaseResponse) GetExpireTimeOk() (*time.Time, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *LeasesReadLeaseResponse) SetExpireTime(v time.Time)`

SetExpireTime sets ExpireTime field to given value.


### HasExpireTime

`func (o *LeasesReadLeaseResponse) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.




### GetId

`func (o *LeasesReadLeaseResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LeasesReadLeaseResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LeasesReadLeaseResponse) SetId(v string)`

SetId sets Id field to given value.


### HasId

`func (o *LeasesReadLeaseResponse) HasId() bool`

HasId returns a boolean if a field has been set.




### GetIssueTime

`func (o *LeasesReadLeaseResponse) GetIssueTime() time.Time`

GetIssueTime returns the IssueTime field if non-nil, zero value otherwise.

### GetIssueTimeOk

`func (o *LeasesReadLeaseResponse) GetIssueTimeOk() (*time.Time, bool)`

GetIssueTimeOk returns a tuple with the IssueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTime

`func (o *LeasesReadLeaseResponse) SetIssueTime(v time.Time)`

SetIssueTime sets IssueTime field to given value.


### HasIssueTime

`func (o *LeasesReadLeaseResponse) HasIssueTime() bool`

HasIssueTime returns a boolean if a field has been set.




### GetLastRenewal

`func (o *LeasesReadLeaseResponse) GetLastRenewal() time.Time`

GetLastRenewal returns the LastRenewal field if non-nil, zero value otherwise.

### GetLastRenewalOk

`func (o *LeasesReadLeaseResponse) GetLastRenewalOk() (*time.Time, bool)`

GetLastRenewalOk returns a tuple with the LastRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRenewal

`func (o *LeasesReadLeaseResponse) SetLastRenewal(v time.Time)`

SetLastRenewal sets LastRenewal field to given value.


### HasLastRenewal

`func (o *LeasesReadLeaseResponse) HasLastRenewal() bool`

HasLastRenewal returns a boolean if a field has been set.




### GetRenewable

`func (o *LeasesReadLeaseResponse) GetRenewable() bool`

GetRenewable returns the Renewable field if non-nil, zero value otherwise.

### GetRenewableOk

`func (o *LeasesReadLeaseResponse) GetRenewableOk() (*bool, bool)`

GetRenewableOk returns a tuple with the Renewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewable

`func (o *LeasesReadLeaseResponse) SetRenewable(v bool)`

SetRenewable sets Renewable field to given value.


### HasRenewable

`func (o *LeasesReadLeaseResponse) HasRenewable() bool`

HasRenewable returns a boolean if a field has been set.




### GetTtl

`func (o *LeasesReadLeaseResponse) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *LeasesReadLeaseResponse) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *LeasesReadLeaseResponse) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *LeasesReadLeaseResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


