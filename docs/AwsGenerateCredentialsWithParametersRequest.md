# AwsGenerateCredentialsWithParametersRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleArn** | Pointer to **string** | ARN of role to assume when credential_type is assumed_role | [optional] 
**RoleSessionName** | Pointer to **string** | Session name to use when assuming role. Max chars: 64 | [optional] 
**Ttl** | Pointer to **string** | Lifetime of the returned credentials in seconds | [optional] [default to "3600"]



## Methods


### NewAwsGenerateCredentialsWithParametersRequest

`func NewAwsGenerateCredentialsWithParametersRequest() *AwsGenerateCredentialsWithParametersRequest`

NewAwsGenerateCredentialsWithParametersRequest instantiates a new AwsGenerateCredentialsWithParametersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsGenerateCredentialsWithParametersRequestWithDefaults

`func NewAwsGenerateCredentialsWithParametersRequestWithDefaults() *AwsGenerateCredentialsWithParametersRequest`

NewAwsGenerateCredentialsWithParametersRequestWithDefaults instantiates a new AwsGenerateCredentialsWithParametersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetRoleArn

`func (o *AwsGenerateCredentialsWithParametersRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsGenerateCredentialsWithParametersRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsGenerateCredentialsWithParametersRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### HasRoleArn

`func (o *AwsGenerateCredentialsWithParametersRequest) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.




### GetRoleSessionName

`func (o *AwsGenerateCredentialsWithParametersRequest) GetRoleSessionName() string`

GetRoleSessionName returns the RoleSessionName field if non-nil, zero value otherwise.

### GetRoleSessionNameOk

`func (o *AwsGenerateCredentialsWithParametersRequest) GetRoleSessionNameOk() (*string, bool)`

GetRoleSessionNameOk returns a tuple with the RoleSessionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleSessionName

`func (o *AwsGenerateCredentialsWithParametersRequest) SetRoleSessionName(v string)`

SetRoleSessionName sets RoleSessionName field to given value.


### HasRoleSessionName

`func (o *AwsGenerateCredentialsWithParametersRequest) HasRoleSessionName() bool`

HasRoleSessionName returns a boolean if a field has been set.




### GetTtl

`func (o *AwsGenerateCredentialsWithParametersRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AwsGenerateCredentialsWithParametersRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AwsGenerateCredentialsWithParametersRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *AwsGenerateCredentialsWithParametersRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


