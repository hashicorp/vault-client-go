# AWSConfigWriteLeaseRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Lease** | Pointer to **string** | Default lease for roles. | [optional] 
**LeaseMax** | Pointer to **string** | Maximum time a credential is valid for. | [optional] 



## Methods


### NewAWSConfigWriteLeaseRequest

`func NewAWSConfigWriteLeaseRequest() *AWSConfigWriteLeaseRequest`

NewAWSConfigWriteLeaseRequest instantiates a new AWSConfigWriteLeaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSConfigWriteLeaseRequestWithDefaults

`func NewAWSConfigWriteLeaseRequestWithDefaults() *AWSConfigWriteLeaseRequest`

NewAWSConfigWriteLeaseRequestWithDefaults instantiates a new AWSConfigWriteLeaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetLease

`func (o *AWSConfigWriteLeaseRequest) GetLease() string`

GetLease returns the Lease field if non-nil, zero value otherwise.

### GetLeaseOk

`func (o *AWSConfigWriteLeaseRequest) GetLeaseOk() (*string, bool)`

GetLeaseOk returns a tuple with the Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease

`func (o *AWSConfigWriteLeaseRequest) SetLease(v string)`

SetLease sets Lease field to given value.


### HasLease

`func (o *AWSConfigWriteLeaseRequest) HasLease() bool`

HasLease returns a boolean if a field has been set.




### GetLeaseMax

`func (o *AWSConfigWriteLeaseRequest) GetLeaseMax() string`

GetLeaseMax returns the LeaseMax field if non-nil, zero value otherwise.

### GetLeaseMaxOk

`func (o *AWSConfigWriteLeaseRequest) GetLeaseMaxOk() (*string, bool)`

GetLeaseMaxOk returns a tuple with the LeaseMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseMax

`func (o *AWSConfigWriteLeaseRequest) SetLeaseMax(v string)`

SetLeaseMax sets LeaseMax field to given value.


### HasLeaseMax

`func (o *AWSConfigWriteLeaseRequest) HasLeaseMax() bool`

HasLeaseMax returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


