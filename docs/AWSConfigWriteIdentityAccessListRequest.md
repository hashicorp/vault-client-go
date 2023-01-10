# AWSConfigWriteIdentityAccessListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisablePeriodicTidy** | Pointer to **bool** | If set to &#39;true&#39;, disables the periodic tidying of the &#39;identity-accesslist/&lt;instance_id&gt;&#39; entries. | [optional] [default to false]
**SafetyBuffer** | Pointer to **int32** | The amount of extra time that must have passed beyond the identity&#39;s expiration, before it is removed from the backend storage. | [optional] [default to 259200]

## Methods

### NewAWSConfigWriteIdentityAccessListRequest

`func NewAWSConfigWriteIdentityAccessListRequest() *AWSConfigWriteIdentityAccessListRequest`

NewAWSConfigWriteIdentityAccessListRequest instantiates a new AWSConfigWriteIdentityAccessListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSConfigWriteIdentityAccessListRequestWithDefaults

`func NewAWSConfigWriteIdentityAccessListRequestWithDefaults() *AWSConfigWriteIdentityAccessListRequest`

NewAWSConfigWriteIdentityAccessListRequestWithDefaults instantiates a new AWSConfigWriteIdentityAccessListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisablePeriodicTidy

`func (o *AWSConfigWriteIdentityAccessListRequest) GetDisablePeriodicTidy() bool`

GetDisablePeriodicTidy returns the DisablePeriodicTidy field if non-nil, zero value otherwise.

### GetDisablePeriodicTidyOk

`func (o *AWSConfigWriteIdentityAccessListRequest) GetDisablePeriodicTidyOk() (*bool, bool)`

GetDisablePeriodicTidyOk returns a tuple with the DisablePeriodicTidy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePeriodicTidy

`func (o *AWSConfigWriteIdentityAccessListRequest) SetDisablePeriodicTidy(v bool)`

SetDisablePeriodicTidy sets DisablePeriodicTidy field to given value.

### HasDisablePeriodicTidy

`func (o *AWSConfigWriteIdentityAccessListRequest) HasDisablePeriodicTidy() bool`

HasDisablePeriodicTidy returns a boolean if a field has been set.

### GetSafetyBuffer

`func (o *AWSConfigWriteIdentityAccessListRequest) GetSafetyBuffer() int32`

GetSafetyBuffer returns the SafetyBuffer field if non-nil, zero value otherwise.

### GetSafetyBufferOk

`func (o *AWSConfigWriteIdentityAccessListRequest) GetSafetyBufferOk() (*int32, bool)`

GetSafetyBufferOk returns a tuple with the SafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyBuffer

`func (o *AWSConfigWriteIdentityAccessListRequest) SetSafetyBuffer(v int32)`

SetSafetyBuffer sets SafetyBuffer field to given value.

### HasSafetyBuffer

`func (o *AWSConfigWriteIdentityAccessListRequest) HasSafetyBuffer() bool`

HasSafetyBuffer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


