# IdentityAliasRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanonicalId** | Pointer to **string** | Entity ID to which this alias belongs to | [optional] 
**EntityId** | Pointer to **string** | Entity ID to which this alias belongs to. This field is deprecated in favor of &#39;canonical_id&#39;. | [optional] 
**Id** | Pointer to **string** | ID of the alias | [optional] 
**MountAccessor** | Pointer to **string** | Mount accessor to which this alias belongs to | [optional] 
**Name** | Pointer to **string** | Name of the alias | [optional] 

## Methods

### NewIdentityAliasRequest

`func NewIdentityAliasRequest() *IdentityAliasRequest`

NewIdentityAliasRequest instantiates a new IdentityAliasRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAliasRequestWithDefaults

`func NewIdentityAliasRequestWithDefaults() *IdentityAliasRequest`

NewIdentityAliasRequestWithDefaults instantiates a new IdentityAliasRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonicalId

`func (o *IdentityAliasRequest) GetCanonicalId() string`

GetCanonicalId returns the CanonicalId field if non-nil, zero value otherwise.

### GetCanonicalIdOk

`func (o *IdentityAliasRequest) GetCanonicalIdOk() (*string, bool)`

GetCanonicalIdOk returns a tuple with the CanonicalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalId

`func (o *IdentityAliasRequest) SetCanonicalId(v string)`

SetCanonicalId sets CanonicalId field to given value.

### HasCanonicalId

`func (o *IdentityAliasRequest) HasCanonicalId() bool`

HasCanonicalId returns a boolean if a field has been set.

### GetEntityId

`func (o *IdentityAliasRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *IdentityAliasRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *IdentityAliasRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *IdentityAliasRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetId

`func (o *IdentityAliasRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityAliasRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityAliasRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityAliasRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMountAccessor

`func (o *IdentityAliasRequest) GetMountAccessor() string`

GetMountAccessor returns the MountAccessor field if non-nil, zero value otherwise.

### GetMountAccessorOk

`func (o *IdentityAliasRequest) GetMountAccessorOk() (*string, bool)`

GetMountAccessorOk returns a tuple with the MountAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountAccessor

`func (o *IdentityAliasRequest) SetMountAccessor(v string)`

SetMountAccessor sets MountAccessor field to given value.

### HasMountAccessor

`func (o *IdentityAliasRequest) HasMountAccessor() bool`

HasMountAccessor returns a boolean if a field has been set.

### GetName

`func (o *IdentityAliasRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityAliasRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityAliasRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityAliasRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


