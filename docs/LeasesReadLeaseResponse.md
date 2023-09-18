# LeasesReadLeaseResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpireTime** | Pointer to **time.Time** | Optional lease expiry time | [optional] 
**Id** | Pointer to **string** | Lease id | [optional] 
**IssueTime** | Pointer to **time.Time** | Timestamp for the lease&#x27;s issue time | [optional] 
**LastRenewal** | Pointer to **time.Time** | Optional Timestamp of the last time the lease was renewed | [optional] 
**Renewable** | Pointer to **bool** | True if the lease is able to be renewed | [optional] 
**Ttl** | Pointer to **int32** | Time to Live set for the lease, returns 0 if unset | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


