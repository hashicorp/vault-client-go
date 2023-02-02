# AzureWriteConfigRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**ClientId** | Pointer to **string** | The OAuth2 client id to connect to Azure. This value can also be provided with the AZURE_CLIENT_ID environment variable. | [optional] 
**ClientSecret** | Pointer to **string** | The OAuth2 client secret to connect to Azure. This value can also be provided with the AZURE_CLIENT_SECRET environment variable. | [optional] 
**Environment** | Pointer to **string** | The Azure environment name. If not provided, AzurePublicCloud is used. This value can also be provided with the AZURE_ENVIRONMENT environment variable. | [optional] 
**PasswordPolicy** | Pointer to **string** | Name of the password policy to use to generate passwords for dynamic credentials. | [optional] 
**RootPasswordTtl** | Pointer to **int32** | The TTL of the root password in Azure. This can be either a number of seconds or a time formatted duration (ex: 24h, 48ds) | [optional] 
**SubscriptionId** | Pointer to **string** | The subscription id for the Azure Active Directory. This value can also be provided with the AZURE_SUBSCRIPTION_ID environment variable. | [optional] 
**TenantId** | Pointer to **string** | The tenant id for the Azure Active Directory. This value can also be provided with the AZURE_TENANT_ID environment variable. | [optional] 



## Methods


### NewAzureWriteConfigRequest

`func NewAzureWriteConfigRequest() *AzureWriteConfigRequest`

NewAzureWriteConfigRequest instantiates a new AzureWriteConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureWriteConfigRequestWithDefaults

`func NewAzureWriteConfigRequestWithDefaults() *AzureWriteConfigRequest`

NewAzureWriteConfigRequestWithDefaults instantiates a new AzureWriteConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetClientId

`func (o *AzureWriteConfigRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureWriteConfigRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AzureWriteConfigRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### HasClientId

`func (o *AzureWriteConfigRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.




### GetClientSecret

`func (o *AzureWriteConfigRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AzureWriteConfigRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AzureWriteConfigRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### HasClientSecret

`func (o *AzureWriteConfigRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.




### GetEnvironment

`func (o *AzureWriteConfigRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AzureWriteConfigRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AzureWriteConfigRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### HasEnvironment

`func (o *AzureWriteConfigRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.




### GetPasswordPolicy

`func (o *AzureWriteConfigRequest) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *AzureWriteConfigRequest) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *AzureWriteConfigRequest) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.


### HasPasswordPolicy

`func (o *AzureWriteConfigRequest) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.




### GetRootPasswordTtl

`func (o *AzureWriteConfigRequest) GetRootPasswordTtl() int32`

GetRootPasswordTtl returns the RootPasswordTtl field if non-nil, zero value otherwise.

### GetRootPasswordTtlOk

`func (o *AzureWriteConfigRequest) GetRootPasswordTtlOk() (*int32, bool)`

GetRootPasswordTtlOk returns a tuple with the RootPasswordTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPasswordTtl

`func (o *AzureWriteConfigRequest) SetRootPasswordTtl(v int32)`

SetRootPasswordTtl sets RootPasswordTtl field to given value.


### HasRootPasswordTtl

`func (o *AzureWriteConfigRequest) HasRootPasswordTtl() bool`

HasRootPasswordTtl returns a boolean if a field has been set.




### GetSubscriptionId

`func (o *AzureWriteConfigRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureWriteConfigRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureWriteConfigRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### HasSubscriptionId

`func (o *AzureWriteConfigRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.




### GetTenantId

`func (o *AzureWriteConfigRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureWriteConfigRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureWriteConfigRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### HasTenantId

`func (o *AzureWriteConfigRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


