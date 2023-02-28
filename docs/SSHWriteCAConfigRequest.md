# SSHWriteCAConfigRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenerateSigningKey** | Pointer to **bool** | Generate SSH key pair internally rather than use the private_key and public_key fields. | [optional] [default to true]
**KeyBits** | Pointer to **int32** | Specifies the desired key bits when generating variable-length keys (such as when key_type&#x3D;\&quot;ssh-rsa\&quot;) or which NIST P-curve to use when key_type&#x3D;\&quot;ec\&quot; (256, 384, or 521). | [optional] [default to 0]
**KeyType** | Pointer to **string** | Specifies the desired key type when generating; could be a OpenSSH key type identifier (ssh-rsa, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, ecdsa-sha2-nistp521, or ssh-ed25519) or an algorithm (rsa, ec, ed25519). | [optional] [default to "ssh-rsa"]
**PrivateKey** | Pointer to **string** | Private half of the SSH key that will be used to sign certificates. | [optional] 
**PublicKey** | Pointer to **string** | Public half of the SSH key that will be used to sign certificates. | [optional] 



## Methods


### NewSSHWriteCAConfigRequest

`func NewSSHWriteCAConfigRequest() *SSHWriteCAConfigRequest`

NewSSHWriteCAConfigRequest instantiates a new SSHWriteCAConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHWriteCAConfigRequestWithDefaults

`func NewSSHWriteCAConfigRequestWithDefaults() *SSHWriteCAConfigRequest`

NewSSHWriteCAConfigRequestWithDefaults instantiates a new SSHWriteCAConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetGenerateSigningKey

`func (o *SSHWriteCAConfigRequest) GetGenerateSigningKey() bool`

GetGenerateSigningKey returns the GenerateSigningKey field if non-nil, zero value otherwise.

### GetGenerateSigningKeyOk

`func (o *SSHWriteCAConfigRequest) GetGenerateSigningKeyOk() (*bool, bool)`

GetGenerateSigningKeyOk returns a tuple with the GenerateSigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSigningKey

`func (o *SSHWriteCAConfigRequest) SetGenerateSigningKey(v bool)`

SetGenerateSigningKey sets GenerateSigningKey field to given value.


### HasGenerateSigningKey

`func (o *SSHWriteCAConfigRequest) HasGenerateSigningKey() bool`

HasGenerateSigningKey returns a boolean if a field has been set.




### GetKeyBits

`func (o *SSHWriteCAConfigRequest) GetKeyBits() int32`

GetKeyBits returns the KeyBits field if non-nil, zero value otherwise.

### GetKeyBitsOk

`func (o *SSHWriteCAConfigRequest) GetKeyBitsOk() (*int32, bool)`

GetKeyBitsOk returns a tuple with the KeyBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyBits

`func (o *SSHWriteCAConfigRequest) SetKeyBits(v int32)`

SetKeyBits sets KeyBits field to given value.


### HasKeyBits

`func (o *SSHWriteCAConfigRequest) HasKeyBits() bool`

HasKeyBits returns a boolean if a field has been set.




### GetKeyType

`func (o *SSHWriteCAConfigRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *SSHWriteCAConfigRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *SSHWriteCAConfigRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### HasKeyType

`func (o *SSHWriteCAConfigRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.




### GetPrivateKey

`func (o *SSHWriteCAConfigRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SSHWriteCAConfigRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SSHWriteCAConfigRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### HasPrivateKey

`func (o *SSHWriteCAConfigRequest) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.




### GetPublicKey

`func (o *SSHWriteCAConfigRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SSHWriteCAConfigRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SSHWriteCAConfigRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### HasPublicKey

`func (o *SSHWriteCAConfigRequest) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


