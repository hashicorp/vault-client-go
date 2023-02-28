# GoogleCloudKMSReencryptRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalAuthenticatedData** | Pointer to **string** | Optional data that, if specified, must also be provided during decryption. | [optional] 
**Ciphertext** | Pointer to **string** | Ciphertext to be re-encrypted to the latest key version. This must be ciphertext that Vault previously generated for this named key. | [optional] 
**KeyVersion** | Pointer to **int32** | Integer version of the crypto key version to use for the new encryption. If unspecified, this defaults to the latest active crypto key version. | [optional] 



## Methods


### NewGoogleCloudKMSReencryptRequest

`func NewGoogleCloudKMSReencryptRequest() *GoogleCloudKMSReencryptRequest`

NewGoogleCloudKMSReencryptRequest instantiates a new GoogleCloudKMSReencryptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudKMSReencryptRequestWithDefaults

`func NewGoogleCloudKMSReencryptRequestWithDefaults() *GoogleCloudKMSReencryptRequest`

NewGoogleCloudKMSReencryptRequestWithDefaults instantiates a new GoogleCloudKMSReencryptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAdditionalAuthenticatedData

`func (o *GoogleCloudKMSReencryptRequest) GetAdditionalAuthenticatedData() string`

GetAdditionalAuthenticatedData returns the AdditionalAuthenticatedData field if non-nil, zero value otherwise.

### GetAdditionalAuthenticatedDataOk

`func (o *GoogleCloudKMSReencryptRequest) GetAdditionalAuthenticatedDataOk() (*string, bool)`

GetAdditionalAuthenticatedDataOk returns a tuple with the AdditionalAuthenticatedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAuthenticatedData

`func (o *GoogleCloudKMSReencryptRequest) SetAdditionalAuthenticatedData(v string)`

SetAdditionalAuthenticatedData sets AdditionalAuthenticatedData field to given value.


### HasAdditionalAuthenticatedData

`func (o *GoogleCloudKMSReencryptRequest) HasAdditionalAuthenticatedData() bool`

HasAdditionalAuthenticatedData returns a boolean if a field has been set.




### GetCiphertext

`func (o *GoogleCloudKMSReencryptRequest) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *GoogleCloudKMSReencryptRequest) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *GoogleCloudKMSReencryptRequest) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.


### HasCiphertext

`func (o *GoogleCloudKMSReencryptRequest) HasCiphertext() bool`

HasCiphertext returns a boolean if a field has been set.




### GetKeyVersion

`func (o *GoogleCloudKMSReencryptRequest) GetKeyVersion() int32`

GetKeyVersion returns the KeyVersion field if non-nil, zero value otherwise.

### GetKeyVersionOk

`func (o *GoogleCloudKMSReencryptRequest) GetKeyVersionOk() (*int32, bool)`

GetKeyVersionOk returns a tuple with the KeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersion

`func (o *GoogleCloudKMSReencryptRequest) SetKeyVersion(v int32)`

SetKeyVersion sets KeyVersion field to given value.


### HasKeyVersion

`func (o *GoogleCloudKMSReencryptRequest) HasKeyVersion() bool`

HasKeyVersion returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


