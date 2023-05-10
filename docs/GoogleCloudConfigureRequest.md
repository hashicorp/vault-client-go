# GoogleCloudConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to **string** | GCP IAM service account credentials JSON with permissions to create new service accounts and set IAM policies | [optional] 
**MaxTtl** | Pointer to **int32** | Maximum time a service account key is valid for. If &lt;&#x3D; 0, will use system default. | [optional] 
**Ttl** | Pointer to **int32** | Default lease for generated keys. If &lt;&#x3D; 0, will use system default. | [optional] 



## Methods


### NewGoogleCloudConfigureRequest

`func NewGoogleCloudConfigureRequest() *GoogleCloudConfigureRequest`

NewGoogleCloudConfigureRequest instantiates a new GoogleCloudConfigureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudConfigureRequestWithDefaults

`func NewGoogleCloudConfigureRequestWithDefaults() *GoogleCloudConfigureRequest`

NewGoogleCloudConfigureRequestWithDefaults instantiates a new GoogleCloudConfigureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetCredentials

`func (o *GoogleCloudConfigureRequest) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GoogleCloudConfigureRequest) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GoogleCloudConfigureRequest) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### HasCredentials

`func (o *GoogleCloudConfigureRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.




### GetMaxTtl

`func (o *GoogleCloudConfigureRequest) GetMaxTtl() int32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *GoogleCloudConfigureRequest) GetMaxTtlOk() (*int32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *GoogleCloudConfigureRequest) SetMaxTtl(v int32)`

SetMaxTtl sets MaxTtl field to given value.


### HasMaxTtl

`func (o *GoogleCloudConfigureRequest) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.




### GetTtl

`func (o *GoogleCloudConfigureRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GoogleCloudConfigureRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GoogleCloudConfigureRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *GoogleCloudConfigureRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


