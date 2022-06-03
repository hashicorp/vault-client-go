# TransitDecryptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | Pointer to **string** | The ciphertext to decrypt, provided as returned by encrypt. | [optional] 
**Context** | Pointer to **string** | Base64 encoded context for key derivation. Required if key derivation is enabled. | [optional] 
**Nonce** | Pointer to **string** | Base64 encoded nonce value used during encryption. Must be provided if convergent encryption is enabled for this key and the key was generated with Vault 0.6.1. Not required for keys created in 0.6.2+. | [optional] 

## Methods

### NewTransitDecryptRequest

`func NewTransitDecryptRequest() *TransitDecryptRequest`

NewTransitDecryptRequest instantiates a new TransitDecryptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitDecryptRequestWithDefaults

`func NewTransitDecryptRequestWithDefaults() *TransitDecryptRequest`

NewTransitDecryptRequestWithDefaults instantiates a new TransitDecryptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphertext

`func (o *TransitDecryptRequest) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *TransitDecryptRequest) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *TransitDecryptRequest) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.

### HasCiphertext

`func (o *TransitDecryptRequest) HasCiphertext() bool`

HasCiphertext returns a boolean if a field has been set.

### GetContext

`func (o *TransitDecryptRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TransitDecryptRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TransitDecryptRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *TransitDecryptRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetNonce

`func (o *TransitDecryptRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *TransitDecryptRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *TransitDecryptRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *TransitDecryptRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


