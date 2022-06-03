# ApproleRoleSecretIdTtlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretIdTtl** | Pointer to **int32** | Duration in seconds after which the issued SecretID should expire. Defaults to 0, meaning no expiration. | [optional] 

## Methods

### NewApproleRoleSecretIdTtlRequest

`func NewApproleRoleSecretIdTtlRequest() *ApproleRoleSecretIdTtlRequest`

NewApproleRoleSecretIdTtlRequest instantiates a new ApproleRoleSecretIdTtlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproleRoleSecretIdTtlRequestWithDefaults

`func NewApproleRoleSecretIdTtlRequestWithDefaults() *ApproleRoleSecretIdTtlRequest`

NewApproleRoleSecretIdTtlRequestWithDefaults instantiates a new ApproleRoleSecretIdTtlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretIdTtl

`func (o *ApproleRoleSecretIdTtlRequest) GetSecretIdTtl() int32`

GetSecretIdTtl returns the SecretIdTtl field if non-nil, zero value otherwise.

### GetSecretIdTtlOk

`func (o *ApproleRoleSecretIdTtlRequest) GetSecretIdTtlOk() (*int32, bool)`

GetSecretIdTtlOk returns a tuple with the SecretIdTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdTtl

`func (o *ApproleRoleSecretIdTtlRequest) SetSecretIdTtl(v int32)`

SetSecretIdTtl sets SecretIdTtl field to given value.

### HasSecretIdTtl

`func (o *ApproleRoleSecretIdTtlRequest) HasSecretIdTtl() bool`

HasSecretIdTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


