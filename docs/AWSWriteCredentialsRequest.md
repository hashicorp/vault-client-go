# AWSWriteCredentialsRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Name** | Pointer to **string** | Name of the role | [optional] 
**RoleArn** | Pointer to **string** | ARN of role to assume when credential_type is assumed_role | [optional] 
**RoleSessionName** | Pointer to **string** | Session name to use when assuming role. Max chars: 64 | [optional] 
**Ttl** | Pointer to **int32** | Lifetime of the returned credentials in seconds | [optional] [default to 3600]



## Methods


### NewAWSWriteCredentialsRequest

`func NewAWSWriteCredentialsRequest() *AWSWriteCredentialsRequest`

NewAWSWriteCredentialsRequest instantiates a new AWSWriteCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSWriteCredentialsRequestWithDefaults

`func NewAWSWriteCredentialsRequestWithDefaults() *AWSWriteCredentialsRequest`

NewAWSWriteCredentialsRequestWithDefaults instantiates a new AWSWriteCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetName

`func (o *AWSWriteCredentialsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AWSWriteCredentialsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AWSWriteCredentialsRequest) SetName(v string)`

SetName sets Name field to given value.


### HasName

`func (o *AWSWriteCredentialsRequest) HasName() bool`

HasName returns a boolean if a field has been set.




### GetRoleArn

`func (o *AWSWriteCredentialsRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AWSWriteCredentialsRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AWSWriteCredentialsRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### HasRoleArn

`func (o *AWSWriteCredentialsRequest) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.




### GetRoleSessionName

`func (o *AWSWriteCredentialsRequest) GetRoleSessionName() string`

GetRoleSessionName returns the RoleSessionName field if non-nil, zero value otherwise.

### GetRoleSessionNameOk

`func (o *AWSWriteCredentialsRequest) GetRoleSessionNameOk() (*string, bool)`

GetRoleSessionNameOk returns a tuple with the RoleSessionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleSessionName

`func (o *AWSWriteCredentialsRequest) SetRoleSessionName(v string)`

SetRoleSessionName sets RoleSessionName field to given value.


### HasRoleSessionName

`func (o *AWSWriteCredentialsRequest) HasRoleSessionName() bool`

HasRoleSessionName returns a boolean if a field has been set.




### GetTtl

`func (o *AWSWriteCredentialsRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AWSWriteCredentialsRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AWSWriteCredentialsRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *AWSWriteCredentialsRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


