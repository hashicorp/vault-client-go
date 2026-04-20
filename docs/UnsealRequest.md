# UnsealRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Specifies a single unseal key share. This is required unless reset is true. | [optional] 
**Reset** | Pointer to **bool** | Specifies if previously-provided unseal keys are discarded and the unseal process is reset. | [optional] 
**Migrate** | Pointer to **bool** | Available in 1.0 - Used to migrate the seal from shamir to autoseal or autoseal to shamir. Must be provided on all unseal key calls. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


