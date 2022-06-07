# GcpkmsConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to **string** | The credentials to use for authenticating to Google Cloud. Leave this blank to use the Default Application Credentials or instance metadata authentication. | [optional] 
**Scopes** | Pointer to **[]string** | The list of full-URL scopes to request when authenticating. By default, this requests https://www.googleapis.com/auth/cloudkms. | [optional] 

## Methods

### NewGcpkmsConfigRequest

`func NewGcpkmsConfigRequest() *GcpkmsConfigRequest`

NewGcpkmsConfigRequest instantiates a new GcpkmsConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpkmsConfigRequestWithDefaults

`func NewGcpkmsConfigRequestWithDefaults() *GcpkmsConfigRequest`

NewGcpkmsConfigRequestWithDefaults instantiates a new GcpkmsConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *GcpkmsConfigRequest) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GcpkmsConfigRequest) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GcpkmsConfigRequest) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *GcpkmsConfigRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetScopes

`func (o *GcpkmsConfigRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GcpkmsConfigRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GcpkmsConfigRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GcpkmsConfigRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


