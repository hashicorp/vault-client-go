# ApproleRoleCustomSecretIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrList** | Pointer to **[]string** | Comma separated string or list of CIDR blocks enforcing secret IDs to be used from specific set of IP addresses. If &#39;bound_cidr_list&#39; is set on the role, then the list of CIDR blocks listed here should be a subset of the CIDR blocks listed on the role. | [optional] 
**Metadata** | Pointer to **string** | Metadata to be tied to the SecretID. This should be a JSON formatted string containing metadata in key value pairs. | [optional] 
**SecretId** | Pointer to **string** | SecretID to be attached to the role. | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token. Should be a subset of the token CIDR blocks listed on the role, if any. | [optional] 

## Methods

### NewApproleRoleCustomSecretIdRequest

`func NewApproleRoleCustomSecretIdRequest() *ApproleRoleCustomSecretIdRequest`

NewApproleRoleCustomSecretIdRequest instantiates a new ApproleRoleCustomSecretIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproleRoleCustomSecretIdRequestWithDefaults

`func NewApproleRoleCustomSecretIdRequestWithDefaults() *ApproleRoleCustomSecretIdRequest`

NewApproleRoleCustomSecretIdRequestWithDefaults instantiates a new ApproleRoleCustomSecretIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrList

`func (o *ApproleRoleCustomSecretIdRequest) GetCidrList() []string`

GetCidrList returns the CidrList field if non-nil, zero value otherwise.

### GetCidrListOk

`func (o *ApproleRoleCustomSecretIdRequest) GetCidrListOk() (*[]string, bool)`

GetCidrListOk returns a tuple with the CidrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrList

`func (o *ApproleRoleCustomSecretIdRequest) SetCidrList(v []string)`

SetCidrList sets CidrList field to given value.

### HasCidrList

`func (o *ApproleRoleCustomSecretIdRequest) HasCidrList() bool`

HasCidrList returns a boolean if a field has been set.

### GetMetadata

`func (o *ApproleRoleCustomSecretIdRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApproleRoleCustomSecretIdRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApproleRoleCustomSecretIdRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApproleRoleCustomSecretIdRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSecretId

`func (o *ApproleRoleCustomSecretIdRequest) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *ApproleRoleCustomSecretIdRequest) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *ApproleRoleCustomSecretIdRequest) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.

### HasSecretId

`func (o *ApproleRoleCustomSecretIdRequest) HasSecretId() bool`

HasSecretId returns a boolean if a field has been set.

### GetTokenBoundCidrs

`func (o *ApproleRoleCustomSecretIdRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *ApproleRoleCustomSecretIdRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *ApproleRoleCustomSecretIdRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.

### HasTokenBoundCidrs

`func (o *ApproleRoleCustomSecretIdRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


