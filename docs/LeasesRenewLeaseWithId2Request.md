# LeasesRenewLeaseWithId2Request


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Increment** | Pointer to **int32** | The desired increment in seconds to the lease | [optional] 
**LeaseId** | Pointer to **string** | The lease identifier to renew. This is included with a lease. | [optional] 



## Methods


### NewLeasesRenewLeaseWithId2Request

`func NewLeasesRenewLeaseWithId2Request() *LeasesRenewLeaseWithId2Request`

NewLeasesRenewLeaseWithId2Request instantiates a new LeasesRenewLeaseWithId2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeasesRenewLeaseWithId2RequestWithDefaults

`func NewLeasesRenewLeaseWithId2RequestWithDefaults() *LeasesRenewLeaseWithId2Request`

NewLeasesRenewLeaseWithId2RequestWithDefaults instantiates a new LeasesRenewLeaseWithId2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetIncrement

`func (o *LeasesRenewLeaseWithId2Request) GetIncrement() int32`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *LeasesRenewLeaseWithId2Request) GetIncrementOk() (*int32, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *LeasesRenewLeaseWithId2Request) SetIncrement(v int32)`

SetIncrement sets Increment field to given value.


### HasIncrement

`func (o *LeasesRenewLeaseWithId2Request) HasIncrement() bool`

HasIncrement returns a boolean if a field has been set.




### GetLeaseId

`func (o *LeasesRenewLeaseWithId2Request) GetLeaseId() string`

GetLeaseId returns the LeaseId field if non-nil, zero value otherwise.

### GetLeaseIdOk

`func (o *LeasesRenewLeaseWithId2Request) GetLeaseIdOk() (*string, bool)`

GetLeaseIdOk returns a tuple with the LeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseId

`func (o *LeasesRenewLeaseWithId2Request) SetLeaseId(v string)`

SetLeaseId sets LeaseId field to given value.


### HasLeaseId

`func (o *LeasesRenewLeaseWithId2Request) HasLeaseId() bool`

HasLeaseId returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


