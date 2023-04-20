# EntityLookUpRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasId** | Pointer to **string** | ID of the alias. | [optional] 
**AliasMountAccessor** | Pointer to **string** | Accessor of the mount to which the alias belongs to. This should be supplied in conjunction with &#x27;alias_name&#x27;. | [optional] 
**AliasName** | Pointer to **string** | Name of the alias. This should be supplied in conjunction with &#x27;alias_mount_accessor&#x27;. | [optional] 
**Id** | Pointer to **string** | ID of the entity. | [optional] 
**Name** | Pointer to **string** | Name of the entity. | [optional] 



## Methods


### NewEntityLookUpRequest

`func NewEntityLookUpRequest() *EntityLookUpRequest`

NewEntityLookUpRequest instantiates a new EntityLookUpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityLookUpRequestWithDefaults

`func NewEntityLookUpRequestWithDefaults() *EntityLookUpRequest`

NewEntityLookUpRequestWithDefaults instantiates a new EntityLookUpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAliasId

`func (o *EntityLookUpRequest) GetAliasId() string`

GetAliasId returns the AliasId field if non-nil, zero value otherwise.

### GetAliasIdOk

`func (o *EntityLookUpRequest) GetAliasIdOk() (*string, bool)`

GetAliasIdOk returns a tuple with the AliasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasId

`func (o *EntityLookUpRequest) SetAliasId(v string)`

SetAliasId sets AliasId field to given value.


### HasAliasId

`func (o *EntityLookUpRequest) HasAliasId() bool`

HasAliasId returns a boolean if a field has been set.




### GetAliasMountAccessor

`func (o *EntityLookUpRequest) GetAliasMountAccessor() string`

GetAliasMountAccessor returns the AliasMountAccessor field if non-nil, zero value otherwise.

### GetAliasMountAccessorOk

`func (o *EntityLookUpRequest) GetAliasMountAccessorOk() (*string, bool)`

GetAliasMountAccessorOk returns a tuple with the AliasMountAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasMountAccessor

`func (o *EntityLookUpRequest) SetAliasMountAccessor(v string)`

SetAliasMountAccessor sets AliasMountAccessor field to given value.


### HasAliasMountAccessor

`func (o *EntityLookUpRequest) HasAliasMountAccessor() bool`

HasAliasMountAccessor returns a boolean if a field has been set.




### GetAliasName

`func (o *EntityLookUpRequest) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *EntityLookUpRequest) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *EntityLookUpRequest) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.


### HasAliasName

`func (o *EntityLookUpRequest) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.




### GetId

`func (o *EntityLookUpRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityLookUpRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityLookUpRequest) SetId(v string)`

SetId sets Id field to given value.


### HasId

`func (o *EntityLookUpRequest) HasId() bool`

HasId returns a boolean if a field has been set.




### GetName

`func (o *EntityLookUpRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityLookUpRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityLookUpRequest) SetName(v string)`

SetName sets Name field to given value.


### HasName

`func (o *EntityLookUpRequest) HasName() bool`

HasName returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


