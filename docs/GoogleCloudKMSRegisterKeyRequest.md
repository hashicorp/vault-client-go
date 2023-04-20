# GoogleCloudKmsRegisterKeyRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CryptoKey** | Pointer to **string** | Full resource ID of the crypto key including the project, location, key ring, and crypto key like \&quot;projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s\&quot;. This crypto key must already exist in Google Cloud KMS unless verify is set to \&quot;false\&quot;. | [optional] 
**Verify** | Pointer to **bool** | Verify that the given Google Cloud KMS crypto key exists and is accessible before creating the storage entry in Vault. Set this to \&quot;false\&quot; if the key will not exist at creation time. | [optional] [default to true]



## Methods


### NewGoogleCloudKmsRegisterKeyRequest

`func NewGoogleCloudKmsRegisterKeyRequest() *GoogleCloudKmsRegisterKeyRequest`

NewGoogleCloudKmsRegisterKeyRequest instantiates a new GoogleCloudKmsRegisterKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudKmsRegisterKeyRequestWithDefaults

`func NewGoogleCloudKmsRegisterKeyRequestWithDefaults() *GoogleCloudKmsRegisterKeyRequest`

NewGoogleCloudKmsRegisterKeyRequestWithDefaults instantiates a new GoogleCloudKmsRegisterKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCryptoKey

`func (o *GoogleCloudKmsRegisterKeyRequest) GetCryptoKey() string`

GetCryptoKey returns the CryptoKey field if non-nil, zero value otherwise.

### GetCryptoKeyOk

`func (o *GoogleCloudKmsRegisterKeyRequest) GetCryptoKeyOk() (*string, bool)`

GetCryptoKeyOk returns a tuple with the CryptoKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoKey

`func (o *GoogleCloudKmsRegisterKeyRequest) SetCryptoKey(v string)`

SetCryptoKey sets CryptoKey field to given value.


### HasCryptoKey

`func (o *GoogleCloudKmsRegisterKeyRequest) HasCryptoKey() bool`

HasCryptoKey returns a boolean if a field has been set.




### GetVerify

`func (o *GoogleCloudKmsRegisterKeyRequest) GetVerify() bool`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *GoogleCloudKmsRegisterKeyRequest) GetVerifyOk() (*bool, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *GoogleCloudKmsRegisterKeyRequest) SetVerify(v bool)`

SetVerify sets Verify field to given value.


### HasVerify

`func (o *GoogleCloudKmsRegisterKeyRequest) HasVerify() bool`

HasVerify returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


