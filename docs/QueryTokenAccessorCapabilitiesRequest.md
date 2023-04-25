# QueryTokenAccessorCapabilitiesRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessor** | Pointer to **string** | Accessor of the token for which capabilities are being queried. | [optional] 
**Path** | Pointer to **[]string** | Use &#x27;paths&#x27; instead. | [optional] 
**Paths** | Pointer to **[]string** | Paths on which capabilities are being queried. | [optional] 



## Methods


### NewQueryTokenAccessorCapabilitiesRequest

`func NewQueryTokenAccessorCapabilitiesRequest() *QueryTokenAccessorCapabilitiesRequest`

NewQueryTokenAccessorCapabilitiesRequest instantiates a new QueryTokenAccessorCapabilitiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryTokenAccessorCapabilitiesRequestWithDefaults

`func NewQueryTokenAccessorCapabilitiesRequestWithDefaults() *QueryTokenAccessorCapabilitiesRequest`

NewQueryTokenAccessorCapabilitiesRequestWithDefaults instantiates a new QueryTokenAccessorCapabilitiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAccessor

`func (o *QueryTokenAccessorCapabilitiesRequest) GetAccessor() string`

GetAccessor returns the Accessor field if non-nil, zero value otherwise.

### GetAccessorOk

`func (o *QueryTokenAccessorCapabilitiesRequest) GetAccessorOk() (*string, bool)`

GetAccessorOk returns a tuple with the Accessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessor

`func (o *QueryTokenAccessorCapabilitiesRequest) SetAccessor(v string)`

SetAccessor sets Accessor field to given value.


### HasAccessor

`func (o *QueryTokenAccessorCapabilitiesRequest) HasAccessor() bool`

HasAccessor returns a boolean if a field has been set.




### GetPath

`func (o *QueryTokenAccessorCapabilitiesRequest) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *QueryTokenAccessorCapabilitiesRequest) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *QueryTokenAccessorCapabilitiesRequest) SetPath(v []string)`

SetPath sets Path field to given value.


### HasPath

`func (o *QueryTokenAccessorCapabilitiesRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.




### GetPaths

`func (o *QueryTokenAccessorCapabilitiesRequest) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *QueryTokenAccessorCapabilitiesRequest) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *QueryTokenAccessorCapabilitiesRequest) SetPaths(v []string)`

SetPaths sets Paths field to given value.


### HasPaths

`func (o *QueryTokenAccessorCapabilitiesRequest) HasPaths() bool`

HasPaths returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


