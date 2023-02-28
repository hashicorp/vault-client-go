# RenewRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Increment** | Pointer to **int32** | The desired increment in seconds to the lease | [optional] 
**LeaseId** | Pointer to **string** | The lease identifier to renew. This is included with a lease. | [optional] 
**UrlLeaseId** | Pointer to **string** | The lease identifier to renew. This is included with a lease. | [optional] 



## Methods


### NewRenewRequest

`func NewRenewRequest() *RenewRequest`

NewRenewRequest instantiates a new RenewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewRequestWithDefaults

`func NewRenewRequestWithDefaults() *RenewRequest`

NewRenewRequestWithDefaults instantiates a new RenewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetIncrement

`func (o *RenewRequest) GetIncrement() int32`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *RenewRequest) GetIncrementOk() (*int32, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *RenewRequest) SetIncrement(v int32)`

SetIncrement sets Increment field to given value.


### HasIncrement

`func (o *RenewRequest) HasIncrement() bool`

HasIncrement returns a boolean if a field has been set.




### GetLeaseId

`func (o *RenewRequest) GetLeaseId() string`

GetLeaseId returns the LeaseId field if non-nil, zero value otherwise.

### GetLeaseIdOk

`func (o *RenewRequest) GetLeaseIdOk() (*string, bool)`

GetLeaseIdOk returns a tuple with the LeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseId

`func (o *RenewRequest) SetLeaseId(v string)`

SetLeaseId sets LeaseId field to given value.


### HasLeaseId

`func (o *RenewRequest) HasLeaseId() bool`

HasLeaseId returns a boolean if a field has been set.




### GetUrlLeaseId

`func (o *RenewRequest) GetUrlLeaseId() string`

GetUrlLeaseId returns the UrlLeaseId field if non-nil, zero value otherwise.

### GetUrlLeaseIdOk

`func (o *RenewRequest) GetUrlLeaseIdOk() (*string, bool)`

GetUrlLeaseIdOk returns a tuple with the UrlLeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlLeaseId

`func (o *RenewRequest) SetUrlLeaseId(v string)`

SetUrlLeaseId sets UrlLeaseId field to given value.


### HasUrlLeaseId

`func (o *RenewRequest) HasUrlLeaseId() bool`

HasUrlLeaseId returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


