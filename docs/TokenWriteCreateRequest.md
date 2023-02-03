# TokenWriteCreateRequest


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


### NewTokenWriteCreateRequest

`func NewTokenWriteCreateRequest() *TokenWriteCreateRequest`

NewTokenWriteCreateRequest instantiates a new TokenWriteCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWriteCreateRequestWithDefaults

`func NewTokenWriteCreateRequestWithDefaults() *TokenWriteCreateRequest`

NewTokenWriteCreateRequestWithDefaults instantiates a new TokenWriteCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetDisplayName

`func (o *TokenWriteCreateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TokenWriteCreateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TokenWriteCreateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### HasDisplayName

`func (o *TokenWriteCreateRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.




### GetEntityAlias

`func (o *TokenWriteCreateRequest) GetEntityAlias() string`

GetEntityAlias returns the EntityAlias field if non-nil, zero value otherwise.

### GetEntityAliasOk

`func (o *TokenWriteCreateRequest) GetEntityAliasOk() (*string, bool)`

GetEntityAliasOk returns a tuple with the EntityAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityAlias

`func (o *TokenWriteCreateRequest) SetEntityAlias(v string)`

SetEntityAlias sets EntityAlias field to given value.


### HasEntityAlias

`func (o *TokenWriteCreateRequest) HasEntityAlias() bool`

HasEntityAlias returns a boolean if a field has been set.




### GetExplicitMaxTtl

`func (o *TokenWriteCreateRequest) GetExplicitMaxTtl() string`

GetExplicitMaxTtl returns the ExplicitMaxTtl field if non-nil, zero value otherwise.

### GetExplicitMaxTtlOk

`func (o *TokenWriteCreateRequest) GetExplicitMaxTtlOk() (*string, bool)`

GetExplicitMaxTtlOk returns a tuple with the ExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitMaxTtl

`func (o *TokenWriteCreateRequest) SetExplicitMaxTtl(v string)`

SetExplicitMaxTtl sets ExplicitMaxTtl field to given value.


### HasExplicitMaxTtl

`func (o *TokenWriteCreateRequest) HasExplicitMaxTtl() bool`

HasExplicitMaxTtl returns a boolean if a field has been set.




### GetId

`func (o *TokenWriteCreateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenWriteCreateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenWriteCreateRequest) SetId(v string)`

SetId sets Id field to given value.


### HasId

`func (o *TokenWriteCreateRequest) HasId() bool`

HasId returns a boolean if a field has been set.




### GetMetadata

`func (o *TokenWriteCreateRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TokenWriteCreateRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TokenWriteCreateRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### HasMetadata

`func (o *TokenWriteCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.




### GetNoDefaultPolicy

`func (o *TokenWriteCreateRequest) GetNoDefaultPolicy() bool`

GetNoDefaultPolicy returns the NoDefaultPolicy field if non-nil, zero value otherwise.

### GetNoDefaultPolicyOk

`func (o *TokenWriteCreateRequest) GetNoDefaultPolicyOk() (*bool, bool)`

GetNoDefaultPolicyOk returns a tuple with the NoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDefaultPolicy

`func (o *TokenWriteCreateRequest) SetNoDefaultPolicy(v bool)`

SetNoDefaultPolicy sets NoDefaultPolicy field to given value.


### HasNoDefaultPolicy

`func (o *TokenWriteCreateRequest) HasNoDefaultPolicy() bool`

HasNoDefaultPolicy returns a boolean if a field has been set.




### GetNoParent

`func (o *TokenWriteCreateRequest) GetNoParent() bool`

GetNoParent returns the NoParent field if non-nil, zero value otherwise.

### GetNoParentOk

`func (o *TokenWriteCreateRequest) GetNoParentOk() (*bool, bool)`

GetNoParentOk returns a tuple with the NoParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoParent

`func (o *TokenWriteCreateRequest) SetNoParent(v bool)`

SetNoParent sets NoParent field to given value.


### HasNoParent

`func (o *TokenWriteCreateRequest) HasNoParent() bool`

HasNoParent returns a boolean if a field has been set.




### GetNumUses

`func (o *TokenWriteCreateRequest) GetNumUses() int32`

GetNumUses returns the NumUses field if non-nil, zero value otherwise.

### GetNumUsesOk

`func (o *TokenWriteCreateRequest) GetNumUsesOk() (*int32, bool)`

GetNumUsesOk returns a tuple with the NumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUses

`func (o *TokenWriteCreateRequest) SetNumUses(v int32)`

SetNumUses sets NumUses field to given value.


### HasNumUses

`func (o *TokenWriteCreateRequest) HasNumUses() bool`

HasNumUses returns a boolean if a field has been set.




### GetPeriod

`func (o *TokenWriteCreateRequest) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TokenWriteCreateRequest) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TokenWriteCreateRequest) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### HasPeriod

`func (o *TokenWriteCreateRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.




### GetPolicies

`func (o *TokenWriteCreateRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *TokenWriteCreateRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *TokenWriteCreateRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### HasPolicies

`func (o *TokenWriteCreateRequest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.




### GetRenewable

`func (o *TokenWriteCreateRequest) GetRenewable() bool`

GetRenewable returns the Renewable field if non-nil, zero value otherwise.

### GetRenewableOk

`func (o *TokenWriteCreateRequest) GetRenewableOk() (*bool, bool)`

GetRenewableOk returns a tuple with the Renewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewable

`func (o *TokenWriteCreateRequest) SetRenewable(v bool)`

SetRenewable sets Renewable field to given value.


### HasRenewable

`func (o *TokenWriteCreateRequest) HasRenewable() bool`

HasRenewable returns a boolean if a field has been set.




### GetTtl

`func (o *TokenWriteCreateRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *TokenWriteCreateRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *TokenWriteCreateRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *TokenWriteCreateRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.




### GetType

`func (o *TokenWriteCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenWriteCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenWriteCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *TokenWriteCreateRequest) HasType() bool`

HasType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


