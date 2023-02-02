# EntityMergeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**ConflictingAliasIdsToKeep** | Pointer to **[]string** | Alias IDs to keep in case of conflicting aliases. Ignored if no conflicting aliases found | [optional] 
**Force** | Pointer to **bool** | Setting this will follow the &#x27;mine&#x27; strategy for merging MFA secrets. If there are secrets of the same type both in entities that are merged from and in entity into which all others are getting merged, secrets in the destination will be unaltered. If not set, this API will throw an error containing all the conflicts. | [optional] 
**FromEntityIds** | Pointer to **[]string** | Entity IDs which need to get merged | [optional] 
**ToEntityId** | Pointer to **string** | Entity ID into which all the other entities need to get merged | [optional] 



## Methods


### NewEntityMergeRequest

`func NewEntityMergeRequest() *EntityMergeRequest`

NewEntityMergeRequest instantiates a new EntityMergeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityMergeRequestWithDefaults

`func NewEntityMergeRequestWithDefaults() *EntityMergeRequest`

NewEntityMergeRequestWithDefaults instantiates a new EntityMergeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetConflictingAliasIdsToKeep

`func (o *EntityMergeRequest) GetConflictingAliasIdsToKeep() []string`

GetConflictingAliasIdsToKeep returns the ConflictingAliasIdsToKeep field if non-nil, zero value otherwise.

### GetConflictingAliasIdsToKeepOk

`func (o *EntityMergeRequest) GetConflictingAliasIdsToKeepOk() (*[]string, bool)`

GetConflictingAliasIdsToKeepOk returns a tuple with the ConflictingAliasIdsToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingAliasIdsToKeep

`func (o *EntityMergeRequest) SetConflictingAliasIdsToKeep(v []string)`

SetConflictingAliasIdsToKeep sets ConflictingAliasIdsToKeep field to given value.


### HasConflictingAliasIdsToKeep

`func (o *EntityMergeRequest) HasConflictingAliasIdsToKeep() bool`

HasConflictingAliasIdsToKeep returns a boolean if a field has been set.




### GetForce

`func (o *EntityMergeRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *EntityMergeRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *EntityMergeRequest) SetForce(v bool)`

SetForce sets Force field to given value.


### HasForce

`func (o *EntityMergeRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.




### GetFromEntityIds

`func (o *EntityMergeRequest) GetFromEntityIds() []string`

GetFromEntityIds returns the FromEntityIds field if non-nil, zero value otherwise.

### GetFromEntityIdsOk

`func (o *EntityMergeRequest) GetFromEntityIdsOk() (*[]string, bool)`

GetFromEntityIdsOk returns a tuple with the FromEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEntityIds

`func (o *EntityMergeRequest) SetFromEntityIds(v []string)`

SetFromEntityIds sets FromEntityIds field to given value.


### HasFromEntityIds

`func (o *EntityMergeRequest) HasFromEntityIds() bool`

HasFromEntityIds returns a boolean if a field has been set.




### GetToEntityId

`func (o *EntityMergeRequest) GetToEntityId() string`

GetToEntityId returns the ToEntityId field if non-nil, zero value otherwise.

### GetToEntityIdOk

`func (o *EntityMergeRequest) GetToEntityIdOk() (*string, bool)`

GetToEntityIdOk returns a tuple with the ToEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToEntityId

`func (o *EntityMergeRequest) SetToEntityId(v string)`

SetToEntityId sets ToEntityId field to given value.


### HasToEntityId

`func (o *EntityMergeRequest) HasToEntityId() bool`

HasToEntityId returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


