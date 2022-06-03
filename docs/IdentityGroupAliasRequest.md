# IdentityGroupAliasRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanonicalId** | Pointer to **string** | ID of the group to which this is an alias. | [optional] 
**Id** | Pointer to **string** | ID of the group alias. | [optional] 
**MountAccessor** | Pointer to **string** | Mount accessor to which this alias belongs to. | [optional] 
**Name** | Pointer to **string** | Alias of the group. | [optional] 

## Methods

### NewIdentityGroupAliasRequest

`func NewIdentityGroupAliasRequest() *IdentityGroupAliasRequest`

NewIdentityGroupAliasRequest instantiates a new IdentityGroupAliasRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityGroupAliasRequestWithDefaults

`func NewIdentityGroupAliasRequestWithDefaults() *IdentityGroupAliasRequest`

NewIdentityGroupAliasRequestWithDefaults instantiates a new IdentityGroupAliasRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonicalId

`func (o *IdentityGroupAliasRequest) GetCanonicalId() string`

GetCanonicalId returns the CanonicalId field if non-nil, zero value otherwise.

### GetCanonicalIdOk

`func (o *IdentityGroupAliasRequest) GetCanonicalIdOk() (*string, bool)`

GetCanonicalIdOk returns a tuple with the CanonicalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalId

`func (o *IdentityGroupAliasRequest) SetCanonicalId(v string)`

SetCanonicalId sets CanonicalId field to given value.

### HasCanonicalId

`func (o *IdentityGroupAliasRequest) HasCanonicalId() bool`

HasCanonicalId returns a boolean if a field has been set.

### GetId

`func (o *IdentityGroupAliasRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityGroupAliasRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityGroupAliasRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityGroupAliasRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMountAccessor

`func (o *IdentityGroupAliasRequest) GetMountAccessor() string`

GetMountAccessor returns the MountAccessor field if non-nil, zero value otherwise.

### GetMountAccessorOk

`func (o *IdentityGroupAliasRequest) GetMountAccessorOk() (*string, bool)`

GetMountAccessorOk returns a tuple with the MountAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountAccessor

`func (o *IdentityGroupAliasRequest) SetMountAccessor(v string)`

SetMountAccessor sets MountAccessor field to given value.

### HasMountAccessor

`func (o *IdentityGroupAliasRequest) HasMountAccessor() bool`

HasMountAccessor returns a boolean if a field has been set.

### GetName

`func (o *IdentityGroupAliasRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityGroupAliasRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityGroupAliasRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityGroupAliasRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


