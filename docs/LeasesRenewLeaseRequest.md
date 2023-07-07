# LeasesRenewLeaseRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Increment** | Pointer to **string** | The desired increment in seconds to the lease | [optional] 
**LeaseId** | Pointer to **string** | The lease identifier to renew. This is included with a lease. | [optional] 
**UrlLeaseId** | Pointer to **string** | The lease identifier to renew. This is included with a lease. | [optional] 



## Methods


### NewLeasesRenewLeaseRequest

`func NewLeasesRenewLeaseRequest() *LeasesRenewLeaseRequest`

NewLeasesRenewLeaseRequest instantiates a new LeasesRenewLeaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeasesRenewLeaseRequestWithDefaults

`func NewLeasesRenewLeaseRequestWithDefaults() *LeasesRenewLeaseRequest`

NewLeasesRenewLeaseRequestWithDefaults instantiates a new LeasesRenewLeaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetIncrement

`func (o *LeasesRenewLeaseRequest) GetIncrement() string`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *LeasesRenewLeaseRequest) GetIncrementOk() (*string, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *LeasesRenewLeaseRequest) SetIncrement(v string)`

SetIncrement sets Increment field to given value.


### HasIncrement

`func (o *LeasesRenewLeaseRequest) HasIncrement() bool`

HasIncrement returns a boolean if a field has been set.




### GetLeaseId

`func (o *LeasesRenewLeaseRequest) GetLeaseId() string`

GetLeaseId returns the LeaseId field if non-nil, zero value otherwise.

### GetLeaseIdOk

`func (o *LeasesRenewLeaseRequest) GetLeaseIdOk() (*string, bool)`

GetLeaseIdOk returns a tuple with the LeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseId

`func (o *LeasesRenewLeaseRequest) SetLeaseId(v string)`

SetLeaseId sets LeaseId field to given value.


### HasLeaseId

`func (o *LeasesRenewLeaseRequest) HasLeaseId() bool`

HasLeaseId returns a boolean if a field has been set.




### GetUrlLeaseId

`func (o *LeasesRenewLeaseRequest) GetUrlLeaseId() string`

GetUrlLeaseId returns the UrlLeaseId field if non-nil, zero value otherwise.

### GetUrlLeaseIdOk

`func (o *LeasesRenewLeaseRequest) GetUrlLeaseIdOk() (*string, bool)`

GetUrlLeaseIdOk returns a tuple with the UrlLeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlLeaseId

`func (o *LeasesRenewLeaseRequest) SetUrlLeaseId(v string)`

SetUrlLeaseId sets UrlLeaseId field to given value.


### HasUrlLeaseId

`func (o *LeasesRenewLeaseRequest) HasUrlLeaseId() bool`

HasUrlLeaseId returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


