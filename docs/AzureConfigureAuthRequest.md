# AzureConfigureAuthRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The OAuth2 client id to connection to Azure. This value can also be provided with the AZURE_CLIENT_ID environment variable. | [optional] 
**ClientSecret** | Pointer to **string** | The OAuth2 client secret to connection to Azure. This value can also be provided with the AZURE_CLIENT_SECRET environment variable. | [optional] 
**Environment** | Pointer to **string** | The Azure environment name. If not provided, AzurePublicCloud is used. This value can also be provided with the AZURE_ENVIRONMENT environment variable. | [optional] 
**Resource** | Pointer to **string** | The resource URL for the vault application in Azure Active Directory. This value can also be provided with the AZURE_AD_RESOURCE environment variable. | [optional] 
**RootPasswordTtl** | Pointer to **int32** | The TTL of the root password in Azure. This can be either a number of seconds or a time formatted duration (ex: 24h, 48ds) | [optional] 
**TenantId** | Pointer to **string** | The tenant id for the Azure Active Directory. This is sometimes referred to as Directory ID in AD. This value can also be provided with the AZURE_TENANT_ID environment variable. | [optional] 



## Methods


### NewAzureConfigureAuthRequest

`func NewAzureConfigureAuthRequest() *AzureConfigureAuthRequest`

NewAzureConfigureAuthRequest instantiates a new AzureConfigureAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureConfigureAuthRequestWithDefaults

`func NewAzureConfigureAuthRequestWithDefaults() *AzureConfigureAuthRequest`

NewAzureConfigureAuthRequestWithDefaults instantiates a new AzureConfigureAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetClientId

`func (o *AzureConfigureAuthRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureConfigureAuthRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AzureConfigureAuthRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### HasClientId

`func (o *AzureConfigureAuthRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.




### GetClientSecret

`func (o *AzureConfigureAuthRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AzureConfigureAuthRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AzureConfigureAuthRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### HasClientSecret

`func (o *AzureConfigureAuthRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.




### GetEnvironment

`func (o *AzureConfigureAuthRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AzureConfigureAuthRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AzureConfigureAuthRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### HasEnvironment

`func (o *AzureConfigureAuthRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.




### GetResource

`func (o *AzureConfigureAuthRequest) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AzureConfigureAuthRequest) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AzureConfigureAuthRequest) SetResource(v string)`

SetResource sets Resource field to given value.


### HasResource

`func (o *AzureConfigureAuthRequest) HasResource() bool`

HasResource returns a boolean if a field has been set.




### GetRootPasswordTtl

`func (o *AzureConfigureAuthRequest) GetRootPasswordTtl() int32`

GetRootPasswordTtl returns the RootPasswordTtl field if non-nil, zero value otherwise.

### GetRootPasswordTtlOk

`func (o *AzureConfigureAuthRequest) GetRootPasswordTtlOk() (*int32, bool)`

GetRootPasswordTtlOk returns a tuple with the RootPasswordTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPasswordTtl

`func (o *AzureConfigureAuthRequest) SetRootPasswordTtl(v int32)`

SetRootPasswordTtl sets RootPasswordTtl field to given value.


### HasRootPasswordTtl

`func (o *AzureConfigureAuthRequest) HasRootPasswordTtl() bool`

HasRootPasswordTtl returns a boolean if a field has been set.




### GetTenantId

`func (o *AzureConfigureAuthRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureConfigureAuthRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureConfigureAuthRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### HasTenantId

`func (o *AzureConfigureAuthRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


