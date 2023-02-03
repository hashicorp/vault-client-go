# EntityBatchDeleteRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**EntityIds** | Pointer to **[]string** | Entity IDs to delete | [optional] 



## Methods


### NewEntityBatchDeleteRequest

`func NewEntityBatchDeleteRequest() *EntityBatchDeleteRequest`

NewEntityBatchDeleteRequest instantiates a new EntityBatchDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityBatchDeleteRequestWithDefaults

`func NewEntityBatchDeleteRequestWithDefaults() *EntityBatchDeleteRequest`

NewEntityBatchDeleteRequestWithDefaults instantiates a new EntityBatchDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetEntityIds

`func (o *EntityBatchDeleteRequest) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *EntityBatchDeleteRequest) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *EntityBatchDeleteRequest) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.


### HasEntityIds

`func (o *EntityBatchDeleteRequest) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


