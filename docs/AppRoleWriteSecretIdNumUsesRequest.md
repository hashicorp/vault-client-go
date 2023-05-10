# AppRoleWriteSecretIdNumUsesRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretIdNumUses** | Pointer to **int32** | Number of times a SecretID can access the role, after which the SecretID will expire. | [optional] 



## Methods


### NewAppRoleWriteSecretIdNumUsesRequest

`func NewAppRoleWriteSecretIdNumUsesRequest() *AppRoleWriteSecretIdNumUsesRequest`

NewAppRoleWriteSecretIdNumUsesRequest instantiates a new AppRoleWriteSecretIdNumUsesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRoleWriteSecretIdNumUsesRequestWithDefaults

`func NewAppRoleWriteSecretIdNumUsesRequestWithDefaults() *AppRoleWriteSecretIdNumUsesRequest`

NewAppRoleWriteSecretIdNumUsesRequestWithDefaults instantiates a new AppRoleWriteSecretIdNumUsesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetSecretIdNumUses

`func (o *AppRoleWriteSecretIdNumUsesRequest) GetSecretIdNumUses() int32`

GetSecretIdNumUses returns the SecretIdNumUses field if non-nil, zero value otherwise.

### GetSecretIdNumUsesOk

`func (o *AppRoleWriteSecretIdNumUsesRequest) GetSecretIdNumUsesOk() (*int32, bool)`

GetSecretIdNumUsesOk returns a tuple with the SecretIdNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdNumUses

`func (o *AppRoleWriteSecretIdNumUsesRequest) SetSecretIdNumUses(v int32)`

SetSecretIdNumUses sets SecretIdNumUses field to given value.


### HasSecretIdNumUses

`func (o *AppRoleWriteSecretIdNumUsesRequest) HasSecretIdNumUses() bool`

HasSecretIdNumUses returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


