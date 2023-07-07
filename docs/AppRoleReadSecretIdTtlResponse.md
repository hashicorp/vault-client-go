# AppRoleReadSecretIdTtlResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretIdTtl** | Pointer to **string** | Duration in seconds after which the issued secret ID should expire. Defaults to 0, meaning no expiration. | [optional] 



## Methods


### NewAppRoleReadSecretIdTtlResponse

`func NewAppRoleReadSecretIdTtlResponse() *AppRoleReadSecretIdTtlResponse`

NewAppRoleReadSecretIdTtlResponse instantiates a new AppRoleReadSecretIdTtlResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRoleReadSecretIdTtlResponseWithDefaults

`func NewAppRoleReadSecretIdTtlResponseWithDefaults() *AppRoleReadSecretIdTtlResponse`

NewAppRoleReadSecretIdTtlResponseWithDefaults instantiates a new AppRoleReadSecretIdTtlResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetSecretIdTtl

`func (o *AppRoleReadSecretIdTtlResponse) GetSecretIdTtl() string`

GetSecretIdTtl returns the SecretIdTtl field if non-nil, zero value otherwise.

### GetSecretIdTtlOk

`func (o *AppRoleReadSecretIdTtlResponse) GetSecretIdTtlOk() (*string, bool)`

GetSecretIdTtlOk returns a tuple with the SecretIdTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretIdTtl

`func (o *AppRoleReadSecretIdTtlResponse) SetSecretIdTtl(v string)`

SetSecretIdTtl sets SecretIdTtl field to given value.


### HasSecretIdTtl

`func (o *AppRoleReadSecretIdTtlResponse) HasSecretIdTtl() bool`

HasSecretIdTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


