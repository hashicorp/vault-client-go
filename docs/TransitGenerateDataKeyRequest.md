# TransitGenerateDataKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bits** | Pointer to **int32** | Number of bits for the key; currently 128, 256, and 512 bits are supported. Defaults to 256. | [optional] [default to 256]
**Context** | Pointer to **string** | Context for key derivation. Required for derived keys. | [optional] 
**KeyVersion** | Pointer to **int32** | The version of the Vault key to use for encryption of the data key. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key. | [optional] 
**Nonce** | Pointer to **string** | Nonce for when convergent encryption v1 is used (only in Vault 0.6.1) | [optional] 

## Methods

### NewTransitGenerateDataKeyRequest

`func NewTransitGenerateDataKeyRequest() *TransitGenerateDataKeyRequest`

NewTransitGenerateDataKeyRequest instantiates a new TransitGenerateDataKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGenerateDataKeyRequestWithDefaults

`func NewTransitGenerateDataKeyRequestWithDefaults() *TransitGenerateDataKeyRequest`

NewTransitGenerateDataKeyRequestWithDefaults instantiates a new TransitGenerateDataKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBits

`func (o *TransitGenerateDataKeyRequest) GetBits() int32`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *TransitGenerateDataKeyRequest) GetBitsOk() (*int32, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *TransitGenerateDataKeyRequest) SetBits(v int32)`

SetBits sets Bits field to given value.

### HasBits

`func (o *TransitGenerateDataKeyRequest) HasBits() bool`

HasBits returns a boolean if a field has been set.

### GetContext

`func (o *TransitGenerateDataKeyRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TransitGenerateDataKeyRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TransitGenerateDataKeyRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *TransitGenerateDataKeyRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetKeyVersion

`func (o *TransitGenerateDataKeyRequest) GetKeyVersion() int32`

GetKeyVersion returns the KeyVersion field if non-nil, zero value otherwise.

### GetKeyVersionOk

`func (o *TransitGenerateDataKeyRequest) GetKeyVersionOk() (*int32, bool)`

GetKeyVersionOk returns a tuple with the KeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersion

`func (o *TransitGenerateDataKeyRequest) SetKeyVersion(v int32)`

SetKeyVersion sets KeyVersion field to given value.

### HasKeyVersion

`func (o *TransitGenerateDataKeyRequest) HasKeyVersion() bool`

HasKeyVersion returns a boolean if a field has been set.

### GetNonce

`func (o *TransitGenerateDataKeyRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *TransitGenerateDataKeyRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *TransitGenerateDataKeyRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *TransitGenerateDataKeyRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


