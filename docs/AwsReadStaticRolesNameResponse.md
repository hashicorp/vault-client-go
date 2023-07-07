# AwsReadStaticRolesNameResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this role. | [optional] 
**RotationPeriod** | Pointer to **string** | Period by which to rotate the backing credential of the adopted user. This can be a Go duration (e.g, &#x27;1m&#x27;, 24h&#x27;), or an integer number of seconds. | [optional] 
**Username** | Pointer to **string** | The IAM user to adopt as a static role. | [optional] 



## Methods


### NewAwsReadStaticRolesNameResponse

`func NewAwsReadStaticRolesNameResponse() *AwsReadStaticRolesNameResponse`

NewAwsReadStaticRolesNameResponse instantiates a new AwsReadStaticRolesNameResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsReadStaticRolesNameResponseWithDefaults

`func NewAwsReadStaticRolesNameResponseWithDefaults() *AwsReadStaticRolesNameResponse`

NewAwsReadStaticRolesNameResponseWithDefaults instantiates a new AwsReadStaticRolesNameResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetName

`func (o *AwsReadStaticRolesNameResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsReadStaticRolesNameResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsReadStaticRolesNameResponse) SetName(v string)`

SetName sets Name field to given value.


### HasName

`func (o *AwsReadStaticRolesNameResponse) HasName() bool`

HasName returns a boolean if a field has been set.




### GetRotationPeriod

`func (o *AwsReadStaticRolesNameResponse) GetRotationPeriod() string`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *AwsReadStaticRolesNameResponse) GetRotationPeriodOk() (*string, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *AwsReadStaticRolesNameResponse) SetRotationPeriod(v string)`

SetRotationPeriod sets RotationPeriod field to given value.


### HasRotationPeriod

`func (o *AwsReadStaticRolesNameResponse) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.




### GetUsername

`func (o *AwsReadStaticRolesNameResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AwsReadStaticRolesNameResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AwsReadStaticRolesNameResponse) SetUsername(v string)`

SetUsername sets Username field to given value.


### HasUsername

`func (o *AwsReadStaticRolesNameResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


