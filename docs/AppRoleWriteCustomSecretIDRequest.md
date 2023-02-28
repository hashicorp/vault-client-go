# AppRoleWriteCustomSecretIDRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrList** | Pointer to **[]string** | Comma separated string or list of CIDR blocks enforcing secret IDs to be used from specific set of IP addresses. If &#x27;bound_cidr_list&#x27; is set on the role, then the list of CIDR blocks listed here should be a subset of the CIDR blocks listed on the role. | [optional] 
**Metadata** | Pointer to **string** | Metadata to be tied to the SecretID. This should be a JSON formatted string containing metadata in key value pairs. | [optional] 
**NumUses** | Pointer to **int32** | Number of times this SecretID can be used, after which the SecretID expires. Overrides secret_id_num_uses role option when supplied. May not be higher than role&#x27;s secret_id_num_uses. | [optional] 
**SecretId** | Pointer to **string** | SecretID to be attached to the role. | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token. Should be a subset of the token CIDR blocks listed on the role, if any. | [optional] 
**Ttl** | Pointer to **int32** | Duration in seconds after which this SecretID expires. Overrides secret_id_ttl role option when supplied. May not be longer than role&#x27;s secret_id_ttl. | [optional] 



## Methods


### NewAppRoleWriteCustomSecretIDRequest

`func NewAppRoleWriteCustomSecretIDRequest() *AppRoleWriteCustomSecretIDRequest`

NewAppRoleWriteCustomSecretIDRequest instantiates a new AppRoleWriteCustomSecretIDRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRoleWriteCustomSecretIDRequestWithDefaults

`func NewAppRoleWriteCustomSecretIDRequestWithDefaults() *AppRoleWriteCustomSecretIDRequest`

NewAppRoleWriteCustomSecretIDRequestWithDefaults instantiates a new AppRoleWriteCustomSecretIDRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCidrList

`func (o *AppRoleWriteCustomSecretIDRequest) GetCidrList() []string`

GetCidrList returns the CidrList field if non-nil, zero value otherwise.

### GetCidrListOk

`func (o *AppRoleWriteCustomSecretIDRequest) GetCidrListOk() (*[]string, bool)`

GetCidrListOk returns a tuple with the CidrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrList

`func (o *AppRoleWriteCustomSecretIDRequest) SetCidrList(v []string)`

SetCidrList sets CidrList field to given value.


### HasCidrList

`func (o *AppRoleWriteCustomSecretIDRequest) HasCidrList() bool`

HasCidrList returns a boolean if a field has been set.




### GetMetadata

`func (o *AppRoleWriteCustomSecretIDRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AppRoleWriteCustomSecretIDRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AppRoleWriteCustomSecretIDRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.


### HasMetadata

`func (o *AppRoleWriteCustomSecretIDRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.




### GetNumUses

`func (o *AppRoleWriteCustomSecretIDRequest) GetNumUses() int32`

GetNumUses returns the NumUses field if non-nil, zero value otherwise.

### GetNumUsesOk

`func (o *AppRoleWriteCustomSecretIDRequest) GetNumUsesOk() (*int32, bool)`

GetNumUsesOk returns a tuple with the NumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUses

`func (o *AppRoleWriteCustomSecretIDRequest) SetNumUses(v int32)`

SetNumUses sets NumUses field to given value.


### HasNumUses

`func (o *AppRoleWriteCustomSecretIDRequest) HasNumUses() bool`

HasNumUses returns a boolean if a field has been set.




### GetSecretId

`func (o *AppRoleWriteCustomSecretIDRequest) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *AppRoleWriteCustomSecretIDRequest) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *AppRoleWriteCustomSecretIDRequest) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.


### HasSecretId

`func (o *AppRoleWriteCustomSecretIDRequest) HasSecretId() bool`

HasSecretId returns a boolean if a field has been set.




### GetTokenBoundCidrs

`func (o *AppRoleWriteCustomSecretIDRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *AppRoleWriteCustomSecretIDRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *AppRoleWriteCustomSecretIDRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.


### HasTokenBoundCidrs

`func (o *AppRoleWriteCustomSecretIDRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.




### GetTtl

`func (o *AppRoleWriteCustomSecretIDRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AppRoleWriteCustomSecretIDRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AppRoleWriteCustomSecretIDRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *AppRoleWriteCustomSecretIDRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


