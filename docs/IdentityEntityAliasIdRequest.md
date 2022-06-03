# IdentityEntityAliasIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanonicalId** | Pointer to **string** | Entity ID to which this alias should be tied to | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** | User provided key-value pairs | [optional] 
**EntityId** | Pointer to **string** | Entity ID to which this alias belongs to. This field is deprecated, use canonical_id. | [optional] 
**MountAccessor** | Pointer to **string** | (Unused) | [optional] 
**Name** | Pointer to **string** | (Unused) | [optional] 

## Methods

### NewIdentityEntityAliasIdRequest

`func NewIdentityEntityAliasIdRequest() *IdentityEntityAliasIdRequest`

NewIdentityEntityAliasIdRequest instantiates a new IdentityEntityAliasIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityEntityAliasIdRequestWithDefaults

`func NewIdentityEntityAliasIdRequestWithDefaults() *IdentityEntityAliasIdRequest`

NewIdentityEntityAliasIdRequestWithDefaults instantiates a new IdentityEntityAliasIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonicalId

`func (o *IdentityEntityAliasIdRequest) GetCanonicalId() string`

GetCanonicalId returns the CanonicalId field if non-nil, zero value otherwise.

### GetCanonicalIdOk

`func (o *IdentityEntityAliasIdRequest) GetCanonicalIdOk() (*string, bool)`

GetCanonicalIdOk returns a tuple with the CanonicalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalId

`func (o *IdentityEntityAliasIdRequest) SetCanonicalId(v string)`

SetCanonicalId sets CanonicalId field to given value.

### HasCanonicalId

`func (o *IdentityEntityAliasIdRequest) HasCanonicalId() bool`

HasCanonicalId returns a boolean if a field has been set.

### GetCustomMetadata

`func (o *IdentityEntityAliasIdRequest) GetCustomMetadata() map[string]interface{}`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *IdentityEntityAliasIdRequest) GetCustomMetadataOk() (*map[string]interface{}, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *IdentityEntityAliasIdRequest) SetCustomMetadata(v map[string]interface{})`

SetCustomMetadata sets CustomMetadata field to given value.

### HasCustomMetadata

`func (o *IdentityEntityAliasIdRequest) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.

### GetEntityId

`func (o *IdentityEntityAliasIdRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *IdentityEntityAliasIdRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *IdentityEntityAliasIdRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *IdentityEntityAliasIdRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetMountAccessor

`func (o *IdentityEntityAliasIdRequest) GetMountAccessor() string`

GetMountAccessor returns the MountAccessor field if non-nil, zero value otherwise.

### GetMountAccessorOk

`func (o *IdentityEntityAliasIdRequest) GetMountAccessorOk() (*string, bool)`

GetMountAccessorOk returns a tuple with the MountAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountAccessor

`func (o *IdentityEntityAliasIdRequest) SetMountAccessor(v string)`

SetMountAccessor sets MountAccessor field to given value.

### HasMountAccessor

`func (o *IdentityEntityAliasIdRequest) HasMountAccessor() bool`

HasMountAccessor returns a boolean if a field has been set.

### GetName

`func (o *IdentityEntityAliasIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityEntityAliasIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityEntityAliasIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityEntityAliasIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


