# KubernetesWriteConfigRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------


**DisableLocalCaJwt** | Pointer to **bool** | Disable defaulting to the local CA certificate and service account JWT when running in a Kubernetes pod. | [optional] [default to false]
**KubernetesCaCert** | Pointer to **string** | PEM encoded CA certificate to use to verify the Kubernetes API server certificate. Defaults to the local pod&#x27;s CA if found. | [optional] 
**KubernetesHost** | Pointer to **string** | Kubernetes API URL to connect to. Defaults to https://$KUBERNETES_SERVICE_HOST:KUBERNETES_SERVICE_PORT if those environment variables are set. | [optional] 
**ServiceAccountJwt** | Pointer to **string** | The JSON web token of the service account used by the secret engine to manage Kubernetes credentials. Defaults to the local pod&#x27;s JWT if found. | [optional] 



## Methods


### NewKubernetesWriteConfigRequest

`func NewKubernetesWriteConfigRequest() *KubernetesWriteConfigRequest`

NewKubernetesWriteConfigRequest instantiates a new KubernetesWriteConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesWriteConfigRequestWithDefaults

`func NewKubernetesWriteConfigRequestWithDefaults() *KubernetesWriteConfigRequest`

NewKubernetesWriteConfigRequestWithDefaults instantiates a new KubernetesWriteConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetDisableLocalCaJwt

`func (o *KubernetesWriteConfigRequest) GetDisableLocalCaJwt() bool`

GetDisableLocalCaJwt returns the DisableLocalCaJwt field if non-nil, zero value otherwise.

### GetDisableLocalCaJwtOk

`func (o *KubernetesWriteConfigRequest) GetDisableLocalCaJwtOk() (*bool, bool)`

GetDisableLocalCaJwtOk returns a tuple with the DisableLocalCaJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableLocalCaJwt

`func (o *KubernetesWriteConfigRequest) SetDisableLocalCaJwt(v bool)`

SetDisableLocalCaJwt sets DisableLocalCaJwt field to given value.


### HasDisableLocalCaJwt

`func (o *KubernetesWriteConfigRequest) HasDisableLocalCaJwt() bool`

HasDisableLocalCaJwt returns a boolean if a field has been set.




### GetKubernetesCaCert

`func (o *KubernetesWriteConfigRequest) GetKubernetesCaCert() string`

GetKubernetesCaCert returns the KubernetesCaCert field if non-nil, zero value otherwise.

### GetKubernetesCaCertOk

`func (o *KubernetesWriteConfigRequest) GetKubernetesCaCertOk() (*string, bool)`

GetKubernetesCaCertOk returns a tuple with the KubernetesCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCaCert

`func (o *KubernetesWriteConfigRequest) SetKubernetesCaCert(v string)`

SetKubernetesCaCert sets KubernetesCaCert field to given value.


### HasKubernetesCaCert

`func (o *KubernetesWriteConfigRequest) HasKubernetesCaCert() bool`

HasKubernetesCaCert returns a boolean if a field has been set.




### GetKubernetesHost

`func (o *KubernetesWriteConfigRequest) GetKubernetesHost() string`

GetKubernetesHost returns the KubernetesHost field if non-nil, zero value otherwise.

### GetKubernetesHostOk

`func (o *KubernetesWriteConfigRequest) GetKubernetesHostOk() (*string, bool)`

GetKubernetesHostOk returns a tuple with the KubernetesHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesHost

`func (o *KubernetesWriteConfigRequest) SetKubernetesHost(v string)`

SetKubernetesHost sets KubernetesHost field to given value.


### HasKubernetesHost

`func (o *KubernetesWriteConfigRequest) HasKubernetesHost() bool`

HasKubernetesHost returns a boolean if a field has been set.




### GetServiceAccountJwt

`func (o *KubernetesWriteConfigRequest) GetServiceAccountJwt() string`

GetServiceAccountJwt returns the ServiceAccountJwt field if non-nil, zero value otherwise.

### GetServiceAccountJwtOk

`func (o *KubernetesWriteConfigRequest) GetServiceAccountJwtOk() (*string, bool)`

GetServiceAccountJwtOk returns a tuple with the ServiceAccountJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountJwt

`func (o *KubernetesWriteConfigRequest) SetServiceAccountJwt(v string)`

SetServiceAccountJwt sets ServiceAccountJwt field to given value.


### HasServiceAccountJwt

`func (o *KubernetesWriteConfigRequest) HasServiceAccountJwt() bool`

HasServiceAccountJwt returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


