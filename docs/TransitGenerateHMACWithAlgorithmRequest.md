# TransitGenerateHMACWithAlgorithmRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Algorithm to use (POST body parameter). Valid values are: * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 Defaults to \&quot;sha2-256\&quot;. | [optional] [default to "sha2-256"]
**Input** | Pointer to **string** | The base64-encoded input data | [optional] 
**KeyVersion** | Pointer to **int32** | The version of the key to use for generating the HMAC. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key. | [optional] 

## Methods

### NewTransitGenerateHMACWithAlgorithmRequest

`func NewTransitGenerateHMACWithAlgorithmRequest() *TransitGenerateHMACWithAlgorithmRequest`

NewTransitGenerateHMACWithAlgorithmRequest instantiates a new TransitGenerateHMACWithAlgorithmRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGenerateHMACWithAlgorithmRequestWithDefaults

`func NewTransitGenerateHMACWithAlgorithmRequestWithDefaults() *TransitGenerateHMACWithAlgorithmRequest`

NewTransitGenerateHMACWithAlgorithmRequestWithDefaults instantiates a new TransitGenerateHMACWithAlgorithmRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *TransitGenerateHMACWithAlgorithmRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *TransitGenerateHMACWithAlgorithmRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *TransitGenerateHMACWithAlgorithmRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *TransitGenerateHMACWithAlgorithmRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetInput

`func (o *TransitGenerateHMACWithAlgorithmRequest) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *TransitGenerateHMACWithAlgorithmRequest) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *TransitGenerateHMACWithAlgorithmRequest) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *TransitGenerateHMACWithAlgorithmRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetKeyVersion

`func (o *TransitGenerateHMACWithAlgorithmRequest) GetKeyVersion() int32`

GetKeyVersion returns the KeyVersion field if non-nil, zero value otherwise.

### GetKeyVersionOk

`func (o *TransitGenerateHMACWithAlgorithmRequest) GetKeyVersionOk() (*int32, bool)`

GetKeyVersionOk returns a tuple with the KeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersion

`func (o *TransitGenerateHMACWithAlgorithmRequest) SetKeyVersion(v int32)`

SetKeyVersion sets KeyVersion field to given value.

### HasKeyVersion

`func (o *TransitGenerateHMACWithAlgorithmRequest) HasKeyVersion() bool`

HasKeyVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


