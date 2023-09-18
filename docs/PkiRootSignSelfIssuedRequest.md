# PkiRootSignSelfIssuedRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | PEM-format self-issued certificate to be signed. | [optional] 
**IssuerRef** | Pointer to **string** | Reference to a existing issuer; either \&quot;default\&quot; for the configured default issuer, an identifier or the name assigned to the issuer. | [optional] [default to "default"]
**RequireMatchingCertificateAlgorithms** | Pointer to **bool** | If true, require the public key algorithm of the signer to match that of the self issued certificate. | [optional] [default to false]





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


