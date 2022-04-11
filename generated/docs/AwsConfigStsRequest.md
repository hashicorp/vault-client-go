# AwsConfigStsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StsRole** | Pointer to **string** | AWS ARN for STS role to be assumed when interacting with the account specified. The Vault server must have permissions to assume this role. | [optional] 

## Methods

### NewAwsConfigStsRequest

`func NewAwsConfigStsRequest() *AwsConfigStsRequest`

NewAwsConfigStsRequest instantiates a new AwsConfigStsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsConfigStsRequestWithDefaults

`func NewAwsConfigStsRequestWithDefaults() *AwsConfigStsRequest`

NewAwsConfigStsRequestWithDefaults instantiates a new AwsConfigStsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStsRole

`func (o *AwsConfigStsRequest) GetStsRole() string`

GetStsRole returns the StsRole field if non-nil, zero value otherwise.

### GetStsRoleOk

`func (o *AwsConfigStsRequest) GetStsRoleOk() (*string, bool)`

GetStsRoleOk returns a tuple with the StsRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsRole

`func (o *AwsConfigStsRequest) SetStsRole(v string)`

SetStsRole sets StsRole field to given value.

### HasStsRole

`func (o *AwsConfigStsRequest) HasStsRole() bool`

HasStsRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


