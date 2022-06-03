# GcpkmsKeysRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | Algorithm to use for encryption, decryption, or signing. The value depends on the key purpose. The value cannot be changed after creation. For a key purpose of \&quot;encrypt_decrypt\&quot;, the valid values are: - symmetric_encryption (default) For a key purpose of \&quot;asymmetric_sign\&quot;, valid values are: - rsa_sign_pss_2048_sha256 - rsa_sign_pss_3072_sha256 - rsa_sign_pss_4096_sha256 - rsa_sign_pkcs1_2048_sha256 - rsa_sign_pkcs1_3072_sha256 - rsa_sign_pkcs1_4096_sha256 - ec_sign_p256_sha256 - ec_sign_p384_sha384 For a key purpose of \&quot;asymmetric_decrypt\&quot;, valid values are: - rsa_decrypt_oaep_2048_sha256 - rsa_decrypt_oaep_3072_sha256 - rsa_decrypt_oaep_4096_sha256 | [optional] 
**CryptoKey** | Pointer to **string** | Name of the crypto key to use. If the given crypto key does not exist, Vault will try to create it. This defaults to the name of the key given to Vault as the parameter if unspecified. | [optional] 
**KeyRing** | Pointer to **string** | Full Google Cloud resource ID of the key ring with the project and location (e.g. projects/my-project/locations/global/keyRings/my-keyring). If the given key ring does not exist, Vault will try to create it during a create operation. | [optional] 
**Labels** | Pointer to **map[string]interface{}** | Arbitrary key&#x3D;value label to apply to the crypto key. To specify multiple labels, specify this argument multiple times (e.g. labels&#x3D;\&quot;a&#x3D;b\&quot; labels&#x3D;\&quot;c&#x3D;d\&quot;). | [optional] 
**ProtectionLevel** | Pointer to **string** | Level of protection to use for the key management. Valid values are \&quot;software\&quot; and \&quot;hsm\&quot;. The default value is \&quot;software\&quot;. The value cannot be changed after creation. | [optional] 
**Purpose** | Pointer to **string** | Purpose of the key. Valid options are \&quot;asymmetric_decrypt\&quot;, \&quot;asymmetric_sign\&quot;, and \&quot;encrypt_decrypt\&quot;. The default value is \&quot;encrypt_decrypt\&quot;. The value cannot be changed after creation. | [optional] 
**RotationPeriod** | Pointer to **int32** | Amount of time between crypto key version rotations. This is specified as a time duration value like 72h (72 hours). The smallest possible value is 24h. This value only applies to keys with a purpose of \&quot;encrypt_decrypt\&quot;. | [optional] 

## Methods

### NewGcpkmsKeysRequest

`func NewGcpkmsKeysRequest() *GcpkmsKeysRequest`

NewGcpkmsKeysRequest instantiates a new GcpkmsKeysRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpkmsKeysRequestWithDefaults

`func NewGcpkmsKeysRequestWithDefaults() *GcpkmsKeysRequest`

NewGcpkmsKeysRequestWithDefaults instantiates a new GcpkmsKeysRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *GcpkmsKeysRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *GcpkmsKeysRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *GcpkmsKeysRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *GcpkmsKeysRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetCryptoKey

`func (o *GcpkmsKeysRequest) GetCryptoKey() string`

GetCryptoKey returns the CryptoKey field if non-nil, zero value otherwise.

### GetCryptoKeyOk

`func (o *GcpkmsKeysRequest) GetCryptoKeyOk() (*string, bool)`

GetCryptoKeyOk returns a tuple with the CryptoKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoKey

`func (o *GcpkmsKeysRequest) SetCryptoKey(v string)`

SetCryptoKey sets CryptoKey field to given value.

### HasCryptoKey

`func (o *GcpkmsKeysRequest) HasCryptoKey() bool`

HasCryptoKey returns a boolean if a field has been set.

### GetKeyRing

`func (o *GcpkmsKeysRequest) GetKeyRing() string`

GetKeyRing returns the KeyRing field if non-nil, zero value otherwise.

### GetKeyRingOk

`func (o *GcpkmsKeysRequest) GetKeyRingOk() (*string, bool)`

GetKeyRingOk returns a tuple with the KeyRing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRing

`func (o *GcpkmsKeysRequest) SetKeyRing(v string)`

SetKeyRing sets KeyRing field to given value.

### HasKeyRing

`func (o *GcpkmsKeysRequest) HasKeyRing() bool`

HasKeyRing returns a boolean if a field has been set.

### GetLabels

`func (o *GcpkmsKeysRequest) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GcpkmsKeysRequest) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GcpkmsKeysRequest) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GcpkmsKeysRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetProtectionLevel

`func (o *GcpkmsKeysRequest) GetProtectionLevel() string`

GetProtectionLevel returns the ProtectionLevel field if non-nil, zero value otherwise.

### GetProtectionLevelOk

`func (o *GcpkmsKeysRequest) GetProtectionLevelOk() (*string, bool)`

GetProtectionLevelOk returns a tuple with the ProtectionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionLevel

`func (o *GcpkmsKeysRequest) SetProtectionLevel(v string)`

SetProtectionLevel sets ProtectionLevel field to given value.

### HasProtectionLevel

`func (o *GcpkmsKeysRequest) HasProtectionLevel() bool`

HasProtectionLevel returns a boolean if a field has been set.

### GetPurpose

`func (o *GcpkmsKeysRequest) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *GcpkmsKeysRequest) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *GcpkmsKeysRequest) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *GcpkmsKeysRequest) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetRotationPeriod

`func (o *GcpkmsKeysRequest) GetRotationPeriod() int32`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *GcpkmsKeysRequest) GetRotationPeriodOk() (*int32, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *GcpkmsKeysRequest) SetRotationPeriod(v int32)`

SetRotationPeriod sets RotationPeriod field to given value.

### HasRotationPeriod

`func (o *GcpkmsKeysRequest) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


