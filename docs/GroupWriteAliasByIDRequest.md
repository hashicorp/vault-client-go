# GroupWriteAliasByIDRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**CanonicalId** | Pointer to **string** | ID of the group to which this is an alias. | [optional] 
**MountAccessor** | Pointer to **string** | Mount accessor to which this alias belongs to. | [optional] 
**Name** | Pointer to **string** | Alias of the group. | [optional] 



## Methods


### NewGroupWriteAliasByIDRequest

`func NewGroupWriteAliasByIDRequest() *GroupWriteAliasByIDRequest`

NewGroupWriteAliasByIDRequest instantiates a new GroupWriteAliasByIDRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWriteAliasByIDRequestWithDefaults

`func NewGroupWriteAliasByIDRequestWithDefaults() *GroupWriteAliasByIDRequest`

NewGroupWriteAliasByIDRequestWithDefaults instantiates a new GroupWriteAliasByIDRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCanonicalId

`func (o *GroupWriteAliasByIDRequest) GetCanonicalId() string`

GetCanonicalId returns the CanonicalId field if non-nil, zero value otherwise.

### GetCanonicalIdOk

`func (o *GroupWriteAliasByIDRequest) GetCanonicalIdOk() (*string, bool)`

GetCanonicalIdOk returns a tuple with the CanonicalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalId

`func (o *GroupWriteAliasByIDRequest) SetCanonicalId(v string)`

SetCanonicalId sets CanonicalId field to given value.


### HasCanonicalId

`func (o *GroupWriteAliasByIDRequest) HasCanonicalId() bool`

HasCanonicalId returns a boolean if a field has been set.




### GetMountAccessor

`func (o *GroupWriteAliasByIDRequest) GetMountAccessor() string`

GetMountAccessor returns the MountAccessor field if non-nil, zero value otherwise.

### GetMountAccessorOk

`func (o *GroupWriteAliasByIDRequest) GetMountAccessorOk() (*string, bool)`

GetMountAccessorOk returns a tuple with the MountAccessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountAccessor

`func (o *GroupWriteAliasByIDRequest) SetMountAccessor(v string)`

SetMountAccessor sets MountAccessor field to given value.


### HasMountAccessor

`func (o *GroupWriteAliasByIDRequest) HasMountAccessor() bool`

HasMountAccessor returns a boolean if a field has been set.




### GetName

`func (o *GroupWriteAliasByIDRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupWriteAliasByIDRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupWriteAliasByIDRequest) SetName(v string)`

SetName sets Name field to given value.


### HasName

`func (o *GroupWriteAliasByIDRequest) HasName() bool`

HasName returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


