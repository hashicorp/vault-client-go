# ApproleRoleSecretIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrList** | Pointer to **[]string** | Comma separated string or list of CIDR blocks enforcing secret IDs to be used from specific set of IP addresses. If &#39;bound_cidr_list&#39; is set on the role, then the list of CIDR blocks listed here should be a subset of the CIDR blocks listed on the role. | [optional] 
**Metadata** | Pointer to **string** | Metadata to be tied to the SecretID. This should be a JSON formatted string containing the metadata in key value pairs. | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 

## Methods

### NewApproleRoleSecretIdRequest

`func NewApproleRoleSecretIdRequest() *ApproleRoleSecretIdRequest`

NewApproleRoleSecretIdRequest instantiates a new ApproleRoleSecretIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproleRoleSecretIdRequestWithDefaults

`func NewApproleRoleSecretIdRequestWithDefaults() *ApproleRoleSecretIdRequest`

NewApproleRoleSecretIdRequestWithDefaults instantiates a new ApproleRoleSecretIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrList

`func (o *ApproleRoleSecretIdRequest) GetCidrList() []string`

GetCidrList returns the CidrList field if non-nil, zero value otherwise.

### GetCidrListOk

`func (o *ApproleRoleSecretIdRequest) GetCidrListOk() (*[]string, bool)`

GetCidrListOk returns a tuple with the CidrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrList

`func (o *ApproleRoleSecretIdRequest) SetCidrList(v []string)`

SetCidrList sets CidrList field to given value.

### HasCidrList

`func (o *ApproleRoleSecretIdRequest) HasCidrList() bool`

HasCidrList returns a boolean if a field has been set.

### GetMetadata

`func (o *ApproleRoleSecretIdRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApproleRoleSecretIdRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApproleRoleSecretIdRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApproleRoleSecretIdRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTokenBoundCidrs

`func (o *ApproleRoleSecretIdRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *ApproleRoleSecretIdRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *ApproleRoleSecretIdRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.

### HasTokenBoundCidrs

`func (o *ApproleRoleSecretIdRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


