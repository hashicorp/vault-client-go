# TransitWriteKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPlaintextBackup** | Pointer to **bool** | Enables taking a backup of the named key in plaintext format. Once set, this cannot be disabled. | [optional] 
**AutoRotatePeriod** | Pointer to **int32** | Amount of time the key should live before being automatically rotated. A value of 0 (default) disables automatic rotation for the key. | [optional] [default to 0]
**Context** | Pointer to **string** | Base64 encoded context for key derivation. When reading a key with key derivation enabled, if the key type supports public keys, this will return the public key for the given context. | [optional] 
**ConvergentEncryption** | Pointer to **bool** | Whether to support convergent encryption. This is only supported when using a key with key derivation enabled and will require all requests to carry both a context and 96-bit (12-byte) nonce. The given nonce will be used in place of a randomly generated nonce. As a result, when the same context and nonce are supplied, the same ciphertext is generated. It is *very important* when using this mode that you ensure that all nonces are unique for a given context. Failing to do so will severely impact the ciphertext&#39;s security. | [optional] 
**Derived** | Pointer to **bool** | Enables key derivation mode. This allows for per-transaction unique keys for encryption operations. | [optional] 
**Exportable** | Pointer to **bool** | Enables keys to be exportable. This allows for all the valid keys in the key ring to be exported. | [optional] 
**KeySize** | Pointer to **int32** | The key size in bytes for the algorithm. Only applies to HMAC and must be no fewer than 32 bytes and no more than 512 | [optional] [default to 0]
**Type** | Pointer to **string** | The type of key to create. Currently, \&quot;aes128-gcm96\&quot; (symmetric), \&quot;aes256-gcm96\&quot; (symmetric), \&quot;ecdsa-p256\&quot; (asymmetric), \&quot;ecdsa-p384\&quot; (asymmetric), \&quot;ecdsa-p521\&quot; (asymmetric), \&quot;ed25519\&quot; (asymmetric), \&quot;rsa-2048\&quot; (asymmetric), \&quot;rsa-3072\&quot; (asymmetric), \&quot;rsa-4096\&quot; (asymmetric) are supported. Defaults to \&quot;aes256-gcm96\&quot;. | [optional] [default to "aes256-gcm96"]

## Methods

### NewTransitWriteKeyRequest

`func NewTransitWriteKeyRequest() *TransitWriteKeyRequest`

NewTransitWriteKeyRequest instantiates a new TransitWriteKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitWriteKeyRequestWithDefaults

`func NewTransitWriteKeyRequestWithDefaults() *TransitWriteKeyRequest`

NewTransitWriteKeyRequestWithDefaults instantiates a new TransitWriteKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPlaintextBackup

`func (o *TransitWriteKeyRequest) GetAllowPlaintextBackup() bool`

GetAllowPlaintextBackup returns the AllowPlaintextBackup field if non-nil, zero value otherwise.

### GetAllowPlaintextBackupOk

`func (o *TransitWriteKeyRequest) GetAllowPlaintextBackupOk() (*bool, bool)`

GetAllowPlaintextBackupOk returns a tuple with the AllowPlaintextBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPlaintextBackup

`func (o *TransitWriteKeyRequest) SetAllowPlaintextBackup(v bool)`

SetAllowPlaintextBackup sets AllowPlaintextBackup field to given value.

### HasAllowPlaintextBackup

`func (o *TransitWriteKeyRequest) HasAllowPlaintextBackup() bool`

HasAllowPlaintextBackup returns a boolean if a field has been set.

### GetAutoRotatePeriod

`func (o *TransitWriteKeyRequest) GetAutoRotatePeriod() int32`

GetAutoRotatePeriod returns the AutoRotatePeriod field if non-nil, zero value otherwise.

### GetAutoRotatePeriodOk

`func (o *TransitWriteKeyRequest) GetAutoRotatePeriodOk() (*int32, bool)`

GetAutoRotatePeriodOk returns a tuple with the AutoRotatePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotatePeriod

`func (o *TransitWriteKeyRequest) SetAutoRotatePeriod(v int32)`

SetAutoRotatePeriod sets AutoRotatePeriod field to given value.

### HasAutoRotatePeriod

`func (o *TransitWriteKeyRequest) HasAutoRotatePeriod() bool`

HasAutoRotatePeriod returns a boolean if a field has been set.

### GetContext

`func (o *TransitWriteKeyRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TransitWriteKeyRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TransitWriteKeyRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *TransitWriteKeyRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetConvergentEncryption

`func (o *TransitWriteKeyRequest) GetConvergentEncryption() bool`

GetConvergentEncryption returns the ConvergentEncryption field if non-nil, zero value otherwise.

### GetConvergentEncryptionOk

`func (o *TransitWriteKeyRequest) GetConvergentEncryptionOk() (*bool, bool)`

GetConvergentEncryptionOk returns a tuple with the ConvergentEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvergentEncryption

`func (o *TransitWriteKeyRequest) SetConvergentEncryption(v bool)`

SetConvergentEncryption sets ConvergentEncryption field to given value.

### HasConvergentEncryption

`func (o *TransitWriteKeyRequest) HasConvergentEncryption() bool`

HasConvergentEncryption returns a boolean if a field has been set.

### GetDerived

`func (o *TransitWriteKeyRequest) GetDerived() bool`

GetDerived returns the Derived field if non-nil, zero value otherwise.

### GetDerivedOk

`func (o *TransitWriteKeyRequest) GetDerivedOk() (*bool, bool)`

GetDerivedOk returns a tuple with the Derived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerived

`func (o *TransitWriteKeyRequest) SetDerived(v bool)`

SetDerived sets Derived field to given value.

### HasDerived

`func (o *TransitWriteKeyRequest) HasDerived() bool`

HasDerived returns a boolean if a field has been set.

### GetExportable

`func (o *TransitWriteKeyRequest) GetExportable() bool`

GetExportable returns the Exportable field if non-nil, zero value otherwise.

### GetExportableOk

`func (o *TransitWriteKeyRequest) GetExportableOk() (*bool, bool)`

GetExportableOk returns a tuple with the Exportable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportable

`func (o *TransitWriteKeyRequest) SetExportable(v bool)`

SetExportable sets Exportable field to given value.

### HasExportable

`func (o *TransitWriteKeyRequest) HasExportable() bool`

HasExportable returns a boolean if a field has been set.

### GetKeySize

`func (o *TransitWriteKeyRequest) GetKeySize() int32`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *TransitWriteKeyRequest) GetKeySizeOk() (*int32, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *TransitWriteKeyRequest) SetKeySize(v int32)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *TransitWriteKeyRequest) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### GetType

`func (o *TransitWriteKeyRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransitWriteKeyRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransitWriteKeyRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransitWriteKeyRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


