# AWSWriteRoleTagDenyListTidySettingsRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SafetyBuffer** | Pointer to **int32** | The amount of extra time that must have passed beyond the roletag expiration, before it is removed from the backend storage. | [optional] [default to 259200]



## Methods


### NewAWSWriteRoleTagDenyListTidySettingsRequest

`func NewAWSWriteRoleTagDenyListTidySettingsRequest() *AWSWriteRoleTagDenyListTidySettingsRequest`

NewAWSWriteRoleTagDenyListTidySettingsRequest instantiates a new AWSWriteRoleTagDenyListTidySettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSWriteRoleTagDenyListTidySettingsRequestWithDefaults

`func NewAWSWriteRoleTagDenyListTidySettingsRequestWithDefaults() *AWSWriteRoleTagDenyListTidySettingsRequest`

NewAWSWriteRoleTagDenyListTidySettingsRequestWithDefaults instantiates a new AWSWriteRoleTagDenyListTidySettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetSafetyBuffer

`func (o *AWSWriteRoleTagDenyListTidySettingsRequest) GetSafetyBuffer() int32`

GetSafetyBuffer returns the SafetyBuffer field if non-nil, zero value otherwise.

### GetSafetyBufferOk

`func (o *AWSWriteRoleTagDenyListTidySettingsRequest) GetSafetyBufferOk() (*int32, bool)`

GetSafetyBufferOk returns a tuple with the SafetyBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyBuffer

`func (o *AWSWriteRoleTagDenyListTidySettingsRequest) SetSafetyBuffer(v int32)`

SetSafetyBuffer sets SafetyBuffer field to given value.


### HasSafetyBuffer

`func (o *AWSWriteRoleTagDenyListTidySettingsRequest) HasSafetyBuffer() bool`

HasSafetyBuffer returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


