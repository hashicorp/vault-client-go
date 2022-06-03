# ApproleLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleId** | Pointer to **string** | Unique identifier of the Role. Required to be supplied when the &#39;bind_secret_id&#39; constraint is set. | [optional] 
**SecretId** | Pointer to **string** | SecretID belong to the App role | [optional] [default to ""]

## Methods

### NewApproleLoginRequest

`func NewApproleLoginRequest() *ApproleLoginRequest`

NewApproleLoginRequest instantiates a new ApproleLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproleLoginRequestWithDefaults

`func NewApproleLoginRequestWithDefaults() *ApproleLoginRequest`

NewApproleLoginRequestWithDefaults instantiates a new ApproleLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleId

`func (o *ApproleLoginRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ApproleLoginRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ApproleLoginRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *ApproleLoginRequest) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetSecretId

`func (o *ApproleLoginRequest) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *ApproleLoginRequest) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *ApproleLoginRequest) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.

### HasSecretId

`func (o *ApproleLoginRequest) HasSecretId() bool`

HasSecretId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


