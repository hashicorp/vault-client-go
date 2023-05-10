# TransitGenerateRandomWithBytesRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **int32** | The number of bytes to generate (POST body parameter). Defaults to 32 (256 bits). | [optional] [default to 32]
**Format** | Pointer to **string** | Encoding format to use. Can be \&quot;hex\&quot; or \&quot;base64\&quot;. Defaults to \&quot;base64\&quot;. | [optional] [default to "base64"]
**Source** | Pointer to **string** | Which system to source random data from, ether \&quot;platform\&quot;, \&quot;seal\&quot;, or \&quot;all\&quot;. | [optional] [default to "platform"]



## Methods


### NewTransitGenerateRandomWithBytesRequest

`func NewTransitGenerateRandomWithBytesRequest() *TransitGenerateRandomWithBytesRequest`

NewTransitGenerateRandomWithBytesRequest instantiates a new TransitGenerateRandomWithBytesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGenerateRandomWithBytesRequestWithDefaults

`func NewTransitGenerateRandomWithBytesRequestWithDefaults() *TransitGenerateRandomWithBytesRequest`

NewTransitGenerateRandomWithBytesRequestWithDefaults instantiates a new TransitGenerateRandomWithBytesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBytes

`func (o *TransitGenerateRandomWithBytesRequest) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *TransitGenerateRandomWithBytesRequest) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *TransitGenerateRandomWithBytesRequest) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### HasBytes

`func (o *TransitGenerateRandomWithBytesRequest) HasBytes() bool`

HasBytes returns a boolean if a field has been set.




### GetFormat

`func (o *TransitGenerateRandomWithBytesRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TransitGenerateRandomWithBytesRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TransitGenerateRandomWithBytesRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### HasFormat

`func (o *TransitGenerateRandomWithBytesRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.




### GetSource

`func (o *TransitGenerateRandomWithBytesRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TransitGenerateRandomWithBytesRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TransitGenerateRandomWithBytesRequest) SetSource(v string)`

SetSource sets Source field to given value.


### HasSource

`func (o *TransitGenerateRandomWithBytesRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


