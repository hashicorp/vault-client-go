# WriteCapabilitiesAccessorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessor** | Pointer to **string** | Accessor of the token for which capabilities are being queried. | [optional] 
**Path** | Pointer to **[]string** | Use &#39;paths&#39; instead. | [optional] 
**Paths** | Pointer to **[]string** | Paths on which capabilities are being queried. | [optional] 

## Methods

### NewWriteCapabilitiesAccessorRequest

`func NewWriteCapabilitiesAccessorRequest() *WriteCapabilitiesAccessorRequest`

NewWriteCapabilitiesAccessorRequest instantiates a new WriteCapabilitiesAccessorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteCapabilitiesAccessorRequestWithDefaults

`func NewWriteCapabilitiesAccessorRequestWithDefaults() *WriteCapabilitiesAccessorRequest`

NewWriteCapabilitiesAccessorRequestWithDefaults instantiates a new WriteCapabilitiesAccessorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessor

`func (o *WriteCapabilitiesAccessorRequest) GetAccessor() string`

GetAccessor returns the Accessor field if non-nil, zero value otherwise.

### GetAccessorOk

`func (o *WriteCapabilitiesAccessorRequest) GetAccessorOk() (*string, bool)`

GetAccessorOk returns a tuple with the Accessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessor

`func (o *WriteCapabilitiesAccessorRequest) SetAccessor(v string)`

SetAccessor sets Accessor field to given value.

### HasAccessor

`func (o *WriteCapabilitiesAccessorRequest) HasAccessor() bool`

HasAccessor returns a boolean if a field has been set.

### GetPath

`func (o *WriteCapabilitiesAccessorRequest) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WriteCapabilitiesAccessorRequest) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WriteCapabilitiesAccessorRequest) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *WriteCapabilitiesAccessorRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPaths

`func (o *WriteCapabilitiesAccessorRequest) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *WriteCapabilitiesAccessorRequest) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *WriteCapabilitiesAccessorRequest) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *WriteCapabilitiesAccessorRequest) HasPaths() bool`

HasPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


