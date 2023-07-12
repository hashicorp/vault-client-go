# AwsConfigureIdentityWhitelistTidyOperationRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisablePeriodicTidy** | Pointer to **bool** | If set to &#x27;true&#x27;, disables the periodic tidying of the &#x27;identity-accesslist/&lt;instance_id&gt;&#x27; entries. | [optional] [default to false]
**SafetyBuffer** | Pointer to **string** | The amount of extra time that must have passed beyond the identity&#x27;s expiration, before it is removed from the backend storage. | [optional] [default to "259200"]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


