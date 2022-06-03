# TransitKeysImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPlaintextBackup** | Pointer to **bool** | Enables taking a backup of the named key in plaintext format. Once set, this cannot be disabled. | [optional] 
**AllowRotation** | Pointer to **bool** | True if the imported key may be rotated within Vault; false otherwise. | [optional] 
**AutoRotatePeriod** | Pointer to **int32** | Amount of time the key should live before being automatically rotated. A value of 0 (default) disables automatic rotation for the key. | [optional] [default to 0]
**Ciphertext** | Pointer to **string** | The base64-encoded ciphertext of the keys. The AES key should be encrypted using OAEP with the wrapping key and then concatenated with the import key, wrapped by the AES key. | [optional] 
**Context** | Pointer to **string** | Base64 encoded context for key derivation. When reading a key with key derivation enabled, if the key type supports public keys, this will return the public key for the given context. | [optional] 
**Derived** | Pointer to **bool** | Enables key derivation mode. This allows for per-transaction unique keys for encryption operations. | [optional] 
**Exportable** | Pointer to **bool** | Enables keys to be exportable. This allows for all the valid keys in the key ring to be exported. | [optional] 
**HashFunction** | Pointer to **string** | The hash function used as a random oracle in the OAEP wrapping of the user-generated, ephemeral AES key. Can be one of \&quot;SHA1\&quot;, \&quot;SHA224\&quot;, \&quot;SHA256\&quot; (default), \&quot;SHA384\&quot;, or \&quot;SHA512\&quot; | [optional] [default to "SHA256"]
**Type** | Pointer to **string** | The type of key being imported. Currently, \&quot;aes128-gcm96\&quot; (symmetric), \&quot;aes256-gcm96\&quot; (symmetric), \&quot;ecdsa-p256\&quot; (asymmetric), \&quot;ecdsa-p384\&quot; (asymmetric), \&quot;ecdsa-p521\&quot; (asymmetric), \&quot;ed25519\&quot; (asymmetric), \&quot;rsa-2048\&quot; (asymmetric), \&quot;rsa-3072\&quot; (asymmetric), \&quot;rsa-4096\&quot; (asymmetric) are supported. Defaults to \&quot;aes256-gcm96\&quot;. | [optional] [default to "aes256-gcm96"]

## Methods

### NewTransitKeysImportRequest

`func NewTransitKeysImportRequest() *TransitKeysImportRequest`

NewTransitKeysImportRequest instantiates a new TransitKeysImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitKeysImportRequestWithDefaults

`func NewTransitKeysImportRequestWithDefaults() *TransitKeysImportRequest`

NewTransitKeysImportRequestWithDefaults instantiates a new TransitKeysImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPlaintextBackup

`func (o *TransitKeysImportRequest) GetAllowPlaintextBackup() bool`

GetAllowPlaintextBackup returns the AllowPlaintextBackup field if non-nil, zero value otherwise.

### GetAllowPlaintextBackupOk

`func (o *TransitKeysImportRequest) GetAllowPlaintextBackupOk() (*bool, bool)`

GetAllowPlaintextBackupOk returns a tuple with the AllowPlaintextBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPlaintextBackup

`func (o *TransitKeysImportRequest) SetAllowPlaintextBackup(v bool)`

SetAllowPlaintextBackup sets AllowPlaintextBackup field to given value.

### HasAllowPlaintextBackup

`func (o *TransitKeysImportRequest) HasAllowPlaintextBackup() bool`

HasAllowPlaintextBackup returns a boolean if a field has been set.

### GetAllowRotation

`func (o *TransitKeysImportRequest) GetAllowRotation() bool`

GetAllowRotation returns the AllowRotation field if non-nil, zero value otherwise.

### GetAllowRotationOk

`func (o *TransitKeysImportRequest) GetAllowRotationOk() (*bool, bool)`

GetAllowRotationOk returns a tuple with the AllowRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRotation

`func (o *TransitKeysImportRequest) SetAllowRotation(v bool)`

SetAllowRotation sets AllowRotation field to given value.

### HasAllowRotation

`func (o *TransitKeysImportRequest) HasAllowRotation() bool`

HasAllowRotation returns a boolean if a field has been set.

### GetAutoRotatePeriod

`func (o *TransitKeysImportRequest) GetAutoRotatePeriod() int32`

GetAutoRotatePeriod returns the AutoRotatePeriod field if non-nil, zero value otherwise.

### GetAutoRotatePeriodOk

`func (o *TransitKeysImportRequest) GetAutoRotatePeriodOk() (*int32, bool)`

GetAutoRotatePeriodOk returns a tuple with the AutoRotatePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotatePeriod

`func (o *TransitKeysImportRequest) SetAutoRotatePeriod(v int32)`

SetAutoRotatePeriod sets AutoRotatePeriod field to given value.

### HasAutoRotatePeriod

`func (o *TransitKeysImportRequest) HasAutoRotatePeriod() bool`

HasAutoRotatePeriod returns a boolean if a field has been set.

### GetCiphertext

`func (o *TransitKeysImportRequest) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *TransitKeysImportRequest) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *TransitKeysImportRequest) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.

### HasCiphertext

`func (o *TransitKeysImportRequest) HasCiphertext() bool`

HasCiphertext returns a boolean if a field has been set.

### GetContext

`func (o *TransitKeysImportRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TransitKeysImportRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TransitKeysImportRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *TransitKeysImportRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetDerived

`func (o *TransitKeysImportRequest) GetDerived() bool`

GetDerived returns the Derived field if non-nil, zero value otherwise.

### GetDerivedOk

`func (o *TransitKeysImportRequest) GetDerivedOk() (*bool, bool)`

GetDerivedOk returns a tuple with the Derived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerived

`func (o *TransitKeysImportRequest) SetDerived(v bool)`

SetDerived sets Derived field to given value.

### HasDerived

`func (o *TransitKeysImportRequest) HasDerived() bool`

HasDerived returns a boolean if a field has been set.

### GetExportable

`func (o *TransitKeysImportRequest) GetExportable() bool`

GetExportable returns the Exportable field if non-nil, zero value otherwise.

### GetExportableOk

`func (o *TransitKeysImportRequest) GetExportableOk() (*bool, bool)`

GetExportableOk returns a tuple with the Exportable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportable

`func (o *TransitKeysImportRequest) SetExportable(v bool)`

SetExportable sets Exportable field to given value.

### HasExportable

`func (o *TransitKeysImportRequest) HasExportable() bool`

HasExportable returns a boolean if a field has been set.

### GetHashFunction

`func (o *TransitKeysImportRequest) GetHashFunction() string`

GetHashFunction returns the HashFunction field if non-nil, zero value otherwise.

### GetHashFunctionOk

`func (o *TransitKeysImportRequest) GetHashFunctionOk() (*string, bool)`

GetHashFunctionOk returns a tuple with the HashFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashFunction

`func (o *TransitKeysImportRequest) SetHashFunction(v string)`

SetHashFunction sets HashFunction field to given value.

### HasHashFunction

`func (o *TransitKeysImportRequest) HasHashFunction() bool`

HasHashFunction returns a boolean if a field has been set.

### GetType

`func (o *TransitKeysImportRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransitKeysImportRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransitKeysImportRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransitKeysImportRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


