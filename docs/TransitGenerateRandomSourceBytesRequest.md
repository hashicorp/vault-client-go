# TransitGenerateRandomSourceBytesRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**Bytes** | Pointer to **int32** | The number of bytes to generate (POST body parameter). Defaults to 32 (256 bits). | [optional] [default to 32]
**Format** | Pointer to **string** | Encoding format to use. Can be \&quot;hex\&quot; or \&quot;base64\&quot;. Defaults to \&quot;base64\&quot;. | [optional] [default to "base64"]



## Methods


### NewTransitGenerateRandomSourceBytesRequest

`func NewTransitGenerateRandomSourceBytesRequest() *TransitGenerateRandomSourceBytesRequest`

NewTransitGenerateRandomSourceBytesRequest instantiates a new TransitGenerateRandomSourceBytesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGenerateRandomSourceBytesRequestWithDefaults

`func NewTransitGenerateRandomSourceBytesRequestWithDefaults() *TransitGenerateRandomSourceBytesRequest`

NewTransitGenerateRandomSourceBytesRequestWithDefaults instantiates a new TransitGenerateRandomSourceBytesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBytes

`func (o *TransitGenerateRandomSourceBytesRequest) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *TransitGenerateRandomSourceBytesRequest) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *TransitGenerateRandomSourceBytesRequest) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### HasBytes

`func (o *TransitGenerateRandomSourceBytesRequest) HasBytes() bool`

HasBytes returns a boolean if a field has been set.




### GetFormat

`func (o *TransitGenerateRandomSourceBytesRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TransitGenerateRandomSourceBytesRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TransitGenerateRandomSourceBytesRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### HasFormat

`func (o *TransitGenerateRandomSourceBytesRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


