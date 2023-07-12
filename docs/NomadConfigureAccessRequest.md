# NomadConfigureAccessRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Nomad server address | [optional] 
**CaCert** | Pointer to **string** | CA certificate to use when verifying Nomad server certificate, must be x509 PEM encoded. | [optional] 
**ClientCert** | Pointer to **string** | Client certificate used for Nomad&#x27;s TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key. | [optional] 
**ClientKey** | Pointer to **string** | Client key used for Nomad&#x27;s TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert. | [optional] 
**MaxTokenNameLength** | Pointer to **int32** | Max length for name of generated Nomad tokens | [optional] 
**Token** | Pointer to **string** | Token for API calls | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


