# TransitRewrapRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | Pointer to **string** | Ciphertext value to rewrap | [optional] 
**Context** | Pointer to **string** | Base64 encoded context for key derivation. Required for derived keys. | [optional] 
**KeyVersion** | Pointer to **int32** | The version of the key to use for encryption. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key. | [optional] 
**Nonce** | Pointer to **string** | Nonce for when convergent encryption is used | [optional] 



## Methods


### NewTransitRewrapRequest

`func NewTransitRewrapRequest() *TransitRewrapRequest`

NewTransitRewrapRequest instantiates a new TransitRewrapRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitRewrapRequestWithDefaults

`func NewTransitRewrapRequestWithDefaults() *TransitRewrapRequest`

NewTransitRewrapRequestWithDefaults instantiates a new TransitRewrapRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCiphertext

`func (o *TransitRewrapRequest) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *TransitRewrapRequest) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *TransitRewrapRequest) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.


### HasCiphertext

`func (o *TransitRewrapRequest) HasCiphertext() bool`

HasCiphertext returns a boolean if a field has been set.




### GetContext

`func (o *TransitRewrapRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TransitRewrapRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TransitRewrapRequest) SetContext(v string)`

SetContext sets Context field to given value.


### HasContext

`func (o *TransitRewrapRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.




### GetKeyVersion

`func (o *TransitRewrapRequest) GetKeyVersion() int32`

GetKeyVersion returns the KeyVersion field if non-nil, zero value otherwise.

### GetKeyVersionOk

`func (o *TransitRewrapRequest) GetKeyVersionOk() (*int32, bool)`

GetKeyVersionOk returns a tuple with the KeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersion

`func (o *TransitRewrapRequest) SetKeyVersion(v int32)`

SetKeyVersion sets KeyVersion field to given value.


### HasKeyVersion

`func (o *TransitRewrapRequest) HasKeyVersion() bool`

HasKeyVersion returns a boolean if a field has been set.




### GetNonce

`func (o *TransitRewrapRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *TransitRewrapRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *TransitRewrapRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### HasNonce

`func (o *TransitRewrapRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


