# EntityWriteAliasByIDRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**CanonicalId** | Pointer to **string** | Entity ID to which this alias should be tied to | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** | User provided key-value pairs | [optional] 
**EntityId** | Pointer to **string** | Entity ID to which this alias belongs to. This field is deprecated, use canonical_id. | [optional] 
**MountAccessor** | Pointer to **string** | (Unused) | [optional] 
**Name** | Pointer to **string** | (Unused) | [optional] 



## Methods


### NewEntityWriteAliasByIDRequest

`func NewEntityWriteAliasByIDRequest() *EntityWriteAliasByIDRequest`

NewEntityWriteAliasByIDRequest instantiates a new EntityWriteAliasByIDRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityWriteAliasByIDRequestWithDefaults

`func NewEntityWriteAliasByIDRequestWithDefaults() *EntityWriteAliasByIDRequest`

NewEntityWriteAliasByIDRequestWithDefaults instantiates a new EntityWriteAliasByIDRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCanonicalId

`func (o *EntityWriteAliasByIDRequest) GetCanonicalId() string`

GetCanonicalId returns the CanonicalId field if non-nil, zero value otherwise.

### GetCanonicalIdOk

`func (o *EntityWriteAliasByIDRequest) GetCanonicalIdOk() (*string, bool)`

GetCanonicalIdOk returns a tuple with the CanonicalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalId

`func (o *EntityWriteAliasByIDRequest) SetCanonicalId(v string)`

SetCanonicalId sets CanonicalId field to given value.


### HasCanonicalId

`func (o *EntityWriteAliasByIDRequest) HasCanonicalId() bool`

HasCanonicalId returns a boolean if a field has been set.




### GetCustomMetadata

`func (o *EntityWriteAliasByIDRequest) GetCustomMetadata() map[string]interface{}`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *EntityWriteAliasByIDRequest) GetCustomMetadataOk() (*map[string]interface{}, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *EntityWriteAliasByIDRequest) SetCustomMetadata(v map[string]interface{})`

SetCustomMetadata sets CustomMetadata field to given value.


### HasCustomMetadata

`func (o *EntityWriteAliasByIDRequest) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.




### GetEntityId

`func (o *EntityWriteAliasByIDRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EntityWriteAliasByIDRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EntityWriteAliasByIDRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### HasEntityId

`func (o *EntityWriteAliasByIDRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.




### GetMountAccessor

`func (o *EntityWriteAliasByIDRequest) GetMountAccessor() string`

GetMountAccessor returns the MountAccessor field if non-nil, zero value otherwise.

### GetMountAccessorOk

`func (o *EntityWriteAliasByIDRequest) GetMountAccessorOk() (*string, bool)`

GetMountAccessorOk returns a tuple with the MountAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountAccessor

`func (o *EntityWriteAliasByIDRequest) SetMountAccessor(v string)`

SetMountAccessor sets MountAccessor field to given value.


### HasMountAccessor

`func (o *EntityWriteAliasByIDRequest) HasMountAccessor() bool`

HasMountAccessor returns a boolean if a field has been set.




### GetName

`func (o *EntityWriteAliasByIDRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityWriteAliasByIDRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityWriteAliasByIDRequest) SetName(v string)`

SetName sets Name field to given value.


### HasName

`func (o *EntityWriteAliasByIDRequest) HasName() bool`

HasName returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


