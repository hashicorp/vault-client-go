# GoogleCloudKMSEncryptRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**AdditionalAuthenticatedData** | Pointer to **string** | Optional base64-encoded data that, if specified, must also be provided to decrypt this payload. | [optional] 
**KeyVersion** | Pointer to **int32** | Integer version of the crypto key version to use for encryption. If unspecified, this defaults to the latest active crypto key version. | [optional] 
**Plaintext** | Pointer to **string** | Plaintext value to be encrypted. This can be a string or binary, but the size is limited. See the Google Cloud KMS documentation for information on size limitations by key types. | [optional] 



## Methods


### NewGoogleCloudKMSEncryptRequest

`func NewGoogleCloudKMSEncryptRequest() *GoogleCloudKMSEncryptRequest`

NewGoogleCloudKMSEncryptRequest instantiates a new GoogleCloudKMSEncryptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudKMSEncryptRequestWithDefaults

`func NewGoogleCloudKMSEncryptRequestWithDefaults() *GoogleCloudKMSEncryptRequest`

NewGoogleCloudKMSEncryptRequestWithDefaults instantiates a new GoogleCloudKMSEncryptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAdditionalAuthenticatedData

`func (o *GoogleCloudKMSEncryptRequest) GetAdditionalAuthenticatedData() string`

GetAdditionalAuthenticatedData returns the AdditionalAuthenticatedData field if non-nil, zero value otherwise.

### GetAdditionalAuthenticatedDataOk

`func (o *GoogleCloudKMSEncryptRequest) GetAdditionalAuthenticatedDataOk() (*string, bool)`

GetAdditionalAuthenticatedDataOk returns a tuple with the AdditionalAuthenticatedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAuthenticatedData

`func (o *GoogleCloudKMSEncryptRequest) SetAdditionalAuthenticatedData(v string)`

SetAdditionalAuthenticatedData sets AdditionalAuthenticatedData field to given value.


### HasAdditionalAuthenticatedData

`func (o *GoogleCloudKMSEncryptRequest) HasAdditionalAuthenticatedData() bool`

HasAdditionalAuthenticatedData returns a boolean if a field has been set.




### GetKeyVersion

`func (o *GoogleCloudKMSEncryptRequest) GetKeyVersion() int32`

GetKeyVersion returns the KeyVersion field if non-nil, zero value otherwise.

### GetKeyVersionOk

`func (o *GoogleCloudKMSEncryptRequest) GetKeyVersionOk() (*int32, bool)`

GetKeyVersionOk returns a tuple with the KeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersion

`func (o *GoogleCloudKMSEncryptRequest) SetKeyVersion(v int32)`

SetKeyVersion sets KeyVersion field to given value.


### HasKeyVersion

`func (o *GoogleCloudKMSEncryptRequest) HasKeyVersion() bool`

HasKeyVersion returns a boolean if a field has been set.




### GetPlaintext

`func (o *GoogleCloudKMSEncryptRequest) GetPlaintext() string`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *GoogleCloudKMSEncryptRequest) GetPlaintextOk() (*string, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *GoogleCloudKMSEncryptRequest) SetPlaintext(v string)`

SetPlaintext sets Plaintext field to given value.


### HasPlaintext

`func (o *GoogleCloudKMSEncryptRequest) HasPlaintext() bool`

HasPlaintext returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


