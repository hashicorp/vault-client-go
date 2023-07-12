# RateLimitQuotasConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableRateLimitAuditLogging** | Pointer to **bool** | If set, starts audit logging of requests that get rejected due to rate limit quota rule violations. | [optional] 
**EnableRateLimitResponseHeaders** | Pointer to **bool** | If set, additional rate limit quota HTTP headers will be added to responses. | [optional] 
**RateLimitExemptPaths** | Pointer to **[]string** | Specifies the list of exempt paths from all rate limit quotas. If empty no paths will be exempt. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


