# AzureLoginRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jwt** | Pointer to **string** | A signed JWT | [optional] 
**ResourceGroupName** | Pointer to **string** | The resource group from the instance. | [optional] 
**ResourceId** | Pointer to **string** | The fully qualified ID of the resource, includingthe resource name and resource type. Use the format, /subscriptions/{guid}/resourceGroups/{resource-group-name}/{resource-provider-namespace}/{resource-type}/{resource-name}. This value is ignored if vm_name or vmss_name is specified. | [optional] 
**Role** | Pointer to **string** | The token role. | [optional] 
**SubscriptionId** | Pointer to **string** | The subscription id for the instance. | [optional] 
**VmName** | Pointer to **string** | The name of the virtual machine. This value is ignored if vmss_name is specified. | [optional] 
**VmssName** | Pointer to **string** | The name of the virtual machine scale set the instance is in. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


