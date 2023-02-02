# AppRoleWriteSecretIDLookupResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**CidrList** | Pointer to **[]string** | List of CIDR blocks enforcing secret IDs to be used from specific set of IP addresses. If &#x27;bound_cidr_list&#x27; is set on the role, then the list of CIDR blocks listed here should be a subset of the CIDR blocks listed on the role. | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**ExpirationTime** | Pointer to **time.Time** |  | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**SecretIdAccessor** | Pointer to **string** | Accessor of the secret ID | [optional] 
**SecretIdNumUses** | Pointer to **int32** | Number of times a secret ID can access the role, after which the secret ID will expire. | [optional] 
**SecretIdTtl** | Pointer to **int32** | Duration in seconds after which the issued secret ID expires. | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | List of CIDR blocks. If set, specifies the blocks of IP addresses which can use the returned token. Should be a subset of the token CIDR blocks listed on the role, if any. | [optional] 



## Methods


### NewAppRoleWriteSecretIDLookupResponse

`func NewAppRoleWriteSecretIDLookupResponse() *AppRoleWriteSecretIDLookupResponse`

NewAppRoleWriteSecretIDLookupResponse instantiates a new AppRoleWriteSecretIDLookupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRoleWriteSecretIDLookupResponseWithDefaults

`func NewAppRoleWriteSecretIDLookupResponseWithDefaults() *AppRoleWriteSecretIDLookupResponse`

NewAppRoleWriteSecretIDLookupResponseWithDefaults instantiates a new AppRoleWriteSecretIDLookupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCidrList

`func (o *AppRoleWriteSecretIDLookupResponse) GetCidrList() []string`

GetCidrList returns the CidrList field if non-nil, zero value otherwise.

### GetCidrListOk

`func (o *AppRoleWriteSecretIDLookupResponse) GetCidrListOk() (*[]string, bool)`

GetCidrListOk returns a tuple with the CidrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrList

`func (o *AppRoleWriteSecretIDLookupResponse) SetCidrList(v []string)`

SetCidrList sets CidrList field to given value.


### HasCidrList

`func (o *AppRoleWriteSecretIDLookupResponse) HasCidrList() bool`

HasCidrList returns a boolean if a field has been set.




### GetCreationTime

`func (o *AppRoleWriteSecretIDLookupResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AppRoleWriteSecretIDLookupResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AppRoleWriteSecretIDLookupResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### HasCreationTime

`func (o *AppRoleWriteSecretIDLookupResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.




### GetExpirationTime

`func (o *AppRoleWriteSecretIDLookupResponse) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *AppRoleWriteSecretIDLookupResponse) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *AppRoleWriteSecretIDLookupResponse) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.


### HasExpirationTime

`func (o *AppRoleWriteSecretIDLookupResponse) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.




### GetLastUpdatedTime

`func (o *AppRoleWriteSecretIDLookupResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AppRoleWriteSecretIDLookupResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AppRoleWriteSecretIDLookupResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### HasLastUpdatedTime

`func (o *AppRoleWriteSecretIDLookupResponse) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.




### GetMetadata

`func (o *AppRoleWriteSecretIDLookupResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AppRoleWriteSecretIDLookupResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AppRoleWriteSecretIDLookupResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### HasMetadata

`func (o *AppRoleWriteSecretIDLookupResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.




### GetSecretIdAccessor

`func (o *AppRoleWriteSecretIDLookupResponse) GetSecretIdAccessor() string`

GetSecretIdAccessor returns the SecretIdAccessor field if non-nil, zero value otherwise.

### GetSecretIdAccessorOk

`func (o *AppRoleWriteSecretIDLookupResponse) GetSecretIdAccessorOk() (*string, bool)`

GetSecretIdAccessorOk returns a tuple with the SecretIdAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdAccessor

`func (o *AppRoleWriteSecretIDLookupResponse) SetSecretIdAccessor(v string)`

SetSecretIdAccessor sets SecretIdAccessor field to given value.


### HasSecretIdAccessor

`func (o *AppRoleWriteSecretIDLookupResponse) HasSecretIdAccessor() bool`

HasSecretIdAccessor returns a boolean if a field has been set.




### GetSecretIdNumUses

`func (o *AppRoleWriteSecretIDLookupResponse) GetSecretIdNumUses() int32`

GetSecretIdNumUses returns the SecretIdNumUses field if non-nil, zero value otherwise.

### GetSecretIdNumUsesOk

`func (o *AppRoleWriteSecretIDLookupResponse) GetSecretIdNumUsesOk() (*int32, bool)`

GetSecretIdNumUsesOk returns a tuple with the SecretIdNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdNumUses

`func (o *AppRoleWriteSecretIDLookupResponse) SetSecretIdNumUses(v int32)`

SetSecretIdNumUses sets SecretIdNumUses field to given value.


### HasSecretIdNumUses

`func (o *AppRoleWriteSecretIDLookupResponse) HasSecretIdNumUses() bool`

HasSecretIdNumUses returns a boolean if a field has been set.




### GetSecretIdTtl

`func (o *AppRoleWriteSecretIDLookupResponse) GetSecretIdTtl() int32`

GetSecretIdTtl returns the SecretIdTtl field if non-nil, zero value otherwise.

### GetSecretIdTtlOk

`func (o *AppRoleWriteSecretIDLookupResponse) GetSecretIdTtlOk() (*int32, bool)`

GetSecretIdTtlOk returns a tuple with the SecretIdTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdTtl

`func (o *AppRoleWriteSecretIDLookupResponse) SetSecretIdTtl(v int32)`

SetSecretIdTtl sets SecretIdTtl field to given value.


### HasSecretIdTtl

`func (o *AppRoleWriteSecretIDLookupResponse) HasSecretIdTtl() bool`

HasSecretIdTtl returns a boolean if a field has been set.




### GetTokenBoundCidrs

`func (o *AppRoleWriteSecretIDLookupResponse) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *AppRoleWriteSecretIDLookupResponse) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *AppRoleWriteSecretIDLookupResponse) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.


### HasTokenBoundCidrs

`func (o *AppRoleWriteSecretIDLookupResponse) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


