# TransitHashRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Algorithm to use (POST body parameter). Valid values are: * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 Defaults to \&quot;sha2-256\&quot;. | [optional] [default to "sha2-256"]
**Format** | Pointer to **string** | Encoding format to use. Can be \&quot;hex\&quot; or \&quot;base64\&quot;. Defaults to \&quot;hex\&quot;. | [optional] [default to "hex"]
**Input** | Pointer to **string** | The base64-encoded input data | [optional] 
**Urlalgorithm** | Pointer to **string** | Algorithm to use (POST URL parameter) | [optional] 

## Methods

### NewTransitHashRequest

`func NewTransitHashRequest() *TransitHashRequest`

NewTransitHashRequest instantiates a new TransitHashRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitHashRequestWithDefaults

`func NewTransitHashRequestWithDefaults() *TransitHashRequest`

NewTransitHashRequestWithDefaults instantiates a new TransitHashRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *TransitHashRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *TransitHashRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *TransitHashRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *TransitHashRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetFormat

`func (o *TransitHashRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TransitHashRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TransitHashRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *TransitHashRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetInput

`func (o *TransitHashRequest) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *TransitHashRequest) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *TransitHashRequest) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *TransitHashRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetUrlalgorithm

`func (o *TransitHashRequest) GetUrlalgorithm() string`

GetUrlalgorithm returns the Urlalgorithm field if non-nil, zero value otherwise.

### GetUrlalgorithmOk

`func (o *TransitHashRequest) GetUrlalgorithmOk() (*string, bool)`

GetUrlalgorithmOk returns a tuple with the Urlalgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlalgorithm

`func (o *TransitHashRequest) SetUrlalgorithm(v string)`

SetUrlalgorithm sets Urlalgorithm field to given value.

### HasUrlalgorithm

`func (o *TransitHashRequest) HasUrlalgorithm() bool`

HasUrlalgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


