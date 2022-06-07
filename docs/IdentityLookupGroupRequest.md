# IdentityLookupGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasId** | Pointer to **string** | ID of the alias. | [optional] 
**AliasMountAccessor** | Pointer to **string** | Accessor of the mount to which the alias belongs to. This should be supplied in conjunction with &#39;alias_name&#39;. | [optional] 
**AliasName** | Pointer to **string** | Name of the alias. This should be supplied in conjunction with &#39;alias_mount_accessor&#39;. | [optional] 
**Id** | Pointer to **string** | ID of the group. | [optional] 
**Name** | Pointer to **string** | Name of the group. | [optional] 

## Methods

### NewIdentityLookupGroupRequest

`func NewIdentityLookupGroupRequest() *IdentityLookupGroupRequest`

NewIdentityLookupGroupRequest instantiates a new IdentityLookupGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityLookupGroupRequestWithDefaults

`func NewIdentityLookupGroupRequestWithDefaults() *IdentityLookupGroupRequest`

NewIdentityLookupGroupRequestWithDefaults instantiates a new IdentityLookupGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliasId

`func (o *IdentityLookupGroupRequest) GetAliasId() string`

GetAliasId returns the AliasId field if non-nil, zero value otherwise.

### GetAliasIdOk

`func (o *IdentityLookupGroupRequest) GetAliasIdOk() (*string, bool)`

GetAliasIdOk returns a tuple with the AliasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasId

`func (o *IdentityLookupGroupRequest) SetAliasId(v string)`

SetAliasId sets AliasId field to given value.

### HasAliasId

`func (o *IdentityLookupGroupRequest) HasAliasId() bool`

HasAliasId returns a boolean if a field has been set.

### GetAliasMountAccessor

`func (o *IdentityLookupGroupRequest) GetAliasMountAccessor() string`

GetAliasMountAccessor returns the AliasMountAccessor field if non-nil, zero value otherwise.

### GetAliasMountAccessorOk

`func (o *IdentityLookupGroupRequest) GetAliasMountAccessorOk() (*string, bool)`

GetAliasMountAccessorOk returns a tuple with the AliasMountAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasMountAccessor

`func (o *IdentityLookupGroupRequest) SetAliasMountAccessor(v string)`

SetAliasMountAccessor sets AliasMountAccessor field to given value.

### HasAliasMountAccessor

`func (o *IdentityLookupGroupRequest) HasAliasMountAccessor() bool`

HasAliasMountAccessor returns a boolean if a field has been set.

### GetAliasName

`func (o *IdentityLookupGroupRequest) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *IdentityLookupGroupRequest) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *IdentityLookupGroupRequest) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *IdentityLookupGroupRequest) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.

### GetId

`func (o *IdentityLookupGroupRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityLookupGroupRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityLookupGroupRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityLookupGroupRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdentityLookupGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityLookupGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityLookupGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityLookupGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


