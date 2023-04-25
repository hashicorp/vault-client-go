# GoogleCloudKmsConfigureKeyRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxVersion** | Pointer to **int32** | Maximum allowed crypto key version. If set to a positive value, key versions greater than the given value are not permitted to be used. If set to 0 or a negative value, there is no maximum key version. | [optional] 
**MinVersion** | Pointer to **int32** | Minimum allowed crypto key version. If set to a positive value, key versions less than the given value are not permitted to be used. If set to 0 or a negative value, there is no minimum key version. This value only affects encryption/re-encryption, not decryption. To restrict old values from being decrypted, increase this value and then perform a trim operation. | [optional] 



## Methods


### NewGoogleCloudKmsConfigureKeyRequest

`func NewGoogleCloudKmsConfigureKeyRequest() *GoogleCloudKmsConfigureKeyRequest`

NewGoogleCloudKmsConfigureKeyRequest instantiates a new GoogleCloudKmsConfigureKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudKmsConfigureKeyRequestWithDefaults

`func NewGoogleCloudKmsConfigureKeyRequestWithDefaults() *GoogleCloudKmsConfigureKeyRequest`

NewGoogleCloudKmsConfigureKeyRequestWithDefaults instantiates a new GoogleCloudKmsConfigureKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetMaxVersion

`func (o *GoogleCloudKmsConfigureKeyRequest) GetMaxVersion() int32`

GetMaxVersion returns the MaxVersion field if non-nil, zero value otherwise.

### GetMaxVersionOk

`func (o *GoogleCloudKmsConfigureKeyRequest) GetMaxVersionOk() (*int32, bool)`

GetMaxVersionOk returns a tuple with the MaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersion

`func (o *GoogleCloudKmsConfigureKeyRequest) SetMaxVersion(v int32)`

SetMaxVersion sets MaxVersion field to given value.


### HasMaxVersion

`func (o *GoogleCloudKmsConfigureKeyRequest) HasMaxVersion() bool`

HasMaxVersion returns a boolean if a field has been set.




### GetMinVersion

`func (o *GoogleCloudKmsConfigureKeyRequest) GetMinVersion() int32`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *GoogleCloudKmsConfigureKeyRequest) GetMinVersionOk() (*int32, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *GoogleCloudKmsConfigureKeyRequest) SetMinVersion(v int32)`

SetMinVersion sets MinVersion field to given value.


### HasMinVersion

`func (o *GoogleCloudKmsConfigureKeyRequest) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


