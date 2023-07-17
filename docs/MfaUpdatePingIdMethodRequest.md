# MfaUpdatePingIdMethodRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodName** | Pointer to **string** | The unique name identifier for this MFA method. | [optional] 
**SettingsFileBase64** | Pointer to **string** | The settings file provided by Ping, Base64-encoded. This must be a settings file suitable for third-party clients, not the PingID SDK or PingFederate. | [optional] 
**UsernameFormat** | Pointer to **string** | A template string for mapping Identity names to MFA method names. Values to subtitute should be placed in {{}}. For example, \&quot;{{alias.name}}@example.com\&quot;. Currently-supported mappings: alias.name: The name returned by the mount configured via the mount_accessor parameter If blank, the Alias&#x27;s name field will be used as-is. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


