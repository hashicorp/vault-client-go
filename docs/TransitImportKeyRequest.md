# TransitImportKeyRequest


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


### NewTransitImportKeyRequest

`func NewTransitImportKeyRequest() *TransitImportKeyRequest`

NewTransitImportKeyRequest instantiates a new TransitImportKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitImportKeyRequestWithDefaults

`func NewTransitImportKeyRequestWithDefaults() *TransitImportKeyRequest`

NewTransitImportKeyRequestWithDefaults instantiates a new TransitImportKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAllowPlaintextBackup

`func (o *TransitImportKeyRequest) GetAllowPlaintextBackup() bool`

GetAllowPlaintextBackup returns the AllowPlaintextBackup field if non-nil, zero value otherwise.

### GetAllowPlaintextBackupOk

`func (o *TransitImportKeyRequest) GetAllowPlaintextBackupOk() (*bool, bool)`

GetAllowPlaintextBackupOk returns a tuple with the AllowPlaintextBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPlaintextBackup

`func (o *TransitImportKeyRequest) SetAllowPlaintextBackup(v bool)`

SetAllowPlaintextBackup sets AllowPlaintextBackup field to given value.


### HasAllowPlaintextBackup

`func (o *TransitImportKeyRequest) HasAllowPlaintextBackup() bool`

HasAllowPlaintextBackup returns a boolean if a field has been set.




### GetAllowRotation

`func (o *TransitImportKeyRequest) GetAllowRotation() bool`

GetAllowRotation returns the AllowRotation field if non-nil, zero value otherwise.

### GetAllowRotationOk

`func (o *TransitImportKeyRequest) GetAllowRotationOk() (*bool, bool)`

GetAllowRotationOk returns a tuple with the AllowRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRotation

`func (o *TransitImportKeyRequest) SetAllowRotation(v bool)`

SetAllowRotation sets AllowRotation field to given value.


### HasAllowRotation

`func (o *TransitImportKeyRequest) HasAllowRotation() bool`

HasAllowRotation returns a boolean if a field has been set.




### GetAutoRotatePeriod

`func (o *TransitImportKeyRequest) GetAutoRotatePeriod() int32`

GetAutoRotatePeriod returns the AutoRotatePeriod field if non-nil, zero value otherwise.

### GetAutoRotatePeriodOk

`func (o *TransitImportKeyRequest) GetAutoRotatePeriodOk() (*int32, bool)`

GetAutoRotatePeriodOk returns a tuple with the AutoRotatePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotatePeriod

`func (o *TransitImportKeyRequest) SetAutoRotatePeriod(v int32)`

SetAutoRotatePeriod sets AutoRotatePeriod field to given value.


### HasAutoRotatePeriod

`func (o *TransitImportKeyRequest) HasAutoRotatePeriod() bool`

HasAutoRotatePeriod returns a boolean if a field has been set.




### GetCiphertext

`func (o *TransitImportKeyRequest) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *TransitImportKeyRequest) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *TransitImportKeyRequest) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.


### HasCiphertext

`func (o *TransitImportKeyRequest) HasCiphertext() bool`

HasCiphertext returns a boolean if a field has been set.




### GetContext

`func (o *TransitImportKeyRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TransitImportKeyRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TransitImportKeyRequest) SetContext(v string)`

SetContext sets Context field to given value.


### HasContext

`func (o *TransitImportKeyRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.




### GetDerived

`func (o *TransitImportKeyRequest) GetDerived() bool`

GetDerived returns the Derived field if non-nil, zero value otherwise.

### GetDerivedOk

`func (o *TransitImportKeyRequest) GetDerivedOk() (*bool, bool)`

GetDerivedOk returns a tuple with the Derived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerived

`func (o *TransitImportKeyRequest) SetDerived(v bool)`

SetDerived sets Derived field to given value.


### HasDerived

`func (o *TransitImportKeyRequest) HasDerived() bool`

HasDerived returns a boolean if a field has been set.




### GetExportable

`func (o *TransitImportKeyRequest) GetExportable() bool`

GetExportable returns the Exportable field if non-nil, zero value otherwise.

### GetExportableOk

`func (o *TransitImportKeyRequest) GetExportableOk() (*bool, bool)`

GetExportableOk returns a tuple with the Exportable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportable

`func (o *TransitImportKeyRequest) SetExportable(v bool)`

SetExportable sets Exportable field to given value.


### HasExportable

`func (o *TransitImportKeyRequest) HasExportable() bool`

HasExportable returns a boolean if a field has been set.




### GetHashFunction

`func (o *TransitImportKeyRequest) GetHashFunction() string`

GetHashFunction returns the HashFunction field if non-nil, zero value otherwise.

### GetHashFunctionOk

`func (o *TransitImportKeyRequest) GetHashFunctionOk() (*string, bool)`

GetHashFunctionOk returns a tuple with the HashFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashFunction

`func (o *TransitImportKeyRequest) SetHashFunction(v string)`

SetHashFunction sets HashFunction field to given value.


### HasHashFunction

`func (o *TransitImportKeyRequest) HasHashFunction() bool`

HasHashFunction returns a boolean if a field has been set.




### GetType

`func (o *TransitImportKeyRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransitImportKeyRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransitImportKeyRequest) SetType(v string)`

SetType sets Type field to given value.


### HasType

`func (o *TransitImportKeyRequest) HasType() bool`

HasType returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


