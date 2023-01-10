# UserpassWriteUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundCidrs** | Pointer to **[]string** | Use \&quot;token_bound_cidrs\&quot; instead. If this and \&quot;token_bound_cidrs\&quot; are both specified, only \&quot;token_bound_cidrs\&quot; will be used. | [optional] 
**MaxTtl** | Pointer to **int32** | Use \&quot;token_max_ttl\&quot; instead. If this and \&quot;token_max_ttl\&quot; are both specified, only \&quot;token_max_ttl\&quot; will be used. | [optional] 
**Password** | Pointer to **string** | Password for this user. | [optional] 
**Policies** | Pointer to **[]string** | Use \&quot;token_policies\&quot; instead. If this and \&quot;token_policies\&quot; are both specified, only \&quot;token_policies\&quot; will be used. | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **int32** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **int32** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#39;default&#39; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **int32** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies | [optional] 
**TokenTtl** | Pointer to **int32** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]
**Ttl** | Pointer to **int32** | Use \&quot;token_ttl\&quot; instead. If this and \&quot;token_ttl\&quot; are both specified, only \&quot;token_ttl\&quot; will be used. | [optional] 

## Methods

### NewUserpassWriteUserRequest

`func NewUserpassWriteUserRequest() *UserpassWriteUserRequest`

NewUserpassWriteUserRequest instantiates a new UserpassWriteUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserpassWriteUserRequestWithDefaults

`func NewUserpassWriteUserRequestWithDefaults() *UserpassWriteUserRequest`

NewUserpassWriteUserRequestWithDefaults instantiates a new UserpassWriteUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundCidrs

`func (o *UserpassWriteUserRequest) GetBoundCidrs() []string`

GetBoundCidrs returns the BoundCidrs field if non-nil, zero value otherwise.

### GetBoundCidrsOk

`func (o *UserpassWriteUserRequest) GetBoundCidrsOk() (*[]string, bool)`

GetBoundCidrsOk returns a tuple with the BoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundCidrs

`func (o *UserpassWriteUserRequest) SetBoundCidrs(v []string)`

SetBoundCidrs sets BoundCidrs field to given value.

### HasBoundCidrs

`func (o *UserpassWriteUserRequest) HasBoundCidrs() bool`

HasBoundCidrs returns a boolean if a field has been set.

### GetMaxTtl

`func (o *UserpassWriteUserRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *UserpassWriteUserRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *UserpassWriteUserRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *UserpassWriteUserRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetPassword

`func (o *UserpassWriteUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserpassWriteUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserpassWriteUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserpassWriteUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPolicies

`func (o *UserpassWriteUserRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *UserpassWriteUserRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *UserpassWriteUserRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *UserpassWriteUserRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetTokenBoundCidrs

`func (o *UserpassWriteUserRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *UserpassWriteUserRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *UserpassWriteUserRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.

### HasTokenBoundCidrs

`func (o *UserpassWriteUserRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.

### GetTokenExplicitMaxTtl

`func (o *UserpassWriteUserRequest) GetTokenExplicitMaxTtl() int32`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *UserpassWriteUserRequest) GetTokenExplicitMaxTtlOk() (*int32, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *UserpassWriteUserRequest) SetTokenExplicitMaxTtl(v int32)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.

### HasTokenExplicitMaxTtl

`func (o *UserpassWriteUserRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.

### GetTokenMaxTtl

`func (o *UserpassWriteUserRequest) GetTokenMaxTtl() int32`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *UserpassWriteUserRequest) GetTokenMaxTtlOk() (*int32, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *UserpassWriteUserRequest) SetTokenMaxTtl(v int32)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.

### HasTokenMaxTtl

`func (o *UserpassWriteUserRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.

### GetTokenNoDefaultPolicy

`func (o *UserpassWriteUserRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *UserpassWriteUserRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *UserpassWriteUserRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.

### HasTokenNoDefaultPolicy

`func (o *UserpassWriteUserRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.

### GetTokenNumUses

`func (o *UserpassWriteUserRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *UserpassWriteUserRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *UserpassWriteUserRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.

### HasTokenNumUses

`func (o *UserpassWriteUserRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.

### GetTokenPeriod

`func (o *UserpassWriteUserRequest) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *UserpassWriteUserRequest) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *UserpassWriteUserRequest) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.

### HasTokenPeriod

`func (o *UserpassWriteUserRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.

### GetTokenPolicies

`func (o *UserpassWriteUserRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *UserpassWriteUserRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *UserpassWriteUserRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.

### HasTokenPolicies

`func (o *UserpassWriteUserRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.

### GetTokenTtl

`func (o *UserpassWriteUserRequest) GetTokenTtl() int32`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *UserpassWriteUserRequest) GetTokenTtlOk() (*int32, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *UserpassWriteUserRequest) SetTokenTtl(v int32)`

SetTokenTtl sets TokenTtl field to given value.

### HasTokenTtl

`func (o *UserpassWriteUserRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.

### GetTokenType

`func (o *UserpassWriteUserRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *UserpassWriteUserRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *UserpassWriteUserRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *UserpassWriteUserRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetTtl

`func (o *UserpassWriteUserRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *UserpassWriteUserRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *UserpassWriteUserRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *UserpassWriteUserRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


