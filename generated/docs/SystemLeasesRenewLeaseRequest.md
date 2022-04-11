# SystemLeasesRenewLeaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Increment** | Pointer to **int32** | The desired increment in seconds to the lease | [optional] 
**LeaseId** | Pointer to **string** | The lease identifier to renew. This is included with a lease. | [optional] 

## Methods

### NewSystemLeasesRenewLeaseRequest

`func NewSystemLeasesRenewLeaseRequest() *SystemLeasesRenewLeaseRequest`

NewSystemLeasesRenewLeaseRequest instantiates a new SystemLeasesRenewLeaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemLeasesRenewLeaseRequestWithDefaults

`func NewSystemLeasesRenewLeaseRequestWithDefaults() *SystemLeasesRenewLeaseRequest`

NewSystemLeasesRenewLeaseRequestWithDefaults instantiates a new SystemLeasesRenewLeaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncrement

`func (o *SystemLeasesRenewLeaseRequest) GetIncrement() int32`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *SystemLeasesRenewLeaseRequest) GetIncrementOk() (*int32, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *SystemLeasesRenewLeaseRequest) SetIncrement(v int32)`

SetIncrement sets Increment field to given value.

### HasIncrement

`func (o *SystemLeasesRenewLeaseRequest) HasIncrement() bool`

HasIncrement returns a boolean if a field has been set.

### GetLeaseId

`func (o *SystemLeasesRenewLeaseRequest) GetLeaseId() string`

GetLeaseId returns the LeaseId field if non-nil, zero value otherwise.

### GetLeaseIdOk

`func (o *SystemLeasesRenewLeaseRequest) GetLeaseIdOk() (*string, bool)`

GetLeaseIdOk returns a tuple with the LeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseId

`func (o *SystemLeasesRenewLeaseRequest) SetLeaseId(v string)`

SetLeaseId sets LeaseId field to given value.

### HasLeaseId

`func (o *SystemLeasesRenewLeaseRequest) HasLeaseId() bool`

HasLeaseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


