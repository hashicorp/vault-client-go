# KubernetesConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableLocalCaJwt** | Pointer to **bool** | Disable defaulting to the local CA certificate and service account JWT when running in a Kubernetes pod. | [optional] [default to false]
**KubernetesCaCert** | Pointer to **string** | PEM encoded CA certificate to use to verify the Kubernetes API server certificate. Defaults to the local pod&#x27;s CA if found. | [optional] 
**KubernetesHost** | Pointer to **string** | Kubernetes API URL to connect to. Defaults to https://$KUBERNETES_SERVICE_HOST:KUBERNETES_SERVICE_PORT if those environment variables are set. | [optional] 
**ServiceAccountJwt** | Pointer to **string** | The JSON web token of the service account used by the secret engine to manage Kubernetes credentials. Defaults to the local pod&#x27;s JWT if found. | [optional] 





[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


