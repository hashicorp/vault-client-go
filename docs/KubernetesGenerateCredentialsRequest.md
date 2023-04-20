# KubernetesGenerateCredentialsRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audiences** | Pointer to **[]string** | The intended audiences of the generated credentials | [optional] 
**ClusterRoleBinding** | Pointer to **bool** | If true, generate a ClusterRoleBinding to grant permissions across the whole cluster instead of within a namespace. Requires the Vault role to have kubernetes_role_type set to ClusterRole. | [optional] 
**KubernetesNamespace** | **string** | The name of the Kubernetes namespace in which to generate the credentials | 
**Ttl** | Pointer to **int32** | The TTL of the generated credentials | [optional] 



## Methods


### NewKubernetesGenerateCredentialsRequest

`func NewKubernetesGenerateCredentialsRequest(kubernetesNamespace string, ) *KubernetesGenerateCredentialsRequest`

NewKubernetesGenerateCredentialsRequest instantiates a new KubernetesGenerateCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesGenerateCredentialsRequestWithDefaults

`func NewKubernetesGenerateCredentialsRequestWithDefaults() *KubernetesGenerateCredentialsRequest`

NewKubernetesGenerateCredentialsRequestWithDefaults instantiates a new KubernetesGenerateCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAudiences

`func (o *KubernetesGenerateCredentialsRequest) GetAudiences() []string`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *KubernetesGenerateCredentialsRequest) GetAudiencesOk() (*[]string, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *KubernetesGenerateCredentialsRequest) SetAudiences(v []string)`

SetAudiences sets Audiences field to given value.


### HasAudiences

`func (o *KubernetesGenerateCredentialsRequest) HasAudiences() bool`

HasAudiences returns a boolean if a field has been set.




### GetClusterRoleBinding

`func (o *KubernetesGenerateCredentialsRequest) GetClusterRoleBinding() bool`

GetClusterRoleBinding returns the ClusterRoleBinding field if non-nil, zero value otherwise.

### GetClusterRoleBindingOk

`func (o *KubernetesGenerateCredentialsRequest) GetClusterRoleBindingOk() (*bool, bool)`

GetClusterRoleBindingOk returns a tuple with the ClusterRoleBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRoleBinding

`func (o *KubernetesGenerateCredentialsRequest) SetClusterRoleBinding(v bool)`

SetClusterRoleBinding sets ClusterRoleBinding field to given value.


### HasClusterRoleBinding

`func (o *KubernetesGenerateCredentialsRequest) HasClusterRoleBinding() bool`

HasClusterRoleBinding returns a boolean if a field has been set.




### GetKubernetesNamespace

`func (o *KubernetesGenerateCredentialsRequest) GetKubernetesNamespace() string`

GetKubernetesNamespace returns the KubernetesNamespace field if non-nil, zero value otherwise.

### GetKubernetesNamespaceOk

`func (o *KubernetesGenerateCredentialsRequest) GetKubernetesNamespaceOk() (*string, bool)`

GetKubernetesNamespaceOk returns a tuple with the KubernetesNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesNamespace

`func (o *KubernetesGenerateCredentialsRequest) SetKubernetesNamespace(v string)`

SetKubernetesNamespace sets KubernetesNamespace field to given value.





### GetTtl

`func (o *KubernetesGenerateCredentialsRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *KubernetesGenerateCredentialsRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *KubernetesGenerateCredentialsRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *KubernetesGenerateCredentialsRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


