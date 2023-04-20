# AwsConfigureIdentityAccessListTidyOperationRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisablePeriodicTidy** | Pointer to **bool** | If set to &#x27;true&#x27;, disables the periodic tidying of the &#x27;identity-accesslist/&lt;instance_id&gt;&#x27; entries. | [optional] [default to false]
**SafetyBuffer** | Pointer to **int32** | The amount of extra time that must have passed beyond the identity&#x27;s expiration, before it is removed from the backend storage. | [optional] [default to 259200]



## Methods


### NewAwsConfigureIdentityAccessListTidyOperationRequest

`func NewAwsConfigureIdentityAccessListTidyOperationRequest() *AwsConfigureIdentityAccessListTidyOperationRequest`

NewAwsConfigureIdentityAccessListTidyOperationRequest instantiates a new AwsConfigureIdentityAccessListTidyOperationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsConfigureIdentityAccessListTidyOperationRequestWithDefaults

`func NewAwsConfigureIdentityAccessListTidyOperationRequestWithDefaults() *AwsConfigureIdentityAccessListTidyOperationRequest`

NewAwsConfigureIdentityAccessListTidyOperationRequestWithDefaults instantiates a new AwsConfigureIdentityAccessListTidyOperationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetDisablePeriodicTidy

`func (o *AwsConfigureIdentityAccessListTidyOperationRequest) GetDisablePeriodicTidy() bool`

GetDisablePeriodicTidy returns the DisablePeriodicTidy field if non-nil, zero value otherwise.

### GetDisablePeriodicTidyOk

`func (o *AwsConfigureIdentityAccessListTidyOperationRequest) GetDisablePeriodicTidyOk() (*bool, bool)`

GetDisablePeriodicTidyOk returns a tuple with the DisablePeriodicTidy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePeriodicTidy

`func (o *AwsConfigureIdentityAccessListTidyOperationRequest) SetDisablePeriodicTidy(v bool)`

SetDisablePeriodicTidy sets DisablePeriodicTidy field to given value.


### HasDisablePeriodicTidy

`func (o *AwsConfigureIdentityAccessListTidyOperationRequest) HasDisablePeriodicTidy() bool`

HasDisablePeriodicTidy returns a boolean if a field has been set.




### GetSafetyBuffer

`func (o *AwsConfigureIdentityAccessListTidyOperationRequest) GetSafetyBuffer() int32`

GetSafetyBuffer returns the SafetyBuffer field if non-nil, zero value otherwise.

### GetSafetyBufferOk

`func (o *AwsConfigureIdentityAccessListTidyOperationRequest) GetSafetyBufferOk() (*int32, bool)`

GetSafetyBufferOk returns a tuple with the SafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyBuffer

`func (o *AwsConfigureIdentityAccessListTidyOperationRequest) SetSafetyBuffer(v int32)`

SetSafetyBuffer sets SafetyBuffer field to given value.


### HasSafetyBuffer

`func (o *AwsConfigureIdentityAccessListTidyOperationRequest) HasSafetyBuffer() bool`

HasSafetyBuffer returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


