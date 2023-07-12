# LdapLibraryConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableCheckInEnforcement** | Pointer to **bool** | Disable the default behavior of requiring that check-ins are performed by the entity that checked them out. | [optional] [default to false]
**MaxTtl** | Pointer to **string** | In seconds, the max amount of time a check-out&#x27;s renewals should last. Defaults to 24 hours. | [optional] [default to "86400"]
**ServiceAccountNames** | Pointer to **[]string** | The username/logon name for the service accounts with which this set will be associated. | [optional] 
**Ttl** | Pointer to **string** | In seconds, the amount of time a check-out should last. Defaults to 24 hours. | [optional] [default to "86400"]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


