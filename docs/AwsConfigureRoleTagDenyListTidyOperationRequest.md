# AwsConfigureRoleTagDenyListTidyOperationRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisablePeriodicTidy** | Pointer to **bool** | If set to &#x27;true&#x27;, disables the periodic tidying of deny listed entries. | [optional] [default to false]
**SafetyBuffer** | Pointer to **string** | The amount of extra time that must have passed beyond the roletag expiration, before it is removed from the backend storage. Defaults to 4320h (180 days). | [optional] [default to "15552000"]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


