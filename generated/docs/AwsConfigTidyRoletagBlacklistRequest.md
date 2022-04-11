# AwsConfigTidyRoletagBlacklistRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisablePeriodicTidy** | Pointer to **bool** | If set to &#39;true&#39;, disables the periodic tidying of deny listed entries. | [optional] [default to false]
**SafetyBuffer** | Pointer to **int32** | The amount of extra time that must have passed beyond the roletag expiration, before it is removed from the backend storage. Defaults to 4320h (180 days). | [optional] [default to 15552000]

## Methods

### NewAwsConfigTidyRoletagBlacklistRequest

`func NewAwsConfigTidyRoletagBlacklistRequest() *AwsConfigTidyRoletagBlacklistRequest`

NewAwsConfigTidyRoletagBlacklistRequest instantiates a new AwsConfigTidyRoletagBlacklistRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsConfigTidyRoletagBlacklistRequestWithDefaults

`func NewAwsConfigTidyRoletagBlacklistRequestWithDefaults() *AwsConfigTidyRoletagBlacklistRequest`

NewAwsConfigTidyRoletagBlacklistRequestWithDefaults instantiates a new AwsConfigTidyRoletagBlacklistRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisablePeriodicTidy

`func (o *AwsConfigTidyRoletagBlacklistRequest) GetDisablePeriodicTidy() bool`

GetDisablePeriodicTidy returns the DisablePeriodicTidy field if non-nil, zero value otherwise.

### GetDisablePeriodicTidyOk

`func (o *AwsConfigTidyRoletagBlacklistRequest) GetDisablePeriodicTidyOk() (*bool, bool)`

GetDisablePeriodicTidyOk returns a tuple with the DisablePeriodicTidy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePeriodicTidy

`func (o *AwsConfigTidyRoletagBlacklistRequest) SetDisablePeriodicTidy(v bool)`

SetDisablePeriodicTidy sets DisablePeriodicTidy field to given value.

### HasDisablePeriodicTidy

`func (o *AwsConfigTidyRoletagBlacklistRequest) HasDisablePeriodicTidy() bool`

HasDisablePeriodicTidy returns a boolean if a field has been set.

### GetSafetyBuffer

`func (o *AwsConfigTidyRoletagBlacklistRequest) GetSafetyBuffer() int32`

GetSafetyBuffer returns the SafetyBuffer field if non-nil, zero value otherwise.

### GetSafetyBufferOk

`func (o *AwsConfigTidyRoletagBlacklistRequest) GetSafetyBufferOk() (*int32, bool)`

GetSafetyBufferOk returns a tuple with the SafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyBuffer

`func (o *AwsConfigTidyRoletagBlacklistRequest) SetSafetyBuffer(v int32)`

SetSafetyBuffer sets SafetyBuffer field to given value.

### HasSafetyBuffer

`func (o *AwsConfigTidyRoletagBlacklistRequest) HasSafetyBuffer() bool`

HasSafetyBuffer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


