# TransitDatakeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bits** | Pointer to **int32** | Number of bits for the key; currently 128, 256, and 512 bits are supported. Defaults to 256. | [optional] [default to 256]
**Context** | Pointer to **string** | Context for key derivation. Required for derived keys. | [optional] 
**KeyVersion** | Pointer to **int32** | The version of the Vault key to use for encryption of the data key. Must be 0 (for latest) or a value greater than or equal to the min_encryption_version configured on the key. | [optional] 
**Nonce** | Pointer to **string** | Nonce for when convergent encryption v1 is used (only in Vault 0.6.1) | [optional] 

## Methods

### NewTransitDatakeyRequest

`func NewTransitDatakeyRequest() *TransitDatakeyRequest`

NewTransitDatakeyRequest instantiates a new TransitDatakeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitDatakeyRequestWithDefaults

`func NewTransitDatakeyRequestWithDefaults() *TransitDatakeyRequest`

NewTransitDatakeyRequestWithDefaults instantiates a new TransitDatakeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBits

`func (o *TransitDatakeyRequest) GetBits() int32`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *TransitDatakeyRequest) GetBitsOk() (*int32, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *TransitDatakeyRequest) SetBits(v int32)`

SetBits sets Bits field to given value.

### HasBits

`func (o *TransitDatakeyRequest) HasBits() bool`

HasBits returns a boolean if a field has been set.

### GetContext

`func (o *TransitDatakeyRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TransitDatakeyRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TransitDatakeyRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *TransitDatakeyRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetKeyVersion

`func (o *TransitDatakeyRequest) GetKeyVersion() int32`

GetKeyVersion returns the KeyVersion field if non-nil, zero value otherwise.

### GetKeyVersionOk

`func (o *TransitDatakeyRequest) GetKeyVersionOk() (*int32, bool)`

GetKeyVersionOk returns a tuple with the KeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersion

`func (o *TransitDatakeyRequest) SetKeyVersion(v int32)`

SetKeyVersion sets KeyVersion field to given value.

### HasKeyVersion

`func (o *TransitDatakeyRequest) HasKeyVersion() bool`

HasKeyVersion returns a boolean if a field has been set.

### GetNonce

`func (o *TransitDatakeyRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *TransitDatakeyRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *TransitDatakeyRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *TransitDatakeyRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


