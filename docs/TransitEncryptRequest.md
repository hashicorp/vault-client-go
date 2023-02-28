# TransitEncryptRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedData** | Pointer to **string** | When using an AEAD cipher mode, such as AES-GCM, this parameter allows passing associated data (AD/AAD) into the encryption function; this data must be passed on subsequent decryption requests but can be transited in plaintext. On successful decryption, both the ciphertext and the associated data are attested not to have been tampered with. | [optional] 
**Context** | Pointer to **string** | Base64 encoded context for key derivation. Required if key derivation is enabled | [optional] 
**ConvergentEncryption** | Pointer to **bool** | This parameter will only be used when a key is expected to be created. Whether to support convergent encryption. This is only supported when using a key with key derivation enabled and will require all requests to carry both a context and 96-bit (12-byte) nonce. The given nonce will be used in place of a randomly generated nonce. As a result, when the same context and nonce are supplied, the same ciphertext is generated. It is *very important* when using this mode that you ensure that all nonces are unique for a given context. Failing to do so will severely impact the ciphertext&#x27;s security. | [optional] 
**KeyVersion** | Pointer to **int32** | The version of the key to use for encryption. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key. | [optional] 
**Nonce** | Pointer to **string** | Base64 encoded nonce value. Must be provided if convergent encryption is enabled for this key and the key was generated with Vault 0.6.1. Not required for keys created in 0.6.2+. The value must be exactly 96 bits (12 bytes) long and the user must ensure that for any given context (and thus, any given encryption key) this nonce value is **never reused**. | [optional] 
**PartialFailureResponseCode** | Pointer to **int32** | Ordinarily, if a batch item fails to encrypt due to a bad input, but other batch items succeed, the HTTP response code is 400 (Bad Request). Some applications may want to treat partial failures differently. Providing the parameter returns the given response code integer instead of a 400 in this case. If all values fail HTTP 400 is still returned. | [optional] 
**Plaintext** | Pointer to **string** | Base64 encoded plaintext value to be encrypted | [optional] 
**Type** | Pointer to **string** | This parameter is required when encryption key is expected to be created. When performing an upsert operation, the type of key to create. Currently, \&quot;aes128-gcm96\&quot; (symmetric) and \&quot;aes256-gcm96\&quot; (symmetric) are the only types supported. Defaults to \&quot;aes256-gcm96\&quot;. | [optional] [default to "aes256-gcm96"]



## Methods


### NewTransitEncryptRequest

`func NewTransitEncryptRequest() *TransitEncryptRequest`

NewTransitEncryptRequest instantiates a new TransitEncryptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitEncryptRequestWithDefaults

`func NewTransitEncryptRequestWithDefaults() *TransitEncryptRequest`

NewTransitEncryptRequestWithDefaults instantiates a new TransitEncryptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAssociatedData

`func (o *TransitEncryptRequest) GetAssociatedData() string`

GetAssociatedData returns the AssociatedData field if non-nil, zero value otherwise.

### GetAssociatedDataOk

`func (o *TransitEncryptRequest) GetAssociatedDataOk() (*string, bool)`

GetAssociatedDataOk returns a tuple with the AssociatedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedData

`func (o *TransitEncryptRequest) SetAssociatedData(v string)`

SetAssociatedData sets AssociatedData field to given value.


### HasAssociatedData

`func (o *TransitEncryptRequest) HasAssociatedData() bool`

HasAssociatedData returns a boolean if a field has been set.




### GetContext

`func (o *TransitEncryptRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TransitEncryptRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TransitEncryptRequest) SetContext(v string)`

SetContext sets Context field to given value.


### HasContext

`func (o *TransitEncryptRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.




### GetConvergentEncryption

`func (o *TransitEncryptRequest) GetConvergentEncryption() bool`

GetConvergentEncryption returns the ConvergentEncryption field if non-nil, zero value otherwise.

### GetConvergentEncryptionOk

`func (o *TransitEncryptRequest) GetConvergentEncryptionOk() (*bool, bool)`

GetConvergentEncryptionOk returns a tuple with the ConvergentEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvergentEncryption

`func (o *TransitEncryptRequest) SetConvergentEncryption(v bool)`

SetConvergentEncryption sets ConvergentEncryption field to given value.


### HasConvergentEncryption

`func (o *TransitEncryptRequest) HasConvergentEncryption() bool`

HasConvergentEncryption returns a boolean if a field has been set.




### GetKeyVersion

`func (o *TransitEncryptRequest) GetKeyVersion() int32`

GetKeyVersion returns the KeyVersion field if non-nil, zero value otherwise.

### GetKeyVersionOk

`func (o *TransitEncryptRequest) GetKeyVersionOk() (*int32, bool)`

GetKeyVersionOk returns a tuple with the KeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersion

`func (o *TransitEncryptRequest) SetKeyVersion(v int32)`

SetKeyVersion sets KeyVersion field to given value.


### HasKeyVersion

`func (o *TransitEncryptRequest) HasKeyVersion() bool`

HasKeyVersion returns a boolean if a field has been set.




### GetNonce

`func (o *TransitEncryptRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *TransitEncryptRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *TransitEncryptRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### HasNonce

`func (o *TransitEncryptRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.




### GetPartialFailureResponseCode

`func (o *TransitEncryptRequest) GetPartialFailureResponseCode() int32`

GetPartialFailureResponseCode returns the PartialFailureResponseCode field if non-nil, zero value otherwise.

### GetPartialFailureResponseCodeOk

`func (o *TransitEncryptRequest) GetPartialFailureResponseCodeOk() (*int32, bool)`

GetPartialFailureResponseCodeOk returns a tuple with the PartialFailureResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialFailureResponseCode

`func (o *TransitEncryptRequest) SetPartialFailureResponseCode(v int32)`

SetPartialFailureResponseCode sets PartialFailureResponseCode field to given value.


### HasPartialFailureResponseCode

`func (o *TransitEncryptRequest) HasPartialFailureResponseCode() bool`

HasPartialFailureResponseCode returns a boolean if a field has been set.




### GetPlaintext

`func (o *TransitEncryptRequest) GetPlaintext() string`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *TransitEncryptRequest) GetPlaintextOk() (*string, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *TransitEncryptRequest) SetPlaintext(v string)`

SetPlaintext sets Plaintext field to given value.


### HasPlaintext

`func (o *TransitEncryptRequest) HasPlaintext() bool`

HasPlaintext returns a boolean if a field has been set.




### GetType

`func (o *TransitEncryptRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransitEncryptRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransitEncryptRequest) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *TransitEncryptRequest) HasType() bool`

HasType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


