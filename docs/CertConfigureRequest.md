# CertConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableBinding** | Pointer to **bool** | If set, during renewal, skips the matching of presented client identity with the client identity used during login. Defaults to false. | [optional] [default to false]
**EnableIdentityAliasMetadata** | Pointer to **bool** | If set, metadata of the certificate including the metadata corresponding to allowed_metadata_extensions will be stored in the alias. Defaults to false. | [optional] [default to false]
**OcspCacheSize** | Pointer to **int32** | The size of the in memory OCSP response cache, shared by all configured certs | [optional] [default to 100]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


