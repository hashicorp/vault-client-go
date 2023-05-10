# TokenCreateAgainstRoleRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Name to associate with this token | [optional] 
**EntityAlias** | Pointer to **string** | Name of the entity alias to associate with this token | [optional] 
**ExplicitMaxTtl** | Pointer to **string** | Explicit Max TTL of this token | [optional] 
**Id** | Pointer to **string** | Value for the token | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary key&#x3D;value metadata to associate with the token | [optional] 
**NoDefaultPolicy** | Pointer to **bool** | Do not include default policy for this token | [optional] 
**NoParent** | Pointer to **bool** | Create the token with no parent | [optional] 
**NumUses** | Pointer to **int32** | Max number of uses for this token | [optional] 
**Period** | Pointer to **string** | Renew period | [optional] 
**Policies** | Pointer to **[]string** | List of policies for the token | [optional] 
**Renewable** | Pointer to **bool** | Allow token to be renewed past its initial TTL up to system/mount maximum TTL | [optional] 
**Ttl** | Pointer to **string** | Time to live for this token | [optional] 
**Type** | Pointer to **string** | Token type | [optional] 



## Methods


### NewTokenCreateAgainstRoleRequest

`func NewTokenCreateAgainstRoleRequest() *TokenCreateAgainstRoleRequest`

NewTokenCreateAgainstRoleRequest instantiates a new TokenCreateAgainstRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCreateAgainstRoleRequestWithDefaults

`func NewTokenCreateAgainstRoleRequestWithDefaults() *TokenCreateAgainstRoleRequest`

NewTokenCreateAgainstRoleRequestWithDefaults instantiates a new TokenCreateAgainstRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetDisplayName

`func (o *TokenCreateAgainstRoleRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TokenCreateAgainstRoleRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TokenCreateAgainstRoleRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### HasDisplayName

`func (o *TokenCreateAgainstRoleRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.




### GetEntityAlias

`func (o *TokenCreateAgainstRoleRequest) GetEntityAlias() string`

GetEntityAlias returns the EntityAlias field if non-nil, zero value otherwise.

### GetEntityAliasOk

`func (o *TokenCreateAgainstRoleRequest) GetEntityAliasOk() (*string, bool)`

GetEntityAliasOk returns a tuple with the EntityAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityAlias

`func (o *TokenCreateAgainstRoleRequest) SetEntityAlias(v string)`

SetEntityAlias sets EntityAlias field to given value.


### HasEntityAlias

`func (o *TokenCreateAgainstRoleRequest) HasEntityAlias() bool`

HasEntityAlias returns a boolean if a field has been set.




### GetExplicitMaxTtl

`func (o *TokenCreateAgainstRoleRequest) GetExplicitMaxTtl() string`

GetExplicitMaxTtl returns the ExplicitMaxTtl field if non-nil, zero value otherwise.

### GetExplicitMaxTtlOk

`func (o *TokenCreateAgainstRoleRequest) GetExplicitMaxTtlOk() (*string, bool)`

GetExplicitMaxTtlOk returns a tuple with the ExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitMaxTtl

`func (o *TokenCreateAgainstRoleRequest) SetExplicitMaxTtl(v string)`

SetExplicitMaxTtl sets ExplicitMaxTtl field to given value.


### HasExplicitMaxTtl

`func (o *TokenCreateAgainstRoleRequest) HasExplicitMaxTtl() bool`

HasExplicitMaxTtl returns a boolean if a field has been set.




### GetId

`func (o *TokenCreateAgainstRoleRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenCreateAgainstRoleRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenCreateAgainstRoleRequest) SetId(v string)`

SetId sets Id field to given value.


### HasId

`func (o *TokenCreateAgainstRoleRequest) HasId() bool`

HasId returns a boolean if a field has been set.




### GetMetadata

`func (o *TokenCreateAgainstRoleRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TokenCreateAgainstRoleRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TokenCreateAgainstRoleRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### HasMetadata

`func (o *TokenCreateAgainstRoleRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.




### GetNoDefaultPolicy

`func (o *TokenCreateAgainstRoleRequest) GetNoDefaultPolicy() bool`

GetNoDefaultPolicy returns the NoDefaultPolicy field if non-nil, zero value otherwise.

### GetNoDefaultPolicyOk

`func (o *TokenCreateAgainstRoleRequest) GetNoDefaultPolicyOk() (*bool, bool)`

GetNoDefaultPolicyOk returns a tuple with the NoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDefaultPolicy

`func (o *TokenCreateAgainstRoleRequest) SetNoDefaultPolicy(v bool)`

SetNoDefaultPolicy sets NoDefaultPolicy field to given value.


### HasNoDefaultPolicy

`func (o *TokenCreateAgainstRoleRequest) HasNoDefaultPolicy() bool`

HasNoDefaultPolicy returns a boolean if a field has been set.




### GetNoParent

`func (o *TokenCreateAgainstRoleRequest) GetNoParent() bool`

GetNoParent returns the NoParent field if non-nil, zero value otherwise.

### GetNoParentOk

`func (o *TokenCreateAgainstRoleRequest) GetNoParentOk() (*bool, bool)`

GetNoParentOk returns a tuple with the NoParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoParent

`func (o *TokenCreateAgainstRoleRequest) SetNoParent(v bool)`

SetNoParent sets NoParent field to given value.


### HasNoParent

`func (o *TokenCreateAgainstRoleRequest) HasNoParent() bool`

HasNoParent returns a boolean if a field has been set.




### GetNumUses

`func (o *TokenCreateAgainstRoleRequest) GetNumUses() int32`

GetNumUses returns the NumUses field if non-nil, zero value otherwise.

### GetNumUsesOk

`func (o *TokenCreateAgainstRoleRequest) GetNumUsesOk() (*int32, bool)`

GetNumUsesOk returns a tuple with the NumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUses

`func (o *TokenCreateAgainstRoleRequest) SetNumUses(v int32)`

SetNumUses sets NumUses field to given value.


### HasNumUses

`func (o *TokenCreateAgainstRoleRequest) HasNumUses() bool`

HasNumUses returns a boolean if a field has been set.




### GetPeriod

`func (o *TokenCreateAgainstRoleRequest) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TokenCreateAgainstRoleRequest) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TokenCreateAgainstRoleRequest) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### HasPeriod

`func (o *TokenCreateAgainstRoleRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.




### GetPolicies

`func (o *TokenCreateAgainstRoleRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *TokenCreateAgainstRoleRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *TokenCreateAgainstRoleRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *TokenCreateAgainstRoleRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.




### GetRenewable

`func (o *TokenCreateAgainstRoleRequest) GetRenewable() bool`

GetRenewable returns the Renewable field if non-nil, zero value otherwise.

### GetRenewableOk

`func (o *TokenCreateAgainstRoleRequest) GetRenewableOk() (*bool, bool)`

GetRenewableOk returns a tuple with the Renewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewable

`func (o *TokenCreateAgainstRoleRequest) SetRenewable(v bool)`

SetRenewable sets Renewable field to given value.


### HasRenewable

`func (o *TokenCreateAgainstRoleRequest) HasRenewable() bool`

HasRenewable returns a boolean if a field has been set.




### GetTtl

`func (o *TokenCreateAgainstRoleRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *TokenCreateAgainstRoleRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *TokenCreateAgainstRoleRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *TokenCreateAgainstRoleRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.




### GetType

`func (o *TokenCreateAgainstRoleRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenCreateAgainstRoleRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenCreateAgainstRoleRequest) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *TokenCreateAgainstRoleRequest) HasType() bool`

HasType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


