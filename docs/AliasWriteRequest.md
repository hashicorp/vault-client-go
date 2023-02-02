# AliasWriteRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**CanonicalId** | Pointer to **string** | Entity ID to which this alias belongs to | [optional] 
**EntityId** | Pointer to **string** | Entity ID to which this alias belongs to. This field is deprecated in favor of &#x27;canonical_id&#x27;. | [optional] 
**Id** | Pointer to **string** | ID of the alias | [optional] 
**MountAccessor** | Pointer to **string** | Mount accessor to which this alias belongs to | [optional] 
**Name** | Pointer to **string** | Name of the alias | [optional] 



## Methods


### NewAliasWriteRequest

`func NewAliasWriteRequest() *AliasWriteRequest`

NewAliasWriteRequest instantiates a new AliasWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAliasWriteRequestWithDefaults

`func NewAliasWriteRequestWithDefaults() *AliasWriteRequest`

NewAliasWriteRequestWithDefaults instantiates a new AliasWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCanonicalId

`func (o *AliasWriteRequest) GetCanonicalId() string`

GetCanonicalId returns the CanonicalId field if non-nil, zero value otherwise.

### GetCanonicalIdOk

`func (o *AliasWriteRequest) GetCanonicalIdOk() (*string, bool)`

GetCanonicalIdOk returns a tuple with the CanonicalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalId

`func (o *AliasWriteRequest) SetCanonicalId(v string)`

SetCanonicalId sets CanonicalId field to given value.


### HasCanonicalId

`func (o *AliasWriteRequest) HasCanonicalId() bool`

HasCanonicalId returns a boolean if a field has been set.




### GetEntityId

`func (o *AliasWriteRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AliasWriteRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AliasWriteRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### HasEntityId

`func (o *AliasWriteRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.




### GetId

`func (o *AliasWriteRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AliasWriteRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AliasWriteRequest) SetId(v string)`

SetId sets Id field to given value.


### HasId

`func (o *AliasWriteRequest) HasId() bool`

HasId returns a boolean if a field has been set.




### GetMountAccessor

`func (o *AliasWriteRequest) GetMountAccessor() string`

GetMountAccessor returns the MountAccessor field if non-nil, zero value otherwise.

### GetMountAccessorOk

`func (o *AliasWriteRequest) GetMountAccessorOk() (*string, bool)`

GetMountAccessorOk returns a tuple with the MountAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountAccessor

`func (o *AliasWriteRequest) SetMountAccessor(v string)`

SetMountAccessor sets MountAccessor field to given value.


### HasMountAccessor

`func (o *AliasWriteRequest) HasMountAccessor() bool`

HasMountAccessor returns a boolean if a field has been set.




### GetName

`func (o *AliasWriteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AliasWriteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AliasWriteRequest) SetName(v string)`

SetName sets Name field to given value.


### HasName

`func (o *AliasWriteRequest) HasName() bool`

HasName returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


