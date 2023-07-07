# LeasesRenewLeaseWithIdRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Increment** | Pointer to **string** | The desired increment in seconds to the lease | [optional] 
**LeaseId** | Pointer to **string** | The lease identifier to renew. This is included with a lease. | [optional] 



## Methods


### NewLeasesRenewLeaseWithIdRequest

`func NewLeasesRenewLeaseWithIdRequest() *LeasesRenewLeaseWithIdRequest`

NewLeasesRenewLeaseWithIdRequest instantiates a new LeasesRenewLeaseWithIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeasesRenewLeaseWithIdRequestWithDefaults

`func NewLeasesRenewLeaseWithIdRequestWithDefaults() *LeasesRenewLeaseWithIdRequest`

NewLeasesRenewLeaseWithIdRequestWithDefaults instantiates a new LeasesRenewLeaseWithIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetIncrement

`func (o *LeasesRenewLeaseWithIdRequest) GetIncrement() string`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *LeasesRenewLeaseWithIdRequest) GetIncrementOk() (*string, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *LeasesRenewLeaseWithIdRequest) SetIncrement(v string)`

SetIncrement sets Increment field to given value.


### HasIncrement

`func (o *LeasesRenewLeaseWithIdRequest) HasIncrement() bool`

HasIncrement returns a boolean if a field has been set.




### GetLeaseId

`func (o *LeasesRenewLeaseWithIdRequest) GetLeaseId() string`

GetLeaseId returns the LeaseId field if non-nil, zero value otherwise.

### GetLeaseIdOk

`func (o *LeasesRenewLeaseWithIdRequest) GetLeaseIdOk() (*string, bool)`

GetLeaseIdOk returns a tuple with the LeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseId

`func (o *LeasesRenewLeaseWithIdRequest) SetLeaseId(v string)`

SetLeaseId sets LeaseId field to given value.


### HasLeaseId

`func (o *LeasesRenewLeaseWithIdRequest) HasLeaseId() bool`

HasLeaseId returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


