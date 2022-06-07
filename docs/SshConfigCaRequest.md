# SshConfigCaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenerateSigningKey** | Pointer to **bool** | Generate SSH key pair internally rather than use the private_key and public_key fields. | [optional] [default to true]
**KeyBits** | Pointer to **int32** | Specifies the desired key bits when generating variable-length keys (such as when key_type&#x3D;\&quot;ssh-rsa\&quot;) or which NIST P-curve to use when key_type&#x3D;\&quot;ec\&quot; (256, 384, or 521). | [optional] [default to 0]
**KeyType** | Pointer to **string** | Specifies the desired key type when generating; could be a OpenSSH key type identifier (ssh-rsa, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, ecdsa-sha2-nistp521, or ssh-ed25519) or an algorithm (rsa, ec, ed25519). | [optional] [default to "ssh-rsa"]
**PrivateKey** | Pointer to **string** | Private half of the SSH key that will be used to sign certificates. | [optional] 
**PublicKey** | Pointer to **string** | Public half of the SSH key that will be used to sign certificates. | [optional] 

## Methods

### NewSshConfigCaRequest

`func NewSshConfigCaRequest() *SshConfigCaRequest`

NewSshConfigCaRequest instantiates a new SshConfigCaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshConfigCaRequestWithDefaults

`func NewSshConfigCaRequestWithDefaults() *SshConfigCaRequest`

NewSshConfigCaRequestWithDefaults instantiates a new SshConfigCaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerateSigningKey

`func (o *SshConfigCaRequest) GetGenerateSigningKey() bool`

GetGenerateSigningKey returns the GenerateSigningKey field if non-nil, zero value otherwise.

### GetGenerateSigningKeyOk

`func (o *SshConfigCaRequest) GetGenerateSigningKeyOk() (*bool, bool)`

GetGenerateSigningKeyOk returns a tuple with the GenerateSigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSigningKey

`func (o *SshConfigCaRequest) SetGenerateSigningKey(v bool)`

SetGenerateSigningKey sets GenerateSigningKey field to given value.

### HasGenerateSigningKey

`func (o *SshConfigCaRequest) HasGenerateSigningKey() bool`

HasGenerateSigningKey returns a boolean if a field has been set.

### GetKeyBits

`func (o *SshConfigCaRequest) GetKeyBits() int32`

GetKeyBits returns the KeyBits field if non-nil, zero value otherwise.

### GetKeyBitsOk

`func (o *SshConfigCaRequest) GetKeyBitsOk() (*int32, bool)`

GetKeyBitsOk returns a tuple with the KeyBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyBits

`func (o *SshConfigCaRequest) SetKeyBits(v int32)`

SetKeyBits sets KeyBits field to given value.

### HasKeyBits

`func (o *SshConfigCaRequest) HasKeyBits() bool`

HasKeyBits returns a boolean if a field has been set.

### GetKeyType

`func (o *SshConfigCaRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *SshConfigCaRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *SshConfigCaRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *SshConfigCaRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetPrivateKey

`func (o *SshConfigCaRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SshConfigCaRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SshConfigCaRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *SshConfigCaRequest) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *SshConfigCaRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SshConfigCaRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SshConfigCaRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *SshConfigCaRequest) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


