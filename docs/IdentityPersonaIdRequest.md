# IdentityPersonaIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** | Entity ID to which this persona should be tied to | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata to be associated with the persona. In CLI, this parameter can be repeated multiple times, and it all gets merged together. For example: vault &lt;command&gt; &lt;path&gt; metadata&#x3D;key1&#x3D;value1 metadata&#x3D;key2&#x3D;value2 | [optional] 
**MountAccessor** | Pointer to **string** | Mount accessor to which this persona belongs to | [optional] 
**Name** | Pointer to **string** | Name of the persona | [optional] 

## Methods

### NewIdentityPersonaIdRequest

`func NewIdentityPersonaIdRequest() *IdentityPersonaIdRequest`

NewIdentityPersonaIdRequest instantiates a new IdentityPersonaIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityPersonaIdRequestWithDefaults

`func NewIdentityPersonaIdRequestWithDefaults() *IdentityPersonaIdRequest`

NewIdentityPersonaIdRequestWithDefaults instantiates a new IdentityPersonaIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *IdentityPersonaIdRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *IdentityPersonaIdRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *IdentityPersonaIdRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *IdentityPersonaIdRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetMetadata

`func (o *IdentityPersonaIdRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IdentityPersonaIdRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IdentityPersonaIdRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IdentityPersonaIdRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMountAccessor

`func (o *IdentityPersonaIdRequest) GetMountAccessor() string`

GetMountAccessor returns the MountAccessor field if non-nil, zero value otherwise.

### GetMountAccessorOk

`func (o *IdentityPersonaIdRequest) GetMountAccessorOk() (*string, bool)`

GetMountAccessorOk returns a tuple with the MountAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountAccessor

`func (o *IdentityPersonaIdRequest) SetMountAccessor(v string)`

SetMountAccessor sets MountAccessor field to given value.

### HasMountAccessor

`func (o *IdentityPersonaIdRequest) HasMountAccessor() bool`

HasMountAccessor returns a boolean if a field has been set.

### GetName

`func (o *IdentityPersonaIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityPersonaIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityPersonaIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityPersonaIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


