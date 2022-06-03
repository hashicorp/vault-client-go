# AzureLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jwt** | Pointer to **string** | A signed JWT | [optional] 
**ResourceGroupName** | Pointer to **string** | The resource group from the instance. | [optional] 
**Role** | Pointer to **string** | The token role. | [optional] 
**SubscriptionId** | Pointer to **string** | The subscription id for the instance. | [optional] 
**VmName** | Pointer to **string** | The name of the virtual machine. This value is ignored if vmss_name is specified. | [optional] 
**VmssName** | Pointer to **string** | The name of the virtual machine scale set the instance is in. | [optional] 

## Methods

### NewAzureLoginRequest

`func NewAzureLoginRequest() *AzureLoginRequest`

NewAzureLoginRequest instantiates a new AzureLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureLoginRequestWithDefaults

`func NewAzureLoginRequestWithDefaults() *AzureLoginRequest`

NewAzureLoginRequestWithDefaults instantiates a new AzureLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwt

`func (o *AzureLoginRequest) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *AzureLoginRequest) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *AzureLoginRequest) SetJwt(v string)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *AzureLoginRequest) HasJwt() bool`

HasJwt returns a boolean if a field has been set.

### GetResourceGroupName

`func (o *AzureLoginRequest) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureLoginRequest) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureLoginRequest) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AzureLoginRequest) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### GetRole

`func (o *AzureLoginRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AzureLoginRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AzureLoginRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AzureLoginRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *AzureLoginRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureLoginRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureLoginRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzureLoginRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetVmName

`func (o *AzureLoginRequest) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *AzureLoginRequest) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *AzureLoginRequest) SetVmName(v string)`

SetVmName sets VmName field to given value.

### HasVmName

`func (o *AzureLoginRequest) HasVmName() bool`

HasVmName returns a boolean if a field has been set.

### GetVmssName

`func (o *AzureLoginRequest) GetVmssName() string`

GetVmssName returns the VmssName field if non-nil, zero value otherwise.

### GetVmssNameOk

`func (o *AzureLoginRequest) GetVmssNameOk() (*string, bool)`

GetVmssNameOk returns a tuple with the VmssName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmssName

`func (o *AzureLoginRequest) SetVmssName(v string)`

SetVmssName sets VmssName field to given value.

### HasVmssName

`func (o *AzureLoginRequest) HasVmssName() bool`

HasVmssName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


