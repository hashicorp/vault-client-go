# AwsWriteStaticRolesNameRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RotationPeriod** | Pointer to **string** | Period by which to rotate the backing credential of the adopted user. This can be a Go duration (e.g, &#x27;1m&#x27;, 24h&#x27;), or an integer number of seconds. | [optional] 
**Username** | Pointer to **string** | The IAM user to adopt as a static role. | [optional] 



## Methods


### NewAwsWriteStaticRolesNameRequest

`func NewAwsWriteStaticRolesNameRequest() *AwsWriteStaticRolesNameRequest`

NewAwsWriteStaticRolesNameRequest instantiates a new AwsWriteStaticRolesNameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsWriteStaticRolesNameRequestWithDefaults

`func NewAwsWriteStaticRolesNameRequestWithDefaults() *AwsWriteStaticRolesNameRequest`

NewAwsWriteStaticRolesNameRequestWithDefaults instantiates a new AwsWriteStaticRolesNameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetRotationPeriod

`func (o *AwsWriteStaticRolesNameRequest) GetRotationPeriod() string`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *AwsWriteStaticRolesNameRequest) GetRotationPeriodOk() (*string, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *AwsWriteStaticRolesNameRequest) SetRotationPeriod(v string)`

SetRotationPeriod sets RotationPeriod field to given value.


### HasRotationPeriod

`func (o *AwsWriteStaticRolesNameRequest) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.




### GetUsername

`func (o *AwsWriteStaticRolesNameRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AwsWriteStaticRolesNameRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AwsWriteStaticRolesNameRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### HasUsername

`func (o *AwsWriteStaticRolesNameRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


