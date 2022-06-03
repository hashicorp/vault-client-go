# SystemConfigUiHeadersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Multivalue** | Pointer to **bool** | Returns multiple values if true | [optional] 
**Values** | Pointer to **[]string** | The values to set the header. | [optional] 

## Methods

### NewSystemConfigUiHeadersRequest

`func NewSystemConfigUiHeadersRequest() *SystemConfigUiHeadersRequest`

NewSystemConfigUiHeadersRequest instantiates a new SystemConfigUiHeadersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemConfigUiHeadersRequestWithDefaults

`func NewSystemConfigUiHeadersRequestWithDefaults() *SystemConfigUiHeadersRequest`

NewSystemConfigUiHeadersRequestWithDefaults instantiates a new SystemConfigUiHeadersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMultivalue

`func (o *SystemConfigUiHeadersRequest) GetMultivalue() bool`

GetMultivalue returns the Multivalue field if non-nil, zero value otherwise.

### GetMultivalueOk

`func (o *SystemConfigUiHeadersRequest) GetMultivalueOk() (*bool, bool)`

GetMultivalueOk returns a tuple with the Multivalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultivalue

`func (o *SystemConfigUiHeadersRequest) SetMultivalue(v bool)`

SetMultivalue sets Multivalue field to given value.

### HasMultivalue

`func (o *SystemConfigUiHeadersRequest) HasMultivalue() bool`

HasMultivalue returns a boolean if a field has been set.

### GetValues

`func (o *SystemConfigUiHeadersRequest) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SystemConfigUiHeadersRequest) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SystemConfigUiHeadersRequest) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *SystemConfigUiHeadersRequest) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


