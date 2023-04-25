# GoogleCloudGenerateRolesetKeyWithParameters2Request


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyAlgorithm** | Pointer to **string** | Private key algorithm for service account key - defaults to KEY_ALG_RSA_2048\&quot; | [optional] [default to "KEY_ALG_RSA_2048"]
**KeyType** | Pointer to **string** | Private key type for service account key - defaults to TYPE_GOOGLE_CREDENTIALS_FILE\&quot; | [optional] [default to "TYPE_GOOGLE_CREDENTIALS_FILE"]
**Ttl** | Pointer to **int32** | Lifetime of the service account key | [optional] 



## Methods


### NewGoogleCloudGenerateRolesetKeyWithParameters2Request

`func NewGoogleCloudGenerateRolesetKeyWithParameters2Request() *GoogleCloudGenerateRolesetKeyWithParameters2Request`

NewGoogleCloudGenerateRolesetKeyWithParameters2Request instantiates a new GoogleCloudGenerateRolesetKeyWithParameters2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudGenerateRolesetKeyWithParameters2RequestWithDefaults

`func NewGoogleCloudGenerateRolesetKeyWithParameters2RequestWithDefaults() *GoogleCloudGenerateRolesetKeyWithParameters2Request`

NewGoogleCloudGenerateRolesetKeyWithParameters2RequestWithDefaults instantiates a new GoogleCloudGenerateRolesetKeyWithParameters2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetKeyAlgorithm

`func (o *GoogleCloudGenerateRolesetKeyWithParameters2Request) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *GoogleCloudGenerateRolesetKeyWithParameters2Request) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *GoogleCloudGenerateRolesetKeyWithParameters2Request) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.


### HasKeyAlgorithm

`func (o *GoogleCloudGenerateRolesetKeyWithParameters2Request) HasKeyAlgorithm() bool`

HasKeyAlgorithm returns a boolean if a field has been set.




### GetKeyType

`func (o *GoogleCloudGenerateRolesetKeyWithParameters2Request) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *GoogleCloudGenerateRolesetKeyWithParameters2Request) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *GoogleCloudGenerateRolesetKeyWithParameters2Request) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### HasKeyType

`func (o *GoogleCloudGenerateRolesetKeyWithParameters2Request) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.




### GetTtl

`func (o *GoogleCloudGenerateRolesetKeyWithParameters2Request) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GoogleCloudGenerateRolesetKeyWithParameters2Request) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GoogleCloudGenerateRolesetKeyWithParameters2Request) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *GoogleCloudGenerateRolesetKeyWithParameters2Request) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


