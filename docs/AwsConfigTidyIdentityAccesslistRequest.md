# AwsConfigTidyIdentityAccesslistRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisablePeriodicTidy** | Pointer to **bool** | If set to &#39;true&#39;, disables the periodic tidying of the &#39;identity-accesslist/&lt;instance_id&gt;&#39; entries. | [optional] [default to false]
**SafetyBuffer** | Pointer to **int32** | The amount of extra time that must have passed beyond the identity&#39;s expiration, before it is removed from the backend storage. | [optional] [default to 259200]

## Methods

### NewAwsConfigTidyIdentityAccesslistRequest

`func NewAwsConfigTidyIdentityAccesslistRequest() *AwsConfigTidyIdentityAccesslistRequest`

NewAwsConfigTidyIdentityAccesslistRequest instantiates a new AwsConfigTidyIdentityAccesslistRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsConfigTidyIdentityAccesslistRequestWithDefaults

`func NewAwsConfigTidyIdentityAccesslistRequestWithDefaults() *AwsConfigTidyIdentityAccesslistRequest`

NewAwsConfigTidyIdentityAccesslistRequestWithDefaults instantiates a new AwsConfigTidyIdentityAccesslistRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisablePeriodicTidy

`func (o *AwsConfigTidyIdentityAccesslistRequest) GetDisablePeriodicTidy() bool`

GetDisablePeriodicTidy returns the DisablePeriodicTidy field if non-nil, zero value otherwise.

### GetDisablePeriodicTidyOk

`func (o *AwsConfigTidyIdentityAccesslistRequest) GetDisablePeriodicTidyOk() (*bool, bool)`

GetDisablePeriodicTidyOk returns a tuple with the DisablePeriodicTidy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePeriodicTidy

`func (o *AwsConfigTidyIdentityAccesslistRequest) SetDisablePeriodicTidy(v bool)`

SetDisablePeriodicTidy sets DisablePeriodicTidy field to given value.

### HasDisablePeriodicTidy

`func (o *AwsConfigTidyIdentityAccesslistRequest) HasDisablePeriodicTidy() bool`

HasDisablePeriodicTidy returns a boolean if a field has been set.

### GetSafetyBuffer

`func (o *AwsConfigTidyIdentityAccesslistRequest) GetSafetyBuffer() int32`

GetSafetyBuffer returns the SafetyBuffer field if non-nil, zero value otherwise.

### GetSafetyBufferOk

`func (o *AwsConfigTidyIdentityAccesslistRequest) GetSafetyBufferOk() (*int32, bool)`

GetSafetyBufferOk returns a tuple with the SafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyBuffer

`func (o *AwsConfigTidyIdentityAccesslistRequest) SetSafetyBuffer(v int32)`

SetSafetyBuffer sets SafetyBuffer field to given value.

### HasSafetyBuffer

`func (o *AwsConfigTidyIdentityAccesslistRequest) HasSafetyBuffer() bool`

HasSafetyBuffer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


