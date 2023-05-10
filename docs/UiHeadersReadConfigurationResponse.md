# UiHeadersReadConfigurationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | returns the first header value when &#x60;multivalue&#x60; request parameter is false | [optional] 
**Values** | Pointer to **[]string** | returns all header values when &#x60;multivalue&#x60; request parameter is true | [optional] 



## Methods


### NewUiHeadersReadConfigurationResponse

`func NewUiHeadersReadConfigurationResponse() *UiHeadersReadConfigurationResponse`

NewUiHeadersReadConfigurationResponse instantiates a new UiHeadersReadConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiHeadersReadConfigurationResponseWithDefaults

`func NewUiHeadersReadConfigurationResponseWithDefaults() *UiHeadersReadConfigurationResponse`

NewUiHeadersReadConfigurationResponseWithDefaults instantiates a new UiHeadersReadConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetValue

`func (o *UiHeadersReadConfigurationResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UiHeadersReadConfigurationResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UiHeadersReadConfigurationResponse) SetValue(v string)`

SetValue sets Value field to given value.


### HasValue

`func (o *UiHeadersReadConfigurationResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.




### GetValues

`func (o *UiHeadersReadConfigurationResponse) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *UiHeadersReadConfigurationResponse) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *UiHeadersReadConfigurationResponse) SetValues(v []string)`

SetValues sets Values field to given value.


### HasValues

`func (o *UiHeadersReadConfigurationResponse) HasValues() bool`

HasValues returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


