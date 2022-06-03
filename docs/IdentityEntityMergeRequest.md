# IdentityEntityMergeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Force** | Pointer to **bool** | Setting this will follow the &#39;mine&#39; strategy for merging MFA secrets. If there are secrets of the same type both in entities that are merged from and in entity into which all others are getting merged, secrets in the destination will be unaltered. If not set, this API will throw an error containing all the conflicts. | [optional] 
**FromEntityIds** | Pointer to **[]string** | Entity IDs which needs to get merged | [optional] 
**ToEntityId** | Pointer to **string** | Entity ID into which all the other entities need to get merged | [optional] 

## Methods

### NewIdentityEntityMergeRequest

`func NewIdentityEntityMergeRequest() *IdentityEntityMergeRequest`

NewIdentityEntityMergeRequest instantiates a new IdentityEntityMergeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityEntityMergeRequestWithDefaults

`func NewIdentityEntityMergeRequestWithDefaults() *IdentityEntityMergeRequest`

NewIdentityEntityMergeRequestWithDefaults instantiates a new IdentityEntityMergeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForce

`func (o *IdentityEntityMergeRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *IdentityEntityMergeRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *IdentityEntityMergeRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *IdentityEntityMergeRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetFromEntityIds

`func (o *IdentityEntityMergeRequest) GetFromEntityIds() []string`

GetFromEntityIds returns the FromEntityIds field if non-nil, zero value otherwise.

### GetFromEntityIdsOk

`func (o *IdentityEntityMergeRequest) GetFromEntityIdsOk() (*[]string, bool)`

GetFromEntityIdsOk returns a tuple with the FromEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEntityIds

`func (o *IdentityEntityMergeRequest) SetFromEntityIds(v []string)`

SetFromEntityIds sets FromEntityIds field to given value.

### HasFromEntityIds

`func (o *IdentityEntityMergeRequest) HasFromEntityIds() bool`

HasFromEntityIds returns a boolean if a field has been set.

### GetToEntityId

`func (o *IdentityEntityMergeRequest) GetToEntityId() string`

GetToEntityId returns the ToEntityId field if non-nil, zero value otherwise.

### GetToEntityIdOk

`func (o *IdentityEntityMergeRequest) GetToEntityIdOk() (*string, bool)`

GetToEntityIdOk returns a tuple with the ToEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToEntityId

`func (o *IdentityEntityMergeRequest) SetToEntityId(v string)`

SetToEntityId sets ToEntityId field to given value.

### HasToEntityId

`func (o *IdentityEntityMergeRequest) HasToEntityId() bool`

HasToEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


