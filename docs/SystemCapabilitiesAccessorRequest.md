# SystemCapabilitiesAccessorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessor** | Pointer to **string** | Accessor of the token for which capabilities are being queried. | [optional] 
**Path** | Pointer to **[]string** | Use &#39;paths&#39; instead. | [optional] 
**Paths** | Pointer to **[]string** | Paths on which capabilities are being queried. | [optional] 

## Methods

### NewSystemCapabilitiesAccessorRequest

`func NewSystemCapabilitiesAccessorRequest() *SystemCapabilitiesAccessorRequest`

NewSystemCapabilitiesAccessorRequest instantiates a new SystemCapabilitiesAccessorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemCapabilitiesAccessorRequestWithDefaults

`func NewSystemCapabilitiesAccessorRequestWithDefaults() *SystemCapabilitiesAccessorRequest`

NewSystemCapabilitiesAccessorRequestWithDefaults instantiates a new SystemCapabilitiesAccessorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessor

`func (o *SystemCapabilitiesAccessorRequest) GetAccessor() string`

GetAccessor returns the Accessor field if non-nil, zero value otherwise.

### GetAccessorOk

`func (o *SystemCapabilitiesAccessorRequest) GetAccessorOk() (*string, bool)`

GetAccessorOk returns a tuple with the Accessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessor

`func (o *SystemCapabilitiesAccessorRequest) SetAccessor(v string)`

SetAccessor sets Accessor field to given value.

### HasAccessor

`func (o *SystemCapabilitiesAccessorRequest) HasAccessor() bool`

HasAccessor returns a boolean if a field has been set.

### GetPath

`func (o *SystemCapabilitiesAccessorRequest) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SystemCapabilitiesAccessorRequest) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SystemCapabilitiesAccessorRequest) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SystemCapabilitiesAccessorRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPaths

`func (o *SystemCapabilitiesAccessorRequest) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *SystemCapabilitiesAccessorRequest) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *SystemCapabilitiesAccessorRequest) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *SystemCapabilitiesAccessorRequest) HasPaths() bool`

HasPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


