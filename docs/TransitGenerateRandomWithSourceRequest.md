# TransitGenerateRandomWithSourceRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **int32** | The number of bytes to generate (POST body parameter). Defaults to 32 (256 bits). | [optional] [default to 32]
**Format** | Pointer to **string** | Encoding format to use. Can be \&quot;hex\&quot; or \&quot;base64\&quot;. Defaults to \&quot;base64\&quot;. | [optional] [default to "base64"]
**Urlbytes** | Pointer to **string** | The number of bytes to generate (POST URL parameter) | [optional] 



## Methods


### NewTransitGenerateRandomWithSourceRequest

`func NewTransitGenerateRandomWithSourceRequest() *TransitGenerateRandomWithSourceRequest`

NewTransitGenerateRandomWithSourceRequest instantiates a new TransitGenerateRandomWithSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGenerateRandomWithSourceRequestWithDefaults

`func NewTransitGenerateRandomWithSourceRequestWithDefaults() *TransitGenerateRandomWithSourceRequest`

NewTransitGenerateRandomWithSourceRequestWithDefaults instantiates a new TransitGenerateRandomWithSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBytes

`func (o *TransitGenerateRandomWithSourceRequest) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *TransitGenerateRandomWithSourceRequest) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *TransitGenerateRandomWithSourceRequest) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### HasBytes

`func (o *TransitGenerateRandomWithSourceRequest) HasBytes() bool`

HasBytes returns a boolean if a field has been set.




### GetFormat

`func (o *TransitGenerateRandomWithSourceRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TransitGenerateRandomWithSourceRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TransitGenerateRandomWithSourceRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### HasFormat

`func (o *TransitGenerateRandomWithSourceRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.




### GetUrlbytes

`func (o *TransitGenerateRandomWithSourceRequest) GetUrlbytes() string`

GetUrlbytes returns the Urlbytes field if non-nil, zero value otherwise.

### GetUrlbytesOk

`func (o *TransitGenerateRandomWithSourceRequest) GetUrlbytesOk() (*string, bool)`

GetUrlbytesOk returns a tuple with the Urlbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlbytes

`func (o *TransitGenerateRandomWithSourceRequest) SetUrlbytes(v string)`

SetUrlbytes sets Urlbytes field to given value.


### HasUrlbytes

`func (o *TransitGenerateRandomWithSourceRequest) HasUrlbytes() bool`

HasUrlbytes returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


