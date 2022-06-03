# IdentityAliasIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanonicalId** | Pointer to **string** | Entity ID to which this alias should be tied to | [optional] 
**EntityId** | Pointer to **string** | Entity ID to which this alias should be tied to. This field is deprecated in favor of &#39;canonical_id&#39;. | [optional] 
**MountAccessor** | Pointer to **string** | Mount accessor to which this alias belongs to | [optional] 
**Name** | Pointer to **string** | Name of the alias | [optional] 

## Methods

### NewIdentityAliasIdRequest

`func NewIdentityAliasIdRequest() *IdentityAliasIdRequest`

NewIdentityAliasIdRequest instantiates a new IdentityAliasIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAliasIdRequestWithDefaults

`func NewIdentityAliasIdRequestWithDefaults() *IdentityAliasIdRequest`

NewIdentityAliasIdRequestWithDefaults instantiates a new IdentityAliasIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonicalId

`func (o *IdentityAliasIdRequest) GetCanonicalId() string`

GetCanonicalId returns the CanonicalId field if non-nil, zero value otherwise.

### GetCanonicalIdOk

`func (o *IdentityAliasIdRequest) GetCanonicalIdOk() (*string, bool)`

GetCanonicalIdOk returns a tuple with the CanonicalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalId

`func (o *IdentityAliasIdRequest) SetCanonicalId(v string)`

SetCanonicalId sets CanonicalId field to given value.

### HasCanonicalId

`func (o *IdentityAliasIdRequest) HasCanonicalId() bool`

HasCanonicalId returns a boolean if a field has been set.

### GetEntityId

`func (o *IdentityAliasIdRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *IdentityAliasIdRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *IdentityAliasIdRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *IdentityAliasIdRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetMountAccessor

`func (o *IdentityAliasIdRequest) GetMountAccessor() string`

GetMountAccessor returns the MountAccessor field if non-nil, zero value otherwise.

### GetMountAccessorOk

`func (o *IdentityAliasIdRequest) GetMountAccessorOk() (*string, bool)`

GetMountAccessorOk returns a tuple with the MountAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountAccessor

`func (o *IdentityAliasIdRequest) SetMountAccessor(v string)`

SetMountAccessor sets MountAccessor field to given value.

### HasMountAccessor

`func (o *IdentityAliasIdRequest) HasMountAccessor() bool`

HasMountAccessor returns a boolean if a field has been set.

### GetName

`func (o *IdentityAliasIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityAliasIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityAliasIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityAliasIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


