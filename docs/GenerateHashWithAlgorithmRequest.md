# GenerateHashWithAlgorithmRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Algorithm to use (POST body parameter). Valid values are: * sha2-224 * sha2-256 * sha2-384 * sha2-512 Defaults to \&quot;sha2-256\&quot;. | [optional] [default to "sha2-256"]
**Format** | Pointer to **string** | Encoding format to use. Can be \&quot;hex\&quot; or \&quot;base64\&quot;. Defaults to \&quot;hex\&quot;. | [optional] [default to "hex"]
**Input** | Pointer to **string** | The base64-encoded input data | [optional] 



## Methods


### NewGenerateHashWithAlgorithmRequest

`func NewGenerateHashWithAlgorithmRequest() *GenerateHashWithAlgorithmRequest`

NewGenerateHashWithAlgorithmRequest instantiates a new GenerateHashWithAlgorithmRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateHashWithAlgorithmRequestWithDefaults

`func NewGenerateHashWithAlgorithmRequestWithDefaults() *GenerateHashWithAlgorithmRequest`

NewGenerateHashWithAlgorithmRequestWithDefaults instantiates a new GenerateHashWithAlgorithmRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAlgorithm

`func (o *GenerateHashWithAlgorithmRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *GenerateHashWithAlgorithmRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *GenerateHashWithAlgorithmRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### HasAlgorithm

`func (o *GenerateHashWithAlgorithmRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.




### GetFormat

`func (o *GenerateHashWithAlgorithmRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GenerateHashWithAlgorithmRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GenerateHashWithAlgorithmRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### HasFormat

`func (o *GenerateHashWithAlgorithmRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.




### GetInput

`func (o *GenerateHashWithAlgorithmRequest) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *GenerateHashWithAlgorithmRequest) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *GenerateHashWithAlgorithmRequest) SetInput(v string)`

SetInput sets Input field to given value.


### HasInput

`func (o *GenerateHashWithAlgorithmRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


