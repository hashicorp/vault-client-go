# AuthTuneConfigurationParametersRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedResponseHeaders** | Pointer to **[]string** | A list of headers to whitelist and allow a plugin to set on responses. | [optional] 
**AuditNonHmacRequestKeys** | Pointer to **[]string** | The list of keys in the request data object that will not be HMAC&#x27;ed by audit devices. | [optional] 
**AuditNonHmacResponseKeys** | Pointer to **[]string** | The list of keys in the response data object that will not be HMAC&#x27;ed by audit devices. | [optional] 
**DefaultLeaseTtl** | Pointer to **string** | The default lease TTL for this mount. | [optional] 
**Description** | Pointer to **string** | User-friendly description for this credential backend. | [optional] 
**ListingVisibility** | Pointer to **string** | Determines the visibility of the mount in the UI-specific listing endpoint. Accepted value are &#x27;unauth&#x27; and &#x27;hidden&#x27;, with the empty default (&#x27;&#x27;) behaving like &#x27;hidden&#x27;. | [optional] 
**MaxLeaseTtl** | Pointer to **string** | The max lease TTL for this mount. | [optional] 
**Options** | Pointer to **map[string]interface{}** | The options to pass into the backend. Should be a json object with string keys and values. | [optional] 
**PassthroughRequestHeaders** | Pointer to **[]string** | A list of headers to whitelist and pass from the request to the plugin. | [optional] 
**PluginVersion** | Pointer to **string** | The semantic version of the plugin to use, or image tag if oci_image is provided. | [optional] 
**TokenType** | Pointer to **string** | The type of token to issue (service or batch). | [optional] 
**UserLockoutConfig** | Pointer to **map[string]interface{}** | The user lockout configuration to pass into the backend. Should be a json object with string keys and values. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


