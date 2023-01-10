# AppRoleWriteRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindSecretId** | Pointer to **bool** | Impose secret_id to be presented when logging in using this role. Defaults to &#39;true&#39;. | [optional] [default to true]
**BoundCidrList** | Pointer to **[]string** | Use \&quot;secret_id_bound_cidrs\&quot; instead. | [optional] 
**LocalSecretIds** | Pointer to **bool** | If set, the secret IDs generated using this role will be cluster local. This can only be set during role creation and once set, it can&#39;t be reset later. | [optional] 
**Period** | Pointer to **int32** | Use \&quot;token_period\&quot; instead. If this and \&quot;token_period\&quot; are both specified, only \&quot;token_period\&quot; will be used. | [optional] 
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
**RoleId** | Pointer to **string** | Identifier of the role. Defaults to a UUID. | [optional] 
**SecretIdBoundCidrs** | Pointer to **[]string** | Comma separated string or list of CIDR blocks. If set, specifies the blocks of IP addresses which can perform the login operation. | [optional] 
**SecretIdNumUses** | Pointer to **int32** | Number of times a SecretID can access the role, after which the SecretID will expire. Defaults to 0 meaning that the the secret_id is of unlimited use. | [optional] 
**SecretIdTtl** | Pointer to **int32** | Duration in seconds after which the issued SecretID should expire. Defaults to 0, meaning no expiration. | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **int32** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **int32** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#39;default&#39; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **int32** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies | [optional] 
**TokenTtl** | Pointer to **int32** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]

## Methods

### NewAppRoleWriteRoleRequest

`func NewAppRoleWriteRoleRequest() *AppRoleWriteRoleRequest`

NewAppRoleWriteRoleRequest instantiates a new AppRoleWriteRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRoleWriteRoleRequestWithDefaults

`func NewAppRoleWriteRoleRequestWithDefaults() *AppRoleWriteRoleRequest`

NewAppRoleWriteRoleRequestWithDefaults instantiates a new AppRoleWriteRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindSecretId

`func (o *AppRoleWriteRoleRequest) GetBindSecretId() bool`

GetBindSecretId returns the BindSecretId field if non-nil, zero value otherwise.

### GetBindSecretIdOk

`func (o *AppRoleWriteRoleRequest) GetBindSecretIdOk() (*bool, bool)`

GetBindSecretIdOk returns a tuple with the BindSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindSecretId

`func (o *AppRoleWriteRoleRequest) SetBindSecretId(v bool)`

SetBindSecretId sets BindSecretId field to given value.

### HasBindSecretId

`func (o *AppRoleWriteRoleRequest) HasBindSecretId() bool`

HasBindSecretId returns a boolean if a field has been set.

### GetBoundCidrList

`func (o *AppRoleWriteRoleRequest) GetBoundCidrList() []string`

GetBoundCidrList returns the BoundCidrList field if non-nil, zero value otherwise.

### GetBoundCidrListOk

`func (o *AppRoleWriteRoleRequest) GetBoundCidrListOk() (*[]string, bool)`

GetBoundCidrListOk returns a tuple with the BoundCidrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCidrList

`func (o *AppRoleWriteRoleRequest) SetBoundCidrList(v []string)`

SetBoundCidrList sets BoundCidrList field to given value.

### HasBoundCidrList

`func (o *AppRoleWriteRoleRequest) HasBoundCidrList() bool`

HasBoundCidrList returns a boolean if a field has been set.

### GetLocalSecretIds

`func (o *AppRoleWriteRoleRequest) GetLocalSecretIds() bool`

GetLocalSecretIds returns the LocalSecretIds field if non-nil, zero value otherwise.

### GetLocalSecretIdsOk

`func (o *AppRoleWriteRoleRequest) GetLocalSecretIdsOk() (*bool, bool)`

GetLocalSecretIdsOk returns a tuple with the LocalSecretIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSecretIds

`func (o *AppRoleWriteRoleRequest) SetLocalSecretIds(v bool)`

SetLocalSecretIds sets LocalSecretIds field to given value.

### HasLocalSecretIds

`func (o *AppRoleWriteRoleRequest) HasLocalSecretIds() bool`

HasLocalSecretIds returns a boolean if a field has been set.

### GetPeriod

`func (o *AppRoleWriteRoleRequest) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AppRoleWriteRoleRequest) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AppRoleWriteRoleRequest) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *AppRoleWriteRoleRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPolicies

`func (o *AppRoleWriteRoleRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *AppRoleWriteRoleRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *AppRoleWriteRoleRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *AppRoleWriteRoleRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetRoleId

`func (o *AppRoleWriteRoleRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *AppRoleWriteRoleRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *AppRoleWriteRoleRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *AppRoleWriteRoleRequest) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetSecretIdBoundCidrs

`func (o *AppRoleWriteRoleRequest) GetSecretIdBoundCidrs() []string`

GetSecretIdBoundCidrs returns the SecretIdBoundCidrs field if non-nil, zero value otherwise.

### GetSecretIdBoundCidrsOk

`func (o *AppRoleWriteRoleRequest) GetSecretIdBoundCidrsOk() (*[]string, bool)`

GetSecretIdBoundCidrsOk returns a tuple with the SecretIdBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdBoundCidrs

`func (o *AppRoleWriteRoleRequest) SetSecretIdBoundCidrs(v []string)`

SetSecretIdBoundCidrs sets SecretIdBoundCidrs field to given value.

### HasSecretIdBoundCidrs

`func (o *AppRoleWriteRoleRequest) HasSecretIdBoundCidrs() bool`

HasSecretIdBoundCidrs returns a boolean if a field has been set.

### GetSecretIdNumUses

`func (o *AppRoleWriteRoleRequest) GetSecretIdNumUses() int32`

GetSecretIdNumUses returns the SecretIdNumUses field if non-nil, zero value otherwise.

### GetSecretIdNumUsesOk

`func (o *AppRoleWriteRoleRequest) GetSecretIdNumUsesOk() (*int32, bool)`

GetSecretIdNumUsesOk returns a tuple with the SecretIdNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdNumUses

`func (o *AppRoleWriteRoleRequest) SetSecretIdNumUses(v int32)`

SetSecretIdNumUses sets SecretIdNumUses field to given value.

### HasSecretIdNumUses

`func (o *AppRoleWriteRoleRequest) HasSecretIdNumUses() bool`

HasSecretIdNumUses returns a boolean if a field has been set.

### GetSecretIdTtl

`func (o *AppRoleWriteRoleRequest) GetSecretIdTtl() int32`

GetSecretIdTtl returns the SecretIdTtl field if non-nil, zero value otherwise.

### GetSecretIdTtlOk

`func (o *AppRoleWriteRoleRequest) GetSecretIdTtlOk() (*int32, bool)`

GetSecretIdTtlOk returns a tuple with the SecretIdTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdTtl

`func (o *AppRoleWriteRoleRequest) SetSecretIdTtl(v int32)`

SetSecretIdTtl sets SecretIdTtl field to given value.

### HasSecretIdTtl

`func (o *AppRoleWriteRoleRequest) HasSecretIdTtl() bool`

HasSecretIdTtl returns a boolean if a field has been set.

### GetTokenBoundCidrs

`func (o *AppRoleWriteRoleRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *AppRoleWriteRoleRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *AppRoleWriteRoleRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.

### HasTokenBoundCidrs

`func (o *AppRoleWriteRoleRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.

### GetTokenExplicitMaxTtl

`func (o *AppRoleWriteRoleRequest) GetTokenExplicitMaxTtl() int32`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *AppRoleWriteRoleRequest) GetTokenExplicitMaxTtlOk() (*int32, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *AppRoleWriteRoleRequest) SetTokenExplicitMaxTtl(v int32)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.

### HasTokenExplicitMaxTtl

`func (o *AppRoleWriteRoleRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.

### GetTokenMaxTtl

`func (o *AppRoleWriteRoleRequest) GetTokenMaxTtl() int32`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *AppRoleWriteRoleRequest) GetTokenMaxTtlOk() (*int32, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *AppRoleWriteRoleRequest) SetTokenMaxTtl(v int32)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.

### HasTokenMaxTtl

`func (o *AppRoleWriteRoleRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.

### GetTokenNoDefaultPolicy

`func (o *AppRoleWriteRoleRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *AppRoleWriteRoleRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *AppRoleWriteRoleRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.

### HasTokenNoDefaultPolicy

`func (o *AppRoleWriteRoleRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.

### GetTokenNumUses

`func (o *AppRoleWriteRoleRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *AppRoleWriteRoleRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *AppRoleWriteRoleRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.

### HasTokenNumUses

`func (o *AppRoleWriteRoleRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.

### GetTokenPeriod

`func (o *AppRoleWriteRoleRequest) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *AppRoleWriteRoleRequest) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *AppRoleWriteRoleRequest) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.

### HasTokenPeriod

`func (o *AppRoleWriteRoleRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.

### GetTokenPolicies

`func (o *AppRoleWriteRoleRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *AppRoleWriteRoleRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *AppRoleWriteRoleRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.

### HasTokenPolicies

`func (o *AppRoleWriteRoleRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.

### GetTokenTtl

`func (o *AppRoleWriteRoleRequest) GetTokenTtl() int32`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *AppRoleWriteRoleRequest) GetTokenTtlOk() (*int32, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *AppRoleWriteRoleRequest) SetTokenTtl(v int32)`

SetTokenTtl sets TokenTtl field to given value.

### HasTokenTtl

`func (o *AppRoleWriteRoleRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.

### GetTokenType

`func (o *AppRoleWriteRoleRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AppRoleWriteRoleRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AppRoleWriteRoleRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *AppRoleWriteRoleRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


