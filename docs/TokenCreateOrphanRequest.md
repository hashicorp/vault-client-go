# TokenCreateOrphanRequest


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
**RoleName** | Pointer to **string** | Name of the role | [optional] 
**Ttl** | Pointer to **string** | Time to live for this token | [optional] 
**Type** | Pointer to **string** | Token type | [optional] 



## Methods


### NewTokenCreateOrphanRequest

`func NewTokenCreateOrphanRequest() *TokenCreateOrphanRequest`

NewTokenCreateOrphanRequest instantiates a new TokenCreateOrphanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCreateOrphanRequestWithDefaults

`func NewTokenCreateOrphanRequestWithDefaults() *TokenCreateOrphanRequest`

NewTokenCreateOrphanRequestWithDefaults instantiates a new TokenCreateOrphanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetDisplayName

`func (o *TokenCreateOrphanRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TokenCreateOrphanRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TokenCreateOrphanRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### HasDisplayName

`func (o *TokenCreateOrphanRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.




### GetEntityAlias

`func (o *TokenCreateOrphanRequest) GetEntityAlias() string`

GetEntityAlias returns the EntityAlias field if non-nil, zero value otherwise.

### GetEntityAliasOk

`func (o *TokenCreateOrphanRequest) GetEntityAliasOk() (*string, bool)`

GetEntityAliasOk returns a tuple with the EntityAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityAlias

`func (o *TokenCreateOrphanRequest) SetEntityAlias(v string)`

SetEntityAlias sets EntityAlias field to given value.


### HasEntityAlias

`func (o *TokenCreateOrphanRequest) HasEntityAlias() bool`

HasEntityAlias returns a boolean if a field has been set.




### GetExplicitMaxTtl

`func (o *TokenCreateOrphanRequest) GetExplicitMaxTtl() string`

GetExplicitMaxTtl returns the ExplicitMaxTtl field if non-nil, zero value otherwise.

### GetExplicitMaxTtlOk

`func (o *TokenCreateOrphanRequest) GetExplicitMaxTtlOk() (*string, bool)`

GetExplicitMaxTtlOk returns a tuple with the ExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitMaxTtl

`func (o *TokenCreateOrphanRequest) SetExplicitMaxTtl(v string)`

SetExplicitMaxTtl sets ExplicitMaxTtl field to given value.


### HasExplicitMaxTtl

`func (o *TokenCreateOrphanRequest) HasExplicitMaxTtl() bool`

HasExplicitMaxTtl returns a boolean if a field has been set.




### GetId

`func (o *TokenCreateOrphanRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenCreateOrphanRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenCreateOrphanRequest) SetId(v string)`

SetId sets Id field to given value.


### HasId

`func (o *TokenCreateOrphanRequest) HasId() bool`

HasId returns a boolean if a field has been set.




### GetMetadata

`func (o *TokenCreateOrphanRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TokenCreateOrphanRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TokenCreateOrphanRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### HasMetadata

`func (o *TokenCreateOrphanRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.




### GetNoDefaultPolicy

`func (o *TokenCreateOrphanRequest) GetNoDefaultPolicy() bool`

GetNoDefaultPolicy returns the NoDefaultPolicy field if non-nil, zero value otherwise.

### GetNoDefaultPolicyOk

`func (o *TokenCreateOrphanRequest) GetNoDefaultPolicyOk() (*bool, bool)`

GetNoDefaultPolicyOk returns a tuple with the NoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDefaultPolicy

`func (o *TokenCreateOrphanRequest) SetNoDefaultPolicy(v bool)`

SetNoDefaultPolicy sets NoDefaultPolicy field to given value.


### HasNoDefaultPolicy

`func (o *TokenCreateOrphanRequest) HasNoDefaultPolicy() bool`

HasNoDefaultPolicy returns a boolean if a field has been set.




### GetNoParent

`func (o *TokenCreateOrphanRequest) GetNoParent() bool`

GetNoParent returns the NoParent field if non-nil, zero value otherwise.

### GetNoParentOk

`func (o *TokenCreateOrphanRequest) GetNoParentOk() (*bool, bool)`

GetNoParentOk returns a tuple with the NoParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoParent

`func (o *TokenCreateOrphanRequest) SetNoParent(v bool)`

SetNoParent sets NoParent field to given value.


### HasNoParent

`func (o *TokenCreateOrphanRequest) HasNoParent() bool`

HasNoParent returns a boolean if a field has been set.




### GetNumUses

`func (o *TokenCreateOrphanRequest) GetNumUses() int32`

GetNumUses returns the NumUses field if non-nil, zero value otherwise.

### GetNumUsesOk

`func (o *TokenCreateOrphanRequest) GetNumUsesOk() (*int32, bool)`

GetNumUsesOk returns a tuple with the NumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUses

`func (o *TokenCreateOrphanRequest) SetNumUses(v int32)`

SetNumUses sets NumUses field to given value.


### HasNumUses

`func (o *TokenCreateOrphanRequest) HasNumUses() bool`

HasNumUses returns a boolean if a field has been set.




### GetPeriod

`func (o *TokenCreateOrphanRequest) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TokenCreateOrphanRequest) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TokenCreateOrphanRequest) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### HasPeriod

`func (o *TokenCreateOrphanRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.




### GetPolicies

`func (o *TokenCreateOrphanRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *TokenCreateOrphanRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *TokenCreateOrphanRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *TokenCreateOrphanRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.




### GetRenewable

`func (o *TokenCreateOrphanRequest) GetRenewable() bool`

GetRenewable returns the Renewable field if non-nil, zero value otherwise.

### GetRenewableOk

`func (o *TokenCreateOrphanRequest) GetRenewableOk() (*bool, bool)`

GetRenewableOk returns a tuple with the Renewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewable

`func (o *TokenCreateOrphanRequest) SetRenewable(v bool)`

SetRenewable sets Renewable field to given value.


### HasRenewable

`func (o *TokenCreateOrphanRequest) HasRenewable() bool`

HasRenewable returns a boolean if a field has been set.




### GetRoleName

`func (o *TokenCreateOrphanRequest) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *TokenCreateOrphanRequest) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *TokenCreateOrphanRequest) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### HasRoleName

`func (o *TokenCreateOrphanRequest) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.




### GetTtl

`func (o *TokenCreateOrphanRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *TokenCreateOrphanRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *TokenCreateOrphanRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *TokenCreateOrphanRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.




### GetType

`func (o *TokenCreateOrphanRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenCreateOrphanRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenCreateOrphanRequest) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *TokenCreateOrphanRequest) HasType() bool`

HasType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


