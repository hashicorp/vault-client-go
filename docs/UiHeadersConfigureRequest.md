# UiHeadersConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Multivalue** | Pointer to **bool** | Returns multiple values if true | [optional] 
**Values** | Pointer to **[]string** | The values to set the header. | [optional] 



## Methods


### NewUiHeadersConfigureRequest

`func NewUiHeadersConfigureRequest() *UiHeadersConfigureRequest`

NewUiHeadersConfigureRequest instantiates a new UiHeadersConfigureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiHeadersConfigureRequestWithDefaults

`func NewUiHeadersConfigureRequestWithDefaults() *UiHeadersConfigureRequest`

NewUiHeadersConfigureRequestWithDefaults instantiates a new UiHeadersConfigureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetMultivalue

`func (o *UiHeadersConfigureRequest) GetMultivalue() bool`

GetMultivalue returns the Multivalue field if non-nil, zero value otherwise.

### GetMultivalueOk

`func (o *UiHeadersConfigureRequest) GetMultivalueOk() (*bool, bool)`

GetMultivalueOk returns a tuple with the Multivalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultivalue

`func (o *UiHeadersConfigureRequest) SetMultivalue(v bool)`

SetMultivalue sets Multivalue field to given value.


### HasMultivalue

`func (o *UiHeadersConfigureRequest) HasMultivalue() bool`

HasMultivalue returns a boolean if a field has been set.




### GetValues

`func (o *UiHeadersConfigureRequest) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *UiHeadersConfigureRequest) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *UiHeadersConfigureRequest) SetValues(v []string)`

SetValues sets Values field to given value.


### HasValues

`func (o *UiHeadersConfigureRequest) HasValues() bool`

HasValues returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


