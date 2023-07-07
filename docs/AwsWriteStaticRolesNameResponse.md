# AwsWriteStaticRolesNameResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this role. | [optional] 
**RotationPeriod** | Pointer to **string** | Period by which to rotate the backing credential of the adopted user. This can be a Go duration (e.g, &#x27;1m&#x27;, 24h&#x27;), or an integer number of seconds. | [optional] 
**Username** | Pointer to **string** | The IAM user to adopt as a static role. | [optional] 



## Methods


### NewAwsWriteStaticRolesNameResponse

`func NewAwsWriteStaticRolesNameResponse() *AwsWriteStaticRolesNameResponse`

NewAwsWriteStaticRolesNameResponse instantiates a new AwsWriteStaticRolesNameResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsWriteStaticRolesNameResponseWithDefaults

`func NewAwsWriteStaticRolesNameResponseWithDefaults() *AwsWriteStaticRolesNameResponse`

NewAwsWriteStaticRolesNameResponseWithDefaults instantiates a new AwsWriteStaticRolesNameResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetName

`func (o *AwsWriteStaticRolesNameResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsWriteStaticRolesNameResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsWriteStaticRolesNameResponse) SetName(v string)`

SetName sets Name field to given value.


### HasName

`func (o *AwsWriteStaticRolesNameResponse) HasName() bool`

HasName returns a boolean if a field has been set.




### GetRotationPeriod

`func (o *AwsWriteStaticRolesNameResponse) GetRotationPeriod() string`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *AwsWriteStaticRolesNameResponse) GetRotationPeriodOk() (*string, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *AwsWriteStaticRolesNameResponse) SetRotationPeriod(v string)`

SetRotationPeriod sets RotationPeriod field to given value.


### HasRotationPeriod

`func (o *AwsWriteStaticRolesNameResponse) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.




### GetUsername

`func (o *AwsWriteStaticRolesNameResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AwsWriteStaticRolesNameResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AwsWriteStaticRolesNameResponse) SetUsername(v string)`

SetUsername sets Username field to given value.


### HasUsername

`func (o *AwsWriteStaticRolesNameResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


