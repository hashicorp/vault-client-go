# GroupUpdateAliasByIdRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanonicalId** | Pointer to **string** | ID of the group to which this is an alias. | [optional] 
**MountAccessor** | Pointer to **string** | Mount accessor to which this alias belongs to. | [optional] 
**Name** | Pointer to **string** | Alias of the group. | [optional] 



## Methods


### NewGroupUpdateAliasByIdRequest

`func NewGroupUpdateAliasByIdRequest() *GroupUpdateAliasByIdRequest`

NewGroupUpdateAliasByIdRequest instantiates a new GroupUpdateAliasByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUpdateAliasByIdRequestWithDefaults

`func NewGroupUpdateAliasByIdRequestWithDefaults() *GroupUpdateAliasByIdRequest`

NewGroupUpdateAliasByIdRequestWithDefaults instantiates a new GroupUpdateAliasByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCanonicalId

`func (o *GroupUpdateAliasByIdRequest) GetCanonicalId() string`

GetCanonicalId returns the CanonicalId field if non-nil, zero value otherwise.

### GetCanonicalIdOk

`func (o *GroupUpdateAliasByIdRequest) GetCanonicalIdOk() (*string, bool)`

GetCanonicalIdOk returns a tuple with the CanonicalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalId

`func (o *GroupUpdateAliasByIdRequest) SetCanonicalId(v string)`

SetCanonicalId sets CanonicalId field to given value.


### HasCanonicalId

`func (o *GroupUpdateAliasByIdRequest) HasCanonicalId() bool`

HasCanonicalId returns a boolean if a field has been set.




### GetMountAccessor

`func (o *GroupUpdateAliasByIdRequest) GetMountAccessor() string`

GetMountAccessor returns the MountAccessor field if non-nil, zero value otherwise.

### GetMountAccessorOk

`func (o *GroupUpdateAliasByIdRequest) GetMountAccessorOk() (*string, bool)`

GetMountAccessorOk returns a tuple with the MountAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountAccessor

`func (o *GroupUpdateAliasByIdRequest) SetMountAccessor(v string)`

SetMountAccessor sets MountAccessor field to given value.


### HasMountAccessor

`func (o *GroupUpdateAliasByIdRequest) HasMountAccessor() bool`

HasMountAccessor returns a boolean if a field has been set.




### GetName

`func (o *GroupUpdateAliasByIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupUpdateAliasByIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupUpdateAliasByIdRequest) SetName(v string)`

SetName sets Name field to given value.


### HasName

`func (o *GroupUpdateAliasByIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


