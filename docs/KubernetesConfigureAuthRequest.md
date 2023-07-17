# KubernetesConfigureAuthRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableIssValidation** | Pointer to **bool** | Disable JWT issuer validation (Deprecated, will be removed in a future release) | [optional] [default to true]
**DisableLocalCaJwt** | Pointer to **bool** | Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod | [optional] [default to false]
**Issuer** | Pointer to **string** | Optional JWT issuer. If no issuer is specified, then this plugin will use kubernetes.io/serviceaccount as the default issuer. (Deprecated, will be removed in a future release) | [optional] 
**KubernetesCaCert** | Pointer to **string** | PEM encoded CA cert for use by the TLS client used to talk with the API. | [optional] 
**KubernetesHost** | Pointer to **string** | Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server. | [optional] 
**PemKeys** | Pointer to **[]string** | Optional list of PEM-formated public keys or certificates used to verify the signatures of kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys. | [optional] 
**TokenReviewerJwt** | Pointer to **string** | A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


