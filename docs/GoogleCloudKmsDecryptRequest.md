# GoogleCloudKmsDecryptRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalAuthenticatedData** | Pointer to **string** | Optional data that was specified during encryption of this payload. | [optional] 
**Ciphertext** | Pointer to **string** | Ciphertext to decrypt as previously returned from an encrypt operation. This must be base64-encoded ciphertext as previously returned from an encrypt operation. | [optional] 
**KeyVersion** | Pointer to **int32** | Integer version of the crypto key version to use for decryption. This is required for asymmetric keys. For symmetric keys, Cloud KMS will choose the correct version automatically. | [optional] 



## Methods


### NewGoogleCloudKmsDecryptRequest

`func NewGoogleCloudKmsDecryptRequest() *GoogleCloudKmsDecryptRequest`

NewGoogleCloudKmsDecryptRequest instantiates a new GoogleCloudKmsDecryptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudKmsDecryptRequestWithDefaults

`func NewGoogleCloudKmsDecryptRequestWithDefaults() *GoogleCloudKmsDecryptRequest`

NewGoogleCloudKmsDecryptRequestWithDefaults instantiates a new GoogleCloudKmsDecryptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAdditionalAuthenticatedData

`func (o *GoogleCloudKmsDecryptRequest) GetAdditionalAuthenticatedData() string`

GetAdditionalAuthenticatedData returns the AdditionalAuthenticatedData field if non-nil, zero value otherwise.

### GetAdditionalAuthenticatedDataOk

`func (o *GoogleCloudKmsDecryptRequest) GetAdditionalAuthenticatedDataOk() (*string, bool)`

GetAdditionalAuthenticatedDataOk returns a tuple with the AdditionalAuthenticatedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAuthenticatedData

`func (o *GoogleCloudKmsDecryptRequest) SetAdditionalAuthenticatedData(v string)`

SetAdditionalAuthenticatedData sets AdditionalAuthenticatedData field to given value.


### HasAdditionalAuthenticatedData

`func (o *GoogleCloudKmsDecryptRequest) HasAdditionalAuthenticatedData() bool`

HasAdditionalAuthenticatedData returns a boolean if a field has been set.




### GetCiphertext

`func (o *GoogleCloudKmsDecryptRequest) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *GoogleCloudKmsDecryptRequest) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *GoogleCloudKmsDecryptRequest) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.


### HasCiphertext

`func (o *GoogleCloudKmsDecryptRequest) HasCiphertext() bool`

HasCiphertext returns a boolean if a field has been set.




### GetKeyVersion

`func (o *GoogleCloudKmsDecryptRequest) GetKeyVersion() int32`

GetKeyVersion returns the KeyVersion field if non-nil, zero value otherwise.

### GetKeyVersionOk

`func (o *GoogleCloudKmsDecryptRequest) GetKeyVersionOk() (*int32, bool)`

GetKeyVersionOk returns a tuple with the KeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersion

`func (o *GoogleCloudKmsDecryptRequest) SetKeyVersion(v int32)`

SetKeyVersion sets KeyVersion field to given value.


### HasKeyVersion

`func (o *GoogleCloudKmsDecryptRequest) HasKeyVersion() bool`

HasKeyVersion returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


