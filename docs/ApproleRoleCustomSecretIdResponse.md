# ApproleRoleCustomSecretIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretId** | Pointer to **string** | Secret ID attached to the role. | [optional] 
**SecretIdAccessor** | Pointer to **string** | Accessor of the secret ID | [optional] 
**SecretIdNumUses** | Pointer to **int32** | Number of times a secret ID can access the role, after which the secret ID will expire. | [optional] 
**SecretIdTtl** | Pointer to **int32** | Duration in seconds after which the issued secret ID expires. | [optional] 

## Methods

### NewApproleRoleCustomSecretIdResponse

`func NewApproleRoleCustomSecretIdResponse() *ApproleRoleCustomSecretIdResponse`

NewApproleRoleCustomSecretIdResponse instantiates a new ApproleRoleCustomSecretIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproleRoleCustomSecretIdResponseWithDefaults

`func NewApproleRoleCustomSecretIdResponseWithDefaults() *ApproleRoleCustomSecretIdResponse`

NewApproleRoleCustomSecretIdResponseWithDefaults instantiates a new ApproleRoleCustomSecretIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretId

`func (o *ApproleRoleCustomSecretIdResponse) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *ApproleRoleCustomSecretIdResponse) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *ApproleRoleCustomSecretIdResponse) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.

### HasSecretId

`func (o *ApproleRoleCustomSecretIdResponse) HasSecretId() bool`

HasSecretId returns a boolean if a field has been set.

### GetSecretIdAccessor

`func (o *ApproleRoleCustomSecretIdResponse) GetSecretIdAccessor() string`

GetSecretIdAccessor returns the SecretIdAccessor field if non-nil, zero value otherwise.

### GetSecretIdAccessorOk

`func (o *ApproleRoleCustomSecretIdResponse) GetSecretIdAccessorOk() (*string, bool)`

GetSecretIdAccessorOk returns a tuple with the SecretIdAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdAccessor

`func (o *ApproleRoleCustomSecretIdResponse) SetSecretIdAccessor(v string)`

SetSecretIdAccessor sets SecretIdAccessor field to given value.

### HasSecretIdAccessor

`func (o *ApproleRoleCustomSecretIdResponse) HasSecretIdAccessor() bool`

HasSecretIdAccessor returns a boolean if a field has been set.

### GetSecretIdNumUses

`func (o *ApproleRoleCustomSecretIdResponse) GetSecretIdNumUses() int32`

GetSecretIdNumUses returns the SecretIdNumUses field if non-nil, zero value otherwise.

### GetSecretIdNumUsesOk

`func (o *ApproleRoleCustomSecretIdResponse) GetSecretIdNumUsesOk() (*int32, bool)`

GetSecretIdNumUsesOk returns a tuple with the SecretIdNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdNumUses

`func (o *ApproleRoleCustomSecretIdResponse) SetSecretIdNumUses(v int32)`

SetSecretIdNumUses sets SecretIdNumUses field to given value.

### HasSecretIdNumUses

`func (o *ApproleRoleCustomSecretIdResponse) HasSecretIdNumUses() bool`

HasSecretIdNumUses returns a boolean if a field has been set.

### GetSecretIdTtl

`func (o *ApproleRoleCustomSecretIdResponse) GetSecretIdTtl() int32`

GetSecretIdTtl returns the SecretIdTtl field if non-nil, zero value otherwise.

### GetSecretIdTtlOk

`func (o *ApproleRoleCustomSecretIdResponse) GetSecretIdTtlOk() (*int32, bool)`

GetSecretIdTtlOk returns a tuple with the SecretIdTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdTtl

`func (o *ApproleRoleCustomSecretIdResponse) SetSecretIdTtl(v int32)`

SetSecretIdTtl sets SecretIdTtl field to given value.

### HasSecretIdTtl

`func (o *ApproleRoleCustomSecretIdResponse) HasSecretIdTtl() bool`

HasSecretIdTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


