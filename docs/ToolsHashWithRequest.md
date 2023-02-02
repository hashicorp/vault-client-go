# ToolsHashWithRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Algorithm** | Pointer to **string** | Algorithm to use (POST body parameter). Valid values are: * sha2-224 * sha2-256 * sha2-384 * sha2-512 Defaults to \&quot;sha2-256\&quot;. | [optional] [default to "sha2-256"]
**Format** | Pointer to **string** | Encoding format to use. Can be \&quot;hex\&quot; or \&quot;base64\&quot;. Defaults to \&quot;hex\&quot;. | [optional] [default to "hex"]
**Input** | Pointer to **string** | The base64-encoded input data | [optional] 



## Methods


### NewToolsHashWithRequest

`func NewToolsHashWithRequest() *ToolsHashWithRequest`

NewToolsHashWithRequest instantiates a new ToolsHashWithRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolsHashWithRequestWithDefaults

`func NewToolsHashWithRequestWithDefaults() *ToolsHashWithRequest`

NewToolsHashWithRequestWithDefaults instantiates a new ToolsHashWithRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAlgorithm

`func (o *ToolsHashWithRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *ToolsHashWithRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *ToolsHashWithRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### HasAlgorithm

`func (o *ToolsHashWithRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.




### GetFormat

`func (o *ToolsHashWithRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ToolsHashWithRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ToolsHashWithRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### HasFormat

`func (o *ToolsHashWithRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.




### GetInput

`func (o *ToolsHashWithRequest) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ToolsHashWithRequest) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ToolsHashWithRequest) SetInput(v string)`

SetInput sets Input field to given value.


### HasInput

`func (o *ToolsHashWithRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


